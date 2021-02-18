//
// csv2xlsx is a command line utility that will convert a CSV file and insert it
// into a named sheet in an Excel Workbook.
//
// @Author R. S. Doiel, <rsdoiel@caltech.edu>
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

	// CaltechLibrary packages
	"github.com/caltechlibrary/cli"
	"github.com/caltechlibrary/datatools"

	// 3rd Party packages
	"github.com/tealeg/xlsx"
)

var (
	description = `
%s will take CSV input and create a new sheet in an Excel Workbook.
If the Workbook does not exist then it is created.
`

	examples = `
Converting a csv to a workbook.

	%s -i data.csv MyWorkbook.xlsx 'My worksheet 1'

This creates a new 'My worksheet 1' in the Excel Workbook
called 'MyWorkbook.xlsx' with the contents of data.csv.

	cat data.csv | %s MyWorkbook.xlsx 'My worksheet 2'

This does the same but the contents of data.csv are piped into
the workbook's 'My worksheet 2' sheet.
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

	// App Specific Options
	workbookName     string
	sheetName        string
	delimiter        string
	lazyQuotes       bool
	trimLeadingSpace bool
)

func csv2XLSXSheet(in *os.File, workbookName string, sheetName string, delimiter string, lazyQuotes, trimLeadingSpace bool) error {
	var workbook *xlsx.File

	// Open the workbook
	if _, err := os.Stat(workbookName); os.IsNotExist(err) == true {
		workbook = xlsx.NewFile()
	} else {
		workbook, err = xlsx.OpenFile(workbookName)
		if err != nil {
			return err
		}
	}
	// Create our worksheet
	sheet, err := workbook.AddSheet(sheetName)
	if err != nil {
		return err
	}

	r := csv.NewReader(in)
	r.LazyQuotes = lazyQuotes
	r.TrimLeadingSpace = trimLeadingSpace

	if delimiter != "" {
		r.Comma = datatools.NormalizeDelimiterRune(delimiter)
	}
	for {
		record, err := r.Read()
		if err == io.EOF {
			break
		}
		cli.OnError(os.Stderr, err, quiet)
		row := sheet.AddRow()
		for _, val := range record {
			cell := row.AddCell()
			cell.Value = val
		}
	}
	// Now write out the changes
	return workbook.Save(workbookName)
}

func main() {
	app := cli.NewCli(datatools.Version)
	appName := app.AppName()

	// Document non-option parameters
	app.SetParams("WORKBOOK_NAME", "SHEET_NAME")

	// Add Help Docs
	app.AddHelp("license", []byte(fmt.Sprintf(datatools.LicenseText, appName, datatools.Version)))
	app.AddHelp("description", []byte(fmt.Sprintf(description, appName)))
	app.AddHelp("examples", []byte(fmt.Sprintf(examples, appName, appName)))

	// Standard Options
	app.BoolVar(&showHelp, "h,help", false, "display help")
	app.BoolVar(&showLicense, "l,license", false, "display license")
	app.BoolVar(&showVersion, "v,version", false, "display version")
	app.BoolVar(&showExamples, "examples", false, "display example(s)")
	app.StringVar(&inputFName, "i,input", "", "input filename (CSV content)")
	app.StringVar(&outputFName, "o,output", "", "output filename")
	app.BoolVar(&generateMarkdown, "generate-markdown", false, "generate markdown documentation")
	app.BoolVar(&generateManPage, "generate-manpage", false, "generate man page")
	app.BoolVar(&quiet, "quiet", false, "suppress error messages")

	// App Specific Options
	app.StringVar(&workbookName, "workbook", "", "Workbook name")
	app.StringVar(&sheetName, "sheet", "", "Sheet name to create/replace")
	app.StringVar(&delimiter, "d,delimiter", "", "set delimiter character (input)")
	app.BoolVar(&lazyQuotes, "use-lazy-quotes", false, "use lazy quotes for CSV input")
	app.BoolVar(&trimLeadingSpace, "trim-leading-space", false, "trim leading space in field(s) for CSV input")

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

	if strings.HasSuffix(outputFName, ".xlsx") {
		workbookName = outputFName
	}

	if len(workbookName) == 0 && len(args) > 0 {
		workbookName = args[0]
		if len(args) > 1 {
			args = args[1:]
		} else {
			args = []string{}
		}
	}
	if len(sheetName) == 0 && len(args) > 0 {
		sheetName = args[0]
	}

	if len(workbookName) == 0 {
		cli.ExitOnError(app.Eout, fmt.Errorf("Missing workbook name"), quiet)
	}
	if len(sheetName) == 0 {
		cli.ExitOnError(app.Eout, fmt.Errorf("Missing sheet name"), quiet)
	}
	err = csv2XLSXSheet(app.In, workbookName, sheetName, delimiter, lazyQuotes, trimLeadingSpace)
	cli.ExitOnError(app.Eout, err, quiet)
}
