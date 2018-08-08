//
// mergepath.go - merge the path variable to avoid duplicates
//
// @Author R. S. Doiel, <rsdoiel@caltech.edu>
//
// Copyright (c) 2018, Caltech
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
	"fmt"
	"os"
	"strings"

	// CaltechLibrary packages
	"github.com/caltechlibrary/cli"
	"github.com/caltechlibrary/datatools"
)

var (
	description = `
%s can merge the new path parts with the existing path with creating duplications.
It can also re-order existing path elements by prefixing or appending to the existing
path and removing the resulting duplicate.
`

	examples = `
This would put your home bin directory at the beginning of your path.

	export PATH=$(%s -p $HOME/bin)
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
	envPath     string
	dir         string
	appendPath  = true
	prependPath = false
	clipPath    = false
)

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
	const (
		pathUsage    = "The path you want to merge with."
		dirUsage     = "The directory you want to add to the path."
		appendUsage  = "Append the directory to the path removing any duplication"
		prependUsage = "Prepend the directory to the path removing any duplication"
		clipUsage    = "Remove a directory from the path"
		helpUsage    = "This help document."
	)

	app := cli.NewCli(datatools.Version)
	appName := app.AppName()

	// Add Help Docs
	app.AddHelp("license", []byte(fmt.Sprintf(datatools.LicenseText, appName, datatools.Version)))
	app.AddHelp("description", []byte(fmt.Sprintf(description, appName)))
	app.AddHelp("examples", []byte(fmt.Sprintf(examples, appName)))

	// Document non-option parameters
	app.AddParams("NEW_PATH_PARTS")

	// Standard Options
	app.BoolVar(&showHelp, "h,help", false, "display help")
	app.BoolVar(&showLicense, "l,license", false, "display license")
	app.BoolVar(&showVersion, "v,version", false, "display version")
	app.BoolVar(&showExamples, "examples", false, "display example(s)")
	app.BoolVar(&generateMarkdown, "generate-markdown", false, "generate markdown documentation")
	app.BoolVar(&generateManPage, "generate-manpage", false, "generate man page")
	app.BoolVar(&quiet, "quiet", false, "suppress error messages")
	app.BoolVar(&newLine, "nl,newline", false, "if true add a trailing newline")

	// App Options
	envPath = "$PATH"
	app.StringVar(&envPath, "e,envpath", envPath, pathUsage)
	app.StringVar(&dir, "d,directory", dir, dirUsage)
	app.BoolVar(&appendPath, "a,append", appendPath, appendUsage)
	app.BoolVar(&prependPath, "p,prepend", prependPath, prependUsage)
	app.BoolVar(&clipPath, "c,clip", clipPath, clipUsage)

	// Parse env and options
	app.Parse()
	args := app.Args()

	// Setup IO
	var err error

	app.Eout = os.Stderr

	/* NOTE: we don't read from stdin
	app.In, err = cli.Open(inputFName, os.Stdin)
	cli.ExitOnError(app.Eout, err, quiet)
	defer cli.CloseFile(inputFName, app.In)
	*/

	app.Out, err = cli.Create(outputFName, os.Stdout)
	cli.ExitOnError(app.Eout, err, quiet)
	defer cli.CloseFile(outputFName, app.Out)

	// Process Options
	if generateMarkdown {
		app.GenerateMarkdown(app.Out)
		os.Exit(0)
	}
	if generateManPage {
		app.GenerateManPage(app.Out)
		os.Exit(0)
	}
	if showHelp || showExamples {
		if len(args) > 0 {
			fmt.Fprintln(app.Out, app.Help(args...))
		} else {
			app.Usage(app.Out)
		}
		os.Exit(0)
	}
	if showLicense {
		fmt.Fprintln(app.Out, app.License())
		os.Exit(0)
	}
	if showVersion {
		fmt.Fprintln(app.Out, app.Version())
		os.Exit(0)
	}
	if newLine {
		eol = "\n"
	}

	if app.NArg() > 0 {
		dir = app.Arg(0)
		if app.NArg() == 2 {
			envPath = app.Arg(1)
		}
	}

	if envPath == "$PATH" {
		envPath = os.Getenv("PATH")
	}
	if dir == "" {
		cli.ExitOnError(app.Eout, fmt.Errorf("Missing directory to add to path"), quiet)
	}
	if clipPath {
		fmt.Fprintf(app.Out, "%s%s", clip(envPath, dir), eol)
		os.Exit(0)
	}
	if prependPath {
		appendPath = false
	}
	if strings.Contains(envPath, dir) {
		envPath = clip(envPath, dir)
	}
	if appendPath {
		fmt.Fprintf(app.Out, "%s:%s%s", envPath, dir, eol)
	} else {
		fmt.Fprintf(app.Out, "%s:%s%s", dir, envPath, eol)
	}
}
