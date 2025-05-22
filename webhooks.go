package gr4vygo

import (
	"crypto/hmac"
	"crypto/sha256"
	"encoding/hex"
	"fmt"
	"strconv"
	"strings"
	"time"
)

// VerifyWebhook verifies a webhook signature
func VerifyWebhook(payload, secret, signatureHeader, timestampHeader string, timestampTolerance int) error {
	// Check for missing headers
	if signatureHeader == "" || timestampHeader == "" {
		return fmt.Errorf("missing header values")
	}

	// Parse timestamp
	timestamp, err := strconv.ParseInt(timestampHeader, 10, 64)
	if err != nil {
		return fmt.Errorf("invalid header timestamp: %w", err)
	}

	// Split signatures by comma
	signatures := strings.Split(signatureHeader, ",")

	// Create expected signature
	message := fmt.Sprintf("%d.%s", timestamp, payload)
	h := hmac.New(sha256.New, []byte(secret))
	h.Write([]byte(message))
	expectedSignature := hex.EncodeToString(h.Sum(nil))

	// Check if expected signature matches any of the provided signatures
	signatureFound := false
	for _, sig := range signatures {
		if sig == expectedSignature {
			signatureFound = true
			break
		}
	}

	if !signatureFound {
		return fmt.Errorf("no matching signature found")
	}

	// Check timestamp tolerance if specified
	if timestampTolerance > 0 {
		currentTime := time.Now().Unix()
		if timestamp < currentTime-int64(timestampTolerance) {
			return fmt.Errorf("timestamp too old")
		}
	}

	return nil
}
