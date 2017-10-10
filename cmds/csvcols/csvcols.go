//
// csvcols - is a command line that takes each argument in order and outputs a line in CSV format.
// It can also take a delimiter and line of text splitting it into a CSV formatted set of columns.
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

	// 3rd Party packages
	"github.com/google/uuid"
)

const (
	//FIXME: maxColumns needs to be calculated from the data rather than being a constant
	maxColumns = 2048
)

var (
	usage = `USAGE: %s [OPTIONS] [ARGS_AS_COL_VALUES]`

	description = `SYNOPSIS

%s converts a set of command line args into columns output in CSV format.
It can also be used CSV input rows and rendering only the column numbers
listed on the commandline (first column is 1 not 0).
`

	examples = `EXAMPLES

Simple usage of building a CSV file one row at a time.

    %s one two three > 3col.csv
    %s 1 2 3 >> 3col.csv
    cat 3col.csv

Example parsing a pipe delimited string into a CSV line

    %s -d "|" "one|two|three" > 3col.csv
    %s -delimiter "|" "1|2|3" >> 3col.csv
    cat 3col.csv

Filter a 10 column CSV file for columns 1,4,6 (left most column is one)

    cat 10col.csv | %s -col 1,4,6 > 3col.csv

Filter a 10 columns CSV file for columns 1,4,6 from file named "10col.csv"

    %s -i 10col.csv -col 1,4,6 > 3col.csv
`

	// Standard Options
	showHelp     bool
	showLicense  bool
	showVersion  bool
	showExamples bool
	inputFName   string
	outputFName  string

	// App Options
	outputColumns string
	prefixUUID    bool
	skipHeaderRow bool
	delimiter     string
)

func selectedColumns(rowNo int, record []string, columnNos []int, prefixUUID bool, skipHeaderRow bool) []string {
	var id string

	result := []string{}
	if prefixUUID == true {
		if rowNo == 0 {
			id = "uuid"

		} else {
			id = uuid.New().String()
		}
		result = append(result, id)
	}
	l := len(record)
	for _, col := range columnNos {
		if col >= 0 && col < l {
			result = append(result, record[col])
		} else {
			// If we don't find the column, story an empty string
			result = append(result, "")
		}
	}
	return result
}

func CSVColumns(in *os.File, out *os.File, columnNos []int, prefixUUID bool, skipHeaderRow bool, delimiter string) {
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
		row := selectedColumns(i, rec, columnNos, prefixUUID, skipHeaderRow)
		if err := w.Write(row); err != nil {
			fmt.Fprintf(os.Stderr, "Error writing record to csv: %s\n", err)
			fmt.Fprintf(os.Stderr, "Row %T %+v\n", row, row)
			os.Exit(1)
		}
	}
	w.Flush()
	if err := w.Error(); err != nil {
		fmt.Fprintf(os.Stderr, "%s\n", err)
		os.Exit(1)
	}
}

func init() {
	// Basic Options
	flag.BoolVar(&showHelp, "h", false, "display help")
	flag.BoolVar(&showHelp, "help", false, "display help")
	flag.BoolVar(&showLicense, "l", false, "display license")
	flag.BoolVar(&showLicense, "license", false, "display license")
	flag.BoolVar(&showVersion, "v", false, "display version")
	flag.BoolVar(&showVersion, "version", false, "display version")
	flag.BoolVar(&showExamples, "example", false, "display example")
	flag.StringVar(&inputFName, "i", "", "input filename")
	flag.StringVar(&inputFName, "input", "", "input filename")
	flag.StringVar(&outputFName, "o", "", "output filename")
	flag.StringVar(&outputFName, "output", "", "output filename")

	// App Options
	flag.StringVar(&outputColumns, "col", "", "output specified columns (e.g. -col 1,12:14,2,4))")
	flag.StringVar(&outputColumns, "cols", "", "output specified columns (e.g. -col 1,12:14,2,4))")
	flag.StringVar(&delimiter, "d", "", "set delimiter character")
	flag.StringVar(&delimiter, "delimiter", "", "set delimiter character")
	flag.BoolVar(&skipHeaderRow, "skip-header-row", true, "skip the header row")
	flag.BoolVar(&prefixUUID, "uuid", false, "add a prefix row with generated UUID cell")
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

	if outputColumns != "" {
		columnNos, err := datatools.ParseRange(outputColumns, maxColumns)
		if err != nil {
			fmt.Fprintf(os.Stderr, "%s\n", err)
			os.Exit(1)
		}
		// NOTE: We need to adjust from humans counting from 1 to counting from zero
		for i := 0; i < len(columnNos); i++ {
			columnNos[i] = columnNos[i] - 1
			if columnNos[i] < 0 {
				columnNos[i] = 0
			}
		}
		CSVColumns(in, out, columnNos, prefixUUID, skipHeaderRow, delimiter)
		os.Exit(0)
	}

	if len(delimiter) > 0 && len(args) == 1 {
		args = strings.Split(args[0], datatools.NormalizeDelimiter(delimiter))
	}

	// Clean up cells removing outer quotes if necessary
	cells := []string{}
	for _, val := range args {
		if strings.HasPrefix(val, "\"") && strings.HasSuffix(val, "\"") {
			val = strings.TrimPrefix(strings.TrimSuffix(val, "\""), "\"")
		}
		cells = append(cells, strings.TrimSpace(val))
	}

	w := csv.NewWriter(out)
	if delimiter != "" {
		w.Comma = datatools.NormalizeDelimiterRune(delimiter)
	}
	if err := w.Write(cells); err != nil {
		log.Fatalf("error writing args as csv, %s", err)
	}
	w.Flush()
	if err := w.Error(); err != nil {
		log.Fatal(err)
	}
}
