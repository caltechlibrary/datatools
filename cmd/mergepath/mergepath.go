// mergepath.go - merge the path variable to avoid duplicates
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
	"flag"
	"fmt"
	"os"
	"path"
	"strings"

	// CaltechLibrary packages
	"github.com/caltechlibrary/datatools"
)

const (
	helpText = `%{app_name}(1) user manual | version {version} {release_hash}
% R. S. Doiel
% {release_date}

# NAME

{app_name}

# SYNOPSIS

{app_name} [OPTIONS] NEW_PATH_PARTS

# DESCRIPTION

{app_name} can merge the new path parts with the existing path with
creating duplications.  It can also re-order existing path elements by
prefixing or appending to the existing path and removing the resulting
duplicate.

# OPTIONS

-help
: display help

-license
: display license

-version
: display version

-a, -append
: Append the directory to the path removing any duplication

-c, -clip
: Remove a directory from the path

-d, -directory
: The directory you want to add to the path.

-e, -envpath
: The path you want to merge with.

-nl, -newline
: if true add a trailing newline

-p, -prepend
: Prepend the directory to the path removing any duplication

-quiet
: suppress error messages


# EXAMPLES

This would put your home bin directory at the beginning of your path.

~~~
	export PATH=$({app_name} -p $HOME/bin)
~~~

{app_name} {version}
`

	pathUsage    = "The path you want to merge with."
	dirUsage     = "The directory you want to add to the path."
	appendUsage  = "Append the directory to the path removing any duplication"
	prependUsage = "Prepend the directory to the path removing any duplication"
	clipUsage    = "Remove a directory from the path"
	helpUsage    = "This help document."
)

var (
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
	envPath     string
	dir         string
	appendPath  = true
	prependPath = false
	clipPath    = false
)

func fmtTxt(src string, appName string, version string) string {
	return strings.ReplaceAll(strings.ReplaceAll(src, "{app_name}", appName), "{version}", version)
}

func clip(envPath string, dir string) string {
	oParts := []string{}
	iParts := strings.Split(envPath, ":")
	for _, v := range iParts {
		if v != dir {
			oParts = append(oParts, v)
		}
	}
	return strings.Join(oParts, ":")
}

func main() {
	appName := path.Base(os.Args[0])

	// Standard Options
	flag.BoolVar(&showHelp, "help", false, "display help")
	flag.BoolVar(&showLicense, "license", false, "display license")
	flag.BoolVar(&showVersion, "version", false, "display version")

	flag.BoolVar(&quiet, "quiet", false, "suppress error messages")
	flag.BoolVar(&newLine, "nl", false, "if true add a trailing newline")
	flag.BoolVar(&newLine, "newline", false, "if true add a trailing newline")

	// App Options
	envPath = "$PATH"
	flag.StringVar(&envPath, "e", envPath, pathUsage)
	flag.StringVar(&envPath, "envpath", envPath, pathUsage)
	flag.StringVar(&dir, "d", dir, dirUsage)
	flag.StringVar(&dir, "directory", dir, dirUsage)
	flag.BoolVar(&appendPath, "a", appendPath, appendUsage)
	flag.BoolVar(&appendPath, "append", appendPath, appendUsage)
	flag.BoolVar(&prependPath, "p", prependPath, prependUsage)
	flag.BoolVar(&prependPath, "prepend", prependPath, prependUsage)
	flag.BoolVar(&clipPath, "c", clipPath, clipUsage)
	flag.BoolVar(&clipPath, "clip", clipPath, clipUsage)

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

	// Process Options
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

	if len(args) > 0 {
		dir = args[0]
		if len(args) == 2 {
			envPath = args[1]
		}
	}

	if envPath == "$PATH" {
		envPath = os.Getenv("PATH")
	}
	if dir == "" {
		fmt.Fprintln(eout, "Missing directory to add to path")
		os.Exit(1)
	}
	if clipPath {
		fmt.Fprintf(out, "%s%s", clip(envPath, dir), eol)
		os.Exit(0)
	}
	if prependPath {
		appendPath = false
	}
	if strings.Contains(envPath, dir) {
		envPath = clip(envPath, dir)
	}
	if appendPath {
		fmt.Fprintf(out, "%s:%s%s", envPath, dir, eol)
	} else {
		fmt.Fprintf(out, "%s:%s%s", dir, envPath, eol)
	}
}
