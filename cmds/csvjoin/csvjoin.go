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
	"encoding/csv"
	"flag"
	"fmt"
	"io"
	"os"
	"path"
	"strings"

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
	verbose         bool
	csv1FName       string
	csv2FName       string
	col1            int
	col2            int
	trimSpaces      bool
	caseSensitive   bool
	useContains     bool
	useLevenshtein  bool
	insertCost      int
	deleteCost      int
	substituteCost  int
	maxEditDistance int
	stopWordsOption string
	allowDuplicates bool
)

// cellsMatch checks if two cells' values match
func cellsMatch(val1, val2 string, stopWords []string) bool {
	if trimSpaces == true {
		val2 = strings.TrimSpace(val2)
	}
	if caseSensitive == false {
		val2 = strings.ToLower(val2)
	}
	if len(stopWords) > 0 {
		val2 = strings.Join(datatools.ApplyStopWords(strings.Split(val2, " "), stopWords), " ")
	}
	switch {
	case useLevenshtein == true:
		distance := datatools.Levenshtein(val2, val1, insertCost, deleteCost, substituteCost, caseSensitive)
		if distance <= maxEditDistance {
			return true
		}
	case useContains == true:
		if strings.Contains(val2, val1) {
			return true
		}
	default:
		if val1 == val2 {
			return true
		}
	}
	return false
}

func scanTable(w *csv.Writer, rowA []string, col1 int, table [][]string, col2 int, stopWords []string) {
	if col1 >= len(rowA) {
		return
	}
	val1 := rowA[col1]
	if trimSpaces == true {
		val1 = strings.TrimSpace(val1)
	}
	if caseSensitive == false {
		val1 = strings.ToLower(val1)
	}
	if len(stopWords) > 0 {
		val1 = strings.Join(datatools.ApplyStopWords(strings.Split(val1, " "), stopWords), " ")
	}
	for i, rowB := range table {
		// Emit a joined row if we have a match
		if col2 < len(rowB) {
			val2 := rowB[col2]
			if cellsMatch(val1, val2, stopWords) == true {
				// We have a match, join the two rows and output
				combinedRows := append(rowA, rowB...)
				if err := w.Write(combinedRows); err != nil {
					fmt.Fprintf(os.Stderr, "Can't write csv row line %d of table 2, %s\n", i, err)
					return
				}
				if allowDuplicates == false {
					return
				}
			}
		}
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
	flag.StringVar(&outputFName, "o", "", "output filename")
	flag.StringVar(&outputFName, "output", "", "output filename")

	// App Options
	flag.BoolVar(&verbose, "verbose", false, "output processing count to stderr")
	flag.StringVar(&csv1FName, "csv1", "", "first CSV filename")
	flag.StringVar(&csv2FName, "csv2", "", "second CSV filename")
	flag.IntVar(&col1, "col1", 0, "column to on join on in first CSV file")
	flag.IntVar(&col2, "col2", 0, "column to on join on in second CSV file")
	flag.BoolVar(&caseSensitive, "case-sensitive", false, "make a case sensitive match (default is case insensitive)")
	flag.BoolVar(&useContains, "contains", false, "match columns based on csv1/col1 contained in csv2/col2")
	flag.BoolVar(&useLevenshtein, "levenshtein", false, "match columns using Levensthein edit distance")
	flag.IntVar(&insertCost, "insert-cost", 1, "insertion cost to use when calculating Levenshtein edit distance")
	flag.IntVar(&deleteCost, "delete-cost", 1, "deletion cost to use when calculating Levenshtein edit distance")
	flag.IntVar(&substituteCost, "substitute-cost", 1, "substitution cost to use when calculating Levenshtein edit distance")
	flag.IntVar(&maxEditDistance, "max-edit-distance", 5, "maximum edit distance for match using Levenshtein distance")
	flag.StringVar(&stopWordsOption, "stop-words", "", "a column delimited list of stop words to ingnore when matching")
	flag.BoolVar(&allowDuplicates, "allow-duplicates", true, "allow duplicates when searching for matches")
	flag.BoolVar(&trimSpaces, "trim-spaces", false, "trim spaces around cell values before comparing")
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

	// FIXME: Should only read the smaller of two files into memory
	// then interate through the other file for matches. This would let you work with larger files.

	// Read in CSV2 to memory then iterate over CSV1 output rows that have
	// matching column's value
	fp1, err := os.Open(csv1FName)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Can't read %s, %s\n", csv1FName, err)
		os.Exit(1)
	}
	defer fp1.Close()
	csv1 := csv.NewReader(fp1)

	fp2, err := os.Open(csv2FName)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Can't read %s, %s\n", csv2FName, err)
		os.Exit(1)
	}
	defer fp2.Close()
	csv2 := csv.NewReader(fp2)

	// Note: we read one of the tables into memory to speed things up and limit disc reads
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

	stopWords := strings.Split(stopWordsOption, ":")
	w := csv.NewWriter(out)
	lineNo := 0 // line number of csv 1 table
	for {
		rowA, err := csv1.Read()
		if err == io.EOF {
			break
		}
		if err != nil {
			fmt.Fprintf(os.Stderr, "%d %s\n", lineNo, err)
		} else {
			if col1 < len(rowA) && rowA[col1] != "" {
				// We are relying on the side effect of writing the CSV output in scanTable
				scanTable(w, rowA, col1, csv2Table, col2, stopWords)
				w.Flush()
				if err := w.Error(); err != nil {
					fmt.Fprintf(os.Stderr, "Can't write CSV at line %d of csv table 1, %s\n", lineNo, err)
				}
			}
			if verbose == true {
				if (lineNo%100) == 0 && lineNo > 0 {
					fmt.Fprintf(os.Stderr, "\n%d rows of %s processed\n", lineNo, csv1FName)
				} else {
					fmt.Fprintf(os.Stderr, ".")
				}
			}
		}
		lineNo++
	}
	w.Flush()
	if err := w.Error(); err != nil {
		fmt.Fprintf(os.Stderr, "Can't write final CSV at line %d lines processed from CSV table 1, %s\n", lineNo+1, err)
	}
}
