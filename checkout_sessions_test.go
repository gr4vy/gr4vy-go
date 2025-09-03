package gr4vygo

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"io"
	"math/rand"
	"net/http"
	"os"
	"path/filepath"
	"strings"
	"testing"
	"time"

	"github.com/google/uuid"
	"github.com/gr4vy/gr4vy-go/models/components"
	"github.com/gr4vy/gr4vy-go/models/operations"
)

// Test configuration
var (
	merchantAccountID string
)

// Adds a custom interceptor that can be added to an HTTP client
// which adds some random data to JSON responses. This helps catch any
// forward compatibility issues in the transaction API.
type jsonInterceptor struct {
	next http.RoundTripper
}

func (i *jsonInterceptor) RoundTrip(req *http.Request) (*http.Response, error) {
	resp, err := i.next.RoundTrip(req)
	if err != nil {
		return nil, err
	}
	if !strings.Contains(resp.Header.Get("Content-Type"), "application/json") {
		return resp, nil
	}
	originalBody, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, fmt.Errorf("failed to read original response body: %w", err)
	}
	resp.Body.Close()
	var data map[string]interface{}
	if err := json.Unmarshal(originalBody, &data); err != nil {
		resp.Body = io.NopCloser(bytes.NewReader(originalBody))
		return resp, nil
	}

	randomKey := fmt.Sprintf("unexpected_field_%d", rand.Intn(1000))
	// This is where we add the unexpected field field
	data[randomKey] = "this is an injected test value"
	modifiedBody, err := json.Marshal(data)
	if err != nil {
		return nil, fmt.Errorf("failed to marshal modified body: %w", err)
	}
	resp.Body = io.NopCloser(bytes.NewReader(modifiedBody))
	resp.ContentLength = int64(len(modifiedBody))

	return resp, nil
}

// loadPrivateKey loads the private key from environment or file
func loadPrivateKey() (string, error) {
	// First try environment variable
	privateKey := os.Getenv("PRIVATE_KEY")
	if privateKey != "" {
		return privateKey, nil
	}

	// Try loading from file
	filename := filepath.Join("..", "..", "private_key.pem")
	data, err := os.ReadFile(filename)
	if err != nil {
		// Try current directory
		filename = "private_key.pem"
		data, err = os.ReadFile(filename)
		if err != nil {
			return "", fmt.Errorf("failed to load private key: %w", err)
		}
	}

	return string(data), nil
}

// setupTestEnvironment sets up the test environment
func setupTestEnvironment(t *testing.T) *Gr4vy {
	// Load private key
	privateKey, err := loadPrivateKey()
	if err != nil {
		t.Skipf("Skipping test: %v", err)
	}

	// Generate JWT token
	token, err := GetToken(privateKey, []JWTScope{ReadAll, WriteAll}, 3600)
	if err != nil {
		t.Fatalf("Failed to generate token: %v", err)
	}

	// Create admin client
	adminClient := New(
		WithServer(ServerSandbox),
		WithID("e2e"),
		WithSecurity(token),
	)

	// Generate random merchant account ID
	merchantAccountID = uuid.New().String()[:8]

	// Create merchant account
	ctx := context.Background()

	merchantAccountCreate := components.MerchantAccountCreate{
		DisplayName: merchantAccountID,
		ID:          merchantAccountID,
	}

	merchantAccount, err := adminClient.MerchantAccounts.Create(ctx, merchantAccountCreate)
	if err != nil {
		t.Fatalf("Failed to create merchant account: %v", err)
	}

	// Create merchant client
	withToken := WithToken(privateKey, []JWTScope{ReadAll, WriteAll}, 3600)

	// Adds a custom interceptor to the HTTP client
	// which adds some random data to JSON responses. This helps catch any
	// forward compatibility issues.
	baseTransport := http.DefaultTransport
	httpClient := &http.Client{
		Timeout: 30 * time.Second,
		Transport: &jsonInterceptor{
			next: baseTransport,
		},
	}

	merchantClient := New(
		WithClient(httpClient),
		WithServer(ServerSandbox),
		WithID("e2e"),
		WithSecuritySource(withToken),
		WithMerchantAccountID(merchantAccount.ID),
	)

	// Setup payment service - adjust Fields type based on your SDK
	paymentServiceCreate := components.PaymentServiceCreate{
		DisplayName:                "Payment service",
		PaymentServiceDefinitionID: "mock-card",
		AcceptedCurrencies:         []string{"USD"},
		AcceptedCountries:          []string{"US"},
		Fields: []components.Field{
			components.Field{
				Key:   "merchant_id",
				Value: merchantAccountID,
			},
		},
	}

	_, err = merchantClient.PaymentServices.Create(ctx, paymentServiceCreate, &merchantAccountID)
	if err != nil {
		t.Fatalf("Failed to create payment service: %v", err)
	}

	return merchantClient
}

// TestMain sets up and tears down the test environment
func TestMain(m *testing.M) {
	// Run tests
	code := m.Run()

	// Exit with test result code
	os.Exit(code)
}

func TestProcessPaymentWithCheckoutSession(t *testing.T) {
	client := setupTestEnvironment(t)
	ctx := context.Background()

	// Create checkout session
	checkoutSession, err := client.CheckoutSessions.Create(ctx, nil, nil)
	if err != nil {
		t.Fatalf("Failed to create checkout session: %v", err)
	}

	if checkoutSession.ID == "" {
		t.Fatal("Checkout session ID is nil")
	}

	// Direct API call to update checkout session fields
	url := fmt.Sprintf("https://api.sandbox.e2e.gr4vy.app/checkout/sessions/%s/fields", checkoutSession.ID)

	updateData := map[string]interface{}{
		"payment_method": map[string]interface{}{
			"method":          "card",
			"number":          "4111111111111111",
			"expiration_date": "11/25",
			"security_code":   "123",
		},
	}

	jsonData, err := json.Marshal(updateData)
	if err != nil {
		t.Fatalf("Failed to marshal update data: %v", err)
	}

	req, err := http.NewRequest("PUT", url, bytes.NewBuffer(jsonData))
	if err != nil {
		t.Fatalf("Failed to create request: %v", err)
	}

	req.Header.Set("Content-Type", "application/json")

	httpClient := &http.Client{Timeout: 5 * time.Second}
	resp, err := httpClient.Do(req)
	if err != nil {
		t.Fatalf("Failed to update checkout session fields: %v", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != 204 {
		body, _ := io.ReadAll(resp.Body)
		t.Fatalf("Expected status code 204, got %d. Body: %s", resp.StatusCode, string(body))
	}

	// Create a transaction using the checkout session
	amount := int64(1299)
	currency := "USD"

	checkoutSessionWithURLPaymentMethodCreate := components.CheckoutSessionWithURLPaymentMethodCreate{
		ID: checkoutSession.ID,
	}
	paymentMethod := components.CreateTransactionCreatePaymentMethodCheckoutSessionWithURLPaymentMethodCreate(checkoutSessionWithURLPaymentMethodCreate)

	transactionCreate := components.TransactionCreate{
		Amount:        amount,
		Currency:      currency,
		PaymentMethod: &paymentMethod,
	}

	transaction, err := client.Transactions.Create(ctx, transactionCreate, nil, nil, nil)

	if err != nil {
		t.Fatalf("Failed to create transaction: %v", err)
	}

	if transaction.ID == "" {
		t.Fatal("Transaction ID is nil")
	}

	if transaction.Status == "" || transaction.Status != components.TransactionStatusAuthorizationSucceeded {
		t.Errorf("Expected status 'authorization_succeeded', got %v", transaction.Status)
	}

	if transaction.Amount == 0 || transaction.Amount != amount {
		t.Errorf("Expected amount %d, got %v", amount, transaction.Amount)
	}
}

func TestHandleErrorOnMissingCardData(t *testing.T) {
	client := setupTestEnvironment(t)
	ctx := context.Background()

	// Create a checkout session
	checkoutSession, err := client.CheckoutSessions.Create(ctx, nil, nil)
	if err != nil {
		t.Fatalf("Failed to create checkout session: %v", err)
	}

	if checkoutSession.ID == "" {
		t.Fatal("Checkout session ID is nil")
	}

	// Attempt to create a transaction with missing card data
	amount := int64(1299)
	currency := "USD"
	checkoutSessionWithURLPaymentMethodCreate := components.CheckoutSessionWithURLPaymentMethodCreate{
		ID: checkoutSession.ID,
	}
	paymentMethod := components.CreateTransactionCreatePaymentMethodCheckoutSessionWithURLPaymentMethodCreate(checkoutSessionWithURLPaymentMethodCreate)

	transactionCreate := components.TransactionCreate{
		Amount:        amount,
		Currency:      currency,
		PaymentMethod: &paymentMethod,
	}

	_, err = client.Transactions.Create(ctx, transactionCreate, nil, nil, nil)

	if err == nil {
		t.Fatal("Expected error for missing card data, but got nil")
	}

	// Check if error contains expected message
	if !strings.Contains(err.Error(), "validation") && !strings.Contains(err.Error(), "400") {
		t.Errorf("Expected validation error, got: %v", err)
	}
}

func TestHandleStoredPaymentMethod(t *testing.T) {
	client := setupTestEnvironment(t)
	ctx := context.Background()

	// Create a card payment method
	number := "4111111111111111"
	expirationDate := "11/25"
	securityCode := "123"

	cardPaymentMethodCreate := components.CardPaymentMethodCreate{
		Number:         number,
		ExpirationDate: expirationDate,
		SecurityCode:   &securityCode,
	}

	request := operations.Body{
		CardPaymentMethodCreate: &cardPaymentMethodCreate,
	}

	paymentMethod, err := client.PaymentMethods.Create(ctx, request, nil)

	if err != nil {
		t.Fatalf("Failed to create payment method: %v", err)
	}

	if paymentMethod.ID == "" {
		t.Fatal("Payment method ID is nil")
	}

	// Create a checkout session
	checkoutSession, err := client.CheckoutSessions.Create(ctx, nil, nil)
	if err != nil {
		t.Fatalf("Failed to create checkout session: %v", err)
	}

	if checkoutSession.ID == "" {
		t.Fatal("Checkout session ID is nil")
	}

	// Direct API call to update checkout session fields
	url := fmt.Sprintf("https://api.sandbox.e2e.gr4vy.app/checkout/sessions/%s/fields", checkoutSession.ID)

	updateData := map[string]interface{}{
		"payment_method": map[string]interface{}{
			"method":        "id",
			"id":            paymentMethod.ID,
			"security_code": "123",
		},
	}

	jsonData, err := json.Marshal(updateData)
	if err != nil {
		t.Fatalf("Failed to marshal update data: %v", err)
	}

	req, err := http.NewRequest("PUT", url, bytes.NewBuffer(jsonData))
	if err != nil {
		t.Fatalf("Failed to create request: %v", err)
	}

	req.Header.Set("Content-Type", "application/json")

	httpClient := &http.Client{Timeout: 5 * time.Second}
	resp, err := httpClient.Do(req)
	if err != nil {
		t.Fatalf("Failed to update checkout session fields: %v", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != 204 {
		body, _ := io.ReadAll(resp.Body)
		t.Fatalf("Expected status code 204, got %d. Body: %s", resp.StatusCode, string(body))
	}

	// Create a transaction using the checkout session
	amount := int64(1299)
	currency := "USD"

	checkoutSessionWithURLPaymentMethodCreate := components.CheckoutSessionWithURLPaymentMethodCreate{
		ID: checkoutSession.ID,
	}
	pm := components.CreateTransactionCreatePaymentMethodCheckoutSessionWithURLPaymentMethodCreate(checkoutSessionWithURLPaymentMethodCreate)

	transactionCreate := components.TransactionCreate{
		Amount:        amount,
		Currency:      currency,
		PaymentMethod: &pm,
	}

	transaction, err := client.Transactions.Create(ctx, transactionCreate, nil, nil, nil)

	if err != nil {
		t.Fatalf("Failed to create transaction: %v", err)
	}

	if transaction.ID == "" {
		t.Fatal("Transaction ID is nil")
	}

	if transaction.Status == "" || transaction.Status != components.TransactionStatusAuthorizationSucceeded {
		t.Errorf("Expected status 'authorization_succeeded', got %v", transaction.Status)
	}

	if transaction.Amount == 0 || transaction.Amount != amount {
		t.Errorf("Expected amount %d, got %v", amount, transaction.Amount)
	}
}
