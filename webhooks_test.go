package gr4vygo

import (
	"strings"
	"testing"
)

// // VerifyWebhook verifies a webhook signature
// func VerifyWebhook(payload, secret, signatureHeader, timestampHeader string, timestampTolerance int) error {
// 	if signatureHeader == "" || timestampHeader == "" {
// 		return fmt.Errorf("missing header values")
// 	}

// 	timestamp, err := strconv.ParseInt(timestampHeader, 10, 64)
// 	if err != nil {
// 		return fmt.Errorf("invalid header timestamp")
// 	}

// 	signatures := strings.Split(signatureHeader, ",")

// 	// Create expected signature
// 	message := fmt.Sprintf("%d.%s", timestamp, payload)
// 	h := hmac.New(sha256.New, []byte(secret))
// 	h.Write([]byte(message))
// 	expectedSignature := hex.EncodeToString(h.Sum(nil))

// 	// Check if expected signature is in the list
// 	signatureFound := false
// 	for _, sig := range signatures {
// 		if sig == expectedSignature {
// 			signatureFound = true
// 			break
// 		}
// 	}

// 	if !signatureFound {
// 		return fmt.Errorf("no matching signature found")
// 	}

// 	// Check timestamp tolerance
// 	if timestampTolerance > 0 && timestamp < time.Now().Unix()-int64(timestampTolerance) {
// 		return fmt.Errorf("timestamp too old")
// 	}

// 	return nil
// }

func TestVerifyWebhookSignature(t *testing.T) {
	secret := "Ik4L-8FH0ihWczctcIPXZRR_8F0fPNgmhEfVBbZ3zNwqQVa1Or4tBz4Pgw2eNaVDod7H56Y268h_wohEUaWbUg"
	signatureHeader := "78aca0c78005107a654a957b8566fa6e0e5e06aea92d7da72a6da9e5a690d013,other"
	timestampHeader := "1744018920"
	payload := "payload"
	timestampTolerance := 0

	err := VerifyWebhook(payload, secret, signatureHeader, timestampHeader, timestampTolerance)
	if err != nil {
		t.Errorf("Expected no error, but got: %v", err)
	}
}

func TestVerifyWebhookOldTimestamp(t *testing.T) {
	secret := "Ik4L-8FH0ihWczctcIPXZRR_8F0fPNgmhEfVBbZ3zNwqQVa1Or4tBz4Pgw2eNaVDod7H56Y268h_wohEUaWbUg"
	signatureHeader := "78aca0c78005107a654a957b8566fa6e0e5e06aea92d7da72a6da9e5a690d013,other"
	timestampHeader := "1744018920"
	payload := "payload"
	timestampTolerance := 60

	err := VerifyWebhook(payload, secret, signatureHeader, timestampHeader, timestampTolerance)
	if err == nil {
		t.Error("Expected 'timestamp too old' error, but got nil")
	} else if !strings.Contains(err.Error(), "timestamp too old") {
		t.Errorf("Expected 'timestamp too old' error, but got: %v", err)
	}
}

func TestVerifyWebhookWrongSignature(t *testing.T) {
	secret := "Ik4L-8FH0ihWczctcIPXZRR_8F0fPNgmhEfVBbZ3zNwqQVa1Or4tBz4Pgw2eNaVDod7H56Y268h_wohEUaWbUg"
	signatureHeader := "other"
	timestampHeader := "1744018920"
	payload := "payload"
	timestampTolerance := 0

	err := VerifyWebhook(payload, secret, signatureHeader, timestampHeader, timestampTolerance)
	if err == nil {
		t.Error("Expected 'no matching signature found' error, but got nil")
	} else if !strings.Contains(err.Error(), "no matching signature found") {
		t.Errorf("Expected 'no matching signature found' error, but got: %v", err)
	}
}

func TestVerifyWebhookInvalidTimestamp(t *testing.T) {
	secret := "Ik4L-8FH0ihWczctcIPXZRR_8F0fPNgmhEfVBbZ3zNwqQVa1Or4tBz4Pgw2eNaVDod7H56Y268h_wohEUaWbUg"
	signatureHeader := "78aca0c78005107a654a957b8566fa6e0e5e06aea92d7da72a6da9e5a690d013,other"
	timestampHeader := "wrong"
	payload := "payload"
	timestampTolerance := 0

	err := VerifyWebhook(payload, secret, signatureHeader, timestampHeader, timestampTolerance)
	if err == nil {
		t.Error("Expected 'invalid header timestamp' error, but got nil")
	} else if !strings.Contains(err.Error(), "invalid header timestamp") {
		t.Errorf("Expected 'invalid header timestamp' error, but got: %v", err)
	}
}

func TestVerifyWebhookMissingSignatureHeader(t *testing.T) {
	secret := "Ik4L-8FH0ihWczctcIPXZRR_8F0fPNgmhEfVBbZ3zNwqQVa1Or4tBz4Pgw2eNaVDod7H56Y268h_wohEUaWbUg"
	signatureHeader := ""
	timestampHeader := "1744018920"
	payload := "payload"
	timestampTolerance := 0

	err := VerifyWebhook(payload, secret, signatureHeader, timestampHeader, timestampTolerance)
	if err == nil {
		t.Error("Expected 'missing header values' error, but got nil")
	} else if !strings.Contains(err.Error(), "missing header values") {
		t.Errorf("Expected 'missing header values' error, but got: %v", err)
	}
}

func TestVerifyWebhookMissingTimestampHeader(t *testing.T) {
	secret := "Ik4L-8FH0ihWczctcIPXZRR_8F0fPNgmhEfVBbZ3zNwqQVa1Or4tBz4Pgw2eNaVDod7H56Y268h_wohEUaWbUg"
	signatureHeader := "78aca0c78005107a654a957b8566fa6e0e5e06aea92d7da72a6da9e5a690d013,other"
	timestampHeader := ""
	payload := "payload"
	timestampTolerance := 0

	err := VerifyWebhook(payload, secret, signatureHeader, timestampHeader, timestampTolerance)
	if err == nil {
		t.Error("Expected 'missing header values' error, but got nil")
	} else if !strings.Contains(err.Error(), "missing header values") {
		t.Errorf("Expected 'missing header values' error, but got: %v", err)
	}
}
