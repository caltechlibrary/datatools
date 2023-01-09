// timefmt formats a date based on the formatting options available with
// Golang's Time.Format
//
// @Author R. S. Doiel, <rsdoiel@caltech.edu>
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
//
package main

import (
	"flag"
	"fmt"
	"os"
	"path"
	"strings"
	"time"

	// CaltechLibrary Packages
	"github.com/caltechlibrary/datatools"
	"github.com/caltechlibrary/datatools/timefmt"
)

var (
	helpText = `---
title: "{app_name} (1) user manual"
author: "R. S. Doiel"
pubDate: 2023-01-09
---

# NAME

{app_name}

# SYNOPSIS

{app_name} [OPTIONS] TIME_STRING_TO_CONVERT

# DESCRIPTION

{app_name} formats the current date or INPUT_DATE based on the output
format provided in options. The default input and  output format is
RFC3339.  Formats are specified based on Golang's time package including
the common constants (e.g. RFC822, RFC1123). 

For details see https://golang.org/pkg/time/#Time.Format.

One additional time layout provided by {app_name} 
 
- mysql "2006-01-02 15:04:05 -0700"

# OPTIONS

-help
: display help

-license
: display license

-version
: display version

-if, -input-format
: Set format for input

-nl, -newline
: if true add a trailing newline

-o, -output
: output filename

-of, -output-format
: Set format for output

-quiet
: suppress error messages

-utc
: timestamps in UTC


# EXAMPLES

Format the date July, 7, 2016 in YYYY-MM-DD format

~~~
    {app_name} -if "2006-01-02" -of "01/02/2006" "2017-12-02"
~~~

Yields "12/02/2017"

Format the MySQL date/time of 8:08am, July 2, 2016

~~~
    {app_name} -input-format mysql -output-format RFC822  "2017-12-02 08:08:08"
~~~

Yields "02 Dec 17 08:08 UTC"

{app_name} {version}
`

	// Standard Options
	showHelp         bool
	showVersion      bool
	showLicense      bool
	showExamples     bool
	outputFName      string
	generateMarkdown bool
	generateManPage  bool
	quiet            bool
	newLine          bool
	eol              string

	// Application Specific Options
	useUTC       bool
	inputFormat  = time.RFC3339
	outputFormat = time.RFC3339
)

func fmtTxt(src string, appName string, version string) string {
	return strings.ReplaceAll(strings.ReplaceAll(src, "{app_name}", appName), "{version}", version)
}

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
	appName := path.Base(os.Args[0])

	// Standard Options
	flag.BoolVar(&showHelp, "help", false, "display help")
	flag.BoolVar(&showLicense, "license", false, "display license")
	flag.BoolVar(&showVersion, "version", false, "display version")

	flag.StringVar(&outputFName, "o,output", "", "output filename")
	flag.StringVar(&outputFName, "o,output", "", "output filename")
	flag.BoolVar(&quiet, "quiet", false, "suppress error messages")
	flag.BoolVar(&newLine, "nl,newline", false, "if true add a trailing newline")
	flag.BoolVar(&newLine, "nl,newline", false, "if true add a trailing newline")

	// Application Options
	flag.BoolVar(&useUTC, "utc", false, "timestamps in UTC")
	flag.StringVar(&inputFormat, "if,input-format", inputFormat, "Set format for input")
	flag.StringVar(&inputFormat, "if,input-format", inputFormat, "Set format for input")
	flag.StringVar(&outputFormat, "of,output-format", outputFormat, "Set format for output")
	flag.StringVar(&outputFormat, "of,output-format", outputFormat, "Set format for output")

	// Parse env and options
	flag.Parse()
	args := flag.Args()

	// Setup IO
	var err error

	out := os.Stdout
	eout := os.Stderr

	if outputFName != "" {
		out, err = os.Create(outputFName)
		if err != nil {
			fmt.Fprintln(eout, err)
			os.Exit(1)
		}
		defer out.Close()
	}

	// Process options
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
			if err != nil {
				fmt.Fprintln(eout, err)
				os.Exit(1)
			}
			if i > 0 && newLine == false {
				fmt.Fprint(out, " ")
			}
			fmt.Fprintf(out, "%s%s", inputDate.Format(outputFormat), eol)
		}
		os.Exit(0)
	}
	inputDate = time.Now()
	fmt.Fprintf(out, "%s%s", inputDate.Format(outputFormat), eol)
}
