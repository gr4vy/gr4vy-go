package processing

import (
	"os"
	"testing"

	"github.com/gr4vy/gr4vy-go/test/harness"
)

func TestMain(m *testing.M) {
	code := m.Run()
	harness.Cleanup()
	os.Exit(code)
}
