package harness

import (
	"errors"

	"github.com/gr4vy/gr4vy-go/models/apierrors"
)

// statusOf extracts the HTTP status code from an SDK error, returning ok=false
// if the error did not carry one (i.e. the request never reached the wire, or
// failed for a non-HTTP reason such as a transport/timeout error).
//
// The Gr4vy Go SDK surfaces errors three ways: the catch-all *apierrors.APIError
// (which exposes StatusCode), a per-status typed struct (Error400, Error404, …,
// each with a Status field), and *apierrors.HTTPValidationError for 422s.
func statusOf(err error) (int, bool) {
	if err == nil {
		return 0, false
	}

	var apiErr *apierrors.APIError
	if errors.As(err, &apiErr) {
		return apiErr.StatusCode, true
	}

	var validationErr *apierrors.HTTPValidationError
	if errors.As(err, &validationErr) {
		return 422, true
	}

	// Typed per-status errors all expose a *int64 Status defaulted to the code.
	deref := func(p *int64) (int, bool) {
		if p == nil {
			return 0, false
		}
		return int(*p), true
	}
	var (
		e400 *apierrors.Error400
		e401 *apierrors.Error401
		e403 *apierrors.Error403
		e404 *apierrors.Error404
		e405 *apierrors.Error405
		e409 *apierrors.Error409
		e425 *apierrors.Error425
		e429 *apierrors.Error429
		e500 *apierrors.Error500
		e502 *apierrors.Error502
		e504 *apierrors.Error504
	)
	switch {
	case errors.As(err, &e400):
		if s, ok := deref(e400.Status); ok {
			return s, true
		}
		return 400, true
	case errors.As(err, &e401):
		if s, ok := deref(e401.Status); ok {
			return s, true
		}
		return 401, true
	case errors.As(err, &e403):
		if s, ok := deref(e403.Status); ok {
			return s, true
		}
		return 403, true
	case errors.As(err, &e404):
		if s, ok := deref(e404.Status); ok {
			return s, true
		}
		return 404, true
	case errors.As(err, &e405):
		if s, ok := deref(e405.Status); ok {
			return s, true
		}
		return 405, true
	case errors.As(err, &e409):
		if s, ok := deref(e409.Status); ok {
			return s, true
		}
		return 409, true
	case errors.As(err, &e425):
		if s, ok := deref(e425.Status); ok {
			return s, true
		}
		return 425, true
	case errors.As(err, &e429):
		if s, ok := deref(e429.Status); ok {
			return s, true
		}
		return 429, true
	case errors.As(err, &e500):
		if s, ok := deref(e500.Status); ok {
			return s, true
		}
		return 500, true
	case errors.As(err, &e502):
		if s, ok := deref(e502.Status); ok {
			return s, true
		}
		return 502, true
	case errors.As(err, &e504):
		if s, ok := deref(e504.Status); ok {
			return s, true
		}
		return 504, true
	}
	return 0, false
}

// Reaches asserts that an operation reached the API and was cleanly rejected
// with a client error (4xx). It is used for endpoints that cannot reach a 2xx
// in the mock environment — live wallet provisioning, real network tokens, a
// payout-capable PSP, a configured 3DS or gift-card service — so we still prove
// the SDK serialises the request correctly and talks to the right endpoint.
//
//   - A 4xx is success: the endpoint was reached and rejected the (intentionally
//     incomplete) input.
//   - A 5xx fails the test: we sent something the API could not handle, which is
//     a real defect to surface rather than hide.
//   - A 1xx–3xx (no error) also counts as reached.
//   - A non-HTTP error (transport/timeout) fails: the request never reached the
//     wire, so the assertion proved nothing.
func Reaches(t TestingT, description string, fn func() error) {
	t.Helper()
	err := fn()
	if err == nil {
		// Reached and accepted — fine for a reach assertion.
		return
	}
	status, ok := statusOf(err)
	if !ok {
		t.Fatalf("%s: did not reach the API (non-HTTP error): %v", description, err)
		return
	}
	if status >= 500 {
		t.Fatalf("%s: reached the API but got a server error (%d): %v", description, status, err)
		return
	}
	if status < 400 {
		// A 1xx–3xx surfaced as an error is unusual but still "reached".
		return
	}
	// 4xx — reached and cleanly rejected. Success.
}
