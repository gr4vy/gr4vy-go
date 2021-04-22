package gr4vy

import (
	"fmt"
	. "github.com/steve-gr4vy/gr4vy/api"
	"github.com/deepmap/oapi-codegen/pkg/securityprovider"
	"github.com/dgrijalva/jwt-go"
	"golang.org/x/crypto/ssh"
	"time"
	"crypto/ecdsa"
	"crypto/sha256"
	"encoding/base64"
	"github.com/twinj/uuid"
	"runtime"
)

func getToken(private_key string, scopes []string, embed map[string]string) (string, error) {
	claims := jwt.MapClaims{
		"iss": fmt.Sprintf("Gr4vy SDK %v - %v", VERSION, runtime.Version()), 
		"nbf": float64(time.Now().Unix()),
		"exp": float64(time.Now().Unix() + 30),
		"scopes": scopes,
		"jti": uuid.NewV4(),
	}

	if embed != nil {
		claims["embed"] = embed
	}

	token := jwt.NewWithClaims(jwt.SigningMethodES512, claims)
	
	parsedKey, err := ssh.ParseRawPrivateKey([]byte(private_key))
	if err != nil {
		return "", err
	}

	var parsedEcdSaKey *ecdsa.PrivateKey
	parsedEcdSaKey = parsedKey.(*ecdsa.PrivateKey)

	kid, err := JWKThumbprint(parsedEcdSaKey.PublicKey)
	if err != nil {
		return "", err
	}

	token.Header["kid"] = kid

	tokenString, err := token.SignedString(parsedKey)
	if err != nil {
		return "", err
	}

	return tokenString, nil
}

func authentication(private_key string, Debug bool) (RequestEditorFn, error) {

	scopes := []string{"*.read", "*.write"}
	tokenString, err := getToken(private_key, scopes, nil)
	if err != nil {
		return nil, err
	}
	if Debug {
		fmt.Printf("Gr4vy - Auth Token - %v\n", tokenString)
	}

	bearerTokenProvider, bearerTokenProviderErr := securityprovider.NewSecurityProviderBearerToken(tokenString)
	if bearerTokenProviderErr != nil {
		return nil, bearerTokenProviderErr
	}

	return bearerTokenProvider.Intercept, nil
}

func jwkEncode(pub ecdsa.PublicKey) (string, error) {
	p := pub.Curve.Params()
	n := p.BitSize / 8
	if p.BitSize%8 != 0 {
		n++
	}
	x := pub.X.Bytes()
	if n > len(x) {
		x = append(make([]byte, n-len(x)), x...)
	}
	y := pub.Y.Bytes()
	if n > len(y) {
		y = append(make([]byte, n-len(y)), y...)
	}
	// Field order is important.
	// See https://tools.ietf.org/html/rfc7638#section-3.3 for details.
	return fmt.Sprintf(`{"crv":"%s","kty":"EC","x":"%s","y":"%s"}`,
		p.Name,
		base64.RawURLEncoding.EncodeToString(x),
		base64.RawURLEncoding.EncodeToString(y),
	), nil
}

func JWKThumbprint(pub ecdsa.PublicKey) (string, error) {
	jwk, err := jwkEncode(pub)
	if err != nil {
		return "", err
	}
	b := sha256.Sum256([]byte(jwk))
	return base64.RawURLEncoding.EncodeToString(b[:]), nil
}