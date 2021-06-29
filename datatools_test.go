//
// datatools package is a collection of Go based command
// line tools for working with CSV, JSON and plain text content
//
// @Author R. S. Doiel, <rsdoiel@caltech.edu>
//
// Copyright (c) 2021, Caltech
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
	"strings"
	"testing"
)

func TestNormalizeDelimiter(t *testing.T) {
	src := `\n`
	expected := "\n"
	result := NormalizeDelimiter(src)
	if result != expected {
		t.Errorf("Expected new line, got %q", result)
	}
	src = `\t`
	expected = "\t"
	result = NormalizeDelimiter(src)
	if result != expected {
		t.Errorf("Expected %q, got %q", expected, result)
	}
}

func TestApplyStopWords(t *testing.T) {
	stopWords := []string{
		"the",
		"a",
		"of",
	}
	src := `The Red Book of Westmarch`
	expected := []string{
		"Red",
		"Book",
		"Westmarch",
	}
	result := ApplyStopWords(strings.Split(src, " "), stopWords)
	for i, val := range expected {
		if result[i] != val {
			t.Errorf("Expected %d %q, got %q", i, val, result[i])
		}
	}

	stopWords = []string{
		"red",
	}
	expected = []string{
		"The",
		"Book",
		"of",
		"Westmarch",
	}
	result = ApplyStopWords(strings.Split(src, " "), stopWords)
	for i, val := range expected {
		if result[i] != val {
			t.Errorf("Expected %d %q, got %q", i, val, result[i])
		}
	}
	src = "Red Book"
	stopWords = []string{
		"red",
	}
	expected = []string{
		"Book",
	}
	result = ApplyStopWords(strings.Split(src, " "), stopWords)
	for i, val := range expected {
		if result[i] != val {
			t.Errorf("Expected %d %q, got %q", i, val, result[i])
		}
	}
}

func TestFilter(t *testing.T) {
	c := rune('R')
	expected := false // Note: Don't filtered out
	result := Filter(c, "", false)
	if expected != result {
		t.Errorf("expected %t, got %t", expected, result)
	}
	c = rune('!')
	expected = true // Note: Filter this out
	result = Filter(c, "", false)
	if expected != result {
		t.Errorf("expected %t, got %t", expected, result)
	}
	expected = false // Note: Don't filter out as we're allow punctation
	result = Filter(c, "", true)
	if expected != result {
		t.Errorf("expected %t, got %t", expected, result)
	}
}
