package tarfs

import (
	"bytes"
	"encoding/hex"
	"io"
	"io/ioutil"
	"strings"
	"testing"

	"github.com/getlantern/testify/assert"
)

func TestRoundTrip(t *testing.T) {
	expectedA, err := ioutil.ReadFile("resources/a.txt")
	if err != nil {
		t.Fatalf("Unable to load expectedA: %v", err)
	}

	expectedB, err := ioutil.ReadFile("localresources/sub/b.txt")
	if err != nil {
		t.Fatalf("Unable to load expectedB: %v", err)
	}

	expectedC, err := ioutil.ReadFile("resources/sub/c.txt")
	if err != nil {
		t.Fatalf("Unable to load expectedC: %v", err)
	}

	tarLiteral := bytes.NewBuffer(nil)
	err = EncodeToTarLiteral("resources", tarLiteral)
	if err != nil {
		t.Fatalf("Unable to encode to tar string: %v", err)
	}

	fs, err := New(tarLiteralToBytes(t, tarLiteral), "localresources")
	if err != nil {
		t.Fatalf("Unable to open filesystem: %v", err)
	}

	a, err := fs.Get("a.txt")
	if assert.NoError(t, err, "a.txt should have loaded") {
		assert.Equal(t, string(expectedA), string(a), "A should have matched expected")
	}

	b, err := fs.Get("sub/b.txt")
	if assert.NoError(t, err, "b.txt should have loaded") {
		assert.Equal(t, string(expectedB), string(b), "B should have matched expected")
	}

	f, err := fs.Open("/nonexistentdirectory/")
	if assert.NoError(t, err, "Opening nonexistent directory should work") {
		fi, err := f.Stat()
		if assert.NoError(t, err, "Should be able to stat directory") {
			assert.True(t, fi.IsDir(), "Nonexistent directory should be a directory")
		}
	}

	f, err = fs.Open("/nonexistentfile")
	assert.Error(t, err, "Opening nonexistent file should fail")

	f, err = fs.Open("/sub//c.txt")
	if assert.NoError(t, err, "Opening existing file with double slash should work") {
		fi, err := f.Stat()
		if assert.NoError(t, err, "Should be able to stat file") {
			if assert.False(t, fi.IsDir(), "File should not be a directory") {
				if assert.Equal(t, len(expectedC), fi.Size(), "File info should report correct size") {
					a := bytes.NewBuffer(nil)
					_, err := io.Copy(a, f)
					if assert.NoError(t, err, "Should be able to read from file") {
						assert.Equal(t, expectedC, a.Bytes(), "Should have read correct data")
					}
				}
			}
		}
	}
}

// tarLiteralToBytes converts a string like []byte {0x4d, 0x5a, 0x90} into a
// byte array.
func tarLiteralToBytes(t *testing.T, bbuf *bytes.Buffer) []byte {
	tarLiteral := string(bbuf.Bytes())
	tarLiteral = strings.Replace(tarLiteral, "[]byte", "", -1)
	tarLiteral = strings.Replace(tarLiteral, "{", "", -1)
	tarLiteral = strings.Replace(tarLiteral, "}", "", -1)
	tarLiteral = strings.Replace(tarLiteral, "\n", "", -1)
	tarLiteral = strings.Replace(tarLiteral, " ", "", -1)

	parts := strings.Split(tarLiteral, ",")
	buf := make([]byte, 0, len(parts)-1)
	for i, s := range parts {
		if i == len(parts)-1 {
			// skip last empty byte
			break
		}
		b, err := hex.DecodeString(s[2:])
		if err != nil {
			t.Fatalf("Unable to decode %v: %v", s, err)
		}
		buf = append(buf, b...)
	}
	return buf
}
