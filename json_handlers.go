package datatools

import (
	"bytes"
	"encoding/json"
	"io"
)

// JSONUnmarshal is a custom JSON decoder so we can treat numbers easier
func JSONUnmarshal(src []byte, data interface{}) error {
	dec := json.NewDecoder(bytes.NewReader(src))
	dec.UseNumber()
	err := dec.Decode(&data)
	if err != nil && err != io.EOF {
		return err
	}
	return nil
}

// JSONMarshal provides provide a custom json encoder to solve a an issue with
// HTML entities getting converted to UTF-8 code points by json.Marshal(), json.MarshalIndent().
func JSONMarshal(data interface{}) ([]byte, error) {
	buf1 := []byte{}
	w1 := bytes.NewBuffer(buf1)
	enc := json.NewEncoder(w1)
	enc.SetIndent("", "")
	enc.SetEscapeHTML(false)
	err := enc.Encode(data)
	if err != nil {
		return nil, err
	}
	src1 := w1.Bytes()

	// compact the record so it takes up only one line.
	buf2 := []byte{}
	w2 := bytes.NewBuffer(buf2)
	err = json.Compact(w2, src1)
	src2 := w2.Bytes()
	return src2, err
}

// JSONMarshalIndent provides provide a custom json encoder to solve a an issue with
// HTML entities getting converted to UTF-8 code points by json.Marshal(), json.MarshalIndent().
func JSONMarshalIndent(data interface{}, prefix string, indent string) ([]byte, error) {
	buf := []byte{}
	w := bytes.NewBuffer(buf)
	enc := json.NewEncoder(w)
	enc.SetEscapeHTML(false)
	enc.SetIndent(prefix, indent)
	err := enc.Encode(data)
	if err != nil {
		return nil, err
	}
	return w.Bytes(), err
}
