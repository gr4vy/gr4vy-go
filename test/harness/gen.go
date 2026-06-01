package harness

import (
	"math/rand"
	"os"
	"strconv"
)

// Runs is the bounded number of iterations each property-based test performs,
// kept small to limit live-sandbox cost.
const Runs = 6

var currencies = []string{"USD", "EUR", "GBP"}

// NewGen returns a deterministic generator. The seed comes from FC_SEED (so a
// failing run can be reproduced) and defaults to a fixed value otherwise. We
// use the standard library rather than an external property-testing dependency
// so the suite stays hermetic.
func NewGen() *rand.Rand {
	seed := int64(1337)
	if s := os.Getenv("FC_SEED"); s != "" {
		if parsed, err := strconv.ParseInt(s, 10, 64); err == nil {
			seed = parsed
		}
	}
	return rand.New(rand.NewSource(seed))
}

// Amount returns a positive minor-unit amount in a realistic range.
func Amount(g *rand.Rand) int64 {
	return int64(g.Intn(99_900) + 100) // 100..100000
}

// Currency returns one of the accepted ISO currency codes.
func Currency(g *rand.Rand) string {
	return currencies[g.Intn(len(currencies))]
}

// Metadata returns a small map mixing camelCase and snake_case keys, exercising
// the API's tolerance of arbitrary string maps.
func Metadata(g *rand.Rand) map[string]string {
	return map[string]string{
		"camelCaseKey":   strconv.Itoa(g.Intn(1000)),
		"snake_case_key": strconv.Itoa(g.Intn(1000)),
	}
}
