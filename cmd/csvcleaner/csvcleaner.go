// csvcleaner provides some basic cleaning function that are applied
// across a csv file.
//
// @author R. S. Doiel, <rsdoiel@library.caltech.edu>
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

	// Caltech Library Packages
	"github.com/caltechlibrary/datatools"
)

var (
	helpText = `%{app_name}(1) user manual | version {version} {release_hash}
% R. S. Doiel
% {release_date}

# NAME

{app_name}

# SYNOPSIS

{app_name} [OPTIONS]

# DESCRIPTION

{app_name} normalizes a CSV file based on the options selected. It
helps to address issues like variable number of columns, leading/trailing
spaces in columns, and non-UTF-8 encoding issues.

By default input is expected from standard in and output is sent to 
standard out (errors to standard error). These can be modified by
appropriate options. The csv file is processed as a stream of rows so 
minimal memory is used to operate on the file.

# OPTIONS

-help
: display help

-license
: display license

-version
: display version

-verbose
: write verbose output to standard error

-comma
: if set use this character in place of a comma for delimiting cells

-comment-char
: if set, rows starting with this character will be ignored as comments

-fields-per-row
: set the number of columns to output right padding empty cells as needed

-i, -input
: input filename

-left-trim
: left trim spaces on CSV out

-o, -output
: output filename

-output-comma
: if set use this character in place of a comma for delimiting output cells

-quiet
: suppress error messages

-reuse
: if false then a new array is allocated for each row processed, if true the array gets reused

-right-trim
: right trim spaces on CSV out

-stop-on-error
: exit on error, useful if you're trying to debug a problematic CSV file

-trim, -trim-spaces
: trim spaces on CSV out

-trim-leading-space
: trim leading space from field(s) for CSV input

-use-crlf
: if set use a charage return and line feed in output

-use-lazy-quotes
: use lazy quotes for CSV input


# EXAMPLES

Normalizing a spread sheet's column count to 5 padding columns as needed per row.

~~~
    cat mysheet.csv | {app_name} -field-per-row=5
~~~

Trim leading spaces from output.

~~~
    cat mysheet.csv | {app_name} -left-trim
~~~

Trim trailing spaces from output.

~~~
    cat mysheet.csv | {app_name} -right-trim
~~~

Trim leading and trailing spaces from output.

~~~
    cat mysheet.csv | {app_name} -trim-space
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
	//newLine              bool
	//eol                  string

	// App Options
	comma            string
	rowComment       string
	fieldsPerRecord  int
	trailingComma    bool
	trimSpace        bool
	trimLeftSpace    bool
	trimRightSpace   bool
	reuseRecord      bool
	commaOut         string
	useCRLF          bool
	stopOnError      bool
	lazyQuotes       bool
	trimLeadingSpace bool

	verbose bool
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
	//app.BoolVar(&newLine, "nl,newline", false, "include trailing newline in output")

	// Application specific options
	flag.IntVar(&fieldsPerRecord, "fields-per-row", 0, "set the number of columns to output right padding empty cells as needed")
	flag.BoolVar(&trimSpace, "trim", false, "trim spaces on CSV out")
	flag.BoolVar(&trimSpace, "trim-spaces", false, "trim spaces on CSV out")
	flag.BoolVar(&trimLeftSpace, "left-trim", false, "left trim spaces on CSV out")
	flag.BoolVar(&trimRightSpace, "right-trim", false, "right trim spaces on CSV out")
	flag.BoolVar(&reuseRecord, "reuse", true, "if false then a new array is allocated for each row processed, if true the array gets reused")
	flag.StringVar(&comma, "comma", "", "if set use this character in place of a comma for delimiting cells")
	flag.StringVar(&rowComment, "comment-char", "", "if set, rows starting with this character will be ignored as comments")
	flag.StringVar(&commaOut, "output-comma", "", "if set use this character in place of a comma for delimiting output cells")
	flag.BoolVar(&useCRLF, "use-crlf", false, "if set use a charage return and line feed in output")
	flag.BoolVar(&stopOnError, "stop-on-error", false, "exit on error, useful if you're trying to debug a problematic CSV file")
	flag.BoolVar(&lazyQuotes, "use-lazy-quotes", false, "use lazy quotes for CSV input")
	flag.BoolVar(&trimLeadingSpace, "trim-leading-space", false, "trim leading space from field(s) for CSV input")

	flag.BoolVar(&verbose, "verbose", false, "write verbose output to standard error")

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
		fmt.Fprintf(out, "%s\n", fmtTxt(helpText, appName, datatools.Version))
		os.Exit(0)
	}
	if showLicense {
		fmt.Fprintf(out, "%s\n", fmtTxt(datatools.LicenseText, appName, datatools.Version))
		os.Exit(0)
	}
	if showVersion {
		fmt.Fprintf(out, "%s %s\n", appName, datatools.Version)
		os.Exit(0)
	}

	// Loop through input CSV, apply options, write to output CSV
	if trimLeftSpace == true && trimRightSpace == true {
		trimSpace = true
	}

	// Setup our CSV reader with any cli options
	var rStr []rune

	r := csv.NewReader(in)
	if comma != "" {
		rStr = []rune(comma)
		if len(rStr) > 0 {
			r.Comma = rStr[0]
		}
	}
	if rowComment != "" {
		rStr = []rune(rowComment)
		if len(rStr) > 0 {
			r.Comment = rStr[0]
		}
	}
	r.FieldsPerRecord = fieldsPerRecord
	r.LazyQuotes = lazyQuotes
	r.TrimLeadingSpace = trimLeadingSpace
	r.ReuseRecord = reuseRecord

	w := csv.NewWriter(out)
	if commaOut != "" {
		rStr = []rune(commaOut)
		if len(rStr) > 0 {
			w.Comma = rStr[0]
		}
	}
	w.UseCRLF = useCRLF

	// i is so we can track row count as we process each streamed in row
	hasError := false
	i := 1
	for {
		row, err := r.Read()
		if err == io.EOF {
			break
		}
		if i == 1 && fieldsPerRecord == 0 {
			fieldsPerRecord = len(row)
		}
		if err != nil {
			serr := fmt.Sprintf("%s", err)
			if strings.HasSuffix(serr, "wrong number of fields in line") == true && fieldsPerRecord >= 0 {
				if verbose {
					fmt.Fprintf(eout, "row %d: expected %d, got %d cells\n", i, fieldsPerRecord, len(row))
				}
			} else {
				hasError = true
				if verbose {
					fmt.Fprintln(eout, err)
				}
			}
		}
		// Trim trailing cells if needed
		if fieldsPerRecord > 0 && len(row) > fieldsPerRecord {
			row = row[0:fieldsPerRecord]
		}
		if fieldsPerRecord > 0 && len(row) < fieldsPerRecord {
			// Append cells if needed
			for len(row) < fieldsPerRecord { //expectedCellCount {
				row = append(row, "")
			}
		}
		if trimSpace || trimLeftSpace || trimRightSpace {
			for i := range row {
				s := row[i]
				switch {
				case trimSpace:
					row[i] = strings.TrimSpace(s)
				case trimRightSpace:
					row[i] = strings.TrimRight(s, " \t\n\r")
				case trimLeftSpace:
					row[i] = strings.TrimLeft(s, " \t\n\r")
				}
			}
		}
		if err := w.Write(row); err != nil {
			fmt.Fprintf(eout, "error writing row %d: %s", i, err)
			hasError = true
		}
		i++
		if verbose == true && (i%100) == 0 {
			fmt.Fprintf(eout, "Processed %d rows\n", i)
		}
		if hasError && stopOnError {
			os.Exit(1)
		}
	}
	// Finally we need to flush any remaining output...
	w.Flush()
	err = w.Error()
	if err != nil {
		fmt.Fprintln(eout, err)
		os.Exit(1)
	}
	if verbose == true {
		fmt.Fprintf(eout, "Processed %d rows\n", i)
	}
	//fmt.Fprintf(app.Out, "%s", eol)
}
