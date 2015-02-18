package tarfs

import (
	"archive/tar"
	"encoding/hex"
	"fmt"
	"io"
	"os"
	"path/filepath"
	"strings"
)

const (
	bytesPerColumn = 12
)

// EncodeToTarLiteral takes the contents of the given directory and writes it to
// the given Writer in the form of a byte array literal, for example
// []byte {0x4d, 0x5a, 0x90}.
func EncodeToTarLiteral(dir string, w io.Writer) error {
	aew := &arrayencodingwriter{w, 0}
	err := aew.start()
	if err != nil {
		return fmt.Errorf("Unable to start byte array literal: %v", err)
	}

	tw := tar.NewWriter(aew)
	defer tw.Close()

	dirPrefix := dir + "/"
	dirPrefixLen := len(dirPrefix)

	err = filepath.Walk(dir, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return fmt.Errorf("Unable to walk to %v: %v", path, err)
		}
		if info.IsDir() {
			return nil
		}
		name := path
		if strings.HasPrefix(name, dirPrefix) {
			name = path[dirPrefixLen:]
		}
		hdr := &tar.Header{
			Name: name,
			Size: info.Size(),
		}
		err = tw.WriteHeader(hdr)
		if err != nil {
			return fmt.Errorf("Unable to write tar header: %v", err)
		}
		file, err := os.Open(path)
		if err != nil {
			return fmt.Errorf("Unable to open file %v: %v", path, err)
		}
		defer file.Close()
		_, err = io.Copy(tw, file)
		if err != nil {
			return fmt.Errorf("Unable to copy file %v to tar: %v", path, err)
		}
		return nil
	})

	if err != nil {
		return err
	}

	err = tw.Close()
	if err != nil {
		return fmt.Errorf("Unable to close tar writer: %v", err)
	}

	return aew.finish()
}

// arrayencodingwriter is a writer that encodes written bytes into a byte array
// literal.
type arrayencodingwriter struct {
	io.Writer
	column int
}

func (w *arrayencodingwriter) start() error {
	_, err := fmt.Fprintf(w.Writer, "[]byte {\n")
	return err
}

func (w *arrayencodingwriter) finish() error {
	_, err := fmt.Fprintf(w.Writer, "\n}")
	return err
}

func (w *arrayencodingwriter) Write(buf []byte) (int, error) {
	n := 0
	for _, b := range buf {
		if w.column == bytesPerColumn {
			// Wrap to next line
			_, err := fmt.Fprintf(w.Writer, "\n")
			if err != nil {
				return 0, err
			}
			w.column = 1
		} else {
			w.column += 1
		}

		_, err := fmt.Fprintf(w.Writer, `0x%v, `, hex.EncodeToString([]byte{b}))
		if err != nil {
			return n, err
		}
		n += 1
	}
	return n, nil
}
