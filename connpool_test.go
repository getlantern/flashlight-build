package connpool

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"math/rand"
	"net"
	"os"
	"os/exec"
	"sync/atomic"
	"testing"
	"time"

	"github.com/getlantern/testify/assert"
	"github.com/getlantern/waitforserver"
)

var (
	msg = []byte("HELLO")
)

func TestIt(t *testing.T) {
	poolSize := 20
	claimTimeout := 1 * time.Second
	fillTime := 100 * time.Millisecond

	addr, err := startTestServer()
	if err != nil {
		t.Fatalf("Unable to start test server: %s", err)
	}
	p := &Pool{
		MinSize:      poolSize,
		ClaimTimeout: claimTimeout,
		Dial: func() (net.Conn, error) {
			return net.DialTimeout("tcp", addr, 15*time.Millisecond)
		},
	}

	fdCountStart := countTCPFiles()

	p.Start()
	// Run another Start() concurrently just to make sure it doesn't muck things up
	go p.Start()

	time.Sleep(fillTime)

	openConns := countTCPFiles() - fdCountStart
	assert.Equal(t, poolSize, openConns, "Pool should initially open the right number of conns")

	// Use more than the pooled connections
	connectAndRead(t, p, poolSize*2)

	time.Sleep(fillTime)
	openConns = countTCPFiles() - fdCountStart
	assert.Equal(t, poolSize, openConns, "Pool should fill itself back up to the right number of conns")

	// Wait for connections to time out
	time.Sleep(claimTimeout * 2)

	// Test our connections again
	connectAndRead(t, p, poolSize*2)

	time.Sleep(fillTime)
	openConns = countTCPFiles() - fdCountStart
	assert.Equal(t, poolSize, openConns, "After pooled conns time out, pool should fill itself back up to the right number of conns")

	// Make a dial function that randomly returns closed connections
	p.Dial = func() (net.Conn, error) {
		conn, err := net.DialTimeout("tcp", addr, 15*time.Millisecond)
		// Close about half of the connections immediately to test closed checking
		if err == nil && rand.Float32() > 0.5 {
			conn.Close()
		}
		return conn, err
	}

	// Make sure we can still get connections and use them
	connectAndRead(t, p, poolSize)

	// Wait for pool to fill again
	time.Sleep(fillTime)

	p.Stop()
	// Run another Stop() concurrently just to make sure it doesn't muck things up
	go p.Stop()

	openConns = countTCPFiles() - fdCountStart
	assert.Equal(t, 0, openConns, "After stopping pool, there should be no more open conns")
}

func TestDialFailure(t *testing.T) {
	fail := int32(1)
	dialAttempts := int32(0)

	addr, err := startTestServer()
	if err != nil {
		t.Fatalf("Unable to start test server: %s", err)
	}
	p := &Pool{
		MinSize:              10,
		RedialDelayIncrement: 10 * time.Millisecond,
		MaxRedialDelay:       100 * time.Millisecond,
		Dial: func() (net.Conn, error) {
			atomic.AddInt32(&dialAttempts, 1)
			if fail == int32(1) {
				return nil, fmt.Errorf("I'm failing!")
			}
			return net.DialTimeout("tcp", addr, 15*time.Millisecond)
		},
	}

	p.Start()
	defer p.Stop()

	// Wait for fill to run for a while with a failing connection
	time.Sleep(1 * time.Second)
	assert.True(t, dialAttempts < 500, fmt.Sprintf("Should have had a small number of dial attempts, but had %d", dialAttempts))

	// Now make connection succeed and verify that it works
	atomic.StoreInt32(&fail, 0)
	time.Sleep(100 * time.Millisecond)
	connectAndRead(t, p, 1)

	// Now make the connection fail again so that when we stop, we're stopping
	// while failing (tests a different code path for stopping)
	atomic.StoreInt32(&fail, 1)
	time.Sleep(100 * time.Millisecond)
}

func connectAndRead(t *testing.T, p *Pool, loops int) {
	for i := 0; i < loops; i++ {
		c, err := p.Get()
		if err != nil {
			t.Fatalf("Error getting connection: %s", err)
		}
		read, err := ioutil.ReadAll(c)
		if err != nil {
			t.Fatalf("Error reading from connection: %s", err)
		}
		assert.Equal(t, msg, read, "Should have received %s from server", string(msg))
		c.Close()
	}
}

func startTestServer() (string, error) {
	l, err := net.Listen("tcp", "localhost:0")
	if err != nil {
		return "", err
	}
	err = waitforserver.WaitForServer("tcp", l.Addr().String(), 1*time.Second)
	if err != nil {
		return "", err
	}
	go func() {
		for {
			c, err := l.Accept()
			if err != nil {
				log.Fatalf("Error listening: %s", err)
			}
			_, err = c.Write(msg)
			if err != nil {
				log.Fatalf("Unable to write message: %s", err)
			}
			c.Close()
		}
	}()
	return l.Addr().String(), nil
}

// see https://groups.google.com/forum/#!topic/golang-nuts/c0AnWXjzNIA
func countTCPFiles() int {
	out, err := exec.Command("lsof", "-p", fmt.Sprintf("%v", os.Getpid())).Output()
	if err != nil {
		log.Fatal(err)
	}
	return bytes.Count(out, []byte("TCP")) - 1
}
