//
// csvrows - is can filter selected rows, out row ranges or turn each command
// line parameter into a CSV row of output.
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
	"flag"
	"fmt"
	"io"
	"log"
	"os"
	"path"
	"strings"

	// Caltech Library packages
	"github.com/caltechlibrary/cli"
	"github.com/caltechlibrary/datatools"
)

const (
	//FIXME: maxRows should be calculated from the data rather than be a constant.
	maxRows = 1000000
)

var (
	usage = `USAGE: %s [OPTIONS] [ARGS_AS_ROW_VALUES]`

	description = `

SYNOPSIS

%s converts a set of command line args into rows of CSV formated output.
It can also be used to filter or list specific rows of CSV input
The first row is 1 not 0. Often row 1 is the header row and %s makes it
easy to output only the data rows.

`

	examples = `

EXAMPLES

Simple usage of building a CSV file one rows at a time.

    %s "First,Second,Third" "one,two,three" > 4rows.csv
    %s "ein,zwei,drei" "1,2,3" >> 4rows.csv
    cat 4row.csv

Example parsing a pipe delimited string into a CSV line

    %s -d "|" "First,Second,Third|one,two,three" > 4rows.csv
    %s -delimiter "|" "ein,zwei,drei|1,2,3" >> 4rows.csv
    cat 4rows.csv

Filter a 10 row CSV file for rows 1,4,6 (top most row is one)

    cat 10row.csv | %s -row 1,4,6 > 3rows.csv

Filter a 10 row CSV file for rows 1,4,6 from file named "10row.csv"

    %s -i 10row.csv -row 1,4,6 > 3rows.csv

`

	// Standard options
	showHelp     bool
	showLicense  bool
	showVersion  bool
	showExamples bool
	inputFName   string
	outputFName  string

	// Application specific options
	validateRows  bool
	showRowCount  bool
	showColCount  bool
	showHeader    bool
	skipHeaderRow bool
	outputRows    string
	delimiter     string
)

func selectedRow(rowNo int, record []string, rowNos []int) []string {
	if len(rowNos) == 0 {
		return record
	}
	for _, i := range rowNos {
		if i == rowNo {
			return record
		}
	}
	return nil
}

func CSVRows(in *os.File, out *os.File, rowNos []int, delimiter string) {
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
		if err != nil {
			fmt.Fprintf(os.Stderr, "%s, %s\n", inputFName, err)
			fmt.Fprintf(os.Stderr, "%T %+v\n", rec, rec)
		}
		row := selectedRow(i, rec, rowNos)
		if row != nil {
			if err := w.Write(row); err != nil {
				fmt.Fprintf(os.Stderr, "Error writing record to csv: %s\n", err)
				fmt.Fprintf(os.Stderr, "Row %T %+v\n", row, row)
				os.Exit(1)
			}
		}
	}
	w.Flush()
	if err := w.Error(); err != nil {
		fmt.Fprintf(os.Stderr, "%s\n", err)
		os.Exit(1)
	}
}

func init() {
	// Standard options
	flag.BoolVar(&showHelp, "h", false, "display help")
	flag.BoolVar(&showHelp, "help", false, "display help")
	flag.BoolVar(&showLicense, "l", false, "display license")
	flag.BoolVar(&showLicense, "license", false, "display license")
	flag.BoolVar(&showVersion, "v", false, "display version")
	flag.BoolVar(&showVersion, "version", false, "display version")
	flag.BoolVar(&showExamples, "example", false, "display example(s)")
	flag.StringVar(&inputFName, "i", "", "input filename")
	flag.StringVar(&inputFName, "input", "", "input filename")
	flag.StringVar(&outputFName, "o", "", "output filename")
	flag.StringVar(&outputFName, "output", "", "output filename")

	// Application specific options
	flag.StringVar(&delimiter, "d", "", "set delimiter character")
	flag.StringVar(&delimiter, "delimiter", "", "set delimiter character")
	flag.StringVar(&outputRows, "row", "", "output specified rows in order (e.g. -row 1,5,2:4))")
	flag.StringVar(&outputRows, "rows", "", "output specified rows in order (e.g. -rows 1,5,2:4))")
	flag.BoolVar(&skipHeaderRow, "skip-header-row", false, "skip the header row (alias for -row 2:")
	flag.BoolVar(&showHeader, "header", false, "display the header row (alias for '-rows 1')")
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
	cfg.OptionText = "OPTIONS"
	cfg.ExampleText = fmt.Sprintf(examples, appName, appName, appName, appName, appName, appName)

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

	in, err := cli.Open(inputFName, os.Stdin)
	if err != nil {
		fmt.Fprintf(os.Stderr, "%s\n", err)
		os.Exit(1)
	}
	defer cli.CloseFile(inputFName, in)

	out, err := cli.Create(outputFName, os.Stdout)
	if err != nil {
		fmt.Fprintf(os.Stderr, "%s\n", err)
		os.Exit(1)
	}
	defer cli.CloseFile(outputFName, out)

	if showHeader == true {
		outputRows = "1"
	}
	if len(args) == 0 && outputRows == "" {
		outputRows = "1:"
		if skipHeaderRow == true {
			outputRows = "2:"
		}
	}

	if outputRows != "" {
		rowNos, err := datatools.ParseRange(outputRows, maxRows)
		if err != nil {
			fmt.Fprintf(os.Stderr, "%s\n", err)
			os.Exit(1)
		}
		// NOTE: We need to adjust from humans counting from 1 to counting from zero
		for i := 0; i < len(rowNos); i++ {
			rowNos[i] = rowNos[i] - 1
			if rowNos[i] < 0 {
				rowNos[i] = 0
			}
		}
		CSVRows(in, out, rowNos, delimiter)
		os.Exit(0)
	}

	if len(delimiter) > 0 && len(args) == 1 {
		args = strings.Split(args[0], datatools.NormalizeDelimiter(delimiter))
	}

	// Clean up cells removing outer quotes if necessary
	w := csv.NewWriter(out)
	if delimiter != "" {
		w.Comma = datatools.NormalizeDelimiterRune(delimiter)
	}
	for _, val := range args {
		r := csv.NewReader(strings.NewReader(val))
		if delimiter != "" {
			r.Comma = datatools.NormalizeDelimiterRune(delimiter)
		}
		record, err := r.Read()
		if err != nil {
			log.Fatalf("%q isn't a CSV row, %s", val, err)
		}
		r = nil
		if err := w.Write(record); err != nil {
			log.Fatalf("error writing args as csv, %s", err)
		}
	}
	w.Flush()
	if err := w.Error(); err != nil {
		log.Fatal(err)
	}
}
