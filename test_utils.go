package gr4vygo

import (
	"bytes"
	"crypto/rand"
	"encoding/hex"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"path/filepath"
	"time"
)

// TestSetup provides utilities for setting up test environments
type TestSetup struct {
	privateKey string
}

// NewTestSetup creates a new test setup instance
func NewTestSetup() (*TestSetup, error) {
	privateKey, err := loadPrivateKey()
	if err != nil {
		return nil, err
	}

	return &TestSetup{
		privateKey: privateKey,
	}, nil
}

// loadPrivateKey loads the private key from an environment variable or a file
func loadPrivateKey() (string, error) {
	if privateKey := os.Getenv("PRIVATE_KEY"); privateKey != "" {
		return privateKey, nil
	}

	wd, err := os.Getwd()
	if err != nil {
		return "", err
	}

	possiblePaths := []string{
		filepath.Join(wd, "private_key.pem"),
		filepath.Join(wd, "../private_key.pem"),
		filepath.Join(wd, "../../private_key.pem"),
		"private_key.pem",
	}

	for _, path := range possiblePaths {
		if content, err := os.ReadFile(path); err == nil {
			return string(content), nil
		}
	}

	return "", fmt.Errorf("private key not found in environment variable PRIVATE_KEY or in any of the expected file locations")
}

// CreateGr4vyClient creates a Gr4vy client instance
func (ts *TestSetup) CreateGr4vyClient(merchantAccountID string) (*Gr4vyClient, error) {
	token, err := GetToken(ts.privateKey, []JWTScope{ReadAll, WriteAll}, 3600, nil, "")
	if err != nil {
		return nil, err
	}

	return NewGr4vyClient("sandbox", "e2e", merchantAccountID, token), nil
}

// SetupEnvironment sets up the test environment
func (ts *TestSetup) SetupEnvironment() (*Gr4vyClient, error) {
	adminClient, err := ts.CreateGr4vyClient("")
	if err != nil {
		return nil, fmt.Errorf("failed to create admin client: %v", err)
	}

	merchantAccountID := generateRandomHex(8)

	err = adminClient.CreateMerchantAccount(merchantAccountID, merchantAccountID)
	if err != nil {
		return nil, fmt.Errorf("failed to create merchant account: %v", err)
	}

	merchantClient, err := ts.CreateGr4vyClient(merchantAccountID)
	if err != nil {
		return nil, fmt.Errorf("failed to create merchant client: %v", err)
	}

	err = merchantClient.CreatePaymentService()
	if err != nil {
		log.Printf("Warning: Failed to create payment service: %v", err)
	}

	return merchantClient, nil
}

// CleanupEnvironment cleans up the test environment
func (ts *TestSetup) CleanupEnvironment() {
	// Currently a no-op
}

// generateRandomHex generates a random hex string
func generateRandomHex(length int) string {
	bytes := make([]byte, length)
	rand.Read(bytes)
	return hex.EncodeToString(bytes)
}

// Test models and client (needed for checkout session tests)
type CheckoutSession struct {
	ID string `json:"id"`
}

type Transaction struct {
	ID     string `json:"id"`
	Status string `json:"status"`
	Amount int    `json:"amount"`
}

type PaymentMethod struct {
	ID string `json:"id"`
}

type CardPaymentMethodCreate struct {
	Number         string `json:"number"`
	ExpirationDate string `json:"expiration_date"`
	SecurityCode   string `json:"security_code"`
}

type Gr4vyClient struct {
	baseURL           string
	merchantAccountID string
	bearerToken       string
	httpClient        *http.Client
}

func NewGr4vyClient(server, id, merchantAccountID, bearerToken string) *Gr4vyClient {
	baseURL := fmt.Sprintf("https://api.%s.%s.gr4vy.app", server, id)
	return &Gr4vyClient{
		baseURL:           baseURL,
		merchantAccountID: merchantAccountID,
		bearerToken:       bearerToken,
		httpClient:        &http.Client{Timeout: 30 * time.Second},
	}
}

func (c *Gr4vyClient) makeRequest(method, path string, body interface{}) (*http.Response, error) {
	url := c.baseURL + path

	var bodyReader io.Reader
	if body != nil {
		jsonBody, err := json.Marshal(body)
		if err != nil {
			return nil, err
		}
		bodyReader = bytes.NewReader(jsonBody)
	}

	req, err := http.NewRequest(method, url, bodyReader)
	if err != nil {
		return nil, err
	}

	req.Header.Set("Authorization", "Bearer "+c.bearerToken)
	req.Header.Set("Content-Type", "application/json")
	if c.merchantAccountID != "" {
		req.Header.Set("gr4vy-merchant-account-id", c.merchantAccountID)
	}

	return c.httpClient.Do(req)
}

func (c *Gr4vyClient) CreateCheckoutSession() (*CheckoutSession, error) {
	resp, err := c.makeRequest("POST", "/checkout/sessions", nil)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusCreated {
		body, _ := io.ReadAll(resp.Body)
		return nil, fmt.Errorf("failed to create checkout session: %s", string(body))
	}

	var session CheckoutSession
	if err := json.NewDecoder(resp.Body).Decode(&session); err != nil {
		return nil, err
	}

	return &session, nil
}

func (c *Gr4vyClient) UpdateCheckoutSessionFields(sessionID string, fields map[string]interface{}) error {
	url := fmt.Sprintf("https://api.sandbox.e2e.gr4vy.app/checkout/sessions/%s/fields", sessionID)

	jsonBody, err := json.Marshal(fields)
	if err != nil {
		return err
	}

	req, err := http.NewRequest("PUT", url, bytes.NewReader(jsonBody))
	if err != nil {
		return err
	}

	req.Header.Set("Content-Type", "application/json")

	client := &http.Client{Timeout: 5 * time.Second}
	resp, err := client.Do(req)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusNoContent {
		body, _ := io.ReadAll(resp.Body)
		return fmt.Errorf("failed to update checkout session fields: %s", string(body))
	}

	return nil
}

func (c *Gr4vyClient) CreateTransaction(amount int, currency string, paymentMethod map[string]interface{}) (*Transaction, error) {
	body := map[string]interface{}{
		"amount":         amount,
		"currency":       currency,
		"payment_method": paymentMethod,
	}

	resp, err := c.makeRequest("POST", "/transactions", body)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	respBody, _ := io.ReadAll(resp.Body)

	if resp.StatusCode != http.StatusCreated {
		return nil, fmt.Errorf("failed to create transaction: %s", string(respBody))
	}

	var transaction Transaction
	if err := json.Unmarshal(respBody, &transaction); err != nil {
		return nil, err
	}

	return &transaction, nil
}

func (c *Gr4vyClient) CreatePaymentMethod(cardData CardPaymentMethodCreate) (*PaymentMethod, error) {
	resp, err := c.makeRequest("POST", "/payment-methods", cardData)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusCreated {
		body, _ := io.ReadAll(resp.Body)
		return nil, fmt.Errorf("failed to create payment method: %s", string(body))
	}

	var paymentMethod PaymentMethod
	if err := json.NewDecoder(resp.Body).Decode(&paymentMethod); err != nil {
		return nil, err
	}

	return &paymentMethod, nil
}

func (c *Gr4vyClient) CreateMerchantAccount(id, displayName string) error {
	body := map[string]interface{}{
		"id":           id,
		"display_name": displayName,
	}

	resp, err := c.makeRequest("POST", "/merchant-accounts", body)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusCreated {
		respBody, _ := io.ReadAll(resp.Body)
		return fmt.Errorf("failed to create merchant account: %s", string(respBody))
	}

	return nil
}

func (c *Gr4vyClient) CreatePaymentService() error {
	body := map[string]interface{}{
		"accepted_countries":            []string{"US"},
		"accepted_currencies":           []string{"USD"},
		"display_name":                  "Payment service",
		"payment_service_definition_id": "mock-card",
		"fields": []map[string]string{
			{"key": "merchant_id", "value": "test"},
		},
	}

	resp, err := c.makeRequest("POST", "/payment-services", body)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusCreated {
		respBody, _ := io.ReadAll(resp.Body)
		return fmt.Errorf("failed to create payment service: %s", string(respBody))
	}

	return nil
}
