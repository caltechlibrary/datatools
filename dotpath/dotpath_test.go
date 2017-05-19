//
// dotpath.go package provides a convient way of mapping JSON dot path notation
// to a decoded JSON datatype (e.g. map, array)
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
	"encoding/json"
	"testing"
)

func TestParse(t *testing.T) {
	keys, err := parse(".name")
	if err != nil {
		t.Errorf("%s", err)
		t.FailNow()
	}
	if len(keys) != 1 {
		t.Errorf("Expected one key for .name, %+v\n", keys)
		t.FailNow()
	}
	_, err = parse("")
	if err == nil {
		t.Errorf("An empty string show throw an error")
		t.FailNow()
	}
}

func TestFindForMaps(t *testing.T) {
	src := []byte(`{"name":"Fred","age":7}`)
	p := ".name"
	keys, err := parse(p)
	if err != nil {
		t.Errorf(`Expected []string{"name"}, got %+v <-- %s`, keys, err)
		t.FailNow()
	}
	data, err := JSONDecode(src)
	if err != nil {
		t.Errorf("JSONDecode error for %q, %s", src, err)
		t.FailNow()
	}

	m := data.(map[string]interface{})

	// Test simple findInMap
	val, err := findInMap(keys, m)
	if err != nil {
		t.Errorf("findInMap(%+v, %+v) returned unexpected error, %s", keys, m, err)
		t.FailNow()
	}
	switch val.(type) {
	case string:
		valS := val.(string)
		if valS != "Fred" {
			t.Errorf(`Expected valS to contain "Fred", got %s`, valS)
		}
	default:
		t.Errorf("Expected val to be a string, %T %+v\n", val, val)
		t.FailNow()
	}

	val, err = findInMap([]string{"age"}, m)
	switch val.(type) {
	case json.Number:
		i, err := val.(json.Number).Int64()
		if err != nil {
			t.Errorf("Failed to convert JSON number to int64, %s", err)
			t.FailNow()
		}
		if i != 7 {
			t.Errorf("Expected val to be an Number with value of 7, %+v %d", i, i)
			t.FailNow()
		}
	default:
		t.Errorf("Expected val to be an int, %T %+v\n", val, val)
		t.FailNow()
	}

	// Test simple find
	val, err = find([]string{"name"}, m)
	switch val.(type) {
	case string:
		valS := val.(string)
		if valS != "Fred" {
			t.Errorf(`Expected valS to contain "Fred", got %s`, valS)
		}
	default:
		t.Errorf("Expected val to be a string, %T %+v from %+v\n", val, val, m)
		t.FailNow()
	}

	val, err = find([]string{"age"}, m)
	switch val.(type) {
	case json.Number:
		i, err := val.(json.Number).Int64()
		if err != nil {
			t.Errorf("Failed to convert json.Number to int64, %s", err)
			t.FailNow()
		}
		if i != 7 {
			t.Errorf("Expected val to be an int64 with value of 7.0, %+v %d", i, i)
			t.FailNow()
		}
	default:
		t.Errorf("Expected val to be an int, %T %+v\n", val, val)
		t.FailNow()
	}

}

func TestFindInArray(t *testing.T) {
	src := []byte(`[1,2,3]`)
	p := `[1]`
	var data interface{}
	data, err := JSONDecode(src)
	if err != nil {
		t.Errorf("%s", err)
		t.FailNow()
	}
	_, err = find([]string{p}, data)
	if err != nil {
		t.Errorf("find(%+v, %+v) returned error %s", []string{p}, data, err)
		t.FailNow()
	}
}

func TestEval(t *testing.T) {
	var (
		err error
		src []byte
	)

	// Check to see if we can fit value of dot when it is a string.
	p := "."
	src = []byte(`"Hello World"`)
	if data, err := JSONDecode(src); err != nil {
		t.Errorf("Can't decode string %q, error %s", src, err)
		t.FailNow()
	} else {
		blob, err := Eval(".", data)
		if err != nil {
			t.Errorf("Eval(%q, %q) -> %s", p, src, err)
			t.FailNow()
		}
		switch blob.(type) {
		case string:
			// We're OK
			if blob.(string) != "Hello World" {
				t.Errorf("Expected %q, got %q", "Hello World", blob)
				t.FailNow()
			}
		default:
			// something went wrong
			t.Errorf("Eval(%q, %q) -> wrong type: %T", p, src, blob)
		}
	}

	// Decode a Number
	p = "."
	src = []byte(`1`)
	if data, err := JSONDecode(src); err != nil {
		t.Errorf("Can't decode string %q, error %s", src, err)
		t.FailNow()
	} else {
		blob, err := Eval(".", data)
		if err != nil {
			t.Errorf("Eval(%q, %q) -> %s", p, src, err)
			t.FailNow()
		}
		switch blob.(type) {
		case json.Number:
			// We're OK
			if i, _ := blob.(json.Number).Int64(); i != int64(1) {
				t.Errorf("Expected %q, got %q", "Hello World", blob)
				t.FailNow()
			}
		default:
			// something went wrong
			t.Errorf("Eval(%q, %q) -> wrong type: %T", p, src, blob)
		}
	}

	// Decode an Object
	p = "."
	src = []byte(`{"greeting": "Hello World"}`)
	if data, err := JSONDecode(src); err != nil {
		t.Errorf("Can't decode string %q, error %s", src, err)
		t.FailNow()
	} else {
		blob, err := Eval(".", data)
		if err != nil {
			t.Errorf("Eval(%q, %q) -> %s", p, src, err)
			t.FailNow()
		}
		switch blob.(type) {
		case map[string]interface{}:
			// We're OK
			m := blob.(map[string]interface{})
			if s, _ := m["greeting"]; s.(string) != "Hello World" {
				t.Errorf("Expected %q, got %q", "Hello World", blob)
				t.FailNow()
			}
		default:
			// something went wrong
			t.Errorf("Eval(%q, %q) -> wrong type: %T", p, src, blob)
		}
	}

	// Decode an Array
	p = "."
	src = []byte(`[1,2,3]`)
	if data, err := JSONDecode(src); err != nil {
		t.Errorf("Can't decode string %q, error %s", src, err)
		t.FailNow()
	} else {
		blob, err := Eval(".", data)
		if err != nil {
			t.Errorf("Eval(%q, %q) -> %s", p, src, err)
			t.FailNow()
		}
		switch blob.(type) {
		case []interface{}:
			// We're OK
			a := blob.([]interface{})
			for i, v := range []int{1, 2, 3} {
				if j, err := a[i].(json.Number).Int64(); err == nil {
					if j != int64(v) {
						t.Errorf("Expected %d, got %d for %+v", v, j, a[i])
					}
				} else {
					t.Errorf("Expected %d, got %+v, %T", v, a[i], a[i])
					t.FailNow()
				}
			}
		default:
			// something went wrong
			t.Errorf("Eval(%q, %q) -> wrong type: %T", p, src, blob)
		}
	}

	// More complex object test
	src = []byte(`{
	"display_name": "Fred Zip",
	"sort_name": "Zip, Fred",
	"count_string": [
		"One",
		"Two",
		"Three"
	],
	"count_number": [ 1, 2, 3 ],
	"name": {
		"first": "Fred",
		"last": "Zip"
	},
	"works": [
		{
			"title":"One",
			"pubdate": { "year": 1992, "month": 10, "day": 23 }
		},
		{
			"title": "Two",
			"pubdate": { "year": 2017, "month": 2, "day": 21 }
		},
		{
			"title": "Three",
			"pubdate": { "year": 2003, "month": 12, "day": 1 }
		}
	]
}`)

	data := map[string]interface{}{}
	blob, err := JSONDecode(src)
	if err != nil {
		t.Errorf("JSONDecode error for %q, %s", src, err)
		t.FailNow()
	}
	data = blob.(map[string]interface{})

	// Check to make sure simple data worked
	if expected, ok := data["display_name"]; ok == true {
		switch expected.(type) {
		case string:
		default:
			t.Errorf("Expected a string for display_name")
			t.FailNow()
		}
		p = ".display_name"
		blob, err := Eval(p, data)
		if err != nil {
			t.Errorf("Eval() returned an error, %s", err)
			t.FailNow()
		}
		switch blob.(type) {
		case string:
		default:
			t.Errorf("Expected display_string to be type string, %T", blob)
		}
		if expected.(string) != blob.(string) {
			t.Errorf("Expected %q, got %q", expected, blob)
			t.FailNow()
		}
	}
	// Now see how we handle object in an object
	if expected, ok := data["works"]; ok == true {
		switch expected.(type) {
		case []interface{}:
			// We have an erray
		default:
			t.Errorf("Expected []interface{}, %T", expected)
			t.FailNow()
		}
		p = ".works"
		blob, err := Eval(p, data)
		if err != nil {
			t.Errorf("Eval() returned an error, %s", err)
			t.FailNow()
		}
		switch blob.(type) {
		case []interface{}:
			if len(blob.([]interface{})) != 3 {
				t.Errorf("Expected length 3, got %d", len(blob.([]interface{})))
				t.FailNow()
			}
		default:
			t.Errorf("Expected []interface{}, %T", blob)
			t.FailNow()
		}
		// Now try .count_string[] and see what happens
		p = ".count_string[]"
		blob, err = Eval(p, data)
		if err != nil {
			t.Errorf("Eval() returned an error, %s", err)
			t.FailNow()
		}
		switch blob.(type) {
		case []interface{}:
			if len(blob.([]interface{})) != 3 {
				t.Errorf("Expected length 3, got %d", len(blob.([]interface{})))
				t.FailNow()
			}
			for i, s := range []string{"One", "Two", "Three"} {
				elem := blob.([]interface{})[i]
				if elem.(string) != s {
					t.Errorf("Expected %q, got %T -> %+v", s, elem, elem)
				}
			}

		default:
			t.Errorf("Expected []interface{}, %T", blob)
			t.FailNow()
		}

		p = ".works[0]"
		blob, err = Eval(p, data)
		if err != nil {
			t.Errorf("Eval() returned an error, %s", err)
			t.FailNow()
		}
		switch blob.(type) {
		case map[string]interface{}:
			if val, ok := blob.(map[string]interface{})["title"]; ok == true {
				switch val.(type) {
				case string:
					if val.(string) != "One" {
						t.Errorf("Expected %q, got %q", "One", val)
					}
				default:
					t.Errorf("Expected string from map[\"title\"], %T", val)
					t.FailNow()
				}
			} else {
				t.Errorf("Missing key %q, %+v", "title", blob)
				t.FailNow()
			}
		default:
			t.Errorf("Expected map[string]interface{}, %T", blob)
			t.FailNow()
		}

		p = ".works[0].title"
		blob, err = Eval(p, data)
		if err != nil {
			t.Errorf("Eval() returned an error, %s", err)
			t.FailNow()
		}
		switch blob.(type) {
		case string:
			if val := blob.(string); val != "One" {
				t.Errorf("Expected %q, got %T -> %+v", "One", blob, blob)
				t.FailNow()
			}
		default:
			t.Errorf("Expected string, %T", blob)
			t.FailNow()
		}

		// Now test the [:] array operation
		p = ".works[:].title"
		blob, err = Eval(p, data)
		if err != nil {
			t.Errorf("Eval() returned an error, %s", err)
			t.FailNow()
		}
		switch blob.(type) {
		case []interface{}:
			obj := blob.([]interface{})
			for i, s := range []string{"One", "Two", "Three"} {
				if i < len(obj) && obj[i].(string) != s {
					t.Errorf("Expected %q, got %T -> %+v", s, obj, obj)
					t.FailNow()
				} else if i >= len(obj) {
					t.Errorf("Missing array value at %d -> %+v", i, obj)
				}
			}
		default:
			t.Errorf("Expected string, %T", blob)
			t.FailNow()
		}

		// Test return a sub-slice
		p = ".works[1:3].pubdate"
		blob, err = Eval(p, data)
		if err != nil {
			t.Errorf("Eval() returned an error, %s", err)
			t.FailNow()
		}
		switch blob.(type) {
		case []interface{}:
			obj := blob.([]interface{})
			if len(obj) != 2 {
				t.Errorf("Expected two elements in array, got %T -> %+v", obj, obj)
				t.FailNow()
			}
			pubdate := obj[0].(map[string]interface{})
			if day, ok := pubdate["day"]; ok == true {
				if dy, err := day.(json.Number).Int64(); err == nil && dy != int64(21) {
					t.Errorf("Expected day == 21, got %d", day)
				} else if err != nil {
					t.Errorf("Error getting a day for %T %+v, %s", obj, obj, err)
				}
			} else {
				t.Errorf("Expected a day, day missing, %T %+v", obj, obj)
			}
			pubdate = obj[1].(map[string]interface{})
			if year, ok := pubdate["year"]; ok == true {
				if yr, err := year.(json.Number).Int64(); err == nil && yr != int64(2003) {
					t.Errorf("Expected year == 2003, got %d", year)
				} else if err != nil {
					t.Errorf("Error getting a year for %T %+v, %s", obj, obj, err)
				}
			} else {
				t.Errorf("Expected a year, year missing, %T %+v", obj, obj)
			}
		default:
			t.Errorf("Expected string, %T", blob)
			t.FailNow()
		}

	} else {
		t.Errorf("Expected data to have %q, %T -> %+v", p, data, data)
		t.FailNow()
	}

}
