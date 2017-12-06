//
// csvcols - is a command line that takes each argument in order and outputs a line in CSV format.
// It can also take a delimiter and line of text splitting it into a CSV formatted set of columns.
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

	// 3rd Party packages
	"github.com/google/uuid"
)

const (
	//FIXME: maxColumns needs to be calculated from the data rather than being a constant
	maxColumns = 2048
)

var (
	description = `
%s converts a set of command line args into columns output in CSV format.
It can also be used CSV input rows and rendering only the column numbers
listed on the commandline (first column is 1 not 0).
`

	examples = `
Simple usage of building a CSV file one row at a time.

    %s one two three > 3col.csv
    %s 1 2 3 >> 3col.csv
    cat 3col.csv

Example parsing a pipe delimited string into a CSV line

    %s -d "|" "one|two|three" > 3col.csv
    %s -delimiter "|" "1|2|3" >> 3col.csv
    cat 3col.csv

Filter a 10 column CSV file for columns 1,4,6 (left most column is one)

    cat 10col.csv | %s -col 1,4,6 > 3col.csv

Filter a 10 columns CSV file for columns 1,4,6 from file named "10col.csv"

    %s -i 10col.csv -col 1,4,6 > 3col.csv
`

	// Standard Options
	showHelp             bool
	showLicense          bool
	showVersion          bool
	showExamples         bool
	inputFName           string
	outputFName          string
	generateMarkdownDocs bool
	quiet                bool
	newLine              bool
	eol                  string

	// App Options
	outputColumns string
	prefixUUID    bool
	skipHeaderRow bool
	delimiter     string
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

func CSVColumns(in *os.File, out *os.File, columnNos []int, prefixUUID bool, skipHeaderRow bool, delimiter string) {
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
	app.AddParams(`[ARGS_AS_COL_VALUES]`)

	// Add Help Docs
	app.AddHelp("license", []byte(fmt.Sprintf(datatools.LicenseText, appName, datatools.Version)))
	app.AddHelp("description", []byte(fmt.Sprintf(description, appName)))
	app.AddHelp("examples", []byte(fmt.Sprintf(examples, appName, appName, appName, appName, appName, appName)))

	// Standard Options
	app.BoolVar(&showHelp, "h,help", false, "display help")
	app.BoolVar(&showLicense, "l,license", false, "display license")
	app.BoolVar(&showVersion, "v,version", false, "display version")
	app.BoolVar(&showExamples, "examples", false, "display example")
	app.StringVar(&inputFName, "i,input", "", "input filename")
	app.StringVar(&outputFName, "o,output", "", "output filename")
	app.BoolVar(&generateMarkdownDocs, "generate-markdown-docs", false, "generate markdown documentation")
	app.BoolVar(&quiet, "quiet", false, "suppress error messages")
	app.BoolVar(&newLine, "nl,newline", true, "include trailing newline in output")

	// App Options
	app.StringVar(&outputColumns, "col,cols", "", "output specified columns (e.g. -col 1,12:14,2,4))")
	app.StringVar(&delimiter, "d,delimiter", "", "set delimiter character")
	app.BoolVar(&skipHeaderRow, "skip-header-row", true, "skip the header row")
	app.BoolVar(&prefixUUID, "uuid", false, "add a prefix row with generated UUID cell")

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
	if newLine {
		eol = "\n"
	}

	if outputColumns != "" {
		columnNos, err := datatools.ParseRange(outputColumns, maxColumns)
		cli.ExitOnError(app.Eout, err, quiet)

		// NOTE: We need to adjust from humans counting from 1 to counting from zero
		for i := 0; i < len(columnNos); i++ {
			columnNos[i] = columnNos[i] - 1
			if columnNos[i] < 0 {
				columnNos[i] = 0
			}
		}
		CSVColumns(app.In, app.Out, columnNos, prefixUUID, skipHeaderRow, delimiter)
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
	if delimiter != "" {
		w.Comma = datatools.NormalizeDelimiterRune(delimiter)
	}
	if err := w.Write(cells); err != nil {
		cli.ExitOnError(app.Eout, fmt.Errorf("error writing args as csv, %s", err), quiet)
	}
	w.Flush()
	if err := w.Error(); err != nil {
		cli.ExitOnError(app.Eout, err, quiet)
	}
	fmt.Fprintf(app.Out, "%s", eol)
}
