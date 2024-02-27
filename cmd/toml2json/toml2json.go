// toml2json is a command line utility that converts an TOML
// to JSON.
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
	"io"
	"os"
	"path"

	// CaltechLibrary packages
	"github.com/caltechlibrary/datatools"

	// 3rd Party packages
	"github.com/BurntSushi/toml"
)

var (
	helpText = `%{app_name}(1) user manual | version {version} {release_hash}
% R. S. Doiel
% {release_date}

# NAME
 
{app_name}

# SYNOPSIS

{app_name} [OPTIONS] [TOML_FILENAME] [JSON_NAME]

# DESCRIPTION

{app_name} is a tool that converts TOML into JSON. It operates
on standard input and writes to standard output.

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

-p, -pretty
: pretty print output

-quiet
: suppress error messages


# EXAMPLES

These would get the file named "my.toml" and save it as my.json

~~~
    {app_name} my.toml > my.json

    {app_name} my.toml my.json

	cat my.toml | {app_name} -i - > my.json
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
	generateMarkdown bool
	generateManPage  bool
	quiet            bool
	newLine          bool
	eol              string

	// Application Options
	prettyPrint bool
)


func toml2JSON(in io.Reader, out io.Writer, printPrint bool) error {
	var (
		src []byte
		err error
	)
	m := map[string]interface{}{}
	_, err = toml.DecodeReader(in, &m)
	if err != nil {
		return err
	}
	if prettyPrint == true {
		src, err = datatools.JSONMarshalIndent(m, "", "    ")
	} else {
		src, err = datatools.JSONMarshal(m)
	}
	fmt.Fprintf(out, "%s", src)
	return nil
}

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

	// App Specific Options
	flag.BoolVar(&prettyPrint, "p", false, "pretty print output")
	flag.BoolVar(&prettyPrint, "pretty", false, "pretty print output")

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
	if newLine {
		eol = "\n"
	}

	err = toml2JSON(in, out, prettyPrint)
	if err != nil {
		fmt.Fprintln(eout, err)
		os.Exit(1)
	}
	fmt.Fprintf(out, "%s", eol)
}
