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
	"flag"
	"fmt"
	"io"
	"os"
	"path"
	"strings"

	// Caltech Library packages
	"github.com/caltechlibrary/cli"
	"github.com/caltechlibrary/datatools"
)

var (
	usage = `USAGE: %s [OPTIONS] TEXT_TO_MATCH`

	description = `

SYNOPSIS

%s processes a CSV file as input returning rows that contain the column
with matched text. Columns are count from one instead of zero. Supports 
exact match as well as some Levenshtein matching.

`

	examples = `

EXAMPLES

Find the rows where the third column matches "The Red Book of Westmarch" exactly

    %s -i books.csv -col=2 "The Red Book of Westmarch"

Find the rows where the third column (colums numbered 0,1,2) matches approximately 
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
	showHelp     bool
	showLicense  bool
	showVersion  bool
	showExamples bool
	inputFName   string
	outputFName  string
	quiet        bool

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
)

func init() {
	// Basic Options
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
	flag.BoolVar(&quiet, "quiet", false, "suppress error messages")

	// App Options
	flag.IntVar(&col, "col", 0, "column to search for match in the CSV file")
	flag.BoolVar(&useContains, "contains", false, "use contains phrase for matching")
	flag.StringVar(&delimiter, "d", "", "set delimiter character")
	flag.StringVar(&delimiter, "delimiter", "", "set delimiter character")
	flag.BoolVar(&useLevenshtein, "levenshtein", false, "use levenshtein matching")
	flag.IntVar(&maxEditDistance, "max-edit-distance", 5, "set the edit distance thresh hold for match, default 0")
	flag.IntVar(&insertCost, "insert-cost", 1, "set the insert cost to use for levenshtein matching")
	flag.IntVar(&deleteCost, "delete-cost", 1, "set the delete cost to use for levenshtein matching")
	flag.IntVar(&substituteCost, "substitute-cost", 1, "set the substitution cost to use for levenshtein matching")
	flag.BoolVar(&caseSensitive, "case-sensitive", false, "perform a case sensitive match (default is false)")
	flag.BoolVar(&appendEditDistance, "append-edit-distance", false, "append column with edit distance found (useful for tuning levenshtein)")
	flag.StringVar(&stopWordsOption, "stop-words", "", "use the colon delimited list of stop words")
	flag.BoolVar(&skipHeaderRow, "skip-header-row", true, "skip the header row")
	flag.BoolVar(&allowDuplicates, "allow-duplicates", true, "allow duplicates when searching for matches")
	flag.BoolVar(&trimSpaces, "trim-spaces", false, "trim spaces around cell values before comparing")
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
	cfg.OptionText = "OPTIONS\n\n"
	cfg.ExampleText = fmt.Sprintf(examples, appName, appName, appName)

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

	if col <= 0 {
		cli.ExitOnError(os.Stderr, fmt.Errorf("Cannot have a zero or negative column reference %d", col), quiet)
	}
	col = col - 1

	if len(args) == 0 {
		cli.ExitOnError(os.Stderr, fmt.Errorf("Missing string to match, try %s --help", appName), quiet)
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

	in, err := cli.Open(inputFName, os.Stdin)
	cli.ExitOnError(os.Stderr, err, quiet)
	defer cli.CloseFile(inputFName, in)
	out, err := cli.Create(outputFName, os.Stdout)
	cli.ExitOnError(os.Stderr, err, quiet)
	defer cli.CloseFile(outputFName, out)

	csvIn := csv.NewReader(in)
	csvOut := csv.NewWriter(out)
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
			cli.OnError(os.Stderr, fmt.Errorf("%d %s", lineNo, err), quiet)
		} else {
			// Find the value we're matching against
			if col < len(record) {
				src := record[col]
				if caseSensitive == false {
					src = strings.ToLower(src)
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
							cli.OnError(os.Stderr, fmt.Errorf("%d %s", lineNo, err), quiet)
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
							cli.OnError(os.Stderr, fmt.Errorf("%d %s", lineNo, err), quiet)
						}
					}
				default:
					if strings.Compare(src, target) == 0 {
						err := csvOut.Write(record)
						if err != nil {
							cli.OnError(os.Stderr, fmt.Errorf("%d %s", lineNo, err), quiet)
						}
					}
				}
				if allowDuplicates == false {
					break
				}
			} else {
				cli.OnError(os.Stderr, fmt.Errorf("%d line skipped, missing column %d", lineNo, col), quiet)
			}
		}
	}
	csvOut.Flush()
	err = csvOut.Error()
	cli.ExitOnError(os.Stderr, err, quiet)
}
