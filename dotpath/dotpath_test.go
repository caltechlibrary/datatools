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
	"encoding/json"
	//"fmt"
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

/*
func TestEval(t *testing.T) {
	t.Error("dotpath.Eval(dotpath, jsonSource) not implemented")
}
*/
