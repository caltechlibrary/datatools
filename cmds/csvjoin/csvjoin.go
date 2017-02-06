//
// csvjoin - is a command line that takes two CSV files and joins them by match a designated column in each.
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
	"bytes"
	"encoding/csv"
	"flag"
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"os"
	"path"

	// My packages
	"github.com/caltechlibrary/cli"
	"github.com/caltechlibrary/datatools"
)

var (
	usage = `USAGE: %s [OPTIONS] CSV1 CSV2 COL1 COL2`

	description = `
SYNOPSIS

%s outputs CSV content based on two CSV files with matching column values.
Each CSV input file has a designated column to match on. The values are
compared as strings.
`

	examples = `
EXAMPLES

Simple usage of building a merged CSV file from data1.csv
and data2.csv where column 1 in data1.csv matches the value in
column 3 of data2.csv with the results being written to 
merged-data.csv..

    %s -csv1=data1.csv -col1=1 \
       -csv2=data2.csv -col2=3 \
	   -output=merged-data.csv
`

	// Standard Options
	showHelp    bool
	showLicense bool
	showVersion bool
	outputFName string

	// App Options
	csv1FName string
	csv2FName string
	col1      int
	col2      int
)

func scanTable(table [][]string, col2 int, val string) ([]string, bool) {
	for _, row := range table {
		if col2 < len(row) && row[col2] == val {
			return row, true
		}
	}
	return []string{}, false
}

func init() {
	// Basic Options
	flag.BoolVar(&showHelp, "h", false, "display help")
	flag.BoolVar(&showHelp, "help", false, "display help")
	flag.BoolVar(&showLicense, "l", false, "display license")
	flag.BoolVar(&showLicense, "license", false, "display license")
	flag.BoolVar(&showVersion, "v", false, "display version")
	flag.BoolVar(&showVersion, "version", false, "display version")
	flag.StringVar(&outputFName, "o", "", "output filename")
	flag.StringVar(&outputFName, "output", "", "output filename")

	// App Options
	flag.StringVar(&csv1FName, "csv1", "", "first CSV filename")
	flag.StringVar(&csv2FName, "csv2", "", "second CSV filename")
	flag.IntVar(&col1, "col1", 0, "column to on join on in first CSV file")
	flag.IntVar(&col2, "col2", 0, "column to on join on in second CSV file")
}

func main() {
	appName := path.Base(os.Args[0])
	flag.Parse()

	// Configuration and command line interation
	cfg := cli.New(appName, appName, fmt.Sprintf(datatools.LicenseText, appName, datatools.Version), datatools.Version)
	cfg.UsageText = fmt.Sprintf(usage, appName)
	cfg.DescriptionText = fmt.Sprintf(description, appName)
	cfg.ExampleText = fmt.Sprintf(examples, appName)

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

	// NOTE: we don't setup inputFName as we need at least two inputs to process the join.
	out, err := cli.Create(outputFName, os.Stdout)
	if err != nil {
		fmt.Fprintf(os.Stderr, "%s\n", err)
		os.Exit(1)
	}
	defer cli.CloseFile(outputFName, out)

	if len(csv1FName) == 0 {
		fmt.Fprintln(os.Stderr, "Missing first CSV filename")
		os.Exit(1)
	}

	if len(csv2FName) == 0 {
		fmt.Fprintln(os.Stderr, "Missing second CSV filename")
		os.Exit(1)
	}

	if col1 < 0 {
		fmt.Fprintf(os.Stderr, "Cannot use a negative column index %d\n", col1)
		os.Exit(1)
	}
	if col2 < 0 {
		fmt.Fprintf(os.Stderr, "Cannot use a negative column index %d\n", col2)
		os.Exit(1)
	}

	// Read in CSV1 and CSV2 then iterate over CSV1 output rows that have
	// matching column's value
	src1, err := ioutil.ReadFile(csv1FName)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Can't read %s, %s\n", csv1FName, err)
		os.Exit(1)
	}
	src2, err := ioutil.ReadFile(csv2FName)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Can't read %s, %s\n", csv2FName, err)
		os.Exit(1)
	}
	csv1 := csv.NewReader(bytes.NewReader(src1))
	csv1Table := [][]string{}
	for {
		record, err := csv1.Read()
		if err == io.EOF {
			break
		}
		if err != nil {
			fmt.Fprintf(os.Stderr, "%s, %s\n", csv1FName, err)
			fmt.Fprintf(os.Stderr, "%T %+v\n", record, record)
		}
		csv1Table = append(csv1Table, record)
	}
	csv2 := csv.NewReader(bytes.NewReader(src2))
	csv2Table := [][]string{}
	for {
		record, err := csv2.Read()
		if err == io.EOF {
			break
		}
		if err != nil {
			fmt.Fprintf(os.Stderr, "%s, %s\n", csv2FName, err)
			fmt.Fprintf(os.Stderr, "%T %+v\n", record, record)
		}
		csv2Table = append(csv2Table, record)
	}

	w := csv.NewWriter(out)
	val := ""
	for _, rowA := range csv1Table {
		if col1 < len(rowA) {
			val = rowA[col1]
			// Name see if we find matching row in table 2
			if rowB, ok := scanTable(csv2Table, col2, val); ok == true {
				// We have
				combinedRows := append(rowA, rowB...)
				if err := w.Write(combinedRows); err != nil {
					log.Fatalf("error wrint args as csv, %s", err)
				}
			}
		}
	}
	w.Flush()
	if err := w.Error(); err != nil {
		log.Fatal(err)
	}
}
