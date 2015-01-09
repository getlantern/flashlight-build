package bindings

import (
	"testing"
)

func TestStartClient(t *testing.T) {
	if err := RunClientProxy("0.0.0.0:8080"); err != nil {
		t.Fatalf("RunClientProxy: %q", err)
	}
}
