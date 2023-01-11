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
	"flag"
	"fmt"
	"io"
	"os"
	"path"
	"strings"

	// CaltechLibrary packages
	"github.com/caltechlibrary/datatools"

	// 3rd Party packages
	"github.com/tealeg/xlsx"
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

{app_name} [OPTIONS] WORKBOOK_NAME SHEET_NAME

# DESCRIPTION

csv2xlsx will take CSV input and create a new sheet in an Excel Workbook.
If the Workbook does not exist then it is created.

# OPTIONS

-help
: display help

-license
: display license

-version
: display version

-d, -delimiter
: set delimiter character (input)

-i, -input
: input filename (CSV content)

-o, -output
: output filename

-quiet
: suppress error messages

-sheet
: Sheet name to create/replace

-trim-leading-space
: trim leading space in field(s) for CSV input

-use-lazy-quotes
: use lazy quotes for CSV input

-workbook
: Workbook name


# EXAMPLES

Converting a csv to a workbook.

~~~
	{app_name} -i data.csv MyWorkbook.xlsx 'My worksheet 1'
~~~

This creates a new 'My worksheet 1' in the Excel Workbook
called 'MyWorkbook.xlsx' with the contents of data.csv.

~~~
	cat data.csv | {app_name} MyWorkbook.xlsx 'My worksheet 2'
~~~

This does the same but the contents of data.csv are piped into
the workbook's 'My worksheet 2' sheet.

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

	// App Specific Options
	workbookName     string
	sheetName        string
	delimiter        string
	lazyQuotes       bool
	trimLeadingSpace bool
)

func fmtTxt(src string, appName string, version string) string {
	return strings.ReplaceAll(strings.ReplaceAll(src, "{app_name}", appName), "{version}", version)
}

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
		if err != nil {
			fmt.Fprintln(os.Stderr, err)
			os.Exit(1)
		}
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
	appName := path.Base(os.Args[0])

	// Standard Options
	flag.BoolVar(&showHelp, "help", false, "display help")
	flag.BoolVar(&showLicense, "license", false, "display license")
	flag.BoolVar(&showVersion, "version", false, "display version")

	flag.StringVar(&inputFName, "i", "", "input filename (CSV content)")
	flag.StringVar(&inputFName, "input", "", "input filename (CSV content)")
	flag.StringVar(&outputFName, "o", "", "output filename")
	flag.StringVar(&outputFName, "output", "", "output filename")
	flag.BoolVar(&quiet, "quiet", false, "suppress error messages")

	// App Specific Options
	flag.StringVar(&workbookName, "workbook", "", "Workbook name")
	flag.StringVar(&sheetName, "sheet", "", "Sheet name to create/replace")
	flag.StringVar(&delimiter, "d", "", "set delimiter character (input)")
	flag.StringVar(&delimiter, "delimiter", "", "set delimiter character (input)")
	flag.BoolVar(&lazyQuotes, "use-lazy-quotes", false, "use lazy quotes for CSV input")
	flag.BoolVar(&trimLeadingSpace, "trim-leading-space", false, "trim leading space in field(s) for CSV input")

	// Parse environment and options
	flag.Parse()
	args := flag.Args()

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
		fmt.Fprintf(out, "%s\n", fmtTxt(helpText, appName, datatools.Version))
		os.Exit(0)
	}
	if showVersion {
		fmt.Fprintf(out, "%s %s\n", appName, datatools.Version)
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
		fmt.Fprintf(eout, "Missing workbook name")
		os.Exit(1)
	}
	if len(sheetName) == 0 {
		fmt.Fprintf(eout, "Missing sheet name")
		os.Exit(1)
	}
	err = csv2XLSXSheet(in, workbookName, sheetName, delimiter, lazyQuotes, trimLeadingSpace)
	if err != nil {
		fmt.Fprintln(eout, err)
		os.Exit(1)
	}
}
