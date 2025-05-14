// csv2json - is a command line that takes CSV input from stdin and
// writes out JSON-L expression. It includes support for using the first
// row as field names or default fieldnames (e.g. col0, col1, col2).
// Additionally it can output the resulting JSON data structures as a
// JSON array or individual JSON blobs (one line per blob).
//
// @author R. S. Doiel, <rsdoiel@caltech.edu>
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
package main

import (
	"encoding/csv"
	"flag"
	"fmt"
	"io"
	"os"
	"path"
	"strings"

	// My packages
	"github.com/caltechlibrary/datatools"
)

var (
	helpText = `%{app_name}(1) user manual | version {version} {release_hash}
% R. S. Doiel
% {release_date}

# NAME

{app_name}

# SYNOPSIS

{app_name} [OPTIIONS]

# DESCRIPTION

csv2jsonl reads CSV from stdin and writes a JSON-L to stdout. JSON output
is one object per line. See https://jsonlines.org.

# OPTIONS

-help
: display help

-license
: display license

-version
: display version

-d, -delimiter
: set the delimter character

-examples
: display example(s)

-fields-per-record
: Set the number of fields expected in the CSV read, -1 to turn off

-i, -input
: input filename

-nl, -newline
: include trailing newline in output

-o, -output
: output filename

-quiet
: suppress error output

-reuse-record
: reuse the backing array

-trim-leading-space
: trim leading space in fields for CSV input

-use-header
: treat the first row as field names

-use-lazy-quotes
: use lazy quotes for for CSV input

-for-dataset COLUMN_NO
: if COLUMN_NO is greater than -1 then, generate a dataset load compatible version of the CSV file
using COLUMN_NO as key.

# EXAMPLES

Convert data1.csv to data1.jsonl using Unix pipes.

~~~
    cat data1.csv | csv2jsonl > data1.jsonl
~~~

Convert data1.csv to JSON line (one object line per blob)

~~~
    csv2jsonl data1.csv
~~~

`

	// Standard Options
	showHelp    bool
	showLicense bool
	showVersion bool
	inputFName  string
	outputFName string
	quiet       bool
	newLine     bool
	eol         string

	// Application Options
	useHeader        bool
	asBlobs          bool
	delimiter        string
	lazyQuotes       bool
	trimLeadingSpace bool
	fieldsPerRecord  int
	reuseRecord      bool
	forDataset       int
)

func main() {
	appName := path.Base(os.Args[0])
	version := datatools.Version
	license := datatools.LicenseText
	releaseDate := datatools.ReleaseDate
	releaseHash := datatools.ReleaseHash

	// Standard Options
	flag.BoolVar(&showHelp, "help", showHelp, "display help")
	flag.BoolVar(&showLicense, "license", showLicense, "display license")
	flag.BoolVar(&showVersion, "version", showVersion, "display version")
	flag.StringVar(&inputFName, "i", "", "input filename")
	flag.StringVar(&inputFName, "input", "", "input filename")
	flag.StringVar(&outputFName, "o", "", "output filename")
	flag.StringVar(&outputFName, "output", "", "output filename")
	flag.BoolVar(&quiet, "quiet", false, "suppress error output")
	flag.BoolVar(&newLine, "nl", true, "include trailing newline in output")
	flag.BoolVar(&newLine, "newline", true, "include trailing newline in output")

	// App Options
	flag.BoolVar(&useHeader, "use-header", true, "treat the first row as field names")
	flag.StringVar(&delimiter, "d", "", "set the delimter character")
	flag.StringVar(&delimiter, "delimiter", "", "set the delimter character")
	flag.BoolVar(&lazyQuotes, "use-lazy-quotes", false, "use lazy quotes for for CSV input")
	flag.BoolVar(&trimLeadingSpace, "trim-leading-space", false, "trim leading space in fields for CSV input")
	flag.BoolVar(&reuseRecord, "reuse-record", false, "reuse the backing array")
	flag.IntVar(&fieldsPerRecord, "fields-per-record", 0, "Set the number of fields expected in the CSV read, -1 to turn off")
	flag.IntVar(&forDataset, "for-dataset", -1, "generate a dataset compatible JSON lines output using column number as key")

	// Parse environment and options
	flag.Parse()

	// Setup IO
	var err error

	in := os.Stdin
	out := os.Stdout
	eout := os.Stderr

	if inputFName != "" && inputFName != "-" {
		in, err = os.Open(inputFName)
		if err != nil {
			fmt.Fprintln(eout, err)
			os.Exit(1)
		}
		defer in.Close()
	}
	if outputFName != "" && outputFName != "-" {
		out, err = os.Create(outputFName)
		if err != nil {
			fmt.Fprintln(eout, err)
			os.Exit(1)
		}
		defer out.Close()
	}

	// Process options
	if showHelp {
		fmt.Fprintf(out, "%s\n", datatools.FmtHelp(helpText, appName, version, releaseDate, releaseHash))
		os.Exit(0)
	}
	if showLicense == true {
		fmt.Fprintf(out, "%s\n", license)
		os.Exit(0)
	}
	if showVersion == true {
		fmt.Fprintf(out, "datatools, %s %s %s\n", appName, version, releaseHash)
		os.Exit(0)
	}
	if newLine {
		eol = "\n"
	}

	rowNo := 0
	fieldNames := []string{}
	r := csv.NewReader(in)
	r.Comment = '#'
	r.FieldsPerRecord = fieldsPerRecord
	r.LazyQuotes = lazyQuotes
	r.TrimLeadingSpace = trimLeadingSpace
	r.ReuseRecord = reuseRecord
	if delimiter != "" {
		r.Comma = datatools.NormalizeDelimiterRune(delimiter)
	}
	if useHeader == true {
		row, err := r.Read()
		if err == io.EOF {
			fmt.Fprintln(eout, "No data")
			os.Exit(1)
		}
		if err != nil {
			fmt.Fprintln(eout, err)
			os.Exit(1)
		}
		for _, val := range row {
			fieldNames = append(fieldNames, strings.TrimSpace(val))
		}
		rowNo++
	}
	hasError := false
	object := map[string]interface{}{}
	for {
		row, err := r.Read()
		if err == io.EOF {
			break
		}
		if err != nil {
			fmt.Fprintln(eout, err)
			os.Exit(1)
		}

		// Pad the fieldnames if necessary
		object = map[string]interface{}{}
		key := ""
		for col, val := range row {
			if col < len(fieldNames) {
				object[fieldNames[col]] = val
			} else {
				object[fmt.Sprintf("col_%d", col)] = val
			}
			if (col == forDataset) {
				key = fmt.Sprintf("%s", val);
			}
		}
		var src []byte
		src, err = datatools.JSONMarshal(object)
		if err != nil {
			if !quiet {
				fmt.Fprintf(eout, "error row %d, %s\n", rowNo, err)
			}
			hasError = true
		}
		if (forDataset >= 0) {
			if (key == "") {
				if !quiet {
					fmt.Fprintf(eout, "error row, mising key value for column %d, row %d\n", forDataset, rowNo)
				}
			}
			fmt.Fprintf(out, `{%q:%q,%q:%s}%s`, "key", key, "object", src, eol)
		} else {
			fmt.Fprintf(out, "%s%s", src, eol)
		}
		rowNo++
	}
	if hasError == true {
		os.Exit(1)
	}
}
