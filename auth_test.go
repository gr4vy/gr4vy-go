package gr4vygo

import (
	"crypto/ecdsa"
	"crypto/x509"
	"encoding/pem"
	"strings"
	"testing"
	"time"

	"github.com/golang-jwt/jwt/v5"
)

const testPrivateKeyPEM = `-----BEGIN PRIVATE KEY-----
MIHuAgEAMBAGByqGSM49AgEGBSuBBAAjBIHWMIHTAgEBBEIBABM9jQu+HT87oIik
O6DiJjYeghr3V+VMBVNU2hCM3X/OAS6TUTylMbnjDnwWdmu7anVSnjvEY1a4KxQ9
WZ8E/PKhgYkDgYYABABRdv5VAtOsGb6THxeK/p7RAARPm6Zwb7FF4sZAYkkSB7h0
2jpj3UHSpyl92BQkiF/xakz7hMMD1A0ZTn5SuXWp3AG9qPHO3eB9WrZhPGYixwyo
XNjhnPEDhmkItKXteke9iBOTOOXB7AFQSh7EXRBmhBs4u3ZlTmrl+8VdBc3+jwAY
rw==
-----END PRIVATE KEY-----`

var testEmbedParams = map[string]interface{}{
	"amount":                    9000,
	"currency":                  "USD",
	"buyer_external_identifier": "user-123",
	"connection_options": map[string]interface{}{
		"stripe-card": map[string]interface{}{
			"stripe_connect": map[string]interface{}{
				"key": "value",
			},
		},
	},
	"metadata": map[string]interface{}{
		"camelCaseKey":   "value1",
		"snake_case_key": "value2",
	},
	"cart_items": []map[string]interface{}{
		{
			"name":        "Joust Duffle Bag",
			"quantity":    1,
			"unit_amount": 9000,
			"tax_amount":  0,
			"categories":  []string{"Gear", "Bags", "Test"},
		},
	},
}

const testCheckoutSessionID = "0ebde6a1-f66c-43ea-bb8b-73751864c604"

func TestGetTokenCreatesValidSignedJWT(t *testing.T) {
	token, err := GetToken(testPrivateKeyPEM, []JWTScope{ReadAll, WriteAll}, 3600)
	if err != nil {
		t.Fatalf("Failed to create token: %v", err)
	}

	// Parse and validate token
	parsedToken, err := jwt.ParseWithClaims(token, &TokenClaims{}, func(token *jwt.Token) (interface{}, error) {
		block, _ := pem.Decode([]byte(testPrivateKeyPEM))
		privateKey, _ := x509.ParsePKCS8PrivateKey(block.Bytes)
		ecdsaKey := privateKey.(*ecdsa.PrivateKey)
		return &ecdsaKey.PublicKey, nil
	})

	if err != nil {
		t.Fatalf("Failed to parse token: %v", err)
	}

	claims := parsedToken.Claims.(*TokenClaims)

	if len(claims.Scopes) != 2 || claims.Scopes[0] != "*.read" || claims.Scopes[1] != "*.write" {
		t.Errorf("Expected scopes [*.read, *.write], got %v", claims.Scopes)
	}

	if claims.IssuedAt == nil {
		t.Error("Expected iat claim to be present")
	}

	if claims.NotBefore == nil {
		t.Error("Expected nbf claim to be present")
	}

	if claims.ExpiresAt == nil {
		t.Error("Expected exp claim to be present")
	}

	if !strings.HasPrefix(claims.Issuer, "speakeasy-sdk/go") {
		t.Errorf("Expected issuer to start with 'speakeasy-sdk/go', got %s", claims.Issuer)
	}
}

// expectedKid is the RFC 7638 JWK thumbprint of testPrivateKeyPEM's public key,
// computed independently (openssl + a standards-compliant JWK library) with the
// EC coordinates zero-padded to the P-521 curve size. testPrivateKeyPEM's X
// coordinate has a leading zero byte, so a thumbprint computed from the
// unpadded big.Int.Bytes() representation would differ from this value — which
// is exactly the bug that caused 401 "No valid API authentication found".
const expectedKid = "va-SLs5AxJNfqKXD8LI5Y38BflpNvjZjY4RSWz66U1w"

func TestGetTokenSetsCorrectKidHeader(t *testing.T) {
	token, err := GetToken(testPrivateKeyPEM, []JWTScope{ReadAll, WriteAll}, 3600)
	if err != nil {
		t.Fatalf("Failed to create token: %v", err)
	}

	parsedToken, _, err := new(jwt.Parser).ParseUnverified(token, &TokenClaims{})
	if err != nil {
		t.Fatalf("Failed to parse token: %v", err)
	}

	kid, ok := parsedToken.Header["kid"].(string)
	if !ok {
		t.Fatalf("Expected kid header to be a string, got %v", parsedToken.Header["kid"])
	}

	if kid != expectedKid {
		t.Errorf("Expected kid %q, got %q (EC coordinates must be zero-padded to the curve size)", expectedKid, kid)
	}
}

func TestCalculateThumbprintZeroPadsCoordinates(t *testing.T) {
	block, _ := pem.Decode([]byte(testPrivateKeyPEM))
	if block == nil {
		t.Fatal("Failed to decode test PEM block")
	}
	privateKey, err := x509.ParsePKCS8PrivateKey(block.Bytes)
	if err != nil {
		t.Fatalf("Failed to parse private key: %v", err)
	}
	ecdsaKey, ok := privateKey.(*ecdsa.PrivateKey)
	if !ok {
		t.Fatalf("Test key is not ECDSA: %T", privateKey)
	}

	// Sanity check: this fixture exercises the bug — its X coordinate is shorter
	// than the curve byte size, so it must be left-padded for a correct thumbprint.
	coordLen := (ecdsaKey.PublicKey.Curve.Params().BitSize + 7) / 8
	if len(ecdsaKey.PublicKey.X.Bytes()) >= coordLen {
		t.Fatalf("test fixture no longer exercises the padding case: X is %d bytes, curve size is %d", len(ecdsaKey.PublicKey.X.Bytes()), coordLen)
	}

	thumbprint, err := calculateThumbprint(ecdsaKey)
	if err != nil {
		t.Fatalf("Failed to calculate thumbprint: %v", err)
	}

	if thumbprint != expectedKid {
		t.Errorf("Expected thumbprint %q, got %q", expectedKid, thumbprint)
	}
}

func TestGetTokenAcceptsOptionalEmbedData(t *testing.T) {
	token, err := GetTokenWithEmbedProperties(testPrivateKeyPEM, []JWTScope{Embed}, 3600, testEmbedParams, "")
	if err != nil {
		t.Fatalf("Failed to create token: %v", err)
	}

	parsedToken, err := jwt.ParseWithClaims(token, &TokenClaims{}, func(token *jwt.Token) (interface{}, error) {
		block, _ := pem.Decode([]byte(testPrivateKeyPEM))
		privateKey, _ := x509.ParsePKCS8PrivateKey(block.Bytes)
		ecdsaKey := privateKey.(*ecdsa.PrivateKey)
		return &ecdsaKey.PublicKey, nil
	})

	if err != nil {
		t.Fatalf("Failed to parse token: %v", err)
	}

	claims := parsedToken.Claims.(*TokenClaims)

	if len(claims.Scopes) != 1 || claims.Scopes[0] != "embed" {
		t.Errorf("Expected scopes [embed], got %v", claims.Scopes)
	}

	if claims.Embed == nil {
		t.Error("Expected embed claim to be present")
	}

	if claims.Embed["amount"] != float64(9000) {
		t.Errorf("Expected embed amount 9000, got %v", claims.Embed["amount"])
	}
}

func TestGetTokenIgnoresEmbedDataWithoutEmbedScope(t *testing.T) {
	token, err := GetTokenWithEmbedProperties(testPrivateKeyPEM, []JWTScope{ReadAll}, 3600, testEmbedParams, "")
	if err != nil {
		t.Fatalf("Failed to create token: %v", err)
	}

	parsedToken, err := jwt.ParseWithClaims(token, &TokenClaims{}, func(token *jwt.Token) (interface{}, error) {
		block, _ := pem.Decode([]byte(testPrivateKeyPEM))
		privateKey, _ := x509.ParsePKCS8PrivateKey(block.Bytes)
		ecdsaKey := privateKey.(*ecdsa.PrivateKey)
		return &ecdsaKey.PublicKey, nil
	})

	if err != nil {
		t.Fatalf("Failed to parse token: %v", err)
	}

	claims := parsedToken.Claims.(*TokenClaims)

	if len(claims.Scopes) != 1 || claims.Scopes[0] != "*.read" {
		t.Errorf("Expected scopes [*.read], got %v", claims.Scopes)
	}

	if claims.Embed != nil {
		t.Error("Expected embed claim to be nil when embed scope is not present")
	}
}

func TestGetEmbedTokenCreatesJWTForEmbed(t *testing.T) {
	token, err := GetEmbedToken(testPrivateKeyPEM, testEmbedParams, "")
	if err != nil {
		t.Fatalf("Failed to create embed token: %v", err)
	}

	parsedToken, err := jwt.ParseWithClaims(token, &TokenClaims{}, func(token *jwt.Token) (interface{}, error) {
		block, _ := pem.Decode([]byte(testPrivateKeyPEM))
		privateKey, _ := x509.ParsePKCS8PrivateKey(block.Bytes)
		ecdsaKey := privateKey.(*ecdsa.PrivateKey)
		return &ecdsaKey.PublicKey, nil
	})

	if err != nil {
		t.Fatalf("Failed to parse token: %v", err)
	}

	claims := parsedToken.Claims.(*TokenClaims)

	if len(claims.Scopes) != 1 || claims.Scopes[0] != "embed" {
		t.Errorf("Expected scopes [embed], got %v", claims.Scopes)
	}

	if claims.Embed == nil {
		t.Error("Expected embed claim to be present")
	}
}

func TestGetEmbedTokenTakesOptionalCheckoutSessionID(t *testing.T) {
	token, err := GetEmbedToken(testPrivateKeyPEM, testEmbedParams, testCheckoutSessionID)
	if err != nil {
		t.Fatalf("Failed to create embed token: %v", err)
	}

	parsedToken, err := jwt.ParseWithClaims(token, &TokenClaims{}, func(token *jwt.Token) (interface{}, error) {
		block, _ := pem.Decode([]byte(testPrivateKeyPEM))
		privateKey, _ := x509.ParsePKCS8PrivateKey(block.Bytes)
		ecdsaKey := privateKey.(*ecdsa.PrivateKey)
		return &ecdsaKey.PublicKey, nil
	})

	if err != nil {
		t.Fatalf("Failed to parse token: %v", err)
	}

	claims := parsedToken.Claims.(*TokenClaims)

	if claims.CheckoutSessionID != testCheckoutSessionID {
		t.Errorf("Expected checkout_session_id %s, got %s", testCheckoutSessionID, claims.CheckoutSessionID)
	}
}

func TestUpdateTokenResignsWithNewSignatureAndExpiration(t *testing.T) {
	originalToken, err := GetToken(testPrivateKeyPEM, nil, 5)
	if err != nil {
		t.Fatalf("Failed to create original token: %v", err)
	}

	time.Sleep(1000 * time.Millisecond)

	newToken, err := UpdateToken(originalToken, testPrivateKeyPEM, nil, 60, nil, "")
	if err != nil {
		t.Fatalf("Failed to update token: %v", err)
	}

	originalParsed, _, err := new(jwt.Parser).ParseUnverified(originalToken, &TokenClaims{})
	if err != nil {
		t.Fatalf("Failed to parse original token: %v", err)
	}

	newParsed, _, err := new(jwt.Parser).ParseUnverified(newToken, &TokenClaims{})
	if err != nil {
		t.Fatalf("Failed to parse new token: %v", err)
	}

	originalClaims := originalParsed.Claims.(*TokenClaims)
	newClaims := newParsed.Claims.(*TokenClaims)

	if len(newClaims.Scopes) != len(originalClaims.Scopes) {
		t.Errorf("Expected scopes to be preserved, got %v vs %v", newClaims.Scopes, originalClaims.Scopes)
	}

	if newClaims.IssuedAt.Unix() <= originalClaims.IssuedAt.Unix() {
		t.Error("Expected new token to have later issued at time")
	}

	if newClaims.ExpiresAt.Unix() <= originalClaims.ExpiresAt.Unix() {
		t.Error("Expected new token to have later expiration time")
	}

	if newClaims.NotBefore.Unix() <= originalClaims.NotBefore.Unix() {
		t.Error("Expected new token to have later not before time")
	}
}

func TestUpdateTokenAllowsEmbedTokenUpdateWithNewParams(t *testing.T) {
	originalToken, err := GetEmbedToken(testPrivateKeyPEM, testEmbedParams, testCheckoutSessionID)
	if err != nil {
		t.Fatalf("Failed to create original embed token: %v", err)
	}

	newEmbedParams := map[string]interface{}{
		"amount":   1299,
		"currency": "USD",
	}

	time.Sleep(100 * time.Millisecond)

	newToken, err := UpdateToken(originalToken, testPrivateKeyPEM, nil, 3600, newEmbedParams, "")
	if err != nil {
		t.Fatalf("Failed to update token: %v", err)
	}

	originalParsed, _, err := new(jwt.Parser).ParseUnverified(originalToken, &TokenClaims{})
	if err != nil {
		t.Fatalf("Failed to parse original token: %v", err)
	}

	newParsed, _, err := new(jwt.Parser).ParseUnverified(newToken, &TokenClaims{})
	if err != nil {
		t.Fatalf("Failed to parse new token: %v", err)
	}

	originalClaims := originalParsed.Claims.(*TokenClaims)
	newClaims := newParsed.Claims.(*TokenClaims)

	if newClaims.Embed["amount"] != float64(1299) {
		t.Errorf("Expected new embed amount 1299, got %v", newClaims.Embed["amount"])
	}

	if newClaims.CheckoutSessionID != originalClaims.CheckoutSessionID {
		t.Errorf("Expected checkout session ID to be preserved: %s vs %s",
			newClaims.CheckoutSessionID, originalClaims.CheckoutSessionID)
	}
}
