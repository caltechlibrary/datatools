//
// xlsx2csv is a command line utility that converts individual
// Excel Workbook Sheets to CSV.
//
// @Author R. S. Doiel
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

	// CaltechLibrary packages
	"github.com/caltechlibrary/cli"
	"github.com/caltechlibrary/datatools"

	// 3rd Party packages
	"github.com/tealeg/xlsx"
)

var (
	description = `
%s is a tool that converts individual Excel Sheets to CSV output.
`

	examples = `
Extract a workbook sheet as a CSV file

    %s MyWorkbook.xlsx "My worksheet 1" > sheet1.csv

This would get the first sheet from the workbook and save it as a CSV file.

    %s -count MyWorkbook.xlsx

This will output the number of sheets in the Workbook.

    %s -sheets MyWorkbook.xlsx

This will display a list of sheet names, one per line.
Putting it all together in a shell script.

	%s -N MyWorkbook.xlsx | while read SHEET_NAME; do
    	CSV_NAME="${SHEET_NAME// /-}.csv"
    	%s -o "${CSV_NAME}" MyWorkbook.xlsx "${SHEET_NAME}" 
	done
`

	// Standard Options
	showHelp     bool
	showLicense  bool
	showVersion  bool
	showExamples bool
	//inputFName           string
	outputFName      string
	generateMarkdown bool
	quiet            bool
	newLine          bool
	eol              string

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

func xlsx2CSV(out io.Writer, workBookName, sheetName string) error {
	xlFile, err := xlsx.OpenFile(workBookName)
	if err != nil {
		return err
	}
	results := [][]string{}
	cells := []string{}
	if sheet, ok := xlFile.Sheet[sheetName]; ok == true {
		for _, row := range sheet.Rows {
			//FIXME: I would be nice to optionally only the columns you wanted to output from the sheet...
			cells = []string{}
			for _, cell := range row.Cells {
				val := cell.String()
				cells = append(cells, val)
			}
			results = append(results, cells)
		}
		w := csv.NewWriter(out)
		for _, record := range results {
			if err := w.Write(record); err != nil {
				return fmt.Errorf("error writing record to csv: %s", err)
			}
		}
		w.Flush()
		if err := w.Error(); err != nil {
			return err
		}
	}
	return fmt.Errorf("%s in worksheet %s", sheetName, workBookName)
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
	app.BoolVar(&generateMarkdown, "generate-markdown", false, "generate markdown documentation")
	app.BoolVar(&newLine, "nl,newline", false, "if true add a trailing newline")

	// App Specific Options
	app.BoolVar(&showSheetCount, "c,count", false, "display number of Workbook sheets")
	app.BoolVar(&showSheetNames, "N,sheets", false, "display the Workbook sheet names")

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
	if generateMarkdown {
		app.GenerateMarkdown(app.Out)
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
			xlsx2CSV(app.Out, workBookName, sheetName)
		}
	}
	fmt.Fprintf(app.Out, "%s", eol)
}
