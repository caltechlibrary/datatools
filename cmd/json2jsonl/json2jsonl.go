// jsonarray2jsonl - reads an JSON array document and renders a JSON lines document.
//
// @author R. S. Doiel, <rsdoiel@caltech.edu>
//
// Copyright (c) 2025, Caltech
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
	"bytes"
	"encoding/json"
	"flag"
	"fmt"
	"io"
	"os"
	"path"

	// My packages
	"github.com/caltechlibrary/datatools"
)

var (
	helpText = `%{app_name}(1) user manual | version {version} {release_hash}
% R. S. Doiel
% {release_date}

# NAME

{app_name}

# SYNOPSIS

{app_name} [OPTIIONS]

# DESCRIPTION

{app_name} reads a JSON array document rending the results as a JSON lines document.
See https://jsonlines.org.

# OPTIONS

-help
: display help

-license
: display license

-version
: display version

-i, -input
: input filename

-o, -output
: output filename

-quiet
: suppress error output

-as-dataset ATTRIBUTE_NAME
: if ATTRIBUTE_NAME is not empty string, find the top level ATTRIBUTE_NAME and use as key value for 
when generating a dataset load compatible version of the JSON array contents.

# EXAMPLES

Convert data1.json containing an array of objects into data1.jsonl using Unix redirection.

~~~
    {app_name} < data1.json > data1.jsonl
~~~

`

	// Standard Options
	showHelp    bool
	showLicense bool
	showVersion bool
	inputFName  string
	outputFName string
	quiet       bool

	// Application Options
	asDatasetKey  string
)

func getKey(obj map[string]interface{}, attrName string) (string, error) {
	if val, ok := obj[attrName].(string); ok {
		return val, nil
	}
	return "", fmt.Errorf("missing %q in object", attrName)
}

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
	flag.StringVar(&inputFName, "i", "", "input filename")
	flag.StringVar(&inputFName, "input", "", "input filename")
	flag.StringVar(&outputFName, "o", "", "output filename")
	flag.StringVar(&outputFName, "output", "", "output filename")
	flag.BoolVar(&quiet, "quiet", false, "suppress error output")

	// App Options
	flag.StringVar(&asDatasetKey, "as-dataset", "", "generate a dataset compatible JSON lines using the top level attribute named in object as key")

	// Parse environment and options
	flag.Parse()

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
	if showLicense == true {
		fmt.Fprintf(out, "%s\n", license)
		os.Exit(0)
	}
	if showVersion == true {
		fmt.Fprintf(out, "datatools, %s %s %s\n", appName, version, releaseHash)
		os.Exit(0)
	}

	// Read in the JSON document containing the array.
	src, err := io.ReadAll(in)
	if err != nil {
		fmt.Fprintf(eout, "%s\n", err)
		os.Exit(2)
	}
	// Make sure we have an array not an object.
	if (! bytes.HasPrefix(src, []byte("["))) {
		fmt.Fprintln(eout, "input must be a JSON array, aborting")
		os.Exit(3)
	}
	jsonObjects := []map[string]interface{}{}
	err = json.Unmarshal(src, &jsonObjects)
	if err != nil {
		fmt.Fprintf(eout, "failed to parse JSON array, %s\n", err)
		os.Exit(4)
	}
	// Iterate over array writing output objects one per line.
	hasError := false
	for i, val := range jsonObjects {
		src, err = json.Marshal(val)
		if err != nil && ! quiet {
			fmt.Fprintf(eout, "failed to marshal index %d in json array\n", i)
			hasError = true
		} else {
			if asDatasetKey == "" {
				fmt.Fprintf(out, "%s\n", src)
			} else {
				key, err := getKey(val, asDatasetKey)
				if err != nil {
					fmt.Fprintf(eout, "failed to find key for %d in json array, skipping")
					hasError = true
				} else {
					fmt.Fprintf(out, "{%q:%q,%q:%s}\n", "key", key, "object", src)
				}
			}
		}
	}

	if hasError == true {
		os.Exit(10)
	}
}
