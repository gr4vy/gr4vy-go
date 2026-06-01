package harness

import (
	"bytes"
	"crypto/rand"
	"encoding/hex"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os"
	"path/filepath"
	"strings"
	"sync"
)

// jsonInterceptor wraps an http.RoundTripper to do two test-only things:
//
//  1. Forward-compatibility: it injects a random unknown field into every JSON
//     *object* response, proving the SDK tolerates fields it does not know
//     about. Disable with GR4VY_NO_INJECT=1 (useful when debugging).
//  2. Endpoint-reach coverage: when GR4VY_TRACK_HTTP=1 it appends the request
//     method + path of every completed exchange to coverage/http/*.jsonl, which
//     scripts/endpoint_coverage reads to report which operations a real HTTP
//     request reached.
//
// Both are best-effort: recording never fails a test, and a non-object body is
// passed through untouched.
type jsonInterceptor struct {
	next   http.RoundTripper
	track  bool
	inject bool
}

func newInterceptor(next http.RoundTripper) *jsonInterceptor {
	return &jsonInterceptor{
		next:   next,
		track:  os.Getenv("GR4VY_TRACK_HTTP") == "1",
		inject: os.Getenv("GR4VY_NO_INJECT") != "1",
	}
}

func (i *jsonInterceptor) RoundTrip(req *http.Request) (*http.Response, error) {
	resp, err := i.next.RoundTrip(req)
	if err != nil {
		return nil, err
	}

	if i.track {
		recordCall(req.Method, req.URL.Path)
	}

	if !i.inject || !strings.Contains(resp.Header.Get("Content-Type"), "application/json") {
		return resp, nil
	}

	originalBody, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, fmt.Errorf("failed to read original response body: %w", err)
	}
	resp.Body.Close()

	// Only inject into JSON objects; arrays/scalars are left as-is.
	var data map[string]interface{}
	if err := json.Unmarshal(originalBody, &data); err != nil {
		resp.Body = io.NopCloser(bytes.NewReader(originalBody))
		return resp, nil
	}

	data["unexpected_field_"+randomToken()] = "this is an injected test value"
	modifiedBody, err := json.Marshal(data)
	if err != nil {
		return nil, fmt.Errorf("failed to marshal modified body: %w", err)
	}
	resp.Body = io.NopCloser(bytes.NewReader(modifiedBody))
	resp.ContentLength = int64(len(modifiedBody))
	resp.Header.Set("Content-Length", fmt.Sprintf("%d", len(modifiedBody)))

	return resp, nil
}

func randomToken() string {
	var b [4]byte
	if _, err := rand.Read(b[:]); err != nil {
		return "0"
	}
	return hex.EncodeToString(b[:])
}

// recordCall appends one {"method","path"} line to a per-process JSONL file
// under coverage/http/. Each process writes its own file (keyed by pid + a
// random suffix) so parallel shards never contend on the same file. Failures
// are swallowed — coverage tracking must never break a test.
var recordMu sync.Mutex

func recordCall(method, path string) {
	recordMu.Lock()
	defer recordMu.Unlock()

	// Go runs each test package in its own working directory, so a relative
	// path would scatter coverage across test/flows, test/processing, etc.
	// CI sets GR4VY_COVERAGE_DIR to one absolute directory so every shard
	// writes to the same place; locally it falls back to ./coverage/http.
	dir := os.Getenv("GR4VY_COVERAGE_DIR")
	if dir == "" {
		dir = filepath.Join("coverage", "http")
	}
	if err := os.MkdirAll(dir, 0o755); err != nil {
		return
	}
	file := filepath.Join(dir, fmt.Sprintf("calls-%d-%s.jsonl", os.Getpid(), randomToken()))
	f, err := os.OpenFile(file, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0o644)
	if err != nil {
		return
	}
	defer f.Close()

	line, err := json.Marshal(map[string]string{"method": method, "path": path})
	if err != nil {
		return
	}
	_, _ = f.Write(append(line, '\n'))
}
