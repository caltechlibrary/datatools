//
// csv2xlsx is a command line utility that will convert a CSV file and insert it
// into a named sheet in an Excel Workbook.
//
// @Author R. S. Doiel, <rsdoiel@caltech.edu>
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
	"flag"
	"fmt"
	"io"
	"os"
	"path"

	// CaltechLibrary packages
	"github.com/caltechlibrary/cli"
	"github.com/caltechlibrary/datatools"

	// 3rd Party packages
	"github.com/tealeg/xlsx"
)

var (
	usage = `USAGE: %s [OPTIONS] WORKBOOK_NAME SHEET_NAME`

	description = `
SYNOPSIS

%s will take CSV input and create a new sheet in an Excel Workbook.
If the Workbook does not exist then it is created. 
`

	examples = `
EXAMPLE

	%s -i data.csv MyWorkbook.xlsx 'My worksheet'

This creates a new 'My worksheet' in the Excel Workbook
called 'MyWorkbook.xlsx' with the contents of data.csv.

	cat data.csv | %s MyWorkbook.xlsx 'My worksheet'

This does the same but the contents of data.csv are piped into
the workbook's sheet.
`

	// Standard Options
	showHelp    bool
	showLicense bool
	showVersion bool
	inputFName  string

	// App Specific Options
	workbookName string
	sheetName    string
	delimiter    string
)

func csv2XLSXSheet(in *os.File, workbookName string, sheetName string, delimiter string) error {
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
	if delimiter != "" {
		r.Comma = datatools.NormalizeDelimiterRune(delimiter)
	}
	for {
		record, err := r.Read()
		if err == io.EOF {
			break
		}
		if err != nil {
			fmt.Fprintf(os.Stderr, "%s\n", err)
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

func init() {
	// Standard Options
	flag.BoolVar(&showHelp, "h", false, "display help")
	flag.BoolVar(&showHelp, "help", false, "display help")
	flag.BoolVar(&showLicense, "l", false, "display license")
	flag.BoolVar(&showLicense, "license", false, "display license")
	flag.BoolVar(&showVersion, "v", false, "display version")
	flag.BoolVar(&showVersion, "version", false, "display version")
	flag.StringVar(&inputFName, "i", "", "input filename (CSV content)")
	flag.StringVar(&inputFName, "input", "", "input filename (CSV content)")

	// App Specific Options
	flag.StringVar(&workbookName, "workbook", "", "Workbook name")
	flag.StringVar(&sheetName, "sheet", "", "Sheet name to create/replace")
	flag.StringVar(&delimiter, "d", "", "set delimiter character (input)")
	flag.StringVar(&delimiter, "delimiter", "", "set delimiter character (input)")
}

func main() {
	appName := path.Base(os.Args[0])
	flag.Parse()
	args := flag.Args()

	// Configuration and command line interation
	cfg := cli.New(appName, appName, fmt.Sprintf(datatools.LicenseText, appName, datatools.Version), datatools.Version)
	cfg.UsageText = fmt.Sprintf(usage, appName)
	cfg.DescriptionText = fmt.Sprintf(description, appName)
	cfg.ExampleText = fmt.Sprintf(examples, appName, appName)

	if showHelp == true {
		fmt.Println(cfg.Usage())
		os.Exit(0)
	}

	if showLicense == true {
		fmt.Println(cfg.License())
		os.Exit(0)
	}

	if showVersion == true {
		fmt.Println(cfg.Version())
		os.Exit(0)
	}

	in, err := cli.Open(inputFName, os.Stdin)
	if err != nil {
		fmt.Fprintf(os.Stderr, "%s\n", err)
		os.Exit(1)
	}
	defer cli.CloseFile(inputFName, in)

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
		fmt.Fprintln(os.Stderr, "Missing workbook name")
		os.Exit(1)
	}
	if len(sheetName) == 0 {
		fmt.Fprintln(os.Stderr, "Missing sheet name")
		os.Exit(1)
	}
	err = csv2XLSXSheet(in, workbookName, sheetName, delimiter)
	if err != nil {
		fmt.Fprintf(os.Stderr, "%s", err)
		os.Exit(1)
	}
}
