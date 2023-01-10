//
// csvrows - is can filter selected rows, out row ranges or turn each command
// line parameter into a CSV row of output.
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
	"os"
	"path"
	"strings"

	// Caltech Library packages
	"github.com/caltechlibrary/datatools"
)

const (
	//FIXME: maxRows should be calculated from the data rather than be a constant.
	maxRows = 1000000
)

var (
	helpText = `---
title: "{app_name}"
author: "R. S. Doiel"
pubDate: 2023-01-06
---

# NAME

{app_name} 

# SYNOPSIS

{app_name} [OPTIONS] [ARGS_AS_ROW_VALUES]

# DESCRIPTION

{app_name} converts a set of command line args into rows of CSV
formated output.  It can also be used to filter or list specific rows
of CSV input The first row is 1 not 0. Often row 1 is the header row 
and {app_name} makes it easy to output only the data rows.

# OPTIONS

-help
: display help

-license
: display license

-version
: display version

-d, -delimiter
: set delimiter character

-header
: display the header row (alias for '-rows 1')

-i, -input
: input filename

-o, -output
: output filename

-quiet
: suppress error messages

-random
: return N randomly selected rows

-row, -rows
: output specified rows in order (e.g. -row 1,5,2-4))

-skip-header-row
: skip the header row (alias for -row 2-

-trim-leading-space
: trim leading space in field(s) for CSV input

-use-lazy-quotes
: use lazy quotes for CSV input


# EXAMPLES

Simple usage of building a CSV file one rows at a time.

~~~
    {app_name} "First,Second,Third" "one,two,three" > 4rows.csv
    {app_name} "ein,zwei,drei" "1,2,3" >> 4rows.csv
    cat 4row.csv
~~~

Example parsing a pipe delimited string into a CSV line

~~~
    {app_name} -d "|" "First,Second,Third|one,two,three" > 4rows.csv
    {app_name} -delimiter "|" "ein,zwei,drei|1,2,3" >> 4rows.csv
    cat 4rows.csv
~~~

Filter a 10 row CSV file for rows 1,4,6 (top most row is one)

~~~
    cat 10row.csv | {app_name} -row 1,4,6 > 3rows.csv
~~~

Filter a 10 row CSV file for rows 1,4,6 from file named "10row.csv"

~~~
    {app_name} -i 10row.csv -row 1,4,6 > 3rows.csv
~~~

Filter 3 randomly selected rows from 10row.csv rendering new CSV with
a header row from 10row.csv.

~~~
	{app_name} -i 10row.csv -header=true -random=3
~~~

{app_name} {version}

`

	// Standard options
	showHelp         bool
	showLicense      bool
	showVersion      bool
	showExamples     bool
	inputFName       string
	outputFName      string
	generateMarkdown bool
	generateManPage  bool
	quiet            bool

	// Application specific options
	validateRows     bool
	showRowCount     bool
	showColCount     bool
	showHeader       bool
	skipHeaderRow    bool
	outputRows       string
	delimiter        string
	randomRows       int
	lazyQuotes       bool
	trimLeadingSpace bool
)

func fmtTxt(src string, appName string, version string) string {
	return strings.ReplaceAll(strings.ReplaceAll(src, "{app_name}", appName), "{version}", version)
}

func main() {
	appName := path.Base(os.Args[0])

	// Standard options
	flag.BoolVar(&showHelp, "help", false, "display help")
	flag.BoolVar(&showLicense, "license", false, "display license")
	flag.BoolVar(&showVersion, "version", false, "display version")

	flag.StringVar(&inputFName, "i", "", "input filename")
	flag.StringVar(&inputFName, "input", "", "input filename")
	flag.StringVar(&outputFName, "o", "", "output filename")
	flag.StringVar(&outputFName, "output", "", "output filename")
	flag.BoolVar(&quiet, "quiet", false, "suppress error messages")

	// Application specific options
	flag.StringVar(&delimiter, "d", "", "set delimiter character")
	flag.StringVar(&delimiter, "delimiter", "", "set delimiter character")
	flag.StringVar(&outputRows, "row", "", "output specified rows in order (e.g. -row 1,5,2-4))")
	flag.StringVar(&outputRows, "rows", "", "output specified rows in order (e.g. -row 1,5,2-4))")
	flag.BoolVar(&skipHeaderRow, "skip-header-row", false, "skip the header row (alias for -row 2-")
	flag.BoolVar(&showHeader, "header", false, "display the header row (alias for '-rows 1')")
	flag.IntVar(&randomRows, "random", 0, "return N randomly selected rows")
	flag.BoolVar(&lazyQuotes, "use-lazy-quotes", false, "use lazy quotes for CSV input")
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
		out, err = os.Create(outputFName)
		if err != nil {
			fmt.Fprintln(eout, err)
			os.Exit(1)
		}
		defer out.Close()
	}


	// Process Options
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

	if randomRows > 0 {
		if err := datatools.CSVRandomRows(in, out, showHeader, randomRows, delimiter, lazyQuotes, trimLeadingSpace); err != nil {
			fmt.Fprintf(eout, "%s, %s\n", inputFName, err)
			os.Exit(1)
		}
		os.Exit(0)
	}

	if outputRows != "" {
		rowNos, err := datatools.ParseRange(outputRows)
		if err != nil {
			fmt.Fprintln(eout, err)
			os.Exit(1)
		}

		// NOTE: We need to adjust from humans counting from 1 to counting from zero
		for i := 0; i < len(rowNos); i++ {
			rowNos[i] = rowNos[i] - 1
			if rowNos[i] < 0 {
				rowNos[i] = 0
			}
		}
		if err := datatools.CSVRows(in, out, showHeader, rowNos, delimiter, lazyQuotes, trimLeadingSpace); err != nil {
			fmt.Fprintf(eout, "%s, %s\n", inputFName, err)
			os.Exit(1)
		}
		os.Exit(0)
	}
	if inputFName != "" {
		if err := datatools.CSVRowsAll(in, out, showHeader, delimiter, lazyQuotes, trimLeadingSpace); err != nil {
			fmt.Fprintf(eout, "%s, %s\n", inputFName, err)
			os.Exit(1)
		}
		os.Exit(0)
	}

	// NOTE: If we're not processing an existing CSV source for input we're turning parameters into CSV rows!
	if len(delimiter) > 0 && len(args) == 1 {
		args = strings.Split(args[0], datatools.NormalizeDelimiter(delimiter))
	}

	// Clean up cells removing outer quotes if necessary
	w := csv.NewWriter(out)
	if delimiter != "" {
		w.Comma = datatools.NormalizeDelimiterRune(delimiter)
	}
	for _, val := range args {
		r := csv.NewReader(strings.NewReader(val))
		r.LazyQuotes = lazyQuotes
		r.TrimLeadingSpace = trimLeadingSpace
		if delimiter != "" {
			r.Comma = datatools.NormalizeDelimiterRune(delimiter)
		}
		record, err := r.Read()
		if err != nil {
			fmt.Fprintln(eout, err)
			os.Exit(1)
		}
		r = nil
		if err := w.Write(record); err != nil {
			fmt.Fprintln(eout, err)
			os.Exit(1)
		}
	}
	w.Flush()
	if err := w.Error(); err != nil {
		fmt.Fprintln(eout, err)
		os.Exit(1)
	}
}
