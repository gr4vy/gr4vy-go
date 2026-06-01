package harness

import "time"

// Until polls fetch until predicate(value) is true or the timeout elapses,
// returning the last fetched value and whether the predicate was satisfied.
// Used for eventually-consistent reads (e.g. a transaction appearing in a list
// shortly after creation). A fetch error is treated as "not yet" and retried.
func Until[T any](timeout, interval time.Duration, fetch func() (T, error), predicate func(T) bool) (T, bool) {
	deadline := time.Now().Add(timeout)
	var last T
	for {
		v, err := fetch()
		if err == nil {
			last = v
			if predicate(v) {
				return v, true
			}
		}
		if time.Now().After(deadline) {
			return last, false
		}
		time.Sleep(interval)
	}
}
