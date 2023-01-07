//
// csvcols - is a command line that takes each argument in order and outputs a line in CSV format.
// It can also take a delimiter and line of text splitting it into a CSV formatted set of columns.
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
//
package main

import (
	"encoding/csv"
	"flag"
	"fmt"
	"io"
	"os"
	"path"
	"strings"

	// Caltech Library packages
	"github.com/caltechlibrary/datatools"

	// 3rd Party packages
	"github.com/google/uuid"
)

var (
	helpText = `---
title: "{app_name} (1) user manual"
author: "R. S. Doiel"
pubDate: 2023-01-06
---

# NAME

{app_name}

# SYNOPSIS

{app_name} [OPTIONS] [ARGS_AS_COL_VALUES]

# DESCRIPTION

{app_name} converts a set of command line args into columns output in CSV format.

It can also be used CSV input rows and rendering only the column numbers
listed on the commandline (first column is 1 not 0).

# OPTIONS

-help
: display help

-license
: display license

-version
: display version

-col, -cols
: output specified columns (e.g. -col 1,12:14,2,4))

-d, -delimiter
: set the input delimiter character

-i, -input
: input filename

-o, -output
: output filename

-od, -output-delimiter
: set the output delimiter character

-quiet
: suppress error messages

-skip-header-row
: skip the header row

-trim-leading-space
: trim leading space in field(s) for CSV input

-use-lazy-quotes
: use lazy quotes on CSV input

-uuid
: add a prefix row with generated UUID cell


# EXAMPLES

Simple usage of building a CSV file one row at a time.

~~~
    {app_name} one two three > 3col.csv
    {app_name} 1 2 3 >> 3col.csv
    cat 3col.csv
~~~

Example parsing a pipe delimited string into a CSV line

~~~
    {app_name} -d "|" "one|two|three" > 3col.csv
    {app_name} -delimiter "|" "1|2|3" >> 3col.csv
    cat 3col.csv
~~~

Using a pipe filter a 3 column CSV for columns 1 and 3 into 2col.csv

~~~
    cat 3col.csv | {app_name} -col 1,3 > 2col.csv
~~~


Using options filter a 3 column CSV file for columns 1,3 into 2col.csv

~~~
    {app_name} -i 3col.csv -col 1,3 -o 2col.csv
~~~

{app_name} {version}

`


	// Standard Options
	showHelp         bool
	showLicense      bool
	showVersion      bool
	showExamples     bool
	inputFName       string
	outputFName      string
	generateMarkdown bool
	generateManPage  bool
	quiet            bool

	// App Options
	outputColumns    string
	prefixUUID       bool
	skipHeaderRow    bool
	delimiter        string
	outputDelimiter  string
	lazyQuotes       bool
	trimLeadingSpace bool
)

func fmtTxt(src string, appName string, version string) string {
	return strings.ReplaceAll(strings.ReplaceAll(src, "{app_name}", appName), "{version}", version)
}

func selectedColumns(rowNo int, record []string, columnNos []int, prefixUUID bool, skipHeaderRow bool) []string {
	var id string

	result := []string{}
	if prefixUUID == true {
		if rowNo == 0 {
			id = "uuid"

		} else {
			id = uuid.New().String()
		}
		result = append(result, id)
	}
	l := len(record)
	for _, col := range columnNos {
		if col >= 0 && col < l {
			result = append(result, record[col])
		} else {
			// If we don't find the column, story an empty string
			result = append(result, "")
		}
	}
	return result
}

func CSVColumns(in *os.File, out *os.File, eout *os.File, columnNos []int, prefixUUID bool, skipHeaderRow bool, delimiterIn string, delimiterOut string, lazyQuotes, trimLeadingSpace bool) {
	var err error

	r := csv.NewReader(in)
	r.LazyQuotes = lazyQuotes
	r.TrimLeadingSpace = trimLeadingSpace

	w := csv.NewWriter(out)
	if delimiterIn != "" {
		r.Comma = datatools.NormalizeDelimiterRune(delimiterIn)
	}
	if delimiterOut != "" {
		w.Comma = datatools.NormalizeDelimiterRune(delimiterOut)
	}
	for i := 0; err != io.EOF; i++ {
		rec, err := r.Read()
		if err == io.EOF {
			break
		}
		if err != nil && ! quiet{
			fmt.Fprintln(eout, err)

		}

		row := selectedColumns(i, rec, columnNos, prefixUUID, skipHeaderRow)
		err = w.Write(row)
		if err != nil {
			fmt.Fprintln(eout, err)
			os.Exit(1)
		}
	}
	w.Flush()
	err = w.Error()
	if err != nil {
		fmt.Fprintln(eout, err)
		os.Exit(1)
	}
}

func main() {
	appName := path.Base(os.Args[0])

	// Standard Options
	flag.BoolVar(&showHelp, "help", false, "display help")
	flag.BoolVar(&showLicense, "license", false, "display license")
	flag.BoolVar(&showVersion, "version", false, "display version")

	flag.StringVar(&inputFName, "i", "", "input filename")
	flag.StringVar(&inputFName, "input", "", "input filename")
	flag.StringVar(&outputFName, "o", "", "output filename")
	flag.StringVar(&outputFName, "output", "", "output filename")
	flag.BoolVar(&quiet, "quiet", false, "suppress error messages")

	// App Options
	flag.StringVar(&outputColumns, "col", "", "output specified columns (e.g. -col 1,12:14,2,4))")
	flag.StringVar(&outputColumns, "cols", "", "output specified columns (e.g. -col 1,12:14,2,4))")
	flag.StringVar(&delimiter, "d", "", "set the input delimiter character")
	flag.StringVar(&delimiter, "delimiter", "", "set the input delimiter character")
	flag.StringVar(&outputDelimiter, "od", "", "set the output delimiter character")
	flag.StringVar(&outputDelimiter, "output-delimiter", "", "set the output delimiter character")
	flag.BoolVar(&skipHeaderRow, "skip-header-row", true, "skip the header row")
	flag.BoolVar(&prefixUUID, "uuid", false, "add a prefix row with generated UUID cell")
	flag.BoolVar(&lazyQuotes, "use-lazy-quotes", false, "use lazy quotes on CSV input")
	flag.BoolVar(&trimLeadingSpace, "trim-leading-space", false, "trim leading space in field(s) for CSV input")

	// Parse env and options
	flag.Parse()
	args := flag.Args()

	// Setup IO
	var err error

	in := os.Stdin
	out := os.Stdout
	eout := os.Stderr


if inputFName != "" {
	in, err = os.Open(inputFName)
	if err != nil {
		fmt.Fprintln(eout, err)
		os.Exit(1)
	}
	defer in.Close()
}

if outputFName != "" {
	out, err := os.Create(outputFName)
	if err != nil {
		fmt.Fprintln(eout, err)
		os.Exit(1)
	}
	defer out.Close()
}

	// Process options
	if showHelp {
		fmt.Fprintf(out, "%s\n", fmtTxt(helpText, appName, datatools.Version))
		os.Exit(0)
	}
	if showLicense {
		fmt.Fprintf(out, "%s\n", datatools.LicenseText)
		os.Exit(0)
	}
	if showVersion {
		fmt.Fprintf(out, "%s %s\n", appName, datatools.Version)
		os.Exit(0)
	}

	if outputColumns != "" {
		columnNos, err := datatools.ParseRange(outputColumns)
		if err != nil {
			fmt.Fprintln(eout, err)
			os.Exit(1)
		}
		// NOTE: We need to adjust from humans counting from 1 to counting from zero
		for i := 0; i < len(columnNos); i++ {
			columnNos[i] = columnNos[i] - 1
			if columnNos[i] < 0 {
				columnNos[i] = 0
			}
		}
		CSVColumns(in, out, eout, columnNos, prefixUUID, skipHeaderRow, delimiter, outputDelimiter, lazyQuotes, trimLeadingSpace)
		os.Exit(0)
	}

	if len(delimiter) > 0 && len(args) == 1 {
		args = strings.Split(args[0], datatools.NormalizeDelimiter(delimiter))
	}

	// Clean up cells removing outer quotes if necessary
	cells := []string{}
	for _, val := range args {
		if strings.HasPrefix(val, "\"") && strings.HasSuffix(val, "\"") {
			val = strings.TrimPrefix(strings.TrimSuffix(val, "\""), "\"")
		}
		cells = append(cells, strings.TrimSpace(val))
	}

	w := csv.NewWriter(out)
	if outputDelimiter != "" {
		w.Comma = datatools.NormalizeDelimiterRune(outputDelimiter)
	}
	if err := w.Write(cells); err != nil {
		fmt.Fprint(eout, "error writing args as csv, %s\n", err)
		os.Exit(1)
	}
	w.Flush()
	if err := w.Error(); err != nil {
		fmt.Fprintln(eout, err)
	}
}
