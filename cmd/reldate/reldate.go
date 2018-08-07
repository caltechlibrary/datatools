//
// Generates a date in YYYY-MM-DD format based on a relative time
// description (e.g. -1 week, +3 years)
//
// @Author R. S. Doiel, <rsdoiel@caltech.edu>
//
// Copyright (c) 2018, Caltech
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
	"io"
	"os"
	"strconv"
	"strings"
	"time"

	// Local package
	"github.com/caltechlibrary/cli"
	"github.com/caltechlibrary/datatools"
	"github.com/caltechlibrary/datatools/reldate"
)

var (
	description = `
%s is a small command line utility which returns the relative date in
YYYY-MM-DD format. This is helpful when scripting various time
relationships. The difference in time returned are determined by
the time increments provided.

Time increments are a positive or negative integer. Time unit can be
either day(s), week(s), month(s), or year(s). Weekday names are
case insentive (e.g. Monday and monday). They can be abbreviated
to the first three letters of the name, e.g. Sunday can be Sun, Monday
can be Mon, Tuesday can be Tue, Wednesday can be Wed, Thursday can
be Thu, Friday can be Fri or Saturday can be Sat.
`

	examples = `
If today was 2014-08-03 and you wanted the date three days in the past try–

    %s 3 days

The output would be

    2014-08-06

TIME UNITS

Supported time units are

+ day(s)
+ week(s)
+ year(s)

Specifying a date to calucate from

%s handles dates in the YYYY-MM-DD format (e.g. March 1, 2014 would be
2014-03-01). By default reldate uses today as the date to calculate relative
time from. If you use the –from option you can it will calculate the
relative date from that specific date.

   %s --from=2014-08-03 3 days

Will yield

    2014-08-06

NEGATIVE INCREMENTS

Command line arguments traditionally start with a dash which we also use to
denote a nagative number. To tell the command line process that to not treat
negative numbers as an “option” precede your time increment and time unit
with a double dash.

    %s --from=2014-08-03 -- -3 days

Will yield

    2014-07-31

RELATIVE WEEK DAYS

You can calculate a date from a weekday name (e.g. Saturday, Monday, Tuesday)
knowning a day (e.g. 2015-02-10 or the current date of the week) occurring in
a week. A common case would be wanting to figure out the Monday date of a week
containing 2015-02-10. The week is presumed to start on Sunday (i.e. 0) and
finish with Saturday (e.g. 6).

    %s --from=2015-02-10 Monday

will yield

    2015-02-09

As that is the Monday of the week containing 2015-02-10. Weekday names case
insensitive and can be the first three letters of the English names or full
English names (e.g. Monday, monday, Mon, mon).
`

	// Standard Options
	showHelp         bool
	showVersion      bool
	showLicense      bool
	showExamples     bool
	outputFName      string
	generateMarkdown bool
	quiet            bool
	newLine          bool
	eol              string

	// Application Options
	endOfMonthFor bool
	relativeTo    string
	relativeT     time.Time
)

func assertOk(eout io.Writer, e error, failMsg string) {
	if e != nil {
		fmt.Fprintf(eout, " %s, %s", failMsg, e)
		os.Exit(1)
	}
}

func main() {
	const (
		relativeToUsage = "Date the relative time is calculated from."
		endOfMonthUsage = "Display the end of month day. E.g. 2012-02-29"
	)
	app := cli.NewCli(datatools.Version)
	appName := app.AppName()

	// Add Help Docs
	app.AddHelp("license", []byte(fmt.Sprintf(datatools.LicenseText, appName, datatools.Version)))
	app.AddHelp("description", []byte(fmt.Sprintf(description, appName)))
	app.AddHelp("examples", []byte(fmt.Sprintf(examples, appName, appName, appName, appName, appName)))

	// Document non-option Params
	app.AddParams("[TIME_DESCRPTION]")

	// Standard Options
	app.BoolVar(&showHelp, "h,help", false, "display help")
	app.BoolVar(&showLicense, "l,license", false, "display license")
	app.BoolVar(&showVersion, "v,version", false, "display version")
	app.BoolVar(&showExamples, "examples", false, "display example(s)")
	app.BoolVar(&quiet, "quiet", false, "suppress error messages")
	app.BoolVar(&newLine, "nl,newline", true, "if true add a trailing newline")
	app.BoolVar(&generateMarkdown, "generate-markdown", false, "generate markdown documentation")

	// App Specific Options
	app.StringVar(&relativeTo, "f,from", relativeTo, relativeToUsage)
	app.BoolVar(&endOfMonthFor, "e,end-of-month", endOfMonthFor, endOfMonthUsage)

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

	// Process Options
	if generateMarkdown {
		app.GenerateMarkdown(app.Out)
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

	argc := app.NArg()
	argv := app.Args()

	var (
		unitString string
	)

	if argc < 1 && endOfMonthFor == false {
		cli.ExitOnError(app.Eout, fmt.Errorf("Missing time increment and units (e.g. +2 days) or weekday name (e.g. Monday, Mon)."), quiet)
	} else if argc > 2 {
		cli.ExitOnError(app.Eout, fmt.Errorf("Too many command line arguments."), quiet)
	}

	relativeT = time.Now()
	if relativeTo != "" {
		relativeT, err = time.Parse(reldate.YYYYMMDD, relativeTo)
		assertOk(app.Eout, err, "Cannot parse the from date.\n")
	}

	if endOfMonthFor == true {
		fmt.Fprintf(app.Out, "%s%s", reldate.EndOfMonth(relativeT), eol)
		os.Exit(0)
	}

	timeInc := 0
	if argc == 2 {
		unitString = strings.ToLower(argv[1])
		timeInc, err = strconv.Atoi(argv[0])
		assertOk(app.Eout, err, "Time increment should be a positive or negative integer.\n")
	} else {
		// We may have a weekday string
		unitString = strings.ToLower(argv[0])
	}
	t, err := reldate.RelativeTime(relativeT, timeInc, unitString)
	assertOk(app.Eout, err, "Did not understand command.")
	fmt.Fprintf(app.Out, "%s%s", t.Format(reldate.YYYYMMDD), eol)
}
