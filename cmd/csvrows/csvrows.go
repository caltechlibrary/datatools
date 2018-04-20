//
// csvrows - is can filter selected rows, out row ranges or turn each command
// line parameter into a CSV row of output.
//
// @author R. S. Doiel, <rsdoiel@caltech.edu>
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
package main

import (
	"encoding/csv"
	"fmt"
	"os"
	"strings"

	// Caltech Library packages
	"github.com/caltechlibrary/cli"
	"github.com/caltechlibrary/datatools"
)

const (
	//FIXME: maxRows should be calculated from the data rather than be a constant.
	maxRows = 1000000
)

var (
	description = `
%s converts a set of command line args into rows of CSV formated output.
It can also be used to filter or list specific rows of CSV input
The first row is 1 not 0. Often row 1 is the header row and %s makes it
easy to output only the data rows.
`

	examples = `
Simple usage of building a CSV file one rows at a time.

    %s "First,Second,Third" "one,two,three" > 4rows.csv
    %s "ein,zwei,drei" "1,2,3" >> 4rows.csv
    cat 4row.csv

Example parsing a pipe delimited string into a CSV line

    %s -d "|" "First,Second,Third|one,two,three" > 4rows.csv
    %s -delimiter "|" "ein,zwei,drei|1,2,3" >> 4rows.csv
    cat 4rows.csv

Filter a 10 row CSV file for rows 1,4,6 (top most row is one)

    cat 10row.csv | %s -row 1,4,6 > 3rows.csv

Filter a 10 row CSV file for rows 1,4,6 from file named "10row.csv"

    %s -i 10row.csv -row 1,4,6 > 3rows.csv

Filter 3 randomly selected rows from 10row.csv rendering new CSV with
a header row from 10row.csv.

	%s -i 10row.csv -header=true -random=3
`

	// Standard options
	showHelp             bool
	showLicense          bool
	showVersion          bool
	showExamples         bool
	inputFName           string
	outputFName          string
	generateMarkdownDocs bool
	quiet                bool

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

func main() {
	app := cli.NewCli(datatools.Version)
	appName := app.AppName()

	// Document non-option parameters
	app.AddParams(`[ARGS_AS_ROW_VALUES]`)

	// Add Help Docs
	app.AddHelp("license", []byte(fmt.Sprintf(datatools.LicenseText, appName, datatools.Version)))
	app.AddHelp("description", []byte(fmt.Sprintf(description, appName)))
	app.AddHelp("examples", []byte(fmt.Sprintf(examples, appName, appName, appName, appName, appName, appName, appName)))

	// Standard options
	app.BoolVar(&showHelp, "h,help", false, "display help")
	app.BoolVar(&showLicense, "l,license", false, "display license")
	app.BoolVar(&showVersion, "v,version", false, "display version")
	app.BoolVar(&showExamples, "examples", false, "display example(s)")
	app.StringVar(&inputFName, "i,input", "", "input filename")
	app.StringVar(&outputFName, "o,output", "", "output filename")
	app.BoolVar(&generateMarkdownDocs, "generate-markdown-docs", false, "generate markdown documentation")
	app.BoolVar(&quiet, "quiet", false, "suppress error messages")

	// Application specific options
	app.StringVar(&delimiter, "d,delimiter", "", "set delimiter character")
	app.StringVar(&outputRows, "row,rows", "", "output specified rows in order (e.g. -row 1,5,2-4))")
	app.BoolVar(&skipHeaderRow, "skip-header-row", false, "skip the header row (alias for -row 2-")
	app.BoolVar(&showHeader, "header", false, "display the header row (alias for '-rows 1')")
	app.IntVar(&randomRows, "random", 0, "return N randomly selected rows")
	app.BoolVar(&lazyQuotes, "use-lazy-quotes", false, "use lazy quotes for CSV input")
	app.BoolVar(&trimLeadingSpace, "trim-leading-space", false, "trim leading space in field(s) for CSV input")

	// Parse env and options
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

	// Process Options
	if generateMarkdownDocs {
		app.GenerateMarkdownDocs(app.Out)
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

	if randomRows > 0 {
		if err := datatools.CSVRandomRows(app.In, app.Out, showHeader, randomRows, delimiter, lazyQuotes, trimLeadingSpace); err != nil {
			fmt.Fprintf(app.Eout, "%s, %s\n", inputFName, err)
			os.Exit(1)
		}
		os.Exit(0)
	}

	if outputRows != "" {
		rowNos, err := datatools.ParseRange(outputRows)
		cli.ExitOnError(app.Eout, err, quiet)

		// NOTE: We need to adjust from humans counting from 1 to counting from zero
		for i := 0; i < len(rowNos); i++ {
			rowNos[i] = rowNos[i] - 1
			if rowNos[i] < 0 {
				rowNos[i] = 0
			}
		}
		if err := datatools.CSVRows(app.In, app.Out, showHeader, rowNos, delimiter, lazyQuotes, trimLeadingSpace); err != nil {
			fmt.Fprintf(app.Eout, "%s, %s\n", inputFName, err)
			os.Exit(1)
		}
		os.Exit(0)
	}
	if inputFName != "" {
		if err := datatools.CSVRowsAll(app.In, app.Out, showHeader, delimiter, lazyQuotes, trimLeadingSpace); err != nil {
			fmt.Fprintf(app.Eout, "%s, %s\n", inputFName, err)
			os.Exit(1)
		}
		os.Exit(0)
	}

	// NOTE: If we're not processing an existing CSV source for input we're turning parameters into CSV rows!
	if len(delimiter) > 0 && len(args) == 1 {
		args = strings.Split(args[0], datatools.NormalizeDelimiter(delimiter))
	}

	// Clean up cells removing outer quotes if necessary
	w := csv.NewWriter(app.Out)
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
		cli.ExitOnError(app.Eout, err, quiet)
		r = nil
		if err := w.Write(record); err != nil {
			cli.ExitOnError(app.Eout, err, quiet)
		}
	}
	w.Flush()
	if err := w.Error(); err != nil {
		cli.ExitOnError(app.Eout, err, quiet)
	}
}
