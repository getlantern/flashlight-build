package runhide

import (
	"testing"
)

func TestOpenURL(t *testing.T) {
	cmd := Command("cmd", "/C", "start", "", "https://www.google.com")
	out, err := cmd.CombinedOutput()
	t.Logf("Error?: %v", err)
	t.Log(out)
}
