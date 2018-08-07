//
// csvjoin - is a command line that takes two CSV files and joins them by match a designated column in each.
//
// @author R. S. Doiel, <rsdoiel@caltech.edu>
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

	// My packages
	"github.com/caltechlibrary/cli"
	"github.com/caltechlibrary/datatools"
)

var (
	description = `
%s outputs CSV content based on two CSV files with matching column values.
Each CSV input file has a designated column to match on. The values are
compared as strings. Columns are counted from one rather than zero.
`

	examples = `
Simple usage of building a merged CSV file from data1.csv
and data2.csv where column 1 in data1.csv matches the value in
column 3 of data2.csv with the results being written to 
merged-data.csv..

    %s -csv1=data1.csv -col1=2 \
       -csv2=data2.csv -col2=4 \
       -output=merged-data.csv
`

	// Standard Options
	showHelp         bool
	showLicense      bool
	showVersion      bool
	showExamples     bool
	outputFName      string
	generateMarkdown bool
	quiet            bool

	// App Options
	verbose          bool
	csv1FName        string
	csv2FName        string
	col1             int
	col2             int
	trimSpaces       bool
	caseSensitive    bool
	useContains      bool
	useLevenshtein   bool
	insertCost       int
	deleteCost       int
	substituteCost   int
	maxEditDistance  int
	stopWordsOption  string
	allowDuplicates  bool
	asInMemory       bool
	delimiter        string
	lazyQuotes       bool
	trimLeadingSpace bool
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

func scanTable(eout io.Writer, w *csv.Writer, rowA []string, col1 int, table [][]string, col2 int, stopWords []string) error {
	if col1 >= len(rowA) {
		return nil
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
					return fmt.Errorf("Can't write csv row line %d of table 2, %s\n", i, err)
				}
				w.Flush()
				if verbose == true {
					fmt.Fprint(eout, "*")
				}
				if err := w.Error(); err != nil {
					return err
				}
				if allowDuplicates == false {
					return nil
				}
			}
		}
	}
	return nil
}

func main() {
	app := cli.NewCli(datatools.Version)
	appName := app.AppName()

	// Documemt non-option parameters
	app.AddParams("CSV1", "CSV2", "COL1", "COL2")

	// Add Help Docs
	app.AddHelp("license", []byte(fmt.Sprintf(datatools.LicenseText, appName, datatools.Version)))
	app.AddHelp("description", []byte(fmt.Sprintf(description, appName)))
	app.AddHelp("examples", []byte(fmt.Sprintf(examples, appName)))

	// Standard Options
	app.BoolVar(&showHelp, "h,help", false, "display help")
	app.BoolVar(&showLicense, "l,license", false, "display license")
	app.BoolVar(&showVersion, "v,version", false, "display version")
	app.BoolVar(&showExamples, "examples", false, "display example(s)")
	app.StringVar(&outputFName, "o,output", "", "output filename")
	app.BoolVar(&generateMarkdown, "generate-markdown", false, "generate markdown documentation")
	app.BoolVar(&quiet, "quiet", false, "supress error messages")

	// App Options
	app.BoolVar(&verbose, "verbose", false, "output processing count to stderr")
	app.StringVar(&csv1FName, "csv1", "", "first CSV filename")
	app.StringVar(&csv2FName, "csv2", "", "second CSV filename")
	app.IntVar(&col1, "col1", 0, "column to on join on in first CSV file")
	app.IntVar(&col2, "col2", 0, "column to on join on in second CSV file")
	app.BoolVar(&caseSensitive, "case-sensitive", false, "make a case sensitive match (default is case insensitive)")
	app.BoolVar(&useContains, "contains", false, "match columns based on csv1/col1 contained in csv2/col2")
	app.BoolVar(&useLevenshtein, "levenshtein", false, "match columns using Levensthein edit distance")
	app.IntVar(&insertCost, "insert-cost", 1, "insertion cost to use when calculating Levenshtein edit distance")
	app.IntVar(&deleteCost, "delete-cost", 1, "deletion cost to use when calculating Levenshtein edit distance")
	app.IntVar(&substituteCost, "substitute-cost", 1, "substitution cost to use when calculating Levenshtein edit distance")
	app.IntVar(&maxEditDistance, "max-edit-distance", 5, "maximum edit distance for match using Levenshtein distance")
	app.StringVar(&stopWordsOption, "stop-words", "", "a column delimited list of stop words to ingnore when matching")
	app.BoolVar(&allowDuplicates, "allow-duplicates", true, "allow duplicates when searching for matches")
	app.BoolVar(&trimSpaces, "trimspaces", false, "trim spaces around cell values before comparing")
	app.BoolVar(&asInMemory, "in-memory", false, "if true read both CSV files")
	app.StringVar(&delimiter, "d,delimiter", "", "set delimiter character")
	app.BoolVar(&lazyQuotes, "use-lazy-quotes", false, "use lazy quotes for CSV input")
	app.BoolVar(&trimLeadingSpace, "trim-leading-space", false, "trim leading space in field(s) for CSV input")

	// Parse env and options
	app.Parse()
	args := app.Args()

	// Setup IO
	var err error

	app.Eout = os.Stderr

	/* NOTE: we don't read from stdin as we need tp csv files
	app.In, err = cli.Open(inputFName, os.Stdin)
	cli.ExitOnError(app.Eout, err, quiet)
	defer cli.CloseFile(inputFName, app.In)
	*/

	app.Out, err = cli.Create(outputFName, os.Stdout)
	cli.ExitOnError(app.Eout, err, quiet)
	defer cli.CloseFile(outputFName, app.Out)

	// Process Options
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

	// NOTE: We are counting columns for humans from 1 rather than zero.
	if col1 <= 0 {
		cli.ExitOnError(app.Eout, fmt.Errorf("col1 must be one or greater, %d\n", col1), quiet)
	}
	if col2 <= 0 {
		cli.ExitOnError(app.Eout, fmt.Errorf("col2 must be one or greater, %d\n", col2), quiet)
	}
	col1--
	col2--

	if len(csv1FName) == 0 {
		cli.ExitOnError(app.Eout, fmt.Errorf("Missing first CSV filename"), quiet)
	}

	if len(csv2FName) == 0 {
		cli.ExitOnError(app.Eout, fmt.Errorf("Missing second CSV filename"), quiet)
	}

	if col1 < 0 {
		cli.ExitOnError(app.Eout, fmt.Errorf("Cannot use a negative column index %d\n", col1), quiet)
	}
	if col2 < 0 {
		cli.ExitOnError(app.Eout, fmt.Errorf("Cannot use a negative column index %d\n", col2), quiet)
	}

	// FIXME: Should only read the smaller of two files into memory
	// then interate through the other file for matches. This would let you work with larger files.

	// Read in CSV2 to memory then iterate over CSV1 output rows that have
	// matching column's value
	fp1, err := os.Open(csv1FName)
	cli.ExitOnError(app.Eout, err, quiet)
	defer fp1.Close()
	csv1 := csv.NewReader(fp1)
	csv1.LazyQuotes = lazyQuotes
	csv1.TrimLeadingSpace = trimLeadingSpace

	fp2, err := os.Open(csv2FName)
	cli.ExitOnError(app.Eout, err, quiet)
	defer fp2.Close()
	csv2 := csv.NewReader(fp2)
	csv2.LazyQuotes = lazyQuotes
	csv2.TrimLeadingSpace = trimLeadingSpace

	w := csv.NewWriter(app.Out)
	if delimiter != "" {
		csv1.Comma = datatools.NormalizeDelimiterRune(delimiter)
		csv2.Comma = datatools.NormalizeDelimiterRune(delimiter)
		w.Comma = datatools.NormalizeDelimiterRune(delimiter)
	}

	// Note: we read one of the tables into memory to speed things up and limit disc reads
	csv2Table := [][]string{}
	for {
		record, err := csv2.Read()
		if err == io.EOF {
			break
		}
		if err != nil {
			cli.OnError(app.Eout, fmt.Errorf("%s, %s (%T %+v)", csv2FName, err, record, record), quiet)
		}
		csv2Table = append(csv2Table, record)
	}

	stopWords := strings.Split(stopWordsOption, ":")
	lineNo := 0 // line number of csv 1 table
	if asInMemory == false {
		for {
			rowA, err := csv1.Read()
			if err == io.EOF {
				break
			}
			if err != nil {
				cli.OnError(app.Eout, fmt.Errorf("%d %s\n", lineNo, err), quiet)
			} else {
				if col1 < len(rowA) && rowA[col1] != "" {
					// We are relying on the side effect of writing the CSV output in scanTable
					if err := scanTable(app.Eout, w, rowA, col1, csv2Table, col2, stopWords); err != nil {
						cli.OnError(app.Eout, fmt.Errorf("Can't write CSV at line %d of csv table 1, %s\n", lineNo, err), quiet)
					}
				}
				if verbose == true {
					if (lineNo%100) == 0 && lineNo > 0 {
						cli.OnError(app.Eout, fmt.Errorf("\n%d rows of %s processed\n", lineNo, csv1FName), quiet)
					}
				}
			}
			lineNo++
		}
	} else {
		csv1Table := [][]string{}

		// Read table 1 into memory
		for {
			record, err := csv1.Read()
			if err == io.EOF {
				break
			}
			if err != nil {
				cli.OnError(app.Eout, fmt.Errorf("%s, %s (%T %+v)", csv1FName, err, record, record), quiet)
			}
			csv1Table = append(csv1Table, record)
		}
		// For each row in table one scan table two.
		for i, rowA := range csv1Table {
			if col1 < len(rowA) && rowA[col1] != "" {
				// We are relying on the side effect of writing the CSV output in scanTable
				if err := scanTable(app.Eout, w, rowA, col1, csv2Table, col2, stopWords); err != nil {
					cli.OnError(app.Eout, fmt.Errorf("Can't write CSV at line %d of csv table 1, %s", lineNo, err), quiet)
				}
			}
			if verbose == true {
				if (lineNo%100) == 0 && lineNo > 0 {
					cli.OnError(app.Eout, fmt.Errorf("%d rows of %s processed", lineNo, csv1FName), quiet)
				}
			}
			lineNo = i
		}
	}
	w.Flush()
	err = w.Error()
	cli.ExitOnError(app.Eout, err, quiet)
}
