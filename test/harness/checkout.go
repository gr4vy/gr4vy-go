package harness

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"time"

	"github.com/gr4vy/gr4vy-go/models/components"
)

// putCheckoutSessionFields performs the raw PUT that secures card details into a
// checkout session. This endpoint is keyed by the session id (used from the
// browser) rather than a bearer token, so it is exercised directly rather than
// through the SDK, mirroring real front-end usage. It must return 204.
func putCheckoutSessionFields(ctx context.Context, sessionID string, paymentMethod map[string]interface{}) error {
	url := fmt.Sprintf("%s/checkout/sessions/%s/fields", APIBaseURL(), sessionID)
	body, err := json.Marshal(map[string]interface{}{"payment_method": paymentMethod})
	if err != nil {
		return err
	}
	req, err := http.NewRequestWithContext(ctx, http.MethodPut, url, bytes.NewReader(body))
	if err != nil {
		return err
	}
	req.Header.Set("Content-Type", "application/json")

	resp, err := (&http.Client{Timeout: 10 * time.Second}).Do(req)
	if err != nil {
		return err
	}
	defer resp.Body.Close()
	if resp.StatusCode != http.StatusNoContent {
		b, _ := io.ReadAll(resp.Body)
		return fmt.Errorf("PUT fields: expected 204, got %d: %s", resp.StatusCode, string(b))
	}
	return nil
}

// PutCard secures a raw card into the checkout session.
func PutCard(ctx context.Context, sessionID string) error {
	return putCheckoutSessionFields(ctx, sessionID, map[string]interface{}{
		"method":          "card",
		"number":          ApprovingCardNumber,
		"expiration_date": CardExpirationDate,
		"security_code":   CardSecurityCode,
	})
}

// PutStoredMethod secures a stored payment-method id into the checkout session.
func PutStoredMethod(ctx context.Context, sessionID, paymentMethodID string) error {
	return putCheckoutSessionFields(ctx, sessionID, map[string]interface{}{
		"method":        "id",
		"id":            paymentMethodID,
		"security_code": CardSecurityCode,
	})
}

// Authorize runs the headline card flow: create a checkout session, secure the
// approving card into it, then create a transaction against the session. It
// returns the authorized transaction.
func Authorize(t TestingT, m *TestMerchant, amount int64, currency string) *components.Transaction {
	t.Helper()
	ctx := context.Background()

	session, err := m.Client.CheckoutSessions.Create(ctx, nil, nil)
	if err != nil {
		t.Fatalf("create checkout session: %v", err)
		return nil
	}
	if err := PutCard(ctx, session.ID); err != nil {
		t.Fatalf("put card fields: %v", err)
		return nil
	}

	pm := components.CreateTransactionCreatePaymentMethodCheckoutSessionWithURLPaymentMethodCreate(
		components.CheckoutSessionWithURLPaymentMethodCreate{ID: session.ID},
	)
	tx, err := m.Client.Transactions.Create(ctx, components.TransactionCreate{
		Amount:        amount,
		Currency:      currency,
		PaymentMethod: &pm,
	}, nil, nil, nil)
	if err != nil {
		t.Fatalf("create transaction: %v", err)
		return nil
	}
	return tx
}
