package gr4vygo

import (
	"strings"
	"testing"
)

const (
	secret             = "Ik4L-8FH0ihWczctcIPXZRR_8F0fPNgmhEfVBbZ3zNwqQVa1Or4tBz4Pgw2eNaVDod7H56Y268h_wohEUaWbUg"
	signatureHeader    = "78aca0c78005107a654a957b8566fa6e0e5e06aea92d7da72a6da9e5a690d013,other"
	timestampHeader    = "1744018920"
	payload            = "payload"
	timestampTolerance = 0
)

func TestVerifyWebhookSignature(t *testing.T) {
	err := VerifyWebhook(payload, secret, signatureHeader, timestampHeader, timestampTolerance)
	if err != nil {
		t.Errorf("Expected no error, but got: %v", err)
	}
}

func TestVerifyWebhookOldTimestamp(t *testing.T) {
	oldTimestampTolerance := 60

	err := VerifyWebhook(payload, secret, signatureHeader, timestampHeader, oldTimestampTolerance)
	if err == nil {
		t.Error("Expected 'timestamp too old' error, but got nil")
	} else if !strings.Contains(err.Error(), "timestamp too old") {
		t.Errorf("Expected 'timestamp too old' error, but got: %v", err)
	}
}

func TestVerifyWebhookWrongSignature(t *testing.T) {
	wrongSignatureHeader := "other"

	err := VerifyWebhook(payload, secret, wrongSignatureHeader, timestampHeader, timestampTolerance)
	if err == nil {
		t.Error("Expected 'no matching signature found' error, but got nil")
	} else if !strings.Contains(err.Error(), "no matching signature found") {
		t.Errorf("Expected 'no matching signature found' error, but got: %v", err)
	}
}

func TestVerifyWebhookInvalidTimestamp(t *testing.T) {
	wrongTimestampHeader := "wrong"

	err := VerifyWebhook(payload, secret, signatureHeader, wrongTimestampHeader, timestampTolerance)
	if err == nil {
		t.Error("Expected 'invalid header timestamp' error, but got nil")
	} else if !strings.Contains(err.Error(), "invalid header timestamp") {
		t.Errorf("Expected 'invalid header timestamp' error, but got: %v", err)
	}
}

func TestVerifyWebhookMissingSignatureHeader(t *testing.T) {
	missingSignatureHeader := ""

	err := VerifyWebhook(payload, secret, missingSignatureHeader, timestampHeader, timestampTolerance)
	if err == nil {
		t.Error("Expected 'missing header values' error, but got nil")
	} else if !strings.Contains(err.Error(), "missing header values") {
		t.Errorf("Expected 'missing header values' error, but got: %v", err)
	}
}

func TestVerifyWebhookMissingTimestampHeader(t *testing.T) {
	missingTimestampHeader := ""

	err := VerifyWebhook(payload, secret, signatureHeader, missingTimestampHeader, timestampTolerance)
	if err == nil {
		t.Error("Expected 'missing header values' error, but got nil")
	} else if !strings.Contains(err.Error(), "missing header values") {
		t.Errorf("Expected 'missing header values' error, but got: %v", err)
	}
}
