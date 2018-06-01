//
// xlsx2json is a command line utility that converts an Excel
// Workboom Sheet into JSON.
//
// @Author R. S. Doiel
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
	"encoding/json"
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
%s is a tool that converts individual Excel Workbook Sheets into
JSON output.
`

	examples = `
This would get the sheet named "Sheet 1" from "MyWorkbook.xlsx" and save as sheet1.json

    %s MyWorkbook.xlsx "My worksheet 1" > sheet1.json

This would get the number of sheets in the workbook

    %s -count MyWorkbook.xlsx

This will output the title of the sheets in the workbook

    %s -sheets MyWorkbook.xlsx

Putting it all together in a shell script and convert all sheets to
into JSON documents..

	%s -N MyWorkbook.xlsx | while read SHEET_NAME; do
    	JSON_NAME="${SHEET_NAME// /-}.json"
    	%s -o "${JSON_NAME}" MyWorkbook.xlsx "$SHEET_NAME"
	done    
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

	// Application Options
	showSheetCount bool
	showSheetNames bool
)

func sheetCount(workBookName string) (int, error) {
	xlFile, err := xlsx.OpenFile(workBookName)
	if err != nil {
		return 0, err
	}
	return len(xlFile.Sheet), nil
}

func sheetNames(workBookName string) ([]string, error) {
	xlFile, err := xlsx.OpenFile(workBookName)
	if err != nil {
		return []string{}, err
	}
	result := []string{}
	for sheetName := range xlFile.Sheet {
		result = append(result, sheetName)
	}
	return result, nil
}

func xlsx2JSON(out io.Writer, workBookName, sheetName string) error {
	xlFile, err := xlsx.OpenFile(workBookName)
	if err != nil {
		return err
	}
	results := [][]string{}
	cells := []string{}
	if sheet, ok := xlFile.Sheet[sheetName]; ok == true {
		for _, row := range sheet.Rows {
			cells = []string{}
			for _, cell := range row.Cells {
				val := cell.String()
				cells = append(cells, val)
			}
			results = append(results, cells)
		}
		src, err := json.Marshal(results)
		if err != nil {
			return err
		}
		fmt.Fprintf(out, "%s", src)
		return nil
	}
	return fmt.Errorf("%s is missing from worksheet %s", sheetName, workBookName)
}

func main() {
	app := cli.NewCli(datatools.Version)
	appName := app.AppName()

	// Add Help Docs
	app.AddHelp("license", []byte(fmt.Sprintf(datatools.LicenseText, appName, datatools.Version)))
	app.AddHelp("description", []byte(fmt.Sprintf(description, appName)))
	app.AddHelp("examples", []byte(fmt.Sprintf(examples, appName, appName, appName, appName, appName)))

	// Document non-option parameters
	app.AddParams("EXCEL_WORKBOOK_NAME", "[SHEET_NAME]")

	// Standard Options
	app.BoolVar(&showHelp, "h,help", false, "display help")
	app.BoolVar(&showLicense, "l,license", false, "display license")
	app.BoolVar(&showVersion, "v,version", false, "display version")
	app.BoolVar(&showExamples, "examples", false, "display example(s)")
	//app.StringVar(&inputFName, "i,input", "", "input filename")
	app.StringVar(&outputFName, "o,output", "", "output filename")
	app.BoolVar(&quiet, "quiet", false, "suppress error messages")
	app.BoolVar(&generateMarkdownDocs, "generate-markdown-docs", false, "generate markdown documentation")
	app.BoolVar(&newLine, "nl,newline", false, "if true add a trailing newline")

	// App Specific Options
	app.BoolVar(&showSheetCount, "c,count", false, "display number of sheets in Excel Workbook")
	app.BoolVar(&showSheetNames, "N,sheets", false, "display sheet names in Excel Workbook")

	// Parse env and options
	app.Parse()
	args := app.Args()

	// Setup IO
	var err error
	app.Eout = os.Stderr

	/* NOTE: this command does not read from stdin
	   app.In, err = cli.Open(inputFName, os.Stdin)
	   cli.ExitOnError(app.Eout, err, quiet)
	   defer cli.CloseFile(inputFName, app.In)
	*/

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

	if len(args) < 1 {
		cli.ExitOnError(app.Eout, fmt.Errorf("Missing Excel Workbook names"), quiet)
	}

	workBookName := args[0]
	if showSheetCount == true {
		cnt, err := sheetCount(workBookName)
		cli.ExitOnError(app.Eout, err, quiet)
		fmt.Fprintf(app.Out, "%d%s", cnt, eol)
		os.Exit(0)
	}

	if showSheetNames == true {
		names, err := sheetNames(workBookName)
		cli.ExitOnError(app.Eout, err, quiet)
		fmt.Fprintf(app.Out, "%s%s", strings.Join(names, "\n"), eol)
		os.Exit(0)
	}

	if len(args) < 2 {
		cli.ExitOnError(app.Eout, fmt.Errorf("Missing worksheet name"), quiet)
	}
	for _, sheetName := range args[1:] {
		if len(sheetName) > 0 {
			err := xlsx2JSON(app.Out, workBookName, sheetName)
			cli.ExitOnError(app.Eout, err, quiet)
		}
	}
	fmt.Fprintf(app.Out, "%s", eol)
}
