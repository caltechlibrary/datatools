//
// xlsx2json.go is a command line utility that converts an Excel
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
	"flag"
	"fmt"
	"os"
	"path"
	"strings"

	// CaltechLibrary packages
	"github.com/caltechlibrary/cli"
	"github.com/caltechlibrary/datatools"

	// 3rd Party packages
	"github.com/tealeg/xlsx"
)

var (
	usage = `USAGE: %s [OPTIONS] EXCEL_WORKBOOK_NAME [SHEET_NAME]`

	description = `

SYNOPSIS

%s is a tool that converts individual Excel Workbook Sheets into
JSON output.

`

	examples = `

EXAMPLES

This would get the sheet named "Sheet 1" from "my-workbook.xlsx" and save as sheet1.json

    %s my-workbook.xlsx "Sheet 1" > sheet1.json

This would get the number of sheets in the workbook

    %s -c my-workbook.xlsx

This will output the title of the sheets in the workbook

    %s -n my-workbook.xlsx

Putting it all together in a shell script and convert all sheets to
into JSON documents..

	 %s -n my-workbook.xlsx | while read SHEET_NAME; do
       %s my-workbook.xlsx "$SHEET_NAME" > \
	       "${SHEET_NAME// /-}.json"
    done

`

	// Standard Options
	showHelp     bool
	showLicense  bool
	showVersion  bool
	showExamples bool
	outputFName  string
	quiet        bool

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

func xlsx2JSON(out *os.File, workBookName, sheetName string) error {
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

func init() {
	// Standard Options
	flag.BoolVar(&showHelp, "h", false, "display help")
	flag.BoolVar(&showHelp, "help", false, "display help")
	flag.BoolVar(&showLicense, "l", false, "display license")
	flag.BoolVar(&showLicense, "license", false, "display license")
	flag.BoolVar(&showVersion, "v", false, "display version")
	flag.BoolVar(&showVersion, "version", false, "display version")
	flag.BoolVar(&showExamples, "example", false, "display example(s)")
	flag.StringVar(&outputFName, "o", "", "output filename")
	flag.StringVar(&outputFName, "output", "", "output filename")

	// App Specific Options
	flag.BoolVar(&showSheetCount, "c", false, "display number of sheets in Excel Workbook")
	flag.BoolVar(&showSheetNames, "n", false, "display sheet names in Excel Workbook")
}

func main() {
	appName := path.Base(os.Args[0])
	flag.Parse()
	args := flag.Args()

	// Configuration and command line interation
	cfg := cli.New(appName, strings.ToUpper(appName), datatools.Version)
	cfg.LicenseText = fmt.Sprintf(datatools.LicenseText, appName, datatools.Version)
	cfg.UsageText = fmt.Sprintf(usage, appName)
	cfg.DescriptionText = fmt.Sprintf(description, appName)
	cfg.OptionText = "OPTIONS\n\n"
	cfg.ExampleText = fmt.Sprintf(examples, appName, appName, appName, appName, appName)

	if showHelp == true {
		if len(args) > 0 {
			fmt.Println(cfg.Help(args...))
		} else {
			fmt.Println(cfg.Usage())
		}
		os.Exit(0)
	}

	if showExamples == true {
		if len(args) > 0 {
			fmt.Println(cfg.Example(args...))
		} else {
			fmt.Println(cfg.ExampleText)
		}
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

	out, err := cli.Create(outputFName, os.Stdout)
	cli.ExitOnError(os.Stderr, err, quiet)
	defer cli.CloseFile(outputFName, out)

	if len(args) < 1 {
		cli.ExitOnError(os.Stderr, fmt.Errorf("Missing Excel Workbook names"), quiet)
	}

	workBookName := args[0]
	if showSheetCount == true {
		cnt, err := sheetCount(workBookName)
		cli.ExitOnError(os.Stderr, err, quiet)
		fmt.Fprintf(out, "%d", cnt)
		os.Exit(0)
	}

	if showSheetNames == true {
		names, err := sheetNames(workBookName)
		cli.ExitOnError(os.Stderr, err, quiet)
		fmt.Fprintln(out, strings.Join(names, "\n"))
		os.Exit(0)
	}

	if len(args) < 2 {
		cli.ExitOnError(os.Stderr, fmt.Errorf("Missing worksheet name"), quiet)
	}
	for _, sheetName := range args[1:] {
		if len(sheetName) > 0 {
			err := xlsx2JSON(out, workBookName, sheetName)
			cli.ExitOnError(os.Stderr, err, quiet)
		}
	}
}
