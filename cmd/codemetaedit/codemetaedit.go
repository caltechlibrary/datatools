// codemetaedit.go is a command line edit for working with a codemeta.json file.
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

    %s ACTION [OPTION] [CODEMETA_JSON]

The tool is for updating a codemeta.json file from the command
line.  It assumings the codemeta.json file is in the current
directory unless a name and path are provided. There are 
four "actions" that can be performed.


ACTIONS

   help            display help
   show            show the value of a field given a json path
   set             set the value of a field given a json path
   delete          removes a a value given a json path

EXAMPLE

Retreive the version number in the codemeta.json file.

	%s show ".version"

Set the "@id" of the first author in the codemeta.json file.

    %s set '.author[0]["@id"]' 'https://orcid.org/0000-0003-0900-6903'

Remove the third author from the codemeta.json file.

    %s delete '.author[2]'

datatools v%s
`, appName, appName, appName, appName, appName, datatools.Version)
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
	if len(args) < 1 {
		usage(appName, 1)
	}
	var err error
	action := args[0]
	switch action {
	case "help":
		usage(appName, 0)
	case "show":
		err = datatools.CodemetaShow(args[1:]...)
	case "set":
		err = datatools.CodemetaSet(args[1:]...)
	case "delete":
		err = datatools.CodemetaDelete(args[1:]...)
	default:
		fmt.Fprintf(os.Stderr, "Unsupported action.\n")
		usage(appName, 1)
	}
	if err != nil {
		fmt.Fprintf(os.Stderr, "ERROR: %s\n", err)
		os.Exit(1)
	}
}
