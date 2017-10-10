//
// range.go - emit a list of integers separated by spaces starting from
// first command line parameter to last command line parameter.
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
	"errors"
	"flag"
	"fmt"
	"math/rand"
	"os"
	"path"
	"strconv"
	"strings"
	"time"

	// CaltechLibrary Packages
	"github.com/caltechlibrary/cli"
	"github.com/caltechlibrary/datatools"
)

var (
	usage = `USAGE: %s [OPTIONS] START_INTEGER END_INTEGER [INCREMENT_INTEGER]`

	description = `SYNOPSIS

%s is a simple utility for shell scripts that emits a list of 
integers starting with the first command line argument and 
ending with the last integer command line argument. It is a 
subset of functionality found in the Unix seq command.

If the first argument is greater than the last then it counts 
down otherwise it counts up.
`

	examples = `EXAMPLES
	
	%s 1 5

Yields 1 2 3 4 5

	%s -- -2 6

Yields -2 -1 0 1 2 3 4 5 6

	%s -increment=2 2 10

Yields 2 4 6 8 10

	%s 10 1

Yields 10 9 8 7 6 5 4 3 2 1

	%s -r 0 10

Yields a random integer from 0 to 10
`

	// Standard Options
	showHelp     bool
	showLicense  bool
	showVersion  bool
	showExamples bool

	// Application Specific Options
	start         int
	end           int
	increment     int
	randomElement bool
)

func init() {
	const (
		startUsage = "The starting integer."
		endUsage   = "The ending integer."
		incUsage   = "The non-zero integer increment value."
	)

	// Standard Options
	flag.BoolVar(&showHelp, "h", false, "display help")
	flag.BoolVar(&showHelp, "help", false, "display help")
	flag.BoolVar(&showLicense, "l", false, "display license")
	flag.BoolVar(&showLicense, "license", false, "display license")
	flag.BoolVar(&showVersion, "v", false, "display version")
	flag.BoolVar(&showVersion, "version", false, "display version")
	flag.BoolVar(&showExamples, "example", false, "display example(s)")

	// App specific options
	flag.IntVar(&start, "start", 0, startUsage)
	flag.IntVar(&start, "s", 0, startUsage)
	flag.IntVar(&end, "end", 0, endUsage)
	flag.IntVar(&end, "e", 0, endUsage)
	flag.IntVar(&increment, "increment", 1, incUsage)
	flag.IntVar(&increment, "i", 1, incUsage)
	flag.BoolVar(&randomElement, "r", false, "Pick a range value from range")
	flag.BoolVar(&randomElement, "random", false, "Pick a range value from range")
}

func assertOk(e error, failMsg string) {
	if e != nil {
		fmt.Fprintf(os.Stderr, " %s\n %s\n", failMsg, e)
		os.Exit(1)
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
	appName := path.Base(os.Args[0])
	flag.Parse()
	args := flag.Args()

	// Configuration and command line interation
	cfg := cli.New(appName, strings.ToUpper(appName), datatools.Version)
	cfg.LicenseText = fmt.Sprintf(datatools.LicenseText, appName, datatools.Version)
	cfg.UsageText = fmt.Sprintf(usage, appName)
	cfg.DescriptionText = fmt.Sprintf(description, appName)
	cfg.OptionText = "OPTIONS"
	cfg.ExampleText = fmt.Sprintf(examples, appName, appName, appName, appName, appName)

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

	argc := flag.NArg()
	argv := flag.Args()

	if argc < 2 {
		fmt.Fprintf(os.Stderr, "Must include start and end integers.")
		os.Exit(1)
	} else if argc > 3 {
		fmt.Fprintf(os.Stderr, "Too many command line arguments.")
		os.Exit(1)
	}

	start, err := strconv.Atoi(argv[0])
	assertOk(err, "Start value must be an integer.")
	end, err := strconv.Atoi(argv[1])
	assertOk(err, "End value must be an integer.")
	if argc == 3 {
		increment, err = strconv.Atoi(argv[2])
	} else if increment == 0 {
		err = errors.New("increment was zero")
	}
	assertOk(err, "Increment must be a non-zero integer.")

	if start == end {
		fmt.Printf("%d", start)
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
				fmt.Printf("%d", i)
			} else {
				fmt.Printf(" %d", i)
			}
		}
	}
	// if randomElement we should an array we can pick the elements from
	if randomElement == true {
		rand.Seed(time.Now().Unix())
		ith = rand.Intn(len(ithArray))
		fmt.Printf("%d", ithArray[ith])
	}
}
