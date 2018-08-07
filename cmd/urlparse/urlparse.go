//
// urlparse - a URL Parser library for use in Bash scripts.
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
	"net/url"
	"os"
	"path"
	"strings"

	// CaltechLibrary Packages
	"github.com/caltechlibrary/cli"
	"github.com/caltechlibrary/datatools"
)

var (
	description = `
%s can parse a URL and return the specific elements
requested (e.g. protocol, hostname, path, query string)
`

	examples = `
With no options returns "http\texample.com\t/my/page.html"

    %s http://example.com/my/page.html

Get protocol. Returns "http".

    %s -protocol http://example.com/my/page.html

Get host or domain name.  Returns "example.com".

    %s -host http://example.com/my/page.html

Get path. Returns "/my/page.html".

    %s -path http://example.com/my/page.html

Get dirname. Returns "my"

    %s -dirname http://example.com/my/page.html

Get basename. Returns "page.html".

    %s -basename http://example.com/my/page.html

Get extension. Returns ".html".

    %s -extname http://example.com/my/page.html

Without options urlparse returns protocol, host and path
fields separated by a tab.
`

	// Standard Options
	showHelp         bool
	showLicense      bool
	showVersion      bool
	showExamples     bool
	inputFName       string
	outputFName      string
	generateMarkdown bool
	quiet            bool
	newLine          bool
	eol              string

	// App Specific Options
	showProtocol  bool
	showHost      bool
	showPort      bool
	showPath      bool
	showDir       bool
	showBase      bool
	showExtension bool
	showMimeType  bool
	envPrefix     = ""
	delimiter     = "\t"
)

func main() {
	const (
		delimiterUsage = "Set the output delimited for parsed display. (defaults to tab)"
		protocolUsage  = "Display the protocol of URL (defaults to http)"
		hostUsage      = "Display the hostname (and port if specified) found in URL."
		pathUsage      = "Display the path after the hostname."
		dirnameUsage   = "Display all but the last element of the path"
		basenameUsage  = "Display the base filename at the end of the path."
		extnameUsage   = "Display the filename extension (e.g. .html)."
	)

	app := cli.NewCli(datatools.Version)
	appName := app.AppName()

	// Add Help Docs
	app.AddHelp("license", []byte(fmt.Sprintf(datatools.LicenseText, appName, datatools.Version)))
	app.AddHelp("description", []byte(fmt.Sprintf(description, appName)))
	app.AddHelp("examples", []byte(fmt.Sprintf(examples, appName, appName, appName, appName, appName, appName, appName)))

	// Document non-option parameters
	app.AddParams("URL_TO_PARSE")

	// Standard Options
	app.BoolVar(&showHelp, "h,help", false, "display help")
	app.BoolVar(&showLicense, "l,license", false, "display license")
	app.BoolVar(&showVersion, "v,version", false, "display version")
	app.BoolVar(&showExamples, "examples", false, "display example(s)")
	app.StringVar(&inputFName, "i,input", "", "input filename")
	app.StringVar(&outputFName, "o,output", "", "output filename")
	app.BoolVar(&generateMarkdown, "generate-markdown", false, "generate markdown documentation")
	app.BoolVar(&quiet, "quiet", false, "suppress error messages")
	app.BoolVar(&newLine, "nl,newline", false, "if true add a trailing newline")

	// App Specific Options
	app.StringVar(&delimiter, "d,delimiter", delimiter, delimiterUsage)
	app.BoolVar(&showProtocol, "P,protocol", false, protocolUsage)
	app.BoolVar(&showHost, "H,host", false, hostUsage)
	app.BoolVar(&showPath, "p,path", false, pathUsage)
	app.BoolVar(&showDir, "dir,dirname", false, dirnameUsage)
	app.BoolVar(&showBase, "base,basename", false, basenameUsage)
	app.BoolVar(&showExtension, "ext,extname", false, extnameUsage)

	// Setup IO
	var err error
	app.Eout = os.Stderr

	app.In, err = cli.Open(inputFName, os.Stdin)
	cli.ExitOnError(app.Eout, err, quiet)
	defer cli.CloseFile(inputFName, app.In)

	app.Out, err = cli.Create(outputFName, os.Stdout)
	cli.ExitOnError(app.Eout, err, quiet)
	defer cli.CloseFile(outputFName, app.Out)

	// Parse env and options
	app.Parse()
	args := app.Args()

	// Process Options
	if generateMarkdown {
		app.GenerateMarkdown(app.Out)
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

	results := []string{}
	urlToParse := app.Arg(0)
	if urlToParse == "" {
		cli.ExitOnError(app.Eout, fmt.Errorf("Missing URL to parse"), quiet)
	}
	u, err := url.Parse(urlToParse)
	cli.ExitOnError(app.Eout, err, quiet)

	useDelim := delimiter
	if showProtocol == true {
		results = append(results, u.Scheme)
	}
	if showHost == true {
		results = append(results, u.Host)
	}
	if showPath == true {
		results = append(results, u.Path)
	}
	if showBase == true {
		results = append(results, path.Base(u.Path))
	}
	if showDir == true {
		results = append(results, path.Dir(u.Path))
	}
	if showExtension == true {
		results = append(results, path.Ext(u.Path))
	}

	if len(results) == 0 {
		fmt.Fprintf(app.Out, "%s%s%s%s%s%s",
			u.Scheme, useDelim, u.Host, useDelim, u.Path, eol)
	} else {
		fmt.Fprintf(app.Out, "%s%s", strings.Join(results, useDelim), eol)
	}
}
