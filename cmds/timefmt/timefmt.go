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
	"flag"
	"fmt"
	"os"
	"path"
	"strings"
	"time"

	// CaltechLibrary Packages
	"github.com/caltechlibrary/cli"
	"github.com/caltechlibrary/datatools"
	"github.com/caltechlibrary/datatools/timefmt"
)

var (
	usage = `USAGE: %s [OPTIONS] TIME_STRING_TO_CONVERT`

	description = `

SYNOPSIS

%s formats the current date or INPUT_DATE based on the output format
provided in options. The default input and  output format is RFC3339. 
Formats are specified based on Golang's time package including the
common constants (e.g. RFC822, RFC1123). 

For details see https://golang.org/pkg/time/#Time.Format.

One additional time layout provided by %s 
 
+ mysql "2006-01-02 15:04:05 -0700"

`

	examples = `

EXAMPLES

    %s -input "2006-01-02" -output "01/02/2006" "2016-07-02"

Yields "07/02/2016"

    %s -input mysql -output RFC822  "2016-07-02 08:08:08"

Yields "02 Jul 16 08:08 UTC"

`

	// Standard Options
	showHelp     bool
	showVersion  bool
	showLicense  bool
	showExamples bool

	// Application Specific Options
	useUTC       bool
	inputFormat  = time.RFC3339
	outputFormat = time.RFC3339
)

func init() {
	// Standard Options
	flag.BoolVar(&showHelp, "h", false, "display help")
	flag.BoolVar(&showHelp, "help", false, "display help")
	flag.BoolVar(&showLicense, "l", false, "display license")
	flag.BoolVar(&showLicense, "license", false, "display license")
	flag.BoolVar(&showVersion, "v", false, "display version")
	flag.BoolVar(&showVersion, "version", false, "display version")
	flag.BoolVar(&showExamples, "example", false, "display example(s)")

	// Application Options
	flag.BoolVar(&useUTC, "utc", false, "timestamps in UTC")
	flag.StringVar(&inputFormat, "input", inputFormat, "Set format for input")
	flag.StringVar(&outputFormat, "output", outputFormat, "Set format for output")
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
	flag.Parse()
	args := flag.Args()

	// Configuration and command line interation
	cfg := cli.New(appName, strings.ToUpper(appName), datatools.Version)
	cfg.LicenseText = fmt.Sprintf(datatools.LicenseText, appName, datatools.Version)
	cfg.UsageText = fmt.Sprintf(usage, appName)
	cfg.DescriptionText = fmt.Sprintf(description, appName, appName)
	cfg.OptionText = "OPTIONS"
	cfg.ExampleText = fmt.Sprintf(examples, appName, appName)

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

	var (
		inputDate time.Time
		err       error
	)

	// Handle constants for formatting
	inputFormat = applyConstants(inputFormat)
	outputFormat = applyConstants(outputFormat)

	if len(args) > 0 {
		for i, dt := range args {
			inputDate, err = time.Parse(inputFormat, dt)
			if err != nil {
				fmt.Fprintf(os.Stderr, "can't read %s, %s\n", dt, err)
				os.Exit(1)
			}
			if i > 0 {
				fmt.Print(" ")
			}
			fmt.Printf("%s", inputDate.Format(outputFormat))
		}
		os.Exit(0)
	}
	inputDate = time.Now()
	fmt.Printf("%s", inputDate.Format(outputFormat))
}
