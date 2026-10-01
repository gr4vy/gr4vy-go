package gr4vygo

import "testing"

// During a secret rotation the header carries several signatures, and any one
// of them may be the valid one.
func TestVerifyWebhookAcceptsAValidSignatureInAnyPosition(t *testing.T) {
	header := "other," + "78aca0c78005107a654a957b8566fa6e0e5e06aea92d7da72a6da9e5a690d013"
	if err := VerifyWebhook(payload, secret, header, timestampHeader, timestampTolerance); err != nil {
		t.Errorf("unexpected error: %v", err)
	}
}

func TestVerifyWebhookRejectsANearlyMatchingSignature(t *testing.T) {
	header := "78aca0c78005107a654a957b8566fa6e0e5e06aea92d7da72a6da9e5a690d014"
	if err := VerifyWebhook(payload, secret, header, timestampHeader, timestampTolerance); err == nil {
		t.Error("expected a signature differing in its last character to be rejected")
	}
}
