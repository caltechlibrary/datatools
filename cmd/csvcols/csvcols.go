//
// csvcols - is a command line that takes each argument in order and outputs a line in CSV format.
// It can also take a delimiter and line of text splitting it into a CSV formatted set of columns.
//
// @author R. S. Doiel, <rsdoiel@caltech.edu>
//
// Copyright (c) 2018, Caltech
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

	// 3rd Party packages
	"github.com/google/uuid"
)

var (
	description = `
%s converts a set of command line args into columns output in CSV format.
It can also be used CSV input rows and rendering only the column numbers
listed on the commandline (first column is 1 not 0).
`

	examples = `
Simple usage of building a CSV file one row at a time.

    csvcols one two three > 3col.csv
    csvcols 1 2 3 >> 3col.csv
    cat 3col.csv

Example parsing a pipe delimited string into a CSV line

    csvcols -d "|" "one|two|three" > 3col.csv
    csvcols -delimiter "|" "1|2|3" >> 3col.csv
    cat 3col.csv

Using a pipe filter a 3 column CSV for columns 1 and 3 into 2col.csv

    cat 3col.csv | csvcols -col 1,3 > 2col.csv

Using options filter a 3 column CSV file for columns 1,3 into 2col.csv

    csvcols -i 3col.csv -col 1,3 -o 2col.csv
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

func CSVColumns(in *os.File, out *os.File, columnNos []int, prefixUUID bool, skipHeaderRow bool, delimiterIn string, delimiterOut string, lazyQuotes, trimLeadingSpace bool) {
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
		cli.OnError(os.Stderr, err, quiet)

		row := selectedColumns(i, rec, columnNos, prefixUUID, skipHeaderRow)
		err = w.Write(row)
		cli.ExitOnError(os.Stderr, err, quiet)
	}
	w.Flush()
	err = w.Error()
	cli.ExitOnError(os.Stderr, err, quiet)
}

func main() {
	app := cli.NewCli(datatools.Version)
	appName := app.AppName()

	// Document non-option parameters
	app.SetParams(`[ARGS_AS_COL_VALUES]`)

	// Add Help Docs
	app.AddHelp("license", []byte(fmt.Sprintf(datatools.LicenseText, appName, datatools.Version)))
	app.AddHelp("description", []byte(fmt.Sprintf(description, appName)))
	app.AddHelp("examples", []byte(examples))

	// Standard Options
	app.BoolVar(&showHelp, "h,help", false, "display help")
	app.BoolVar(&showLicense, "l,license", false, "display license")
	app.BoolVar(&showVersion, "v,version", false, "display version")
	app.BoolVar(&showExamples, "examples", false, "display example")
	app.StringVar(&inputFName, "i,input", "", "input filename")
	app.StringVar(&outputFName, "o,output", "", "output filename")
	app.BoolVar(&generateMarkdown, "generate-markdown", false, "generate markdown documentation")
	app.BoolVar(&generateManPage, "generate-manpage", false, "generate man page")
	app.BoolVar(&quiet, "quiet", false, "suppress error messages")

	// App Options
	app.StringVar(&outputColumns, "col,cols", "", "output specified columns (e.g. -col 1,12:14,2,4))")
	app.StringVar(&delimiter, "d,delimiter", "", "set the input delimiter character")
	app.StringVar(&outputDelimiter, "od,output-delimiter", "", "set the output delimiter character")
	app.BoolVar(&skipHeaderRow, "skip-header-row", true, "skip the header row")
	app.BoolVar(&prefixUUID, "uuid", false, "add a prefix row with generated UUID cell")
	app.BoolVar(&lazyQuotes, "use-lazy-quotes", false, "use lazy quotes on CSV input")
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

	if outputColumns != "" {
		columnNos, err := datatools.ParseRange(outputColumns)
		cli.ExitOnError(app.Eout, err, quiet)

		// NOTE: We need to adjust from humans counting from 1 to counting from zero
		for i := 0; i < len(columnNos); i++ {
			columnNos[i] = columnNos[i] - 1
			if columnNos[i] < 0 {
				columnNos[i] = 0
			}
		}
		CSVColumns(app.In, app.Out, columnNos, prefixUUID, skipHeaderRow, delimiter, outputDelimiter, lazyQuotes, trimLeadingSpace)
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

	w := csv.NewWriter(app.Out)
	if outputDelimiter != "" {
		w.Comma = datatools.NormalizeDelimiterRune(outputDelimiter)
	}
	if err := w.Write(cells); err != nil {
		cli.ExitOnError(app.Eout, fmt.Errorf("error writing args as csv, %s", err), quiet)
	}
	w.Flush()
	if err := w.Error(); err != nil {
		cli.ExitOnError(app.Eout, err, quiet)
	}
}
