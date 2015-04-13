package archiver

import (
	"bytes"
	"encoding/gob"
	"encoding/json"
	"os"
)

type Saver interface {
	Load([]interface{}) error
	Save([]interface{}) error
}

type JSONSaver struct {
	Data []byte
}

func (s *JSONSaver) Load(d *interface{}) (err error) {
	s.Data, err = json.Marshal(d)
	return
}

func (s *JSONSaver) Save(d interface{}) error {
	return json.Unmarshal(s.Data, d)
}

type GobFileSaver struct {
	FileName string
}

func (s *GobFileSaver) Load(d []interface{}) error {
	f, err := os.Open(s.FileName)
	if err != nil {
		return err
	}
	var buf bytes.Buffer
	if _, err := buf.ReadFrom(f); err != nil {
		return err
	}

	enc := gob.NewDecoder(&buf)
	if err := enc.Decode(d); err != nil {
		return err
	}
	return nil
}

func (s *GobFileSaver) Save(d []interface{}) error {
	var buf bytes.Buffer
	enc := gob.NewEncoder(&buf)

	if err := enc.Encode(d); err != nil {
		return err
	}
	f, err := os.Open(s.FileName)
	if err != nil {
		return err
	}
	if _, err := buf.WriteTo(f); err != nil {
		return err
	}

	return nil
}
