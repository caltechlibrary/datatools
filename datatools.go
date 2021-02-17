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
	"bufio"
	"bytes"
	"encoding/csv"
	"encoding/json"
	"fmt"
	"io"
	"strings"
	"unicode"

	// 3rd Party packages
	"github.com/dexyk/stringosim"
)

const (
	Version = `v0.0.26`

	LicenseText = `
%s %s

Copyright (c) 2019, Caltech
All rights not granted herein are expressly reserved by Caltech.

Redistribution and use in source and binary forms, with or without modification, are permitted provided that the following conditions are met:

1. Redistributions of source code must retain the above copyright notice, this list of conditions and the following disclaimer.

2. Redistributions in binary form must reproduce the above copyright notice, this list of conditions and the following disclaimer in the documentation and/or other materials provided with the distribution.

3. Neither the name of the copyright holder nor the names of its contributors may be used to endorse or promote products derived from this software without specific prior written permission.

THIS SOFTWARE IS PROVIDED BY THE COPYRIGHT HOLDERS AND CONTRIBUTORS "AS IS" AND ANY EXPRESS OR IMPLIED WARRANTIES, INCLUDING, BUT NOT LIMITED TO, THE IMPLIED WARRANTIES OF MERCHANTABILITY AND FITNESS FOR A PARTICULAR PURPOSE ARE DISCLAIMED. IN NO EVENT SHALL THE COPYRIGHT HOLDER OR CONTRIBUTORS BE LIABLE FOR ANY DIRECT, INDIRECT, INCIDENTAL, SPECIAL, EXEMPLARY, OR CONSEQUENTIAL DAMAGES (INCLUDING, BUT NOT LIMITED TO, PROCUREMENT OF SUBSTITUTE GOODS OR SERVICES; LOSS OF USE, DATA, OR PROFITS; OR BUSINESS INTERRUPTION) HOWEVER CAUSED AND ON ANY THEORY OF LIABILITY, WHETHER IN CONTRACT, STRICT LIABILITY, OR TORT (INCLUDING NEGLIGENCE OR OTHERWISE) ARISING IN ANY WAY OUT OF THE USE OF THIS SOFTWARE, EVEN IF ADVISED OF THE POSSIBILITY OF SUCH DAMAGE.
`

	// Constants for datatools functions
	AsDelimited = iota
	AsCSV       = iota
	AsJSON      = iota
)

// NormalizeDelimiters handles the messy translation from a format string
// received as an option in the cli to something useful to pass to Join.
func NormalizeDelimiter(s string) string {
	if strings.Contains(s, `\n`) {
		s = strings.Replace(s, `\n`, "\n", -1)
	}
	if strings.Contains(s, `\t`) {
		s = strings.Replace(s, `\t`, "\t", -1)
	}
	return s
}

// NormalizeDelimiterRune take a delimiter string and returns a single Rune
func NormalizeDelimiterRune(s string) rune {
	runes := []rune(NormalizeDelimiter(s))
	if len(runes) > 0 {
		return runes[0]
	}
	return ','
}

// ApplyStopWords takes a list of words (array of strings) and
// removes any occurrences of the stop words return a revised list of
// words.
func ApplyStopWords(fields []string, stopWords []string) []string {
	var results []string
	for _, field := range fields {
		skip := false
		s := strings.ToLower(fmt.Sprintf("%s", field))
		for _, term := range stopWords {
			if strings.Compare(s, term) == 0 {
				skip = true
			}
		}
		if skip == false {
			results = append(results, field)
		}
	}
	return results
}

// Filter filters out characters from string. By default it allows
// letters and numbers through with options for allow punctuation
// and other specific characters. Returns true if matches filter, false otherwise
func Filter(c rune, allowableCharacters string, allowPunctuation bool) bool {
	result := !unicode.IsLetter(c) && !unicode.IsNumber(c)
	if allowPunctuation == true {
		result = result && !unicode.IsPunct(c)
	}
	if len(allowableCharacters) > 0 {
		result = result && !strings.ContainsRune(allowableCharacters, c)
	}
	return result
}

// CSVMarshal takes a list of strings and returns a byte array
// of CSV formated output.
func CSVMarshal(fields []string) ([]byte, error) {
	records := [][]string{}
	row := []string{}

	// Turns fields into a 2D array
	for _, col := range fields {
		row = append(row, string(col))
	}
	records = append(records, row)

	buf := new(bytes.Buffer)
	w := csv.NewWriter(buf)
	w.WriteAll(records)
	if err := w.Error(); err != nil {
		return buf.Bytes(), err
	}
	return buf.Bytes(), nil
}

// Options is the data structure to configure the Text2Fields parser
type Options struct {
	AllowCharacters  string
	AllowPunctuation bool
	ToLower          bool
	ToUpper          bool
	StopWords        []string
	Delimiter        string
	Format           int
}

// Text2Fields process a io.Reader as input and returns byte array of fields and error
// Options provides the configuration to apply
func Text2Fields(r *bufio.Reader, options *Options) ([]byte, error) {
	var (
		fields []string
		words  []string
	)
	for done := false; done == false; {
		// Read in line and convert to byte array
		s, err := r.ReadString('\n')
		if err == io.EOF {
			done = true
		}

		// Preprocess input if needed (i.e. lower/upper case input text)
		if options.ToLower == true {
			s = strings.ToLower(s)
		}
		if options.ToUpper == true {
			s = strings.ToUpper(s)
		}

		// Split into fields appling filter
		fields = strings.FieldsFunc(s, func(c rune) bool {
			return Filter(c, options.AllowCharacters, options.AllowPunctuation)
		})

		// Convert to an array of strings for rendering
		fields = ApplyStopWords(fields, options.StopWords)
		for _, field := range fields {
			words = append(words, field)
		}
	}

	// Output fields as JSON, CSV or delimited
	switch options.Format {
	case AsCSV:
		return CSVMarshal(words)
	case AsJSON:
		return json.Marshal(words)
	default:
		return []byte(strings.Join(words, options.Delimiter)), nil
	}
}

// Levenshtein does a fuzzy match on two strings.
func Levenshtein(src string, target string, insertCost int, deleteCost int, substituteCost int, caseSensitive bool) int {
	var caseInsensitive bool
	// NOTE: flipped the termonology for readability in tools using this function.
	if caseSensitive == true {
		caseInsensitive = false
	} else {
		caseInsensitive = true
	}
	return stringosim.Levenshtein([]rune(src), []rune(target), stringosim.LevenshteinSimilarityOptions{
		InsertCost:      insertCost,
		DeleteCost:      deleteCost,
		SubstituteCost:  substituteCost,
		CaseInsensitive: caseInsensitive,
	})
}

// EnglishTitle - uses an improve capitalization rules for English titles.
// This is based on the approach suggested in the Go language Cookbook:
//     http://golangcookbook.com/chapters/strings/title/
func EnglishTitle(s string) string {
	words := strings.Fields(s)
	smallwords := " a an on the to of in "

	for index, word := range words {
		if strings.Contains(smallwords, " "+word+" ") && index != 0 {
			words[index] = word
		} else {
			words[index] = strings.Title(word)
		}
	}
	return strings.Join(words, " ")
}
