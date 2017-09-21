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
	usage = `USAGE: %s [OPTIONS]`

	description = `
SYNOPSIS

%s reads CSV from stdin and writes a JSON to stdout. JSON output
can be either an array of JSON blobs or one JSON blob (row as object)
per line.
`

	examples = `
EXAMPLES

Convert data1.csv to data1.json using Unix pipes.

    cat data1.csv | %s > data1.json

Convert data1.csv to JSON blobs, one line per blob

    %s -as-blobs -i data1.csv
`

	// Standard Options
	showHelp    bool
	showLicense bool
	showVersion bool
	inputFName  string
	outputFName string

	// Application Options
	useHeader bool
	asBlobs   bool
	delimiter string
)

func init() {
	// Standard Options
	flag.BoolVar(&showHelp, "h", false, "display help")
	flag.BoolVar(&showHelp, "help", false, "display help")
	flag.BoolVar(&showLicense, "l", false, "display license")
	flag.BoolVar(&showLicense, "license", false, "display license")
	flag.BoolVar(&showVersion, "v", false, "display version")
	flag.BoolVar(&showVersion, "version", false, "display version")
	flag.StringVar(&inputFName, "i", "", "input filename")
	flag.StringVar(&inputFName, "input", "", "input filename")
	flag.StringVar(&outputFName, "o", "", "output filename")
	flag.StringVar(&outputFName, "output", "", "output filename")

	// App Options
	flag.BoolVar(&useHeader, "use-header", true, "treat the first row as field names")
	flag.BoolVar(&asBlobs, "as-blobs", false, "output as one JSON blob per line")
	flag.StringVar(&delimiter, "d", "", "set the delimter character")
	flag.StringVar(&delimiter, "delimiter", "", "set the delimter character")
}

func main() {
	appName := path.Base(os.Args[0])
	flag.Parse()

	// Configuration and command line interation
	cfg := cli.New(appName, appName, fmt.Sprintf(datatools.LicenseText, appName, datatools.Version), datatools.Version)
	cfg.UsageText = fmt.Sprintf(usage, appName)
	cfg.DescriptionText = fmt.Sprintf(description, appName)
	cfg.ExampleText = fmt.Sprintf(examples, appName, appName)

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

	rowNo := 0
	fieldNames := []string{}
	r := csv.NewReader(in)
	if delimiter != "" {
		r.Comma = datatools.NormalizeDelimiterRune(delimiter)
	}
	if useHeader == true {
		row, err := r.Read()
		if err == io.EOF {
			fmt.Fprintf(os.Stderr, "No data\n")
			os.Exit(1)
		}
		if err != nil {
			fmt.Fprintf(os.Stderr, "%s\n", err)
			os.Exit(1)
		}
		for _, val := range row {
			fieldNames = append(fieldNames, strings.TrimSpace(val))
		}
		rowNo++
	}
	arrayOfObjects := []string{}
	object := map[string]interface{}{}
	for {
		row, err := r.Read()
		if err == io.EOF {
			break
		}
		if err != nil {
			fmt.Fprintf(os.Stderr, "%s\n", err)
			os.Exit(1)
		}
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
			fmt.Fprintf(os.Stderr, "error row %d, %s\n", rowNo, err)
		}
		if asBlobs == true {
			fmt.Fprintf(out, "%s\n", src)
		} else {
			arrayOfObjects = append(arrayOfObjects, string(src))
		}
		rowNo++
	}
	if asBlobs == false {
		fmt.Fprintf(out, "[%s]\n", strings.Join(arrayOfObjects, ","))
	}
}
