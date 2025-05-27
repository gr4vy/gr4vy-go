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

	if !strings.HasPrefix(claims.Issuer, "Gr4vy Python SDK") {
		t.Errorf("Expected issuer to start with 'Gr4vy Python SDK', got %s", claims.Issuer)
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
