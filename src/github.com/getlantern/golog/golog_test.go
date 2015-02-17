package golog

import (
	"bytes"
	"io"
	"io/ioutil"
	"os"
	"testing"
	"time"

	"github.com/getlantern/testify/assert"
)

var (
	expectedLog      = "myprefix: Hello world\nmyprefix: Hello 5\n"
	expectedTraceLog = expectedLog + "myprefix: Gravy\nmyprefix: TraceWriter closed due to unexpected error: EOF\n"
)

func TestDebug(t *testing.T) {
	out := bytes.NewBuffer(nil)
	SetOutputs(ioutil.Discard, out)
	l := LoggerFor("myprefix")
	l.Debug("Hello world")
	l.Debugf("Hello %d", 5)

	assert.Equal(t, expectedLog, string(out.Bytes()), "Logged information didn't match expected")
}

func TestError(t *testing.T) {
	out := bytes.NewBuffer(nil)
	SetOutputs(out, ioutil.Discard)
	l := LoggerFor("myprefix")
	l.Error("Hello world")
	l.Errorf("Hello %d", 5)

	assert.Equal(t, expectedLog, string(out.Bytes()), "Logged information didn't match expected")
}

func TestTraceEnabled(t *testing.T) {
	originalTrace := os.Getenv("TRACE")
	err := os.Setenv("TRACE", "true")
	if err != nil {
		t.Fatalf("Unable to set trace to true")
	}
	defer os.Setenv("TRACE", originalTrace)

	out := bytes.NewBuffer(nil)
	SetOutputs(ioutil.Discard, out)
	l := LoggerFor("myprefix")
	l.Trace("Hello world")
	l.Tracef("Hello %d", 5)
	tw := l.TraceOut()
	tw.Write([]byte("Gravy\n"))
	tw.(io.Closer).Close()

	// Give trace writer a moment to catch up
	time.Sleep(50 * time.Millisecond)

	assert.Equal(t, expectedTraceLog, string(out.Bytes()), "Logged information didn't match expected")
}

func TestTraceDisabled(t *testing.T) {
	originalTrace := os.Getenv("TRACE")
	err := os.Setenv("TRACE", "false")
	if err != nil {
		t.Fatalf("Unable to set trace to false")
	}
	defer os.Setenv("TRACE", originalTrace)

	out := bytes.NewBuffer(nil)
	SetOutputs(ioutil.Discard, out)
	l := LoggerFor("myprefix")
	l.Trace("Hello world")
	l.Tracef("Hello %d", 5)
	l.TraceOut().Write([]byte("Gravy\n"))

	// Give trace writer a moment to catch up
	time.Sleep(50 * time.Millisecond)

	assert.Equal(t, "", string(out.Bytes()), "Nothing should have been logged")
}
