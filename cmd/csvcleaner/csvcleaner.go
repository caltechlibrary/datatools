//
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
//
package main

import (
	"encoding/csv"
	"fmt"
	"io"
	"os"
	"strings"

	// Caltech Library Packages
	"github.com/caltechlibrary/cli"
	"github.com/caltechlibrary/datatools"
)

var (
	description = `
%s normalizes a CSV file based on the options selected. It
helps to address issues like variable number of columns, leading/trailing
spaces in columns, and non-UTF-8 encoding issues.

By default input is expected from standard in and output is sent to 
standard out (errors to standard error). These can be modified by
appropriate options. The csv file is processed as a stream of rows so 
minimal memory is used to operate on the file. 
`

	examples = `
Normalizing a spread sheet's column count to 5 padding columns as needed per row.

    cat mysheet.csv | %s -field-per-row=5

Trim leading spaces from output.

    cat mysheet.csv | %s -left-trim

Trim trailing spaces from output.

    cat mysheet.csv | %s -right-trim

Trim leading and trailing spaces from output.

    cat mysheet.csv | %s -trim-space
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

func main() {
	app := cli.NewCli(datatools.Version)
	appName := app.AppName()

	// Add Help Docs
	app.AddHelp("license", []byte(fmt.Sprintf(datatools.LicenseText, appName, datatools.Version)))
	app.AddHelp("description", []byte(fmt.Sprintf(description, appName)))
	app.AddHelp("examples", []byte(fmt.Sprintf(examples, appName, appName, appName, appName)))

	// Standard options
	app.BoolVar(&showHelp, "h,help", false, "display help")
	app.BoolVar(&showLicense, "l,license", false, "display license")
	app.BoolVar(&showVersion, "v,version", false, "display version")
	app.BoolVar(&showExamples, "examples", false, "display example(s)")
	app.StringVar(&inputFName, "i,input", "", "input filename")
	app.StringVar(&outputFName, "o,output", "", "output filename")
	app.BoolVar(&generateMarkdown, "generate-markdown", false, "generation markdown documentation")
	app.BoolVar(&generateManPage, "generate-manpage", false, "generation man page")
	app.BoolVar(&quiet, "quiet", false, "suppress error messages")
	//app.BoolVar(&newLine, "nl,newline", false, "include trailing newline in output")

	// Application specific options
	app.IntVar(&fieldsPerRecord, "fields-per-row", 0, "set the number of columns to output right padding empty cells as needed")
	app.BoolVar(&trimSpace, "trim,trim-spaces", false, "trim spaces on CSV out")
	app.BoolVar(&trimLeftSpace, "left-trim", false, "left trim spaces on CSV out")
	app.BoolVar(&trimRightSpace, "right-trim", false, "right trim spaces on CSV out")
	app.BoolVar(&reuseRecord, "reuse", true, "if false then a new array is allocated for each row processed, if true the array gets reused")
	app.StringVar(&comma, "comma", "", "if set use this character in place of a comma for delimiting cells")
	app.StringVar(&rowComment, "comment-char", "", "if set, rows starting with this character will be ignored as comments")
	app.StringVar(&commaOut, "output-comma", "", "if set use this character in place of a comma for delimiting output cells")
	app.BoolVar(&useCRLF, "use-crlf", false, "if set use a charage return and line feed in output")
	app.BoolVar(&stopOnError, "stop-on-error", false, "exit on error, useful if you're trying to debug a problematic CSV file")
	app.BoolVar(&lazyQuotes, "use-lazy-quotes", false, "use lazy quotes for CSV input")
	app.BoolVar(&trimLeadingSpace, "trim-leading-space", false, "trim leading space from field(s) for CSV input")

	app.BoolVar(&verbose, "V,verbose", false, "write verbose output to standard error")

	// Parse environment and options
	app.Parse()
	args := app.Args()

	// Setup IO
	var err error

	app.Eout = os.Stderr
	app.In, err = cli.Open(inputFName, os.Stdin)
	cli.ExitOnError(app.Eout, err, quiet)
	defer cli.CloseFile(inputFName, app.In)

	app.Out, err = cli.Create(outputFName, os.Stdout)
	cli.ExitOnError(app.Eout, err, quiet)
	defer cli.CloseFile(outputFName, app.Out)

	// Process options
	if generateMarkdown {
		app.GenerateMarkdown(app.Out)
		os.Exit(0)
	}
	if generateManPage {
		app.GenerateManPage(app.Out)
		os.Exit(0)
	}
	if showHelp || showExamples {
		if len(args) > 0 {
			fmt.Fprintln(app.Out, app.Help(args...))
		} else {
			app.Usage(app.Out)
		}
		os.Exit(0)
	}
	if showLicense {
		fmt.Fprintln(app.Out, app.License())
		os.Exit(0)
	}
	if showVersion {
		fmt.Fprintln(app.Out, app.Version())
		os.Exit(0)
	}

	// Loop through input CSV, apply options, write to output CSV
	if trimLeftSpace == true && trimRightSpace == true {
		trimSpace = true
	}

	// Setup our CSV reader with any cli options
	var rStr []rune

	r := csv.NewReader(app.In)
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

	w := csv.NewWriter(app.Out)
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
					cli.OnError(app.Eout, fmt.Errorf("row %d: expected %d, got %d cells\n", i, fieldsPerRecord, len(row)), quiet)
				}
			} else {
				hasError = true
				if verbose {
					cli.OnError(app.Eout, fmt.Errorf("%s", err), quiet)
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
			cli.OnError(app.Eout, fmt.Errorf("error writing row %d: %s", i, err), quiet)
			hasError = true
		}
		i++
		if verbose == true && (i%100) == 0 {
			cli.OnError(app.Eout, fmt.Errorf("Processed %d rows\n", i), quiet)
		}
		if hasError && stopOnError {
			os.Exit(1)
		}
	}
	// Finally we need to flush any remaining output...
	w.Flush()
	err = w.Error()
	cli.ExitOnError(app.Eout, err, quiet)
	if verbose == true {
		cli.OnError(app.Eout, fmt.Errorf("Processed %d rows\n", i), quiet)
	}
	//fmt.Fprintf(app.Out, "%s", eol)
}
