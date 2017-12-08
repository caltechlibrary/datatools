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
	"io"
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
	validateRows  bool
	showRowCount  bool
	showColCount  bool
	showHeader    bool
	skipHeaderRow bool
	outputRows    string
	delimiter     string
)

func selectedRow(rowNo int, record []string, rowNos []int) []string {
	if len(rowNos) == 0 {
		return record
	}
	for _, i := range rowNos {
		if i == rowNo {
			return record
		}
	}
	return nil
}

func CSVRows(in io.Reader, out io.Writer, eout io.Writer, rowNos []int, delimiter string) {
	var err error

	r := csv.NewReader(in)
	w := csv.NewWriter(out)
	if delimiter != "" {
		r.Comma = datatools.NormalizeDelimiterRune(delimiter)
		w.Comma = datatools.NormalizeDelimiterRune(delimiter)
	}
	for i := 0; err != io.EOF; i++ {
		rec, err := r.Read()
		if err == io.EOF {
			break
		}
		if err != nil {
			fmt.Fprintf(eout, "%s, %s (%T %+v)", inputFName, err, rec, rec)
			os.Exit(1)
		}
		row := selectedRow(i, rec, rowNos)
		if row != nil {
			if err := w.Write(row); err != nil {
				fmt.Fprintf(eout, "Error writing record to csv: %s (Row %T %+v)", err, row, row)
				os.Exit(1)
			}
		}
	}
	w.Flush()
	err = w.Error()
	if err != nil {
		fmt.Fprintf(eout, "%s\n", err)
		os.Exit(1)
	}
}

func main() {
	app := cli.NewCli(datatools.Version)
	appName := app.AppName()

	// Document non-option parameters
	app.AddParams(`[ARGS_AS_ROW_VALUES]`)

	// Add Help Docs
	app.AddHelp("license", []byte(fmt.Sprintf(datatools.LicenseText, appName, datatools.Version)))
	app.AddHelp("description", []byte(fmt.Sprintf(description, appName)))
	app.AddHelp("examples", []byte(fmt.Sprintf(examples, appName, appName, appName, appName, appName, appName)))

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
	app.StringVar(&outputRows, "row,rows", "", "output specified rows in order (e.g. -row 1,5,2:4))")
	app.BoolVar(&skipHeaderRow, "skip-header-row", false, "skip the header row (alias for -row 2:")
	app.BoolVar(&showHeader, "header", false, "display the header row (alias for '-rows 1')")

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

	if showHeader == true {
		outputRows = "1"
	}
	if len(args) == 0 && outputRows == "" {
		outputRows = "1:"
		if skipHeaderRow == true {
			outputRows = "2:"
		}
	}

	if outputRows != "" {
		rowNos, err := datatools.ParseRange(outputRows, maxRows)
		cli.ExitOnError(app.Eout, err, quiet)

		// NOTE: We need to adjust from humans counting from 1 to counting from zero
		for i := 0; i < len(rowNos); i++ {
			rowNos[i] = rowNos[i] - 1
			if rowNos[i] < 0 {
				rowNos[i] = 0
			}
		}
		CSVRows(app.In, app.Out, app.Eout, rowNos, delimiter)
		os.Exit(0)
	}

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
