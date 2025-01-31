// csvfind - is a command line that takes CSV files in returns the rows that match a column value.
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
	"runtime"
	"strings"

	// Caltech Library packages
	"github.com/caltechlibrary/datatools"
)

var (
	helpText = `%{app_name}(1) user manual | version {version} {release_hash}
% R. S. Doiel
% {release_date}

# NAME

{app_name}

# SYNOPSIS

{app_name} [OPTIONS] TEXT_TO_MATCH

# DESCRIPTION

{app_name} processes a CSV file as input returning rows that contain
the column with matched text. Columns are counted from one instead of
zero. Supports exact match as well as some Levenshtein matching.

# OPTIONS

-help
: display help

-license
: display license

-version
: display version


-allow-duplicates
: allow duplicates when searching for matches

-append-edit-distance
: append column with edit distance found (useful for tuning levenshtein)

-case-sensitive
: perform a case sensitive match (default is false)

-col, -cols
: column to search for match in the CSV file

-contains
: use contains phrase for matching

-d, -delimiter
: set delimiter character

-delete-cost
: set the delete cost to use for levenshtein matching

-i, -input
: input filename

-insert-cost
: set the insert cost to use for levenshtein matching

-levenshtein
: use levenshtein matching

-max-edit-distance
: set the edit distance thresh hold for match, default 0

-nl, -newline
: include trailing newline from output for end of file (EOF)

-crlf
: use CRLF for end of line (EOL) on write, defaults to true for Windows

-o, -output
: output filename

-quiet
: suppress error messages

-skip-header-row
: skip the header row

-stop-words
: use the colon delimited list of stop words

-substitute-cost
: set the substitution cost to use for levenshtein matching

-trim-leading-space
: trim leadings space in field(s) for CSV input

-trimspace, -trimspaces
: trim spaces around cell values before comparing

-use-lazy-quotes
: use lazy quotes on CSV input


# EXAMPLES

Find the rows where the third column matches "The Red Book of Westmarch"
exactly

~~~
    {app_name} -i books.csv -col=2 "The Red Book of Westmarch"
~~~

Find the rows where the third column (colums numbered 1,2,3) matches
approximately "The Red Book of Westmarch"

~~~
    {app_name} -i books.csv -col=2 -levenshtein \
       -insert-cost=1 -delete-cost=1 -substitute-cost=3 \
       -max-edit-distance=50 -append-edit-distance \
       "The Red Book of Westmarch"
~~~

In this example we've appended the edit distance to see how close the
matches are.

You can also search for phrases in columns.

~~~
    {app_name} -i books.csv -col=2 -contains "Red Book"
~~~

{app_name} {version}
`

	// Standard Options
	showHelp         bool
	showLicense      bool
	showVersion      bool
	showExamples     bool
	inputFName       string
	outputFName      string
	generateMarkdown bool
	generateManPage  bool
	quiet            bool
	newLine          bool
	eol              string

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
	useCRLF            bool
)

func main() {
	appName := path.Base(os.Args[0])
	version := datatools.Version
	license := datatools.LicenseText
	releaseDate := datatools.ReleaseDate
	releaseHash := datatools.ReleaseHash
	useCRLF = (runtime.GOOS == "windows")

	// Basic Options
	flag.BoolVar(&showHelp, "help", false, "display help")
	flag.BoolVar(&showLicense, "license", false, "display license")
	flag.BoolVar(&showVersion, "version", false, "display version")

	flag.StringVar(&inputFName, "i", "", "input filename")
	flag.StringVar(&inputFName, "input", "", "input filename")
	flag.StringVar(&outputFName, "o", "", "output filename")
	flag.StringVar(&outputFName, "output", "", "output filename")
	flag.BoolVar(&quiet, "quiet", false, "suppress error messages")
	flag.BoolVar(&newLine, "nl", true, "include trailing newline from output")
	flag.BoolVar(&newLine, "newline", true, "include trailing newline from output")

	// App Options
	flag.IntVar(&col, "col", 0, "column to search for match in the CSV file")
	flag.IntVar(&col, "cols", 0, "column to search for match in the CSV file")
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
	flag.BoolVar(&trimSpaces, "trimspace", false, "trim spaces around cell values before comparing")
	flag.BoolVar(&trimSpaces, "trimspaces", false, "trim spaces around cell values before comparing")
	flag.BoolVar(&lazyQuotes, "use-lazy-quotes", false, "use lazy quotes on CSV input")
	flag.BoolVar(&trimLeadingSpace, "trim-leading-space", false, "trim leadings space in field(s) for CSV input")
	flag.BoolVar(&useCRLF, "crlf", useCRLF, "use CRLF for end of line (EOL) on write")

	// Parse env and options
	flag.Parse()
	args := flag.Args()

	// Setup IO
	var err error

	in := os.Stdin
	out := os.Stdout
	eout := os.Stderr

	if inputFName != "" && inputFName != "-" {
		in, err = os.Open(inputFName)
		if err != nil {
			fmt.Fprintln(eout, err)
			os.Exit(1)
		}
		defer in.Close()

	}

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

	if col <= 0 {
		fmt.Fprintf(eout, "Cannot have a zero or negative column reference %d\n", col)
		os.Exit(1)
	}
	col = col - 1

	if len(args) == 0 {
		fmt.Fprintf(eout, "Missing string to match, try %s --help\n", appName)
		os.Exit(1)
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

	csvIn := csv.NewReader(in)
	csvIn.LazyQuotes = lazyQuotes
	csvIn.TrimLeadingSpace = trimLeadingSpace
	csvOut := csv.NewWriter(out)
	csvOut.UseCRLF = useCRLF
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
			if !quiet {
				fmt.Fprintf(eout, "%d %s\n", lineNo, err)
			}
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
							if !quiet {
								fmt.Fprintf(eout, "%d %s\n", lineNo, err)
							}
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
							if !quiet {
								fmt.Fprintf(eout, "%d %s\n", lineNo, err)
							}
						}
					}
				default:
					if strings.Compare(src, target) == 0 {
						err := csvOut.Write(record)
						if err != nil {
							if !quiet {
								fmt.Fprintf(eout, "%d %s\n", lineNo, err)
							}
						}
					}
				}
				if allowDuplicates == false {
					break
				}
			} else {
				if !quiet {
					fmt.Fprintf(eout, "%d line skipped, missing column %d\n", lineNo, col)
				}
			}
		}
	}
	csvOut.Flush()
	err = csvOut.Error()
	if err != nil {
		fmt.Fprintln(eout, err)
	}
	fmt.Fprintf(out, "%s", eol)
}
