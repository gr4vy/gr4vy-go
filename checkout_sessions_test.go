package gr4vygo

import (
	"strings"
	"testing"
)

var testClient *Gr4vyClient

func setupClient(t *testing.T) *Gr4vyClient {
	if testClient != nil {
		return testClient
	}

	setup, err := NewTestSetup()
	if err != nil {
		t.Skipf("Skipping integration test: failed to create test setup: %v", err)
	}

	client, err := setup.SetupEnvironment()
	if err != nil {
		t.Skipf("Skipping integration test: failed to setup environment: %v", err)
	}

	testClient = client
	return testClient
}

func TestProcessPaymentWithCheckoutSession(t *testing.T) {
	client := setupClient(t)

	checkoutSession, err := client.CreateCheckoutSession()
	if err != nil {
		t.Fatalf("Failed to create checkout session: %v", err)
	}

	if checkoutSession.ID == "" {
		t.Fatal("Checkout session ID should not be empty")
	}

	cardFields := map[string]interface{}{
		"payment_method": map[string]interface{}{
			"method":          "card",
			"number":          "4111111111111111",
			"expiration_date": "11/25",
			"security_code":   "123",
		},
	}

	err = client.UpdateCheckoutSessionFields(checkoutSession.ID, cardFields)
	if err != nil {
		t.Fatalf("Failed to update checkout session fields: %v", err)
	}

	paymentMethod := map[string]interface{}{
		"method": "checkout-session",
		"id":     checkoutSession.ID,
	}

	transaction, err := client.CreateTransaction(1299, "USD", paymentMethod)
	if err != nil {
		t.Fatalf("Failed to create transaction: %v", err)
	}

	if transaction.ID == "" {
		t.Error("Transaction ID should not be empty")
	}

	if transaction.Status != "authorization_succeeded" {
		t.Errorf("Expected transaction status 'authorization_succeeded', got '%s'", transaction.Status)
	}

	if transaction.Amount != 1299 {
		t.Errorf("Expected transaction amount 1299, got %d", transaction.Amount)
	}
}

func TestHandleErrorOnMissingCardData(t *testing.T) {
	client := setupClient(t)

	checkoutSession, err := client.CreateCheckoutSession()
	if err != nil {
		t.Fatalf("Failed to create checkout session: %v", err)
	}

	if checkoutSession.ID == "" {
		t.Fatal("Checkout session ID should not be empty")
	}

	paymentMethod := map[string]interface{}{
		"method": "checkout-session",
		"id":     checkoutSession.ID,
	}

	_, err = client.CreateTransaction(1299, "USD", paymentMethod)
	if err == nil {
		t.Fatal("Expected error when creating transaction without card data")
	}

	if !strings.Contains(err.Error(), "Request failed validation") &&
		!strings.Contains(err.Error(), "failed to create transaction") {
		t.Errorf("Expected validation error, got: %v", err)
	}
}

func TestHandleStoredPaymentMethod(t *testing.T) {
	client := setupClient(t)

	cardData := CardPaymentMethodCreate{
		Number:         "4111111111111111",
		ExpirationDate: "11/25",
		SecurityCode:   "123",
	}

	paymentMethod, err := client.CreatePaymentMethod(cardData)
	if err != nil {
		t.Fatalf("Failed to create payment method: %v", err)
	}

	if paymentMethod.ID == "" {
		t.Fatal("Payment method ID should not be empty")
	}

	checkoutSession, err := client.CreateCheckoutSession()
	if err != nil {
		t.Fatalf("Failed to create checkout session: %v", err)
	}

	if checkoutSession.ID == "" {
		t.Fatal("Checkout session ID should not be empty")
	}

	storedFields := map[string]interface{}{
		"payment_method": map[string]interface{}{
			"method":        "id",
			"id":            paymentMethod.ID,
			"security_code": "123",
		},
	}

	err = client.UpdateCheckoutSessionFields(checkoutSession.ID, storedFields)
	if err != nil {
		t.Fatalf("Failed to update checkout session fields: %v", err)
	}

	transactionPaymentMethod := map[string]interface{}{
		"method": "checkout-session",
		"id":     checkoutSession.ID,
	}

	transaction, err := client.CreateTransaction(1299, "USD", transactionPaymentMethod)
	if err != nil {
		t.Fatalf("Failed to create transaction: %v", err)
	}

	if transaction.ID == "" {
		t.Error("Transaction ID should not be empty")
	}

	if transaction.Status != "authorization_succeeded" {
		t.Errorf("Expected transaction status 'authorization_succeeded', got '%s'", transaction.Status)
	}

	if transaction.Amount != 1299 {
		t.Errorf("Expected transaction amount 1299, got %d", transaction.Amount)
	}
}
