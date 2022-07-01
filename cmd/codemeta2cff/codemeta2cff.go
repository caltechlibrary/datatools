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
//
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

func usage(appName string, exitCode int) {
	out := os.Stderr
	if exitCode == 0 {
		out = os.Stdout
	}
	fmt.Fprintf(out, `
USAGE: %s

    %s [CODEMETA_JSON CITATION_CFF]

Reads codemeta.json file and writes CITATION.cff. By default
it assume both are in the current directory.  You can also
provide the name and path to both files.

OPTIONS

    -h, -help      display help

EXAMPLE

Generating the CITATION.cff from the codemeta.json file the current
working directory.

    %s

Specifying the full paths.

	%s /opt/local/myproject/codemeta.json /opt/local/myproject/CITATION.cff

datatools v%s
`, appName, appName, appName, appName, datatools.Version)
	os.Exit(exitCode)
}

func main() {
	// command line name and options support
	appName := path.Base(os.Args[0])
	help, version := false, false
	args := []string{}
	// Setup to parse command line
	flag.BoolVar(&help, "h", false, "display help")
	flag.BoolVar(&help, "help", false, "display help")
	flag.BoolVar(&version, "version", false, "display version")
	flag.Parse()

	args = flag.Args()

	// Process options and run report
	if help {
		usage(appName, 0)
	}
	if version {
		fmt.Printf("datatools, %s v%s\n", appName, datatools.Version)
		os.Exit(0)
	}
	codemeta, citation := "codemeta.json", "CITATION.cff"
	if len(args) == 2 {
		codemeta, citation = args[0], args[1]
	} else if len(args) == 1 {
		codemeta = args[0]
	} else if len(args) != 0 {
		fmt.Fprintf(os.Stderr, "Unexpected parameters: %q\n", strings.Join(os.Args, " "))
		usage(appName, 1)
	}
	err := datatools.CodemetaToCitationCff(codemeta, citation)
	if err != nil {
		fmt.Fprintf(os.Stderr, "ERROR: %s\n", err)
		os.Exit(1)
	}
}
