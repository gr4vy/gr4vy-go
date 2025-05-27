package gr4vygo

import (
	"context"
	"crypto/ecdsa"
	"crypto/sha256"
	"crypto/x509"
	"encoding/base64"
	"encoding/json"
	"encoding/pem"
	"fmt"
	"time"

	"github.com/golang-jwt/jwt/v5"
	"github.com/google/uuid"
	"github.com/gr4vy/gr4vy-go/models/components"
)

// JWTScope represents the available JWT scopes
type JWTScope string

const (
	ReadAll                          JWTScope = "*.read"
	WriteAll                         JWTScope = "*.write"
	Embed                            JWTScope = "embed"
	AntiFraudServiceDefinitionsRead  JWTScope = "anti-fraud-service-definitions.read"
	AntiFraudServiceDefinitionsWrite JWTScope = "anti-fraud-service-definitions.write"
	AntiFraudServicesRead            JWTScope = "anti-fraud-services.read"
	AntiFraudServicesWrite           JWTScope = "anti-fraud-services.write"
	AuditLogsRead                    JWTScope = "audit-logs.read"
	BuyersRead                       JWTScope = "buyers.read"
	BuyersWrite                      JWTScope = "buyers.write"
	BuyersBillingDetailsRead         JWTScope = "buyers.billing-details.read"
	BuyersBillingDetailsWrite        JWTScope = "buyers.billing-details.write"
	CardSchemeDefinitionsRead        JWTScope = "card-scheme-definitions.read"
	CheckoutSessionsRead             JWTScope = "checkout-sessions.read"
	CheckoutSessionsWrite            JWTScope = "checkout-sessions.write"
	ConnectionsRead                  JWTScope = "connections.read"
	ConnectionsWrite                 JWTScope = "connections.write"
	DigitalWalletsRead               JWTScope = "digital-wallets.read"
	DigitalWalletsWrite              JWTScope = "digital-wallets.write"
	FlowsRead                        JWTScope = "flows.read"
	FlowsWrite                       JWTScope = "flows.write"
	GiftCardServiceDefinitionsRead   JWTScope = "gift-card-service-definitions.read"
	GiftCardServicesRead             JWTScope = "gift-card-services.read"
	GiftCardServicesWrite            JWTScope = "gift-card-services.write"
	GiftCardsRead                    JWTScope = "gift-cards.read"
	GiftCardsWrite                   JWTScope = "gift-cards.write"
	MerchantAccountRead              JWTScope = "merchant-accounts.reads"
	MerchantAccountWrite             JWTScope = "merchant-accounts.write"
	PaymentMethodDefinitionsRead     JWTScope = "payment-method-definitions.read"
	PaymentMethodRead                JWTScope = "payment-methods.read"
	PaymentMethodWrite               JWTScope = "payment-methods.write"
	PaymentOptionsRead               JWTScope = "payment-options.read"
	PaymentServiceDefinitionsRead    JWTScope = "payment-service-definitions.read"
	PaymentServicesRead              JWTScope = "payment-services.read"
	PaymentServicesWrite             JWTScope = "payment-services.write"
	ReportsRead                      JWTScope = "reports.read"
	ReportsWrite                     JWTScope = "reports.write"
	TransactionsRead                 JWTScope = "transactions.read"
	TransactionsWrite                JWTScope = "transactions.write"
	VaultForwardWrite                JWTScope = "vault-forward.write"
)

// TokenClaims represents the JWT claims structure
type TokenClaims struct {
	Scopes            []string               `json:"scopes"`
	Issuer            string                 `json:"iss"`
	IssuedAt          *jwt.NumericDate       `json:"iat"`
	NotBefore         *jwt.NumericDate       `json:"nbf"`
	ExpiresAt         *jwt.NumericDate       `json:"exp"`
	ID                string                 `json:"jti"`
	CheckoutSessionID string                 `json:"checkout_session_id,omitempty"`
	Embed             map[string]interface{} `json:"embed,omitempty"`
	jwt.RegisteredClaims
}

// TokenFunc is a function that returns a JWT token
type TokenFunc func() string

// GetToken generates a JWT token with the specified parameters
func GetToken(privateKeyPEM string, scopes []JWTScope, expiresIn int, embedParams map[string]interface{}, checkoutSessionID string) (string, error) {
	// Parse private key
	block, _ := pem.Decode([]byte(privateKeyPEM))
	if block == nil {
		return "", fmt.Errorf("failed to parse PEM block containing the private key")
	}

	privateKey, err := x509.ParsePKCS8PrivateKey(block.Bytes)
	if err != nil {
		return "", fmt.Errorf("failed to parse private key: %v", err)
	}

	ecdsaKey, ok := privateKey.(*ecdsa.PrivateKey)
	if !ok {
		return "", fmt.Errorf("private key is not ECDSA")
	}

	// Set default scopes if none provided
	if len(scopes) == 0 {
		scopes = []JWTScope{ReadAll, WriteAll}
	}

	// Convert scopes to strings
	scopeStrings := make([]string, len(scopes))
	for i, scope := range scopes {
		scopeStrings[i] = string(scope)
	}

	now := time.Now().UTC()

	// Create claims
	claims := TokenClaims{
		Scopes:    scopeStrings,
		Issuer:    "Gr4vy Python SDK", // Keep the same issuer as Python version
		IssuedAt:  jwt.NewNumericDate(now),
		NotBefore: jwt.NewNumericDate(now),
		ExpiresAt: jwt.NewNumericDate(now.Add(time.Duration(expiresIn) * time.Second)),
		ID:        uuid.New().String(),
	}

	if checkoutSessionID != "" {
		claims.CheckoutSessionID = checkoutSessionID
	}

	// Add embed params if embed scope is present
	hasEmbedScope := false
	for _, scope := range scopes {
		if scope == Embed {
			hasEmbedScope = true
			break
		}
	}

	if hasEmbedScope && embedParams != nil {
		claims.Embed = embedParams
	}

	// Create token with kid header (thumbprint)
	token := jwt.NewWithClaims(jwt.SigningMethodES512, claims)

	// Calculate thumbprint for kid header
	thumbprint, err := calculateThumbprint(ecdsaKey)
	if err != nil {
		return "", fmt.Errorf("failed to calculate thumbprint: %v", err)
	}

	token.Header["kid"] = thumbprint

	// Sign token
	tokenString, err := token.SignedString(ecdsaKey)
	if err != nil {
		return "", fmt.Errorf("failed to sign token: %v", err)
	}

	return tokenString, nil
}

// GetEmbedToken generates a token specifically for embed use
func GetEmbedToken(privateKeyPEM string, embedParams map[string]interface{}, checkoutSessionID string) (string, error) {
	return GetToken(privateKeyPEM, []JWTScope{Embed}, 3600, embedParams, checkoutSessionID)
}

// UpdateToken updates an existing token with a new signature, and optionally new data
func UpdateToken(token string, privateKeyPEM string, scopes []JWTScope, expiresIn int, embedParams map[string]interface{}, checkoutSessionID string) (string, error) {
	// Parse the existing token without signature verification to get claims
	parsedToken, _, err := new(jwt.Parser).ParseUnverified(token, &TokenClaims{})
	if err != nil {
		return "", fmt.Errorf("failed to parse existing token: %v", err)
	}

	claims := parsedToken.Claims.(*TokenClaims)

	// Use previous scopes if none provided
	previousScopes := make([]JWTScope, len(claims.Scopes))
	for i, scope := range claims.Scopes {
		previousScopes[i] = JWTScope(scope)
	}

	if len(scopes) == 0 {
		scopes = previousScopes
	}

	// Use previous embed params if none provided
	if embedParams == nil {
		embedParams = claims.Embed
	}

	// Use previous checkout session ID if none provided
	if checkoutSessionID == "" {
		checkoutSessionID = claims.CheckoutSessionID
	}

	return GetToken(privateKeyPEM, scopes, expiresIn, embedParams, checkoutSessionID)
}

// WithToken generates a function that creates a new token for every API request
func WithToken(privateKeyPEM string, scopes []JWTScope, expiresIn int, embedParams map[string]interface{}, checkoutSessionID string) func(ctx context.Context) (components.Security, error) {
	return func(ctx context.Context) (components.Security, error) {
		token, err := GetToken(privateKeyPEM, scopes, expiresIn, embedParams, checkoutSessionID)
		if err != nil {
			return components.Security{}, err
		}
		return components.Security{BearerAuth: &token}, nil
	}
}

// calculateThumbprint calculates the JWK thumbprint for the kid header
func calculateThumbprint(privateKey *ecdsa.PrivateKey) (string, error) {
	// Create JWK representation for ES512 (P-521 curve)
	jwk := map[string]interface{}{
		"kty": "EC",
		"crv": "P-521",
		"x":   base64.RawURLEncoding.EncodeToString(privateKey.PublicKey.X.Bytes()),
		"y":   base64.RawURLEncoding.EncodeToString(privateKey.PublicKey.Y.Bytes()),
	}

	// Create canonical JSON (sorted keys, no spaces)
	jsonBytes, err := json.Marshal(jwk)
	if err != nil {
		return "", err
	}

	// Calculate SHA256 hash
	hash := sha256.Sum256(jsonBytes)

	// Return base64url encoded hash
	return base64.RawURLEncoding.EncodeToString(hash[:]), nil
}
