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

// VerifyWebhook verifies the authenticity of a webhook request by checking its signature and timestamp.
//
// Parameters:
//   - payload: The raw webhook payload as a string.
//   - secret: The webhook secret used to sign the payload.
//   - signatureHeader: The value of the X-Gr4vy-Signature header from the webhook request.
//   - timestampHeader: The value of the X-Gr4vy-Timestamp header from the webhook request.
//   - timestampTolerance: The maximum allowed difference (in seconds) between the current time and the timestamp in the header. Set to 0 to disable timestamp validation.
//
// Returns:
//   - error: Returns an error if verification fails due to missing headers, invalid timestamp, signature mismatch, or timestamp being too old. Returns nil if verification succeeds.
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
