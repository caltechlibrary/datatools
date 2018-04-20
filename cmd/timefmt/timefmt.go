// datefmt formats a date based on the formatting options available with
// Golang's Time.Format
//
// @Author R. S. Doiel, <rsdoiel@caltech.edu>
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
	"fmt"
	"os"
	"strings"
	"time"

	// CaltechLibrary Packages
	"github.com/caltechlibrary/cli"
	"github.com/caltechlibrary/datatools"
	"github.com/caltechlibrary/datatools/timefmt"
)

var (
	description = `
%s formats the current date or INPUT_DATE based on the output format
provided in options. The default input and  output format is RFC3339. 
Formats are specified based on Golang's time package including the
common constants (e.g. RFC822, RFC1123). 

For details see https://golang.org/pkg/time/#Time.Format.

One additional time layout provided by %s 
 
+ mysql "2006-01-02 15:04:05 -0700"
`

	examples = `
Format the date July, 7, 2016 in YYYY-MM-DD format

    %s -if "2006-01-02" -of "01/02/2006" "2017-12-02"

Yields "12/02/2017"

Format the MySQL date/time of 8:08am, July 2, 2016

    %s -input-format mysql -output-format RFC822  "2017-12-02 08:08:08"

Yields "02 Dec 17 08:08 UTC"
`

	// Standard Options
	showHelp             bool
	showVersion          bool
	showLicense          bool
	showExamples         bool
	outputFName          string
	generateMarkdownDocs bool
	quiet                bool
	newLine              bool
	eol                  string

	// Application Specific Options
	useUTC       bool
	inputFormat  = time.RFC3339
	outputFormat = time.RFC3339
)

func applyConstants(s string) string {
	switch strings.ToLower(s) {
	case "ansic":
		s = time.ANSIC
	case "unixdate":
		s = time.UnixDate
	case "rubydate":
		s = time.RubyDate
	case "rfc822":
		s = time.RFC822
	case "rfc822z":
		s = time.RFC822Z
	case "rfc850":
		s = time.RFC850
	case "rfc1123":
		s = time.RFC1123
	case "rfc1123z":
		s = time.RFC1123Z
	case "rfc3339":
		s = time.RFC3339
	case "RFC3339Nano":
		s = time.RFC3339Nano
	case "kitchen":
		s = time.Kitchen
	case "stamp":
		s = time.Stamp
	case "stampmilli":
		s = time.StampMilli
	case "stampmicro":
		s = time.StampMicro
	case "stampnano":
		s = time.StampNano
	case "mysql":
		s = timefmt.MySQL
	}
	return s
}

func main() {
	app := cli.NewCli(datatools.Version)
	appName := app.AppName()

	// Add Help Docs
	app.AddHelp("license", []byte(fmt.Sprintf(datatools.LicenseText, appName, datatools.Version)))
	app.AddHelp("description", []byte(fmt.Sprintf(description, appName, appName)))
	app.AddHelp("examples", []byte(fmt.Sprintf(examples, appName, appName)))

	// Document non-option parameters
	app.AddParams("TIME_STRING_TO_CONVERT")

	// Standard Options
	app.BoolVar(&showHelp, "h,help", false, "display help")
	app.BoolVar(&showLicense, "l,license", false, "display license")
	app.BoolVar(&showVersion, "v,version", false, "display version")
	app.BoolVar(&showExamples, "examples", false, "display example(s)")
	//app.StringVar(&inputFName, "i,input", "", "input filename")
	app.StringVar(&outputFName, "o,output", "", "output filename")
	app.BoolVar(&quiet, "quiet", false, "suppress error messages")
	app.BoolVar(&generateMarkdownDocs, "generate-markdown-docs", false, "generate markdown documentation")
	app.BoolVar(&newLine, "nl,newline", false, "if true add a trailing newline")

	// Application Options
	app.BoolVar(&useUTC, "utc", false, "timestamps in UTC")
	app.StringVar(&inputFormat, "if,input-format", inputFormat, "Set format for input")
	app.StringVar(&outputFormat, "of,output-format", outputFormat, "Set format for output")

	// Parse env and options
	app.Parse()
	args := app.Args()

	// Setup IO
	var err error
	app.Eout = os.Stderr

	/* NOTE: this command does not read from stdin
	   app.In, err = cli.Open(inputFName, os.Stdin)
	   cli.ExitOnError(app.Eout, err, quiet)
	   defer cli.CloseFile(inputFName, app.In)
	*/

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

	var (
		inputDate time.Time
	)

	// Handle constants for formatting
	inputFormat = applyConstants(inputFormat)
	outputFormat = applyConstants(outputFormat)

	if len(args) > 0 {
		for i, dt := range args {
			inputDate, err = time.Parse(inputFormat, dt)
			cli.ExitOnError(app.Eout, err, quiet)
			if i > 0 && newLine == false {
				fmt.Fprint(app.Out, " ")
			}
			fmt.Fprintf(app.Out, "%s%s", inputDate.Format(outputFormat), eol)
		}
		os.Exit(0)
	}
	inputDate = time.Now()
	fmt.Fprintf(app.Out, "%s%s", inputDate.Format(outputFormat), eol)
}
