//
// urldecode.go is a simple command line utility to decode a string in a URL friendly way.
//
// @author R. S. Doiel, <rsdoiel@caltech.edu>
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
	"io/ioutil"
	"net/url"
	"os"
	"path"
	"strings"

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

{app_name} [OPTIONS] [URL_ENCODED_STRING]

# DESCRIPTION

{app_name} is a simple command line utility to URL decode content. By default
it reads from standard input and writes to standard out.  You can
also specifty the string to decode as a command line parameter.

You can provide the URL encoded string as a command line parameter or if none
present it will be read from standard input.

# OPTIONS

-help
: display help

-license
: display license

-version
: display version

-query
: use query escape (pluses for spaces)

-newline
: Append a trailing newline

# EXAMPLE

echo 'This%20is%20the%20string%20to%20encode%20&%20nothing%20else%0A' | {app_name}

would yield (without the double quotes)

	"This is the string to encode & nothing else!" 

`

	// Standard Options
	showHelp         bool
	showLicense      bool
	showVersion      bool

	// App Options
	useQueryUnescape bool
	newLine bool
)

func main() {
	appName := path.Base(os.Args[0])
	version := datatools.Version
	license := datatools.LicenseText
	releaseDate := datatools.ReleaseDate
	releaseHash := datatools.ReleaseHash

	// Standard Options
	flag.BoolVar(&showHelp, "help", showHelp, "display help")
	flag.BoolVar(&showLicense, "license", showLicense, "display license")
	flag.BoolVar(&showVersion, "version", showVersion, "display version")
	flag.BoolVar(&newLine, "newline", false, "append a newline character")
	flag.BoolVar(&useQueryUnescape, "query", false, "use query escape (pluses for spaces)")
 
	// Parse environment and options
	flag.Parse()
	args := flag.Args()

	// Setup IO
	var err error

	in := os.Stdin
	out := os.Stdout
	eout := os.Stderr

	// Handle the default options
	if showHelp {
			fmt.Fprintf(out, "%s\n", datatools.FmtHelp(helpText, appName, version, releaseDate, releaseHash))
			os.Exit(0)
	}
	if showLicense == true {
			fmt.Fprintf(out, "%s\n", license)
			os.Exit(0)
	}
	if showVersion == true {
			fmt.Fprintf(out, "datatools, %s %s %s\n", appName, version, releaseHash)
			os.Exit(0)
	}

	nl := "\n"
	if newLine == false {
		nl = ""
	}

	var (
		src string
		s   string
	)

	if len(args) > 0 {
		src = strings.Join(args, " ")
	} else {
		buf, err := ioutil.ReadAll(in)
		if err != nil {
			fmt.Fprintf(eout, "%s\n", err)
			os.Exit(1)
		}
		src = fmt.Sprintf("%s", buf)
	}
	if useQueryUnescape {
		s, err = url.QueryUnescape(src)
	} else {
		s, err = url.PathUnescape(src)
	}
	if err != nil {
		fmt.Fprintf(eout, "%s\n", err)
		os.Exit(1)
	}
	fmt.Fprintf(out, "%s%s", s, nl)
}
