//
// xlsx2json is a command line utility that converts an Excel
// Workboom Sheet into JSON.
//
// @Author R. S. Doiel
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
	"encoding/json"
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
	helpText =  `---
title: "{app_name} (1) user manual"
author: "R. S. Doiel"
pubDate: 2023-01-09
---

# NAME

{app_name}

# SYNOPSIS

{app_name} [OPTIONS] EXCEL_WORKBOOK_NAME [SHEET_NAME]

# DESCRIPTION

{app_name} is a tool that converts individual Excel Workbook Sheets into
JSON output.

# OPTIONS

-help
: display help

-license
: display license

-version
: display version

-N, -sheets
: display sheet names in Excel Workbook

-c, -count
: display number of sheets in Excel Workbook

-nl, -newline
: if true add a trailing newline

-o, -output
: output filename

-quiet
: suppress error messages


# EXAMPLES

This would get the sheet named "Sheet 1" from "MyWorkbook.xlsx" and save as sheet1.json

~~~
    {app_name} MyWorkbook.xlsx "My worksheet 1" > sheet1.json
~~~

This would get the number of sheets in the workbook

~~~
    {app_name} -count MyWorkbook.xlsx
~~~

This will output the title of the sheets in the workbook

~~~
    {app_name} -sheets MyWorkbook.xlsx
~~~

Putting it all together in a shell script and convert all sheets to
into JSON documents..

~~~
	{app_name} -N MyWorkbook.xlsx | while read SHEET_NAME; do
    	JSON_NAME="${SHEET_NAME// /-}.json"
    	{app_name} -o "${JSON_NAME}" MyWorkbook.xlsx "$SHEET_NAME"
	done
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
	newLine          bool
	eol              string

	// Application Options
	showSheetCount bool
	showSheetNames bool
)

func fmtTxt(src string, appName string, version string) string {
	return strings.ReplaceAll(strings.ReplaceAll(src, "{app_name}", appName), "{version}", version)
}

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
	appName := path.Base(os.Args[0])

	// Standard Options
	flag.BoolVar(&showHelp, "help", false, "display help")
	flag.BoolVar(&showLicense, "license", false, "display license")
	flag.BoolVar(&showVersion, "version", false, "display version")

	flag.StringVar(&outputFName, "o", "", "output filename")
	flag.StringVar(&outputFName, "output", "", "output filename")
	flag.BoolVar(&quiet, "quiet", false, "suppress error messages")
	

	flag.BoolVar(&newLine, "nl", false, "if true add a trailing newline")
	flag.BoolVar(&newLine, "newline", false, "if true add a trailing newline")

	// App Specific Options
	flag.BoolVar(&showSheetCount, "c", false, "display number of sheets in Excel Workbook")
	flag.BoolVar(&showSheetCount, "count", false, "display number of sheets in Excel Workbook")
	flag.BoolVar(&showSheetNames, "N", false, "display sheet names in Excel Workbook")
	flag.BoolVar(&showSheetNames, "sheets", false, "display sheet names in Excel Workbook")

	// Parse env and options
	flag.Parse()
	args := flag.Args()

	// Setup IO
	var err error

	out := os.Stdout
	eout := os.Stderr

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
		fmt.Fprintf(out, "%s\n", datatools.LicenseText)
		os.Exit(0)
	}
	if showVersion {
		fmt.Fprintf(out, "%s %s\n", appName, datatools.Version)
		os.Exit(0)
	}
	if newLine {
		eol = "\n"
	}

	if len(args) < 1 {
		fmt.Fprintln(eout, "Missing Excel Workbook names")
		os.Exit(1)
	}

	workBookName := args[0]
	if showSheetCount == true {
		cnt, err := sheetCount(workBookName)
		if err != nil {
			fmt.Fprintln(eout, err)
			os.Exit(1)
		}
		fmt.Fprintf(out, "%d%s", cnt, eol)
		os.Exit(0)
	}

	if showSheetNames == true {
		names, err := sheetNames(workBookName)
		if err != nil {
			fmt.Fprintln(eout, err)
			os.Exit(1)
		}
		fmt.Fprintf(out, "%s%s", strings.Join(names, "\n"), eol)
		os.Exit(0)
	}

	if len(args) < 2 {
		fmt.Fprintln(eout, "Missing worksheet name")
		os.Exit(1)
	}
	for _, sheetName := range args[1:] {
		if len(sheetName) > 0 {
			err := xlsx2JSON(out, workBookName, sheetName)
			if err != nil {
				fmt.Fprintln(eout, err)
				os.Exit(1)
			}
		}
	}
	fmt.Fprintf(out, "%s", eol)
}
