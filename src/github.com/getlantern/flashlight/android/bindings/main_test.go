package flashlight

import (
	"testing"
)

func TestStartClient(t *testing.T) {

	// TODO: Fire up a http.Client in another goroutine to test the proxy.

	// TODO: Find a way to stop the proxy.

	if err := RunClientProxy("0.0.0.0:8080"); err != nil {
		t.Fatalf("RunClientProxy: %q", err)
	}
}
