// Command endpoint_coverage reports how many SDK operations were reached by a
// real HTTP request during the E2E suite.
//
// It builds the operation catalogue directly from the generated SDK source
// (each resource method builds its request with url.JoinPath/utils.GenerateURL
// for the path and http.NewRequestWithContext for the method), then matches it
// against the method+path lines the test HTTP client recorded to
// coverage/http/*.jsonl when GR4VY_TRACK_HTTP=1. Newly generated but untested
// endpoints surface as "not reached".
//
// It is a report, not a gate: it always exits 0.
package main

import (
	"bufio"
	"encoding/json"
	"fmt"
	"os"
	"path/filepath"
	"regexp"
	"sort"
	"strings"
)

type operation struct {
	Method string
	Path   string
}

var (
	// "/transactions" or "/transactions/{transaction_id}/capture"
	pathRe   = regexp.MustCompile(`(?:url\.JoinPath\(baseURL,|utils\.GenerateURL\(ctx, baseURL,)\s*"(/[^"]*)"`)
	methodRe = regexp.MustCompile(`http\.NewRequestWithContext\(ctx,\s*"([A-Z]+)"`)
	paramRe  = regexp.MustCompile(`\{[^}]+\}`)
)

func main() {
	root := "."
	if len(os.Args) > 1 {
		root = os.Args[1]
	}

	catalogue := buildCatalogue(root)
	reached := matchReached(catalogue, readCalls(filepath.Join(root, "coverage", "http")))

	report := renderReport(catalogue, reached)

	outDir := filepath.Join(root, "coverage")
	_ = os.MkdirAll(outDir, 0o755)
	outFile := filepath.Join(outDir, "endpoint-coverage.md")
	if err := os.WriteFile(outFile, []byte(report), 0o644); err != nil {
		fmt.Fprintf(os.Stderr, "warning: could not write %s: %v\n", outFile, err)
	}
	fmt.Print(report)
}

// buildCatalogue scans the root-level generated resource files for path+method
// pairs. The path literal appears first; the HTTP method follows on a later
// line in the same method body, so we pair each pending path with the next
// method we see.
func buildCatalogue(root string) []operation {
	entries, err := os.ReadDir(root)
	if err != nil {
		fmt.Fprintf(os.Stderr, "warning: could not read %s: %v\n", root, err)
		return nil
	}

	seen := map[operation]bool{}
	var ops []operation
	for _, e := range entries {
		name := e.Name()
		if e.IsDir() || !strings.HasSuffix(name, ".go") || strings.HasSuffix(name, "_test.go") {
			continue
		}
		data, err := os.ReadFile(filepath.Join(root, name))
		if err != nil {
			continue
		}
		var pendingPath string
		scanner := bufio.NewScanner(strings.NewReader(string(data)))
		scanner.Buffer(make([]byte, 1024*1024), 1024*1024)
		for scanner.Scan() {
			line := scanner.Text()
			if mp := pathRe.FindStringSubmatch(line); mp != nil {
				pendingPath = mp[1]
				continue
			}
			if mm := methodRe.FindStringSubmatch(line); mm != nil && pendingPath != "" {
				op := operation{Method: mm[1], Path: pendingPath}
				if !seen[op] {
					seen[op] = true
					ops = append(ops, op)
				}
				pendingPath = ""
			}
		}
	}
	sort.Slice(ops, func(i, j int) bool {
		if ops[i].Path != ops[j].Path {
			return ops[i].Path < ops[j].Path
		}
		return ops[i].Method < ops[j].Method
	})
	return ops
}

func readCalls(dir string) []operation {
	files, err := filepath.Glob(filepath.Join(dir, "*.jsonl"))
	if err != nil || len(files) == 0 {
		return nil
	}
	var calls []operation
	for _, f := range files {
		data, err := os.ReadFile(f)
		if err != nil {
			continue
		}
		scanner := bufio.NewScanner(strings.NewReader(string(data)))
		for scanner.Scan() {
			line := strings.TrimSpace(scanner.Text())
			if line == "" {
				continue
			}
			var c struct {
				Method string `json:"method"`
				Path   string `json:"path"`
			}
			if err := json.Unmarshal([]byte(line), &c); err == nil && c.Path != "" {
				calls = append(calls, operation{Method: strings.ToUpper(c.Method), Path: c.Path})
			}
		}
	}
	return calls
}

// matchReached attributes each recorded call to the single most-specific
// catalogue operation it matches, so a call to /transactions/{id}/capture is
// not also credited to /transactions/{id}.
func matchReached(catalogue, calls []operation) map[operation]bool {
	type compiled struct {
		op    operation
		re    *regexp.Regexp
		score int
	}
	var compiledOps []compiled
	for _, op := range catalogue {
		segments := strings.Split(op.Path, "/")
		params := len(paramRe.FindAllString(op.Path, -1))
		// More literal segments and fewer params => more specific.
		score := len(segments)*100 - params*10 + len(op.Path)
		// Swap {param} for a sentinel before quoting (QuoteMeta would escape
		// the braces), then swap the quoted sentinel for a path-segment match.
		const sentinel = "\x00PARAM\x00"
		pattern := regexp.QuoteMeta(paramRe.ReplaceAllString(op.Path, sentinel))
		pattern = strings.ReplaceAll(pattern, regexp.QuoteMeta(sentinel), "[^/]+")
		compiledOps = append(compiledOps, compiled{
			op:    op,
			re:    regexp.MustCompile("^" + pattern + "$"),
			score: score,
		})
	}

	reached := map[operation]bool{}
	for _, call := range calls {
		var best *compiled
		for i := range compiledOps {
			c := &compiledOps[i]
			if c.op.Method != call.Method {
				continue
			}
			if c.re.MatchString(call.Path) && (best == nil || c.score > best.score) {
				best = c
			}
		}
		if best != nil {
			reached[best.op] = true
		}
	}
	return reached
}

func renderReport(catalogue []operation, reached map[operation]bool) string {
	hit := 0
	for _, op := range catalogue {
		if reached[op] {
			hit++
		}
	}
	total := len(catalogue)
	pct := 0.0
	if total > 0 {
		pct = float64(hit) / float64(total) * 100
	}

	var b strings.Builder
	fmt.Fprintf(&b, "## Endpoint-reach coverage\n\n")
	fmt.Fprintf(&b, "**%d / %d** operations reached by a real HTTP request (%.0f%%).\n\n", hit, total, pct)

	var missing []operation
	for _, op := range catalogue {
		if !reached[op] {
			missing = append(missing, op)
		}
	}
	if len(missing) == 0 {
		fmt.Fprintf(&b, "Every operation in the SDK was reached. 🎉\n")
		return b.String()
	}
	fmt.Fprintf(&b, "<details>\n<summary>%d not reached</summary>\n\n", len(missing))
	for _, op := range missing {
		fmt.Fprintf(&b, "- `%s %s`\n", op.Method, op.Path)
	}
	fmt.Fprintf(&b, "\n</details>\n")
	return b.String()
}
