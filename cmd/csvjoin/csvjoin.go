// csvjoin - is a command line that takes two CSV files and joins them by match a designated column in each.
//
// @author R. S. Doiel, <rsdoiel@caltech.edu>
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
	"strings"

	// My packages
	"github.com/caltechlibrary/datatools"
)

var (
	helpText = `%{app_name}(1) user manual | version {version} {release_hash}
% R. S. Doiel
% {release_date}

# NAME

{app_name}

# SYNOPSIS

{app_name} [OPTIONS] CSV1 CSV2 COL1 COL2

# DESCRIPTION

{app_name} outputs CSV content based on two CSV files with matching
column values.  Each CSV input file has a designated column to match
on. The values are compared as strings. Columns are counted from one
rather than zero.

# OPTIONS

-help
: display help

-license
: display license

-version
: display version

-allow-duplicates
: allow duplicates when searching for matches

-case-sensitive
: make a case sensitive match (default is case insensitive)

-col1
: column to on join on in first CSV file

-col2
: column to on join on in second CSV file

-contains
: match columns based on csv1/col1 contained in csv2/col2

-csv1
: first CSV filename

-csv2
: second CSV filename

-d, -delimiter
: set delimiter character

-delete-cost
: deletion cost to use when calculating Levenshtein edit distance

-in-memory
: if true read both CSV files

-insert-cost
: insertion cost to use when calculating Levenshtein edit distance

-levenshtein
: match columns using Levensthein edit distance

-max-edit-distance
: maximum edit distance for match using Levenshtein distance

-o, -output
: output filename

-quiet
: supress error messages

-stop-words
: a column delimited list of stop words to ingnore when matching

-substitute-cost
: substitution cost to use when calculating Levenshtein edit distance

-trim-leading-space
: trim leading space in field(s) for CSV input

-trimspaces
: trim spaces around cell values before comparing

-use-lazy-quotes
: use lazy quotes for CSV input

-verbose
: output processing count to stderr


# EXAMPLES

Simple usage of building a merged CSV file from data1.csv
and data2.csv where column 1 in data1.csv matches the value in
column 3 of data2.csv with the results being written to 
merged-data.csv..

~~~
    {app_name} -csv1=data1.csv -col1=2 \
       -csv2=data2.csv -col2=4 \
       -output=merged-data.csv
~~~

{app_name} {version}

`

	// Standard Options
	showHelp         bool
	showLicense      bool
	showVersion      bool
	showExamples     bool
	outputFName      string
	generateMarkdown bool
	generateManPage  bool
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

func fmtTxt(src string, appName string, version string) string {
	return strings.ReplaceAll(strings.ReplaceAll(src, "{app_name}", appName), "{version}", version)
}

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
	appName := path.Base(os.Args[0])

	// Standard Options
	flag.BoolVar(&showHelp, "help", false, "display help")
	flag.BoolVar(&showLicense, "license", false, "display license")
	flag.BoolVar(&showVersion, "version", false, "display version")

	flag.StringVar(&outputFName, "o", "", "output filename")
	flag.StringVar(&outputFName, "output", "", "output filename")
	flag.BoolVar(&quiet, "quiet", false, "supress error messages")

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
	flag.BoolVar(&trimSpaces, "trimspaces", false, "trim spaces around cell values before comparing")
	flag.BoolVar(&asInMemory, "in-memory", false, "if true read both CSV files")
	flag.StringVar(&delimiter, "d", "", "set delimiter character")
	flag.StringVar(&delimiter, "delimiter", "", "set delimiter character")
	flag.BoolVar(&lazyQuotes, "use-lazy-quotes", false, "use lazy quotes for CSV input")
	flag.BoolVar(&trimLeadingSpace, "trim-leading-space", false, "trim leading space in field(s) for CSV input")

	// Parse env and options
	flag.Parse()

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

	// Process Options
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

	// NOTE: We are counting columns for humans from 1 rather than zero.
	if col1 <= 0 {
		fmt.Fprintf(eout, "col1 must be one or greater, %d\n", col1)
		os.Exit(1)
	}
	if col2 <= 0 {
		fmt.Fprintf(eout, "col2 must be one or greater, %d\n", col2)
		os.Exit(1)
	}
	col1--
	col2--

	if len(csv1FName) == 0 {
		fmt.Fprintln(eout, "Missing first CSV filename")
		os.Exit(1)
	}

	if len(csv2FName) == 0 {
		fmt.Fprintln(eout, "Missing second CSV filename")
		os.Exit(1)
	}

	if col1 < 0 {
		fmt.Fprintf(eout, "Cannot use a negative column index %d\n", col1)
		os.Exit(1)
	}
	if col2 < 0 {
		fmt.Fprintf(eout, "Cannot use a negative column index %d\n", col2)
		os.Exit(1)
	}

	// FIXME: Should only read the smaller of two files into memory
	// then interate through the other file for matches. This would let you work with larger files.

	// Read in CSV2 to memory then iterate over CSV1 output rows that have
	// matching column's value
	fp1, err := os.Open(csv1FName)
	if err != nil {
		fmt.Fprintln(eout, err)
		os.Exit(1)
	}
	defer fp1.Close()
	csv1 := csv.NewReader(fp1)
	csv1.LazyQuotes = lazyQuotes
	csv1.TrimLeadingSpace = trimLeadingSpace

	fp2, err := os.Open(csv2FName)
	if err != nil {
		fmt.Fprintln(eout, err)
		os.Exit(1)
	}
	defer fp2.Close()
	csv2 := csv.NewReader(fp2)
	csv2.LazyQuotes = lazyQuotes
	csv2.TrimLeadingSpace = trimLeadingSpace

	w := csv.NewWriter(out)
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
			if !quiet {
				fmt.Fprintf(eout, "%s, %s (%T %+v)\n", csv2FName, err, record, record)
			}
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
				if !quiet {
					fmt.Fprintf(eout, "%d %s\n", lineNo, err)
				}
			} else {
				if col1 < len(rowA) && rowA[col1] != "" {
					// We are relying on the side effect of writing the CSV output in scanTable
					if err := scanTable(eout, w, rowA, col1, csv2Table, col2, stopWords); err != nil {
						if !quiet {
							fmt.Fprintf(eout, "Can't write CSV at line %d of csv table 1, %s\n", lineNo, err)
						}
					}
				}
				if verbose == true {
					if (lineNo%100) == 0 && lineNo > 0 {
						if !quiet {
							fmt.Fprintf(eout, "\n%d rows of %s processed\n", lineNo, csv1FName)
						}
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
				if !quiet {
					fmt.Fprintf(eout, "%s, %s (%T %+v)\n", csv1FName, err, record, record)
				}
			}
			csv1Table = append(csv1Table, record)
		}
		// For each row in table one scan table two.
		for i, rowA := range csv1Table {
			if col1 < len(rowA) && rowA[col1] != "" {
				// We are relying on the side effect of writing the CSV output in scanTable
				if err := scanTable(eout, w, rowA, col1, csv2Table, col2, stopWords); err != nil {
					if !quiet {
						fmt.Fprintf(eout, "Can't write CSV at line %d of csv table 1, %s\n", lineNo, err)
					}
				}
			}
			if verbose == true {
				if (lineNo%100) == 0 && lineNo > 0 {
					if !quiet {
						fmt.Fprintf(eout, "%d rows of %s processed\n", lineNo, csv1FName)
					}
				}
			}
			lineNo = i
		}
	}
	w.Flush()
	err = w.Error()
	if err != nil {
		fmt.Fprintln(eout, err)
		os.Exit(1)
	}
}
