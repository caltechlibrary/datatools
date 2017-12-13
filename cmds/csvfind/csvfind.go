//
// csvfind - is a command line that takes CSV files in returns the rows that match a column value.
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
	"fmt"
	"io"
	"os"
	"strings"

	// Caltech Library packages
	"github.com/caltechlibrary/cli"
	"github.com/caltechlibrary/datatools"
)

var (
	description = `
%s processes a CSV file as input returning rows that contain the column
with matched text. Columns are counted from one instead of zero. Supports 
exact match as well as some Levenshtein matching.
`

	examples = `
Find the rows where the third column matches "The Red Book of Westmarch" exactly

    %s -i books.csv -col=2 "The Red Book of Westmarch"

Find the rows where the third column (colums numbered 1,2,3) matches approximately 
"The Red Book of Westmarch"

    %s -i books.csv -col=2 -levenshtein \
       -insert-cost=1 -delete-cost=1 -substitute-cost=3 \
       -max-edit-distance=50 -append-edit-distance \
       "The Red Book of Westmarch"

In this example we've appended the edit distance to see how close the matches are.

You can also search for phrases in columns.

    %s -i books.csv -col=2 -contains "Red Book"
`

	// Standard Options
	showHelp             bool
	showLicense          bool
	showVersion          bool
	showExamples         bool
	inputFName           string
	outputFName          string
	generateMarkdownDocs bool
	quiet                bool
	newLine              bool
	eol                  string

	// App Options
	skipHeaderRow      bool
	col                int
	useContains        bool
	useLevenshtein     bool
	insertCost         int
	deleteCost         int
	substituteCost     int
	caseSensitive      bool
	maxEditDistance    int
	appendEditDistance bool
	stopWordsOption    string
	trimSpaces         bool
	allowDuplicates    bool
	delimiter          string
	lazyQuotes         bool
	trimLeadingSpace   bool
)

func main() {
	app := cli.NewCli(datatools.Version)
	appName := app.AppName()

	// Document non-option parameters
	app.AddParams(`TEXT_TO_MATCH`)

	// Add Help Docs
	app.AddHelp("license", []byte(fmt.Sprintf(datatools.LicenseText, appName, datatools.Version)))
	app.AddHelp("description", []byte(fmt.Sprintf(description, appName)))
	app.AddHelp("examples", []byte(fmt.Sprintf(examples, appName, appName, appName)))

	// Basic Options
	app.BoolVar(&showHelp, "h,help", false, "display help")
	app.BoolVar(&showLicense, "l,license", false, "display license")
	app.BoolVar(&showVersion, "v,version", false, "display version")
	app.BoolVar(&showExamples, "examples", false, "display example(s)")
	app.StringVar(&inputFName, "i,input", "", "input filename")
	app.StringVar(&outputFName, "o,output", "", "output filename")
	app.BoolVar(&quiet, "quiet", false, "suppress error messages")
	app.BoolVar(&newLine, "nl,newline", true, "include trailing newline from output")
	app.BoolVar(&generateMarkdownDocs, "generate-markdown-docs", false, "generation markdown documentation")

	// App Options
	app.IntVar(&col, "col,cols", 0, "column to search for match in the CSV file")
	app.BoolVar(&useContains, "contains", false, "use contains phrase for matching")
	app.StringVar(&delimiter, "d,delimiter", "", "set delimiter character")
	app.BoolVar(&useLevenshtein, "levenshtein", false, "use levenshtein matching")
	app.IntVar(&maxEditDistance, "max-edit-distance", 5, "set the edit distance thresh hold for match, default 0")
	app.IntVar(&insertCost, "insert-cost", 1, "set the insert cost to use for levenshtein matching")
	app.IntVar(&deleteCost, "delete-cost", 1, "set the delete cost to use for levenshtein matching")
	app.IntVar(&substituteCost, "substitute-cost", 1, "set the substitution cost to use for levenshtein matching")
	app.BoolVar(&caseSensitive, "case-sensitive", false, "perform a case sensitive match (default is false)")
	app.BoolVar(&appendEditDistance, "append-edit-distance", false, "append column with edit distance found (useful for tuning levenshtein)")
	app.StringVar(&stopWordsOption, "stop-words", "", "use the colon delimited list of stop words")
	app.BoolVar(&skipHeaderRow, "skip-header-row", true, "skip the header row")
	app.BoolVar(&allowDuplicates, "allow-duplicates", true, "allow duplicates when searching for matches")
	app.BoolVar(&trimSpaces, "trimspace,trimspaces", false, "trim spaces around cell values before comparing")
	app.BoolVar(&lazyQuotes, "use-lazy-quotes", false, "use lazy quotes on CSV input")
	app.BoolVar(&trimLeadingSpace, "trim-leading-space", false, "trim leadings space in field(s) for CSV input")

	// Parse env and options
	app.Parse()
	args := app.Args()

	// Setup IO
	var err error

	app.Eout = app.Eout

	app.In, err = cli.Open(inputFName, os.Stdin)
	cli.ExitOnError(app.Eout, err, quiet)
	defer cli.CloseFile(inputFName, app.In)

	app.Out, err = cli.Create(outputFName, os.Stdout)
	cli.ExitOnError(app.Eout, err, quiet)
	defer cli.CloseFile(outputFName, app.Out)

	// Process options
	if generateMarkdownDocs {
		app.GenerateMarkdownDocs(app.Out)
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

	if col <= 0 {
		cli.ExitOnError(app.Eout, fmt.Errorf("Cannot have a zero or negative column reference %d", col), quiet)
	}
	col = col - 1

	if len(args) == 0 {
		cli.ExitOnError(app.Eout, fmt.Errorf("Missing string to match, try %s --help", appName), quiet)
	}

	target := args[0]
	stopWords := []string{}

	// NOTE: If we're doing a case Insensitive search (the default) the lower case everything before matching
	if caseSensitive == false {
		target = strings.ToLower(target)
		stopWordsOption = strings.ToLower(stopWordsOption)
	}
	if len(stopWordsOption) > 0 {
		stopWords = strings.Split(stopWordsOption, ":")
		target = strings.Join(datatools.ApplyStopWords(strings.Split(target, " "), stopWords), " ")
	}

	csvIn := csv.NewReader(app.In)
	csvIn.LazyQuotes = lazyQuotes
	csvIn.TrimLeadingSpace = trimLeadingSpace
	csvOut := csv.NewWriter(app.Out)
	if delimiter != "" {
		csvIn.Comma = datatools.NormalizeDelimiterRune(delimiter)
		csvOut.Comma = datatools.NormalizeDelimiterRune(delimiter)
	}
	lineNo := 0
	if skipHeaderRow == true {
		_, _ = csvIn.Read()
	}
	for {
		lineNo++
		record, err := csvIn.Read()
		if err == io.EOF {
			break
		}
		if err != nil {
			cli.OnError(app.Eout, fmt.Errorf("%d %s", lineNo, err), quiet)
		} else {
			// Find the value we're matching against
			if col < len(record) {
				src := record[col]
				if caseSensitive == false {
					src = strings.ToLower(src)
				}
				if trimSpaces == true {
					src = strings.TrimSpace(src)
				}
				if len(stopWords) > 0 {
					// Split into fields applying datatools filter
					fields := strings.FieldsFunc(src, func(c rune) bool {
						return datatools.Filter(c, "", false)
					})
					// Convert to an array of strings back into a space separted string
					src = strings.Join(datatools.ApplyStopWords(fields, stopWords), " ")
				}
				switch {
				case useContains:
					if strings.Contains(src, target) {
						err := csvOut.Write(record)
						if err != nil {
							cli.OnError(app.Eout, fmt.Errorf("%d %s", lineNo, err), quiet)
						}
					}
				case useLevenshtein == true:
					distance := datatools.Levenshtein(src, target, insertCost, deleteCost, substituteCost, caseSensitive)
					if distance <= maxEditDistance {
						if appendEditDistance == true {
							record = append(record, fmt.Sprintf("%d", distance))
						}
						err := csvOut.Write(record)
						if err != nil {
							cli.OnError(app.Eout, fmt.Errorf("%d %s", lineNo, err), quiet)
						}
					}
				default:
					if strings.Compare(src, target) == 0 {
						err := csvOut.Write(record)
						if err != nil {
							cli.OnError(app.Eout, fmt.Errorf("%d %s", lineNo, err), quiet)
						}
					}
				}
				if allowDuplicates == false {
					break
				}
			} else {
				cli.OnError(app.Eout, fmt.Errorf("%d line skipped, missing column %d", lineNo, col), quiet)
			}
		}
	}
	csvOut.Flush()
	err = csvOut.Error()
	cli.ExitOnError(app.Eout, err, quiet)
	fmt.Fprintf(app.Out, "%s", eol)
}
