//
// datatools package is a collection of Go based command
// line tools for working with JSON content
//
// @Author R. S. Doiel, <rsdoiel@caltech.edu>
//
// Copyright (c) 2019, Caltech
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
package datatools

import (
	"testing"
)

// TestParseRange tests basic range expressions
func TestParseRange(t *testing.T) {
	expr := "1"
	expected := []int{1}
	result, err := ParseRange(expr)
	if err != nil {
		t.Errorf("expected (%s) no errors, got %s", expr, err)
	}
	for i, got := range result {
		if got != expected[i] {
			t.Errorf("expected (%s) %d, got %d", expr, expected[i], got)
		}
	}

	expr = "1,2,3"
	expected = []int{1, 2, 3}
	result, err = ParseRange(expr)
	if err != nil {
		t.Errorf("expected (%s) no errors, got %s", expr, err)
	}
	for i, got := range result {
		if got != expected[i] {
			t.Errorf("expected (%s) %d, got %d", expr, expected[i], got)
		}
	}

	expr = "1-3"
	result, err = ParseRange(expr)
	if err != nil {
		t.Errorf("expected (%s) no errors, got %s", expr, err)
	}
	for i, got := range result {
		if got != expected[i] {
			t.Errorf("expected (%s) %d, got %d", expr, expected[i], got)
		}
	}

	expr = "1, 2,8 - 10"
	expected = []int{1, 2, 8, 9, 10}
	result, err = ParseRange(expr)
	if err != nil {
		t.Errorf("expected (%s) no errors, got %s", expr, err)
	}
	for i, got := range result {
		if got != expected[i] {
			t.Errorf("expected (%s) %d, got %d", expr, expected[i], got)
		}
	}

}

func TestParseRangeOriginal(t *testing.T) {
	src := `1`
	expected := []int{1}
	result, err := ParseRange(src)
	if err != nil {
		t.Errorf("ParseRange failed, %s", err)
		t.FailNow()
	}
	for i, val := range expected {
		if i >= len(result) {
			t.Errorf("item %d: expected %d, missing element in result d", i, val)
		} else {
			if result[i] != val {
				t.Errorf("item %d: expected %d, got %d", i, val, result[i])
			}
		}
	}

	src = `1:3`
	expected = []int{1, 2, 3}
	result, err = ParseRange(src)
	if err != nil {
		t.Errorf("ParseRange failed, %s", err)
		t.FailNow()
	}
	for i, val := range expected {
		if i >= len(result) {
			t.Errorf("item %d: expected %d, missing element in result d", i, val)
		} else {
			if result[i] != val {
				t.Errorf("item %d: expected %d, got %d", i, val, result[i])
			}
		}
	}

	src = `1,4:6,10`
	expected = []int{
		1,
		4,
		5,
		6,
		10,
	}

	result, err = ParseRange(src)
	if err != nil {
		t.Errorf("ParseRange failed, %s", err)
		t.FailNow()
	}
	for i, val := range expected {
		if i >= len(result) {
			t.Errorf("item %d: expected %d, missing element in result d", i, val)
		} else {
			if result[i] != val {
				t.Errorf("item %d: expected %d, got %d", i, val, result[i])
			}
		}
	}
}
