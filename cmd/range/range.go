//
// range - emit a list of integers separated by spaces starting from
// first command line parameter to last command line parameter.
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
	"errors"
	"fmt"
	"io"
	"math/rand"
	"os"
	"strconv"
	"time"

	// CaltechLibrary Packages
	"github.com/caltechlibrary/cli"
	"github.com/caltechlibrary/datatools"
)

var (
	description = `
%s is a simple utility for shell scripts that emits a list of 
integers starting with the first command line argument and 
ending with the last integer command line argument. It is a 
subset of functionality found in the Unix seq command.

If the first argument is greater than the last then it counts 
down otherwise it counts up.
`

	examples = `
Create a range of integers one through five

	%s 1 5

Yields 1 2 3 4 5

Create a range of integer negative two to six

	%s -- -2 6

Yields -2 -1 0 1 2 3 4 5 6

Create a range of even integers two to ten

	%s -increment=2 2 10

Yields 2 4 6 8 10

Create a descending range of integers ten down to one

	%s 10 1

Yields 10 9 8 7 6 5 4 3 2 1


Pick a random integer between zero and ten

	%s -random 0 10

Yields a random integer from 0 to 10
`

	// Standard Options
	showHelp         bool
	showLicense      bool
	showVersion      bool
	showExamples     bool
	outputFName      string
	generateMarkdown bool
	generateManPage  bool
	quiet            bool
	newLine          bool
	eol              string

	// Application Specific Options
	start         int
	end           int
	increment     int
	randomElement bool
)

func assertOk(eout io.Writer, e error, failMsg string) {
	if e != nil {
		fmt.Fprintf(eout, " %s, %s", failMsg, e)
		os.Exit(0)
	}
}

func inRange(i, start, end int) bool {
	if start <= end && i <= end {
		return true
	}
	if start >= end && i >= end {
		return true
	}
	return false
}

func main() {
	const (
		startUsage = "The starting integer."
		endUsage   = "The ending integer."
		incUsage   = "The non-zero integer increment value."
	)
	app := cli.NewCli(datatools.Version)
	appName := app.AppName()

	// Add Help Docs
	app.AddHelp("license", []byte(fmt.Sprintf(datatools.LicenseText, appName, datatools.Version)))
	app.AddHelp("description", []byte(fmt.Sprintf(description, appName)))
	app.AddHelp("examples", []byte(fmt.Sprintf(examples, appName, appName, appName, appName, appName)))

	// Document non-option parameters
	app.SetParams("START_INTEGER", "END_INTEGER", "[INCREMENT_INTEGER]")

	// Standard Options
	app.BoolVar(&showHelp, "h,help", false, "display help")
	app.BoolVar(&showLicense, "l,license", false, "display license")
	app.BoolVar(&showVersion, "v,version", false, "display version")
	app.BoolVar(&showExamples, "examples", false, "display example(s)")
	app.BoolVar(&generateMarkdown, "generate-markdown", false, "generate markdown documentation")
	app.BoolVar(&generateManPage, "generate-manpage", false, "generate man page")
	app.BoolVar(&quiet, "quiet", false, "suppress error messages")
	app.BoolVar(&newLine, "nl,newline", false, "if true add a trailing newline")

	// App specific options
	app.IntVar(&start, "s,start", 0, startUsage)
	app.IntVar(&end, "e,end", 0, endUsage)
	app.IntVar(&increment, "inc,increment", 1, incUsage)
	app.BoolVar(&randomElement, "random", false, "Pick a range value from range")

	// Parse env and options
	app.Parse()
	args := app.Args()

	// Setup IO
	var err error

	app.Eout = os.Stderr

	/* NOTE: we don't read from stdin as we need tp csv files
	app.In, err = cli.Open(inputFName, os.Stdin)
	cli.ExitOnError(app.Eout, err, quiet)
	defer cli.CloseFile(inputFName, app.In)
	*/

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

	if argc < 2 {
		cli.ExitOnError(app.Eout, fmt.Errorf("Must include start and end integers."), quiet)
	} else if argc > 3 {
		cli.ExitOnError(app.Eout, fmt.Errorf("Too many command line arguments."), quiet)
	}

	start, err := strconv.Atoi(argv[0])
	assertOk(app.Eout, err, "Start value must be an integer.")
	end, err := strconv.Atoi(argv[1])
	assertOk(app.Eout, err, "End value must be an integer.")
	if argc == 3 {
		increment, err = strconv.Atoi(argv[2])
	} else if increment == 0 {
		err = errors.New("increment was zero")
	}
	assertOk(app.Eout, err, "Increment must be a non-zero integer.")

	if start == end {
		fmt.Fprintf(app.Out, "%d%s", start, eol)
		os.Exit(0)
	}

	// Normalize to a positive value.
	if start <= end && increment < 0 {
		increment = increment * -1
	}
	if start > end && increment > 0 {
		increment = increment * -1
	}

	// If randonElement than generate range and pick the ith random element from range
	var (
		ithArray []int
		ith      = 0
	)

	// Now count up or down as appropriate.
	for i := start; inRange(i, start, end) == true; i = i + increment {
		if randomElement == true {
			ithArray = append(ithArray, i)
		} else {
			if i == start {
				fmt.Fprintf(app.Out, "%d%s", i, eol)
			} else {
				fmt.Fprintf(app.Out, " %d%s", i, eol)
			}
		}
	}
	// if randomElement we should an array we can pick the elements from
	if randomElement == true {
		rand.Seed(time.Now().Unix())
		ith = rand.Intn(len(ithArray))
		fmt.Fprintf(app.Out, "%d%s", ithArray[ith], eol)
	}
}
