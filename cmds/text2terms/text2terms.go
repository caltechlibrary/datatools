//
// fields - is a command line that takes a string of text and turns it into JSON array or CSV
// row.  Options included support to exclude punctuation and change case.
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
	"bufio"
	"flag"
	"fmt"
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

%s reads a line of text from stdin and writing fields as JSON array, CSV row, or delimited text to stdout. 
Additional options include ignoring punctation, changing case or allowing special characters. Stardard
delimiter for output is a space.
`

	examples = `
EXAMPLES

Convert sentence into a JSON array of words.

    echo "The cat jumpted over the shifty fox." | %s -json 

Convert each word into a column in a CSV row.

    echo "The cat jumpted over the shifty fox." | %s -csv
`

	// Standard Options
	showHelp    bool
	showLicense bool
	showVersion bool
	inputFName  string
	outputFName string

	// Application Options
	asCSV            bool
	asJSON           bool
	allowPunctuation bool
	allowCharacters  string
	stopWords        string
	toLower          bool
	toUpper          bool
	delimiter        string
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
	flag.StringVar(&delimiter, "delimiter", " ", "use this delimiter for output and stop words (default is space)")
	flag.BoolVar(&asCSV, "csv", false, "output as a CSV row")
	flag.BoolVar(&asJSON, "json", false, "output as a JSON array")
	flag.BoolVar(&toLower, "to-lower", false, "lower case the input string")
	flag.BoolVar(&toUpper, "to-upper", false, "upper case the input string")
	flag.BoolVar(&allowPunctuation, "allow-punctuation", false, "allow punctuation (i.e. allows letters, numbers and punctuation)")
	flag.StringVar(&allowCharacters, "allow-characters", "", "also allow these characters")
	flag.StringVar(&stopWords, "stop-words", "", "a colon delimited list of stop words to ignore (case insensitive)")
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

	// populate options
	opts := new(datatools.Options)
	switch {
	case asCSV == true:
		opts.Format = datatools.AsCSV
	case asJSON == true:
		opts.Format = datatools.AsJSON
	default:
		opts.Format = datatools.AsDelimited
		opts.Delimiter = delimiter
	}
	opts.AllowPunctuation = allowPunctuation
	if len(allowCharacters) > 0 {
		opts.AllowCharacters = allowCharacters
	}
	if len(stopWords) > 0 {
		opts.StopWords = strings.Split(stopWords, ":")
	}
	opts.ToLower = toLower
	opts.ToUpper = toUpper

	// Map in the cli options to opts struct
	buf := bufio.NewReader(in)
	src, err := datatools.Text2Fields(buf, opts)
	if err != nil {
		fmt.Fprintf(os.Stderr, "%s", err)
	} else {
		fmt.Fprintf(out, "%s\n", src)
	}
}
