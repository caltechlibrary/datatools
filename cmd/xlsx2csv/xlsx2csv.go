// xlsx2csv is a command line utility that converts individual
// Excel Workbook Sheets to CSV.
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
package main

import (
	"encoding/csv"
	"flag"
	"fmt"
	"io"
	"os"
	"path"
	"runtime"
	"strings"

	// CaltechLibrary packages
	"github.com/caltechlibrary/datatools"

	// 3rd Party packages
	"github.com/tealeg/xlsx"
)

var (
	helpText = `%{app_name}(1) user manual | version {version} {release_hash}
% R. S. Doiel
% {release_date}

# NAME

{app_name}

# SYNOPSIS

{app_name} [OPTIONS] EXCEL_WORKBOOK_NAME [SHEET_NAME]

# DESCRIPTION

{app_name} is a tool that converts individual Excel Sheets to CSV output.

# OPTIONS

-help
: display help

-license
: display license

-version
: display version

-N, -sheets
: display the Workbook sheet names

-c, -count
: display number of Workbook sheets

-nl, -newline
: add a trailing newline to the end of file (EOF)

-crlf
: use CRLF for end of line (EOL). On Windows this option is the default.
Set to false to use a LF on Windows. On other OS LF is the default.

-o, -output
: output filename

-quiet
: suppress error messages


# EXAMPLES

Extract a workbook sheet as a CSV file

~~~
    {app_name} MyWorkbook.xlsx "My worksheet 1" > sheet1.csv
~~~

This would get the first sheet from the workbook and save it as a CSV file.

~~~
    {app_name} -count MyWorkbook.xlsx
~~~


This will output the number of sheets in the Workbook.

~~~
    {app_name} -sheets MyWorkbook.xlsx
~~~

This will display a list of sheet names, one per line.
Putting it all together in a shell script.

~~~
	{app_name} -N MyWorkbook.xlsx | while read SHEET_NAME; do
    	CSV_NAME="${SHEET_NAME// /-}.csv"
    	{app_name} -o "${CSV_NAME}" MyWorkbook.xlsx "${SHEET_NAME}" 
	done
~~~

{app_name} {version}

`

	// Standard Options
	showHelp     bool
	showLicense  bool
	showVersion  bool
	showExamples bool

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

func xlsx2CSV(out io.Writer, workBookName, sheetName string, useCRLF bool) error {
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
		w.UseCRLF = useCRLF
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
	appName := path.Base(os.Args[0])
	version := datatools.Version
	license := datatools.LicenseText
	releaseDate := datatools.ReleaseDate
	releaseHash := datatools.ReleaseHash
	useCRLF := (runtime.GOOS == "windows")

	// Standard Options
	flag.BoolVar(&showHelp, "help", false, "display help")
	flag.BoolVar(&showLicense, "license", false, "display license")
	flag.BoolVar(&showVersion, "version", false, "display version")

	flag.StringVar(&outputFName, "o", "", "output filename")
	flag.StringVar(&outputFName, "output", "", "output filename")
	flag.BoolVar(&quiet, "quiet", false, "suppress error messages")

	flag.BoolVar(&newLine, "nl", false, "add a trailing newline to end of file (EOF)")
	flag.BoolVar(&newLine, "newline", false, "add a trailing newline")

	flag.BoolVar(&useCRLF, "crlf", useCRLF, "use CRLF for end of line (EOL) to end of file (EOF)")

	// App Specific Options
	flag.BoolVar(&showSheetCount, "c", false, "display number of Workbook sheets")
	flag.BoolVar(&showSheetCount, "count", false, "display number of Workbook sheets")
	flag.BoolVar(&showSheetNames, "N", false, "display the Workbook sheet names")
	flag.BoolVar(&showSheetNames, "sheets", false, "display the Workbook sheet names")

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
		fmt.Fprintf(out, "%s\n", datatools.FmtHelp(helpText, appName, version, releaseDate, releaseHash))
		os.Exit(0)
	}
	if showLicense {
		fmt.Fprintf(out, "%s\n", license)
		os.Exit(0)
	}
	if showVersion {
		fmt.Fprintf(out, "datatools, %s %s %s\n", appName, version, releaseHash)
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
		if err != nil {
			fmt.Fprintln(eout, err)
			os.Exit(1)
		}
	}
	for _, sheetName := range args[1:] {
		if len(sheetName) > 0 {
			xlsx2CSV(out, workBookName, sheetName, useCRLF)
		}
	}
	fmt.Fprintf(out, "%s", eol)
}
