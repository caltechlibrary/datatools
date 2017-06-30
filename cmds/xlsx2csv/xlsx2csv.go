//
// xlsx2csv.go is a command line utility that converts individual
// Excel Workbook Sheets to CSV.
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
	"encoding/csv"
	"flag"
	"fmt"
	"io"
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

%s is a tool that converts individual Excel Sheets to CSV output.
`

	examples = `
EXAMPLE

    %s my-workbook.xlsx "Sheet 1" > sheet1.csv

This would get the first sheet from the workbook and save it as a CSV file.

    %s -c my-workbook.xlsx

This will output the number of sheets in the Workbook.

    %s -n my-workbook.xlsx

This will display a list of sheet names, one per line.
Putting it all together in a shell script.

    for SHEET_NAME in $(%s -n my-workbook.xlsx); do
       %s my-workbook.xlsx "$SHEET_NAME" > \
	       "${SHEET_NAME// /-}.csv"
    done
`

	// Standard Options
	showHelp    bool
	showLicense bool
	showVersion bool
	outputFName string

	// Application Options
	showSheetCount bool
	showSheetNames bool
	verbose        bool
	permissive     bool
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
	for sheetName, _ := range xlFile.Sheet {
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
		for i, row := range sheet.Rows {
			//FIXME: I would be nice to optionally only the columns you wanted to output from the sheet...
			cells = []string{}
			for j, cell := range row.Cells {
				val, err := cell.String()
				if err != nil {
					if permissive == true {
						cells = append(cells, fmt.Sprintf("%s", val))
					} else {
						cells = append(cells, fmt.Sprintf("%s", err))
					}
					if verbose == true {
						fmt.Fprintf(os.Stderr, "row %d, col %d %s\n", i, j, err)
					}
				} else {
					cells = append(cells, val)
				}
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

func init() {
	// Standard Options
	flag.BoolVar(&showHelp, "h", false, "display help")
	flag.BoolVar(&showHelp, "help", false, "display help")
	flag.BoolVar(&showLicense, "l", false, "display license")
	flag.BoolVar(&showLicense, "license", false, "display license")
	flag.BoolVar(&showVersion, "v", false, "display version")
	flag.BoolVar(&showVersion, "version", false, "display version")
	flag.StringVar(&outputFName, "o", "", "output filename")
	flag.StringVar(&outputFName, "output", "", "output filename")

	// App Specific Options
	flag.BoolVar(&showSheetCount, "c", false, "display number of sheets in Excel Workbook")
	flag.BoolVar(&showSheetNames, "n", false, "display sheet names in Excel W9rkbook")
	flag.BoolVar(&verbose, "verbose", false, "output cell level errors")
	flag.BoolVar(&verbose, "permissive", false, "ignore cell level errors")
}

func main() {
	appName := path.Base(os.Args[0])
	flag.Parse()
	args := flag.Args()

	// Configuration and command line interation
	cfg := cli.New(appName, appName, fmt.Sprintf(datatools.LicenseText, appName, datatools.Version), datatools.Version)
	cfg.UsageText = fmt.Sprintf(usage, appName)
	cfg.DescriptionText = fmt.Sprintf(description, appName)
	cfg.ExampleText = fmt.Sprintf(examples, appName, appName, appName, appName, appName)

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

	out, err := cli.Create(outputFName, os.Stdout)
	if err != nil {
		fmt.Fprintf(os.Stderr, "%s\n", err)
		os.Exit(1)
	}
	defer cli.CloseFile(outputFName, out)

	if len(args) < 1 {
		fmt.Println(cfg.Usage())
		fmt.Fprintln(os.Stderr, "Missing Excel Workbook names")
		os.Exit(1)
	}

	workBookName := args[0]
	if showSheetCount == true {
		if cnt, err := sheetCount(workBookName); err == nil {
			fmt.Printf("%d", cnt)
			os.Exit(0)
		} else {
			fmt.Fprintf(os.Stderr, "%s, %s\n", workBookName, err)
			os.Exit(1)
		}
	}

	if showSheetNames == true {
		if names, err := sheetNames(workBookName); err == nil {
			fmt.Println(strings.Join(names, "\n"))
			os.Exit(0)
		} else {
			fmt.Fprintf(os.Stderr, "%s, %s\n", workBookName, err)
			os.Exit(1)
		}
	}

	if len(args) < 2 {
		fmt.Fprintln(os.Stderr, "Missing worksheet name")
		os.Exit(1)
	}
	for _, sheetName := range args[1:] {
		if len(sheetName) > 0 {
			xlsx2CSV(out, workBookName, sheetName)
		}
	}
}
