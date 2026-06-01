package harness

import (
	"context"
	"errors"
	"fmt"
	"net/http"
	"os"
	"path/filepath"
	"sync"
	"time"

	"github.com/google/uuid"
	gr4vygo "github.com/gr4vy/gr4vy-go"
	"github.com/gr4vy/gr4vy-go/models/components"
)

// The E2E sandbox. WithServer(sandbox)+WithID("e2e") resolves to this host.
const (
	serverID   = "e2e"
	apiBaseURL = "https://api.sandbox.e2e.gr4vy.app"
)

// ErrNoPrivateKey is returned by SetupMerchant when no signing key is available.
// Tests treat it as "skip" rather than "fail": fork PRs and laptops without the
// sandbox key should skip the live suite, while a missing key on a trusted ref
// is caught in CI (see ci.yaml), not here.
var ErrNoPrivateKey = errors.New("no private key (set PRIVATE_KEY or add private_key.pem)")

var defaultScopes = []gr4vygo.JWTScope{gr4vygo.ReadAll, gr4vygo.WriteAll}

// APIBaseURL is exported so helpers that bypass the SDK (raw checkout-session
// field PUTs) can target the same host.
func APIBaseURL() string { return apiBaseURL }

// loadPrivateKey reads the signing key from PRIVATE_KEY, falling back to a
// private_key.pem at the repo root (the test packages run from test/<shard>, so
// we also probe parent directories).
func loadPrivateKey() (string, error) {
	if key := os.Getenv("PRIVATE_KEY"); key != "" {
		return key, nil
	}
	candidates := []string{
		"private_key.pem",
		filepath.Join("..", "..", "private_key.pem"),
		filepath.Join("..", "private_key.pem"),
	}
	for _, c := range candidates {
		if data, err := os.ReadFile(c); err == nil {
			return string(data), nil
		}
	}
	return "", ErrNoPrivateKey
}

// newClient builds an SDK client pointed at the sandbox, signing every request
// with a fresh JWT and routing through the forward-compat/recording transport.
// When merchantAccountID is non-empty the client is bound to it.
func newClient(privateKey, merchantAccountID string) *gr4vygo.Gr4vy {
	httpClient := &http.Client{
		Timeout:   30 * time.Second,
		Transport: newInterceptor(http.DefaultTransport),
	}
	opts := []gr4vygo.SDKOption{
		gr4vygo.WithClient(httpClient),
		gr4vygo.WithServer(gr4vygo.ServerSandbox),
		gr4vygo.WithID(serverID),
		gr4vygo.WithSecuritySource(gr4vygo.WithToken(privateKey, defaultScopes, 3600)),
	}
	if merchantAccountID != "" {
		opts = append(opts, gr4vygo.WithMerchantAccountID(merchantAccountID))
	}
	return gr4vygo.New(opts...)
}

// SetupMerchant provisions one isolated merchant account (random id) plus a
// mock-card payment service, and returns a client bound to it. Returns
// ErrNoPrivateKey when no key is available so callers can skip.
func SetupMerchant() (*TestMerchant, error) {
	privateKey, err := loadPrivateKey()
	if err != nil {
		return nil, err
	}

	ctx := context.Background()
	merchantAccountID := uuid.New().String()[:8]

	// Admin client (not merchant-scoped) creates the merchant account.
	admin := newClient(privateKey, "")
	account, err := admin.MerchantAccounts.Create(ctx, components.MerchantAccountCreate{
		ID:          merchantAccountID,
		DisplayName: merchantAccountID,
	})
	if err != nil {
		return nil, fmt.Errorf("create merchant account: %w", err)
	}

	merchant := newClient(privateKey, account.ID)

	// A deterministic mock-card service so card flows reach a real 2xx.
	_, err = merchant.PaymentServices.Create(ctx, components.PaymentServiceCreate{
		DisplayName:                "E2E mock card",
		PaymentServiceDefinitionID: "mock-card",
		AcceptedCurrencies:         []string{"USD"},
		AcceptedCountries:          []string{"US"},
		Fields: []components.Field{
			{Key: "merchant_id", Value: account.ID},
		},
	}, &account.ID)
	if err != nil {
		return nil, fmt.Errorf("create mock-card payment service: %w", err)
	}

	return &TestMerchant{
		Client:            merchant,
		MerchantAccountID: account.ID,
		PrivateKey:        privateKey,
	}, nil
}

// One merchant per test package (= per CI shard). Because `go test` runs each
// package in its own process, this package-level once yields exactly one
// merchant per shard with no cross-shard sharing.
var (
	sharedOnce     sync.Once
	sharedMerchant *TestMerchant
	sharedErr      error
)

// Merchant returns the shard's shared merchant, provisioning it on first use.
// If no signing key is available the calling test is skipped; any other setup
// failure fails the test.
func Merchant(t TestingT) *TestMerchant {
	t.Helper()
	sharedOnce.Do(func() { sharedMerchant, sharedErr = SetupMerchant() })
	if sharedErr != nil {
		if errors.Is(sharedErr, ErrNoPrivateKey) {
			t.Skipf("skipping live E2E: %v", sharedErr)
		}
		t.Fatalf("merchant setup failed: %v", sharedErr)
	}
	return sharedMerchant
}

// TestingT is the subset of *testing.T the harness relies on, kept small so
// helpers stay easy to reason about (and testable).
type TestingT interface {
	Helper()
	Skipf(format string, args ...any)
	Fatalf(format string, args ...any)
	Errorf(format string, args ...any)
}
