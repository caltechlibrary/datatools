// codemeta2cff.go converts a codemeta.json file to CITATION.cff.
//
// Author: R. S. Doiel <rsdoiel@caltech.edu>
//
// Copyright (c) 2021, Caltech
// All rights not granted herein are expressly reserved by Caltech.
//
// Redistribution and use in source and binary forms, with or without
// modification, are permitted provided that the following conditions are met:
//
// 1. Redistributions of source code must retain the above copyright notice,
// this list of conditions and the following disclaimer.
//
// 2. Redistributions in binary form must reproduce the above copyright notice,
// this list of conditions and the following disclaimer in the documentation
// and/or other materials provided with the distribution.
//
// 3. Neither the name of the copyright holder nor the names of its contributors
// may be used to endorse or promote products derived from this software without
// specific prior written permission.
//
// THIS SOFTWARE IS PROVIDED BY THE COPYRIGHT HOLDERS AND CONTRIBUTORS "AS IS"
// AND ANY EXPRESS OR IMPLIED WARRANTIES, INCLUDING, BUT NOT LIMITED TO, THE
// IMPLIED WARRANTIES OF MERCHANTABILITY AND FITNESS FOR A PARTICULAR PURPOSE
// ARE DISCLAIMED. IN NO EVENT SHALL THE COPYRIGHT HOLDER OR CONTRIBUTORS BE
// LIABLE FOR ANY DIRECT, INDIRECT, INCIDENTAL, SPECIAL, EXEMPLARY, OR
// CONSEQUENTIAL DAMAGES (INCLUDING, BUT NOT LIMITED TO, PROCUREMENT OF
// SUBSTITUTE GOODS OR SERVICES; LOSS OF USE, DATA, OR PROFITS; OR BUSINESS
// INTERRUPTION) HOWEVER CAUSED AND ON ANY THEORY OF LIABILITY, WHETHER IN
// CONTRACT, STRICT LIABILITY, OR TORT (INCLUDING NEGLIGENCE OR OTHERWISE)
// ARISING IN ANY WAY OUT OF THE USE OF THIS SOFTWARE, EVEN IF ADVISED OF THE
// POSSIBILITY OF SUCH DAMAGE.
package main

import (
	"flag"
	"fmt"
	"os"
	"path"
	"strings"

	// Caltech Library Package
	"github.com/caltechlibrary/datatools"
)

var (
	helpText = `%{app_name}(1) user manual | version {version} {release_hash}
% R. S. Doiel
% {release_date}

# NAME

{app_name}

# SYSNOPSIS

{app_name} [OPTIONS] [CODEMETA_JSON CITATION_CFF]

# DESCRIPTION

Reads codemeta.json file and writes CITATION.cff. By default
it assume both are in the current directory.  You can also
provide the name and path to both files.

# OPTIONS

-help
: display help

# EXAMPLE

Generating the CITATION.cff from the codemeta.json file the current
working directory.

~~~
{app_name}
~~~

Specifying the full paths.

~~~
{app_name} /opt/local/myproject/codemeta.json /opt/local/myproject/CITATION.cff
~~~

`
)

func main() {

	// command line name and options support
	appName := path.Base(os.Args[0])
	version := datatools.Version
	license := datatools.LicenseText
	releaseDate := datatools.ReleaseDate
	releaseHash := datatools.ReleaseHash

	showHelp, showVersion, showLicense := false, false, false
	// Setup to parse command line
	flag.BoolVar(&showHelp, "help", false, "display help")
	flag.BoolVar(&showLicense, "license", false, "display license")
	flag.BoolVar(&showVersion, "version", false, "display version")
	flag.Parse()

	args := flag.Args()

	//in := os.Stdin
	out := os.Stdout
	eout := os.Stderr

	// Process options and run report
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
	codemeta, citation := "codemeta.json", "CITATION.cff"
	if len(args) == 2 {
		codemeta, citation = args[0], args[1]
	} else if len(args) == 1 {
		codemeta = args[0]
	} else if len(args) != 0 {
		fmt.Fprintf(eout, "Unexpected parameters: %q\n", strings.Join(os.Args, " "))
		os.Exit(1)
	}
	err := datatools.CodemetaToCitationCff(codemeta, citation)
	if err != nil {
		fmt.Fprintf(eout, "ERROR: %s\n", err)
		os.Exit(1)
	}
}
