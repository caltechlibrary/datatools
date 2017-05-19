//
// dotpath.go package provides a convient way of mapping JSON dot path notation
// to a nested map structure.
//
// @author R. S. Doiel, <rsdoiel@library.caltech.edu>
//
// Copyright (c) 2017, Caltech
// All rights not granted herein are expressly reserved by Caltech.
//
// Redistribution and use in source and binary forms, with or without modification, are permitted provided that the following conditions are met:
//
// 1. Redistributions of source code must retain the above copyright notice, this list of conditions and the following disclaimer.
//
// 2. Redistributions in binary form must reproduce the above copyright notice, this list of conditions and the following disclaimer in the documentation and/or other materials provided with the distribution.
//
// 3. Neither the name of the copyright holder nor the names of its contributors may be used to endorse or promote products derived from this software without specific prior written permission.
//
// THIS SOFTWARE IS PROVIDED BY THE COPYRIGHT HOLDERS AND CONTRIBUTORS "AS IS" AND ANY EXPRESS OR IMPLIED WARRANTIES, INCLUDING, BUT NOT LIMITED TO, THE IMPLIED WARRANTIES OF MERCHANTABILITY AND FITNESS FOR A PARTICULAR PURPOSE ARE DISCLAIMED. IN NO EVENT SHALL THE COPYRIGHT HOLDER OR CONTRIBUTORS BE LIABLE FOR ANY DIRECT, INDIRECT, INCIDENTAL, SPECIAL, EXEMPLARY, OR CONSEQUENTIAL DAMAGES (INCLUDING, BUT NOT LIMITED TO, PROCUREMENT OF SUBSTITUTE GOODS OR SERVICES; LOSS OF USE, DATA, OR PROFITS; OR BUSINESS INTERRUPTION) HOWEVER CAUSED AND ON ANY THEORY OF LIABILITY, WHETHER IN CONTRACT, STRICT LIABILITY, OR TORT (INCLUDING NEGLIGENCE OR OTHERWISE) ARISING IN ANY WAY OUT OF THE USE OF THIS SOFTWARE, EVEN IF ADVISED OF THE POSSIBILITY OF SUCH DAMAGE.
//
package dotpath

import (
	"bytes"
	"encoding/json"
	"fmt"
	"strconv"
	"strings"
)

// JSONDecode decodes JSON using the json.Number instead of float64 for numeric values.
func JSONDecode(buf []byte) (interface{}, error) {
	var result interface{}
	d := json.NewDecoder(bytes.NewReader(buf))
	d.UseNumber()
	if err := d.Decode(&result); err != nil {
		return nil, err
	}
	return result, nil
}

// EvalJSON takes a dot path plus JSON encoded as byte array and returns the value in the dot path or error
func EvalJSON(p string, src []byte) (interface{}, error) {
	data, err := JSONDecode(src)
	if err != nil {
		return nil, err
	}
	return Eval(p, data)
}

// Eval takes a dot path and interface (either a map[string]interface{} or []interface) and
// returns a value from the dot ath or error
func Eval(p string, data interface{}) (interface{}, error) {
	// Parse the dotpath into an array representing map keys or array indexes
	if p == "." {
		return data, nil
	}
	keys, err := parse(p)
	if err != nil {
		return nil, err
	}
	return find(keys, data)
}

// parse takes a dot path notation string and returns a list of keys to traverse
func parse(p string) ([]string, error) {
	delimiters := []rune{'.', '[', ']', '"'}

	nextField := func(c rune) bool {
		for _, d := range delimiters {
			if c == d {
				return true
			}
		}
		return false
	}

	if strings.HasPrefix(p, ".") {
		return strings.FieldsFunc(p, nextField), nil
	} else if strings.Contains(p, ".") {
		return strings.FieldsFunc(p, nextField), nil
	}
	return nil, fmt.Errorf("%q is an invalid dot path", p)
}

// normalizeKey handles square bracket case for map keys
func normalizeKey(s string) string {
	if strings.HasPrefix(s, "[") && strings.HasSuffix(s, "]") {
		l := len(s)
		s = strings.TrimSpace(s[1 : l-1]) // removing brackets
		if strings.HasPrefix(s, `"`) && strings.HasSuffix(s, `"`) {
			l = len(s)
			s = s[1 : l-1]
		}
	}
	return s
}

// find evals the zero'th element in key path and based on the notation calls either
// findInMap or findInArray else returns error type not known
func find(p []string, v interface{}) (interface{}, error) {
	l := len(p)
	if l < 1 {
		return nil, fmt.Errorf("dot path exhausted")
	}
	key := normalizeKey(p[0])
	if key != p[0] {
		p[0] = key
	}
	switch v.(type) {
	case map[string]interface{}:
		return findInMap(p, v.(map[string]interface{}))
	case []interface{}:
		return findInArray(p, v.([]interface{}))
	}
	return nil, fmt.Errorf("invalid dot path key %q in %+v of type %T", p[0], v, v)
}

// findInMap takes an array of strings that represent path elements (e.g. keys in a map structure)
// and either returns the value found at the end of the path OR an error if not found
func findInMap(p []string, m map[string]interface{}) (interface{}, error) {
	if len(p) > 0 {
		if val, ok := m[p[0]]; ok == true {
			if len(p) == 1 {
				return val, nil
			}
			// else recursively eval next element in key path
			return find(p[1:], val)
		}
	}
	return nil, fmt.Errorf("value not found")
}

// findInArray takes a path element and returns the value found at the point in the array
// or an error if element not present.
func findInArray(p []string, a []interface{}) (interface{}, error) {
	if len(p) > 0 {
		if strings.Contains(p[0], ":") == true {
			pts := strings.Split(p[0], ":")
			if len(pts) != 2 {
				return nil, fmt.Errorf("%q is an invalid range", p[0])
			}
			var (
				i, j int
				err  error
			)
			if strings.TrimSpace(pts[0]) == "" {
				i = 0
			} else {
				i, err = strconv.Atoi(pts[0])
			}
			if err != nil {
				return nil, fmt.Errorf("error parsing start of range %q, %s", p[0], err)
			}
			if strings.TrimSpace(pts[1]) == "" {
				j = len(a)
			} else {
				j, err = strconv.Atoi(pts[1])
			}
			if err != nil {
				return nil, fmt.Errorf("error parsing end of range %q, %s", p[0], err)
			}
			if len(p) > 1 {
				v := []interface{}{}
				for _, sVal := range a[i:j] {
					if d, err := find(p[1:], sVal); err != nil {
						return nil, fmt.Errorf("Can't find %q in %+v", p[1], sVal)
					} else {
						v = append(v, d)
					}
				}
				return v, nil
			}
			return a[i:j], nil
		}
		i, err := strconv.Atoi(p[0])
		if err != nil {
			return nil, fmt.Errorf("Can't parse array index %q", p[0])
		}
		if i < 0 || i >= len(a) {
			return nil, fmt.Errorf("index %d is out of bounds", i)
		}
		if len(p) == 1 {
			if i < len(a) && i >= 0 {
				return a[i], nil
			}
		}
		return find(p[1:], a[i])
	}
	return nil, fmt.Errorf("value not found")
}
