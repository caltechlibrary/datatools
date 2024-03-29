// jsonobjects2csv is a command line utility that converts a JSON list of objects to CSV.
//
// @Author R. S. Doiel
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
	"flag"
	"fmt"
	"os"
	"path"

	// CaltechLibrary packages
	"github.com/caltechlibrary/datatools"
)

var (
	helpText = `%{app_name}(1) user manual | version {version} {release_hash}
% R. S. Doiel
% {release_date}

# NAME

{app_name}

# SYNOPSIS

{app_name} [OPTIONS] [JSON_FILENAME] [YAML_FILENAME]

# DESCRIPTION

{app_name} is a tool that converts a JSON list of objects into CSV output.

{app_name} will take JSON expressing a list of objects and turn them into a CSV
representation. If the object's attributes include other objects or arrays they
are rendered as YAML in the cell of the csv output.

# OPTIONS

-help
: display help

-license
: display license

-version
: display version

-nl, -newline
: if true add a trailing newline

-o, -output
: output filename

-quiet
: suppress error messages

-delimiter
: set the CSV column delimiter for output

-show-header
: set whether or not to output a header row at start of outout.

-i FILENAME
: Use  FILENAME for input, "-" will be interpreted as standard input

-o FILENAME
: Use FILENAME for ouput, "-" will be interpreted as standard output


# EXAMPLES

Used by typing into standard in (press Ctrl-d to end your input).

~~~shell
	{app_name}
	[
	  {"one": 1, "two": 2},
	  {"one": 10, "two": 20},
    ]
	^D
~~~

This should yield the following.

~~~text
one,two
1,2
10,20
~~~

These would get the file named "my_list.json" and save it as my.csv

~~~shell
    {app_name} my_list.json > my.csv

	{app_name} my_list.json my.csv

	cat my_list.json | {app_name} -i - > my.csv
~~~

{app_name} {version}

`

	// Standard Options
	showHelp         bool
	showLicense      bool
	showVersion      bool
	showExamples     bool
	inputFName       string
	outputFName      string
	quiet            bool
	newLine          bool
)


func main() {
	appName := path.Base(os.Args[0])
	version := datatools.Version
	license := datatools.LicenseText
	releaseDate := datatools.ReleaseDate
	releaseHash := datatools.ReleaseHash

	// Standard Options
	flag.BoolVar(&showHelp, "help", false, "display help")
	flag.BoolVar(&showLicense, "license", false, "display license")
	flag.BoolVar(&showVersion, "version", false, "display version")

	flag.StringVar(&outputFName, "o", "", "output filename")
	flag.StringVar(&outputFName, "output", "", "output filename")
	flag.BoolVar(&quiet, "quiet", false, "suppress error messages")
	flag.BoolVar(&newLine, "nl", false, "if true add a trailing newline")
	flag.BoolVar(&newLine, "newline", false, "if true add a trailing newline")

	// App Options
	showHeader, delimiter := true, ""
	flag.BoolVar(&showHeader, "use-header", showHeader, "display a header row in CSV output, default is true")
	flag.StringVar(&delimiter, "delimiter", delimiter, "set delimiter, default is comma")

	// Parse env and options
	flag.Parse()
	args := flag.Args()

	// Handle case of input/output filenames provided without -i, -o
	if len(args) > 0 {
		inputFName = args[0]
		if len(args) > 1 {
			outputFName = args[1]
		}
	}

	// Setup IO
	var err error

	in := os.Stdin
	out := os.Stdout
	eout := os.Stderr

	if inputFName != "" && inputFName != "-" {
		in, err = os.Open(inputFName)
		if err != nil {
			fmt.Fprintln(eout, err)
			os.Exit(1)
		}
		defer in.Close()
	}

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
		fmt.Fprintf(out, "%s\n", datatools.FmtHelp(helpText, appName, version, releaseDate, releaseHash))
		os.Exit(0)
	}
	if showLicense {
		fmt.Fprintf(out, "%s\n", license)
		os.Exit(0)
	}
	if showVersion {
		fmt.Fprintf(out, "datatools, %s %s %s\n", appName, version, releaseHash)
		os.Exit(0)
	}
	err = datatools.JSONObjectsToCSV(in, out, eout, quiet, showHeader, delimiter)
	if err != nil {
		fmt.Fprintln(eout, err)
		os.Exit(1)
	}
	eol := ""
	if newLine {
		eol = "\n"
	}
	fmt.Fprintf(out, "%s", eol)
}
