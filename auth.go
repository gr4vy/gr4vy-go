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
	ApiLogsRead                      JWTScope = "api-logs.read"
	ApiLogsWrite                     JWTScope = "api-logs.write"
	ApplePayCertificatesRead         JWTScope = "apple-pay-certificates.read"
	ApplePayCertificatesWrite        JWTScope = "apple-pay-certificates.write"
	AuditLogsRead                    JWTScope = "audit-logs.read"
	AuditLogsWrite                   JWTScope = "audit-logs.write"
	BuyersRead                       JWTScope = "buyers.read"
	BuyersWrite                      JWTScope = "buyers.write"
	BuyersBillingDetailsRead         JWTScope = "buyers.billing-details.read"
	BuyersBillingDetailsWrite        JWTScope = "buyers.billing-details.write"
	CardSchemeDefinitionsRead        JWTScope = "card-scheme-definitions.read"
	CardSchemeDefinitionsWrite       JWTScope = "card-scheme-definitions.write"
	CheckoutSessionsRead             JWTScope = "checkout-sessions.read"
	CheckoutSessionsWrite            JWTScope = "checkout-sessions.write"
	ConnectionsRead                  JWTScope = "connections.read"
	ConnectionsWrite                 JWTScope = "connections.write"
	DigitalWalletsRead               JWTScope = "digital-wallets.read"
	DigitalWalletsWrite              JWTScope = "digital-wallets.write"
	FlowsRead                        JWTScope = "flows.read"
	FlowsWrite                       JWTScope = "flows.write"
	GiftCardServiceDefinitionsRead   JWTScope = "gift-card-service-definitions.read"
	GiftCardServiceDefinitionsWrite  JWTScope = "gift-card-service-definitions.write"
	GiftCardServicesRead             JWTScope = "gift-card-services.read"
	GiftCardServicesWrite            JWTScope = "gift-card-services.write"
	GiftCardsRead                    JWTScope = "gift-cards.read"
	GiftCardsWrite                   JWTScope = "gift-cards.write"
	MerchantAccountsRead             JWTScope = "merchant-accounts.read"
	MerchantAccountsWrite            JWTScope = "merchant-accounts.write"
	PaymentLinksRead                 JWTScope = "payment-links.read"
	PaymentLinksWrite                JWTScope = "payment-links.write"
	PaymentMethodDefinitionsRead     JWTScope = "payment-method-definitions.read"
	PaymentMethodDefinitionsWrite    JWTScope = "payment-method-definitions.write"
	PaymentMethodsRead               JWTScope = "payment-methods.read"
	PaymentMethodsWrite              JWTScope = "payment-methods.write"
	PaymentOptionsRead               JWTScope = "payment-options.read"
	PaymentOptionsWrite              JWTScope = "payment-options.write"
	PaymentServiceDefinitionsRead    JWTScope = "payment-service-definitions.read"
	PaymentServiceDefinitionsWrite   JWTScope = "payment-service-definitions.write"
	PaymentServicesRead              JWTScope = "payment-services.read"
	PaymentServicesWrite             JWTScope = "payment-services.write"
	PayoutsRead                      JWTScope = "payouts.read"
	PayoutsWrite                     JWTScope = "payouts.write"
	ReportsRead                      JWTScope = "reports.read"
	ReportsWrite                     JWTScope = "reports.write"
	RolesRead                        JWTScope = "roles.read"
	RolesWrite                       JWTScope = "roles.write"
	TransactionsRead                 JWTScope = "transactions.read"
	TransactionsWrite                JWTScope = "transactions.write"
	UsersMeRead                      JWTScope = "users.me.read"
	UsersMeWrite                     JWTScope = "users.me.write"
	VaultForwardRead                 JWTScope = "vault-forward.read"
	VaultForwardWrite                JWTScope = "vault-forward.write"
	VaultForwardAuthenticationsRead  JWTScope = "vault-forward-authentications.read"
	VaultForwardAuthenticationsWrite JWTScope = "vault-forward-authentications.write"
	VaultForwardConfigsRead          JWTScope = "vault-forward-configs.read"
	VaultForwardConfigsWrite         JWTScope = "vault-forward-configs.write"
	VaultForwardDefinitionsRead      JWTScope = "vault-forward-definitions.read"
	VaultForwardDefinitionsWrite     JWTScope = "vault-forward-definitions.write"
	WebhookSubscriptionsRead         JWTScope = "webhook-subscriptions.read"
	WebhookSubscriptionsWrite        JWTScope = "webhook-subscriptions.write"
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

// GetToken generates a JWT token signed with the provided ECDSA private key.
//   - privateKeyPEM: PEM-encoded ECDSA private key string.
//   - scopes: List of JWTScope values to include in the token.
//   - expiresIn: Token expiration in seconds from now.
//
// Returns the signed JWT token string or an error.
func GetToken(privateKeyPEM string, scopes []JWTScope, expiresIn int) (string, error) {
	return getToken(privateKeyPEM, scopes, expiresIn, nil, "")
}

// GetTokenWithEmbedProperties generates a JWT token with embed parameters and an optional checkout session ID.
//   - privateKeyPEM: PEM-encoded ECDSA private key string.
//   - scopes: List of JWTScope values to include in the token.
//   - expiresIn: Token expiration in seconds from now.
//   - embedParams: Optional map of embed parameters to include in the token.
//   - checkoutSessionID: Optional checkout session ID to include in the token.
//
// Returns the signed JWT token string or an error.
func GetTokenWithEmbedProperties(privateKeyPEM string, scopes []JWTScope, expiresIn int, embedParams map[string]interface{}, checkoutSessionID string) (string, error) {
	return getToken(privateKeyPEM, scopes, expiresIn, embedParams, checkoutSessionID)
}

// GetToken generates a JWT token with the specified parameters
func getToken(privateKeyPEM string, scopes []JWTScope, expiresIn int, embedParams map[string]interface{}, checkoutSessionID string) (string, error) {
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

	sdk := New()
	issuer := sdk.sdkConfiguration.UserAgent

	// Create claims
	claims := TokenClaims{
		Scopes:    scopeStrings,
		Issuer:    issuer,
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

// GetEmbedToken generates a JWT token specifically for embed use, with embed parameters and an optional checkout session ID.
//   - privateKeyPEM: PEM-encoded ECDSA private key string.
//   - embedParams: Map of embed parameters to include in the token.
//   - checkoutSessionID: Optional checkout session ID to include in the token.
//
// Returns the signed JWT token string or an error.
func GetEmbedToken(privateKeyPEM string, embedParams map[string]interface{}, checkoutSessionID string) (string, error) {
	return getToken(privateKeyPEM, []JWTScope{Embed}, 3600, embedParams, checkoutSessionID)
}

// UpdateToken updates an existing JWT token with a new signature and optionally new data.
//   - token: The existing JWT token string.
//   - privateKeyPEM: PEM-encoded ECDSA private key string.
//   - scopes: Optional list of JWTScope values to include in the new token. If empty, uses previous scopes.
//   - expiresIn: Token expiration in seconds from now.
//   - embedParams: Optional map of embed parameters. If nil, uses previous embed params.
//   - checkoutSessionID: Optional checkout session ID. If empty, uses previous value.
//
// Returns the updated signed JWT token string or an error.
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

	return getToken(privateKeyPEM, scopes, expiresIn, embedParams, checkoutSessionID)
}

// WithToken returns a function that generates a new JWT token for every API request, suitable for use with the SDK's WithSecuritySource option.
//   - privateKeyPEM: PEM-encoded ECDSA private key string.
//   - scopes: List of JWTScope values to include in the token.
//   - expiresIn: Token expiration in seconds from now.
//
// Returns a function that generates a Security struct with a fresh token.
func WithToken(privateKeyPEM string, scopes []JWTScope, expiresIn int) func(ctx context.Context) (components.Security, error) {
	return withToken(privateKeyPEM, scopes, expiresIn, nil, "")
}

func WithTokenWithEmbedProperties(privateKeyPEM string, scopes []JWTScope, expiresIn int, embedParams map[string]interface{}, checkoutSessionID string) func(ctx context.Context) (components.Security, error) {
	return withToken(privateKeyPEM, scopes, expiresIn, embedParams, checkoutSessionID)
}

// WithToken generates a function that creates a new token for every API request
func withToken(privateKeyPEM string, scopes []JWTScope, expiresIn int, embedParams map[string]interface{}, checkoutSessionID string) func(ctx context.Context) (components.Security, error) {
	return func(ctx context.Context) (components.Security, error) {
		token, err := getToken(privateKeyPEM, scopes, expiresIn, embedParams, checkoutSessionID)
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
