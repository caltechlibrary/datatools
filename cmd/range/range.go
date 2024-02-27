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
package main

import (
	"errors"
	"flag"
	"fmt"
	"io"
	"math/rand"
	"os"
	"path"
	"strconv"
	"strings"
	"time"

	// CaltechLibrary Packages
	"github.com/caltechlibrary/datatools"
)

var (
	helpText = `%{app_name}(1) user manual | version {version} {release_hash}
% R. S. Doiel
% {release_date}

# NAME

{app_name}

# SYNOPSIS

{app_name} [OPTIONS] START_INTEGER END_INTEGER [INCREMENT_INTEGER]

# DESCRIPTION

{app_name} is a simple utility for shell scripts that emits a list of 
integers starting with the first command line argument and 
ending with the last integer command line argument. It is a 
subset of functionality found in the Unix seq command.

If the first argument is greater than the last then it counts 
down otherwise it counts up.

# OPTIONS

-help
: display help

-license
: display license

-version
: display version

-e, -end
: The ending integer.

-inc, -increment
: The non-zero integer increment value.

-nl, -newline
: if true add a trailing newline

-quiet
: suppress error messages

-random
: Pick a range value from range

-s, -start
: The starting integer.


# EXAMPLES

Create a range of integers one through five

~~~
	{app_name} 1 5
~~~

Yields 1 2 3 4 5

Create a range of integer negative two to six

~~~
	{app_name} -- -2 6
~~~

Yields -2 -1 0 1 2 3 4 5 6

Create a range of even integers two to ten

~~~
	{app_name} -increment=2 2 10
~~~

Yields 2 4 6 8 10

Create a descending range of integers ten down to one

~~~
	{app_name} 10 1
~~~

Yields 10 9 8 7 6 5 4 3 2 1


Pick a random integer between zero and ten

~~~
	{app_name} -random 0 10
~~~

Yields a random integer from 0 to 10

{app_name} 1.2.1
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

func fmtTxt(src string, appName string, version string) string {
	return strings.ReplaceAll(strings.ReplaceAll(src, "{app_name}", appName), "{version}", version)
}

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

	appName := path.Base(os.Args[0])

	// Standard Options
	flag.BoolVar(&showHelp, "help", false, "display help")
	flag.BoolVar(&showLicense, "license", false, "display license")
	flag.BoolVar(&showVersion, "version", false, "display version")

	flag.BoolVar(&quiet, "quiet", false, "suppress error messages")
	flag.BoolVar(&newLine, "nl", false, "if true add a trailing newline")
	flag.BoolVar(&newLine, "newline", false, "if true add a trailing newline")

	// App specific options
	flag.IntVar(&start, "s", 0, startUsage)
	flag.IntVar(&start, "start", 0, startUsage)
	flag.IntVar(&end, "e", 0, endUsage)
	flag.IntVar(&end, "end", 0, endUsage)
	flag.IntVar(&increment, "inc", 1, incUsage)
	flag.IntVar(&increment, "increment", 1, incUsage)
	flag.BoolVar(&randomElement, "random", false, "Pick a range value from range")

	// Parse env and options
	flag.Parse()
	args := flag.Args()

	// Setup IO
	var err error

	out := os.Stdout
	eout := os.Stderr

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

	argc := len(args)

	if argc < 2 {
		fmt.Fprintln(eout, "Must include start and end integers.")
		os.Exit(1)
	} else if argc > 3 {
		fmt.Fprintln(eout, "Too many command line arguments.")
		os.Exit(1)
	}

	start, err := strconv.Atoi(args[0])
	assertOk(eout, err, "Start value must be an integer.")
	end, err := strconv.Atoi(args[1])
	assertOk(eout, err, "End value must be an integer.")
	if argc == 3 {
		increment, err = strconv.Atoi(args[2])
	} else if increment == 0 {
		err = errors.New("increment was zero")
	}
	assertOk(eout, err, "Increment must be a non-zero integer.")

	if start == end {
		fmt.Fprintf(out, "%d%s", start, eol)
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
				fmt.Fprintf(out, "%d%s", i, eol)
			} else {
				fmt.Fprintf(out, " %d%s", i, eol)
			}
		}
	}
	// if randomElement we should an array we can pick the elements from
	if randomElement == true {
		rand.Seed(time.Now().Unix())
		ith = rand.Intn(len(ithArray))
		fmt.Fprintf(out, "%d%s", ithArray[ith], eol)
	}
}
