//
// csv2json - is a command line that takes CSV input from stdin and
// writes out JSON expression. It includes support for using the first
// row as field names or default fieldnames (e.g. col0, col1, col2).
// Additionally it can output the resulting JSON data structures as a
// JSON array or individual JSON blobs (one line per blob).
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
	"encoding/json"
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
%s reads CSV from stdin and writes a JSON to stdout. JSON output
can be either an array of JSON blobs or one JSON blob (row as object)
per line.
`

	examples = `
Convert data1.csv to data1.json using Unix pipes.

    cat data1.csv | %s > data1.json

Convert data1.csv to JSON blobs, one line per blob

    %s -as-blobs -i data1.csv
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

	// Application Options
	useHeader        bool
	asBlobs          bool
	delimiter        string
	lazyQuotes       bool
	trimLeadingSpace bool
)

func main() {
	app := cli.NewCli(datatools.Version)
	appName := app.AppName()

	// Add Help Docs
	app.AddHelp("license", []byte(fmt.Sprintf(datatools.LicenseText, appName, datatools.Version)))
	app.AddHelp("description", []byte(fmt.Sprintf(description, appName)))
	app.AddHelp("examples", []byte(fmt.Sprintf(examples, appName, appName)))

	// Standard Options
	app.BoolVar(&showHelp, "h,help", false, "display help")
	app.BoolVar(&showLicense, "l,license", false, "display license")
	app.BoolVar(&showVersion, "v,version", false, "display version")
	app.BoolVar(&showExamples, "examples", false, "display example(s)")
	app.StringVar(&inputFName, "i,input", "", "input filename")
	app.StringVar(&outputFName, "o,output", "", "output filename")
	app.BoolVar(&generateMarkdown, "generate-markdown", false, "generation markdown documentation")
	app.BoolVar(&generateManPage, "generate-manpage", false, "generation man page")
	app.BoolVar(&quiet, "quiet", false, "suppress error output")
	app.BoolVar(&newLine, "nl,newline", true, "include trailing newline in output")

	// App Options
	app.BoolVar(&useHeader, "use-header", true, "treat the first row as field names")
	app.BoolVar(&asBlobs, "as-blobs", false, "output as one JSON blob per line")
	app.StringVar(&delimiter, "d,delimiter", "", "set the delimter character")
	app.BoolVar(&lazyQuotes, "use-lazy-quotes", false, "use lazy quotes for for CSV input")
	app.BoolVar(&trimLeadingSpace, "trim-leading-space", false, "trim leading space in fields for CSV input")

	// Parse environment and options
	app.Parse()
	args := app.Args()

	// Setup IO
	var err error

	app.Eout = os.Stderr

	app.In, err = cli.Open(inputFName, os.Stdin)
	cli.ExitOnError(app.Eout, err, quiet)
	defer cli.CloseFile(inputFName, app.In)

	app.Out, err = cli.Create(outputFName, os.Stdout)
	cli.ExitOnError(app.Eout, err, quiet)
	defer cli.CloseFile(outputFName, app.Out)

	// Process options
	if generateMarkdown {
		app.GenerateMarkdown(app.Out)
		os.Exit(0)
	}
	if generateManPage {
		app.GenerateManPage(app.Out)
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
	if showLicense == true {
		fmt.Fprintln(app.Out, app.License())
		os.Exit(0)
	}
	if showVersion == true {
		fmt.Fprintln(app.Out, app.Version())
		os.Exit(0)
	}
	if newLine {
		eol = "\n"
	}

	rowNo := 0
	fieldNames := []string{}
	r := csv.NewReader(app.In)
	r.LazyQuotes = lazyQuotes
	r.TrimLeadingSpace = trimLeadingSpace
	if delimiter != "" {
		r.Comma = datatools.NormalizeDelimiterRune(delimiter)
	}
	if useHeader == true {
		row, err := r.Read()
		if err == io.EOF {
			cli.ExitOnError(app.Eout, fmt.Errorf("No data"), quiet)
		}
		cli.ExitOnError(app.Eout, err, quiet)
		for _, val := range row {
			fieldNames = append(fieldNames, strings.TrimSpace(val))
		}
		rowNo++
	}
	hasError := false
	arrayOfObjects := []string{}
	object := map[string]interface{}{}
	for {
		row, err := r.Read()
		if err == io.EOF {
			break
		}
		cli.ExitOnError(app.Eout, err, quiet)

		// Pad the fieldnames if necessary
		object = map[string]interface{}{}
		for col, val := range row {
			if col < len(fieldNames) {
				object[fieldNames[col]] = val
			} else {
				object[fmt.Sprintf("col_%d", col)] = val
			}
		}
		src, err := json.Marshal(object)
		if err != nil {
			cli.OnError(app.Eout, fmt.Errorf("error row %d, %s\n", rowNo, err), quiet)
			hasError = true
		}
		if asBlobs == true {
			fmt.Fprintf(app.Out, "%s%s", src, eol)
		} else {
			arrayOfObjects = append(arrayOfObjects, string(src))
		}
		rowNo++
	}
	if asBlobs == false {
		fmt.Fprintf(app.Out, "[%s]%s", strings.Join(arrayOfObjects, ","), eol)
	}
	if hasError == true {
		os.Exit(1)
	}
}
