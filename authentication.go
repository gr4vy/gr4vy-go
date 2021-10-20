package gr4vy

import (
	"fmt"
	"github.com/golang-jwt/jwt"
	"golang.org/x/crypto/ssh"
	"time"
	"crypto/ecdsa"
	"crypto/sha256"
	"encoding/base64"
	"encoding/json"
	"github.com/twinj/uuid"
	"runtime"
)

func getEmbedToken(private_key string, embed EmbedParams) (string, error) {
	claims := jwt.MapClaims{
		"iss": fmt.Sprintf("Gr4vy SDK %v - %v", VERSION, runtime.Version()), 
		"nbf": float64(time.Now().Unix()),
		"exp": float64(time.Now().Unix() + 3000),
		"scopes": []string{"embed"},
		"jti": uuid.NewV4(),
	}

	var inInterface map[string]interface{}
    inrec, _ := json.Marshal(embed)
    json.Unmarshal(inrec, &inInterface)

	claims["embed"] = inInterface

	return getTokenWithClaims(private_key, claims)
}

func getToken(private_key string, scopes []string) (string, error) {
	claims := jwt.MapClaims{
		"iss": fmt.Sprintf("Gr4vy SDK %v - %v", VERSION, runtime.Version()), 
		"nbf": float64(time.Now().Unix()),
		"exp": float64(time.Now().Unix() + 3000),
		"scopes": scopes,
		"jti": uuid.NewV4(),
	}

	return getTokenWithClaims(private_key, claims)
}

func getTokenWithClaims(private_key string, claims jwt.MapClaims) (string, error) {
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

func authentication(private_key string, Debug bool) (string, error) {

	scopes := []string{"*.read", "*.write"}
	tokenString, err := getToken(private_key, scopes)
	if err != nil {
		return "", err
	}
	if Debug {
		fmt.Printf("Gr4vy - Auth Token - %v\n", tokenString)
	}

	return tokenString, nil
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