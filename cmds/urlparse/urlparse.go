//
// urlparse.go - a URL Parser library for use in Bash scripts.
// @Author R. S. Doiel, <rsdoiel@caltech.edu>
//
// Copyright (c) 2017, Caltech
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
	"net/url"
	"os"
	"path"
	"strings"

	// CaltechLibrary Packages
	"github.com/caltechlibrary/cli"
	"github.com/caltechlibrary/datatools"
)

var (
	usage = `USAGE: %s [OPTIONS] URL_TO_PARSE`

	description = `

SYNOPSIS

%s can parse a URL and return the specific elements
requested (e.g. protocol, hostname, path, query string)

`

	examples = `

EXAMPLE

With no options returns "http\texample.com\t/my/page.html"

    %s http://example.com/my/page.html

Get protocol. Returns "http".

    %s --protocol http://example.com/my/page.html

Get host or domain name.  Returns "example.com".

    %s --host http://example.com/my/page.html

Get path. Returns "/my/page.html".

    %s --path http://example.com/my/page.html

Get basename. Returns "page.html".

    %s --basename http://example.com/my/page.html

Get extension. Returns ".html".

    %s --extension http://example.com/my/page.html

Without options urlparse returns protocol, host and path
fields separated by a tab.

`

	// Standard Options
	showHelp     bool
	showLicense  bool
	showVersion  bool
	showExamples bool

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

func init() {
	const (
		delimiterUsage = "Set the output delimited for parsed display. (defaults to tab)"
		protocolUsage  = "Display the protocol of URL (defaults to http)"
		hostUsage      = "Display the hostname (and port if specified) found in URL."
		pathUsage      = "Display the path after the hostname."
		dirUsage       = "Display all but the last element of the path"
		basenameUsage  = "Display the base filename at the end of the path."
		extensionUsage = "Display the filename extension (e.g. .html)."
	)

	// Standard Options
	flag.BoolVar(&showHelp, "h", false, "display help")
	flag.BoolVar(&showHelp, "help", false, "display help")
	flag.BoolVar(&showLicense, "l", false, "display license")
	flag.BoolVar(&showLicense, "license", false, "display license")
	flag.BoolVar(&showVersion, "v", false, "display verison")
	flag.BoolVar(&showVersion, "version", false, "display version")
	flag.BoolVar(&showExamples, "example", false, "display example(s)")

	// App Specific Options
	flag.StringVar(&delimiter, "delimiter", delimiter, delimiterUsage)
	flag.StringVar(&delimiter, "D", delimiter, delimiterUsage)
	flag.BoolVar(&showProtocol, "protocol", false, protocolUsage)
	flag.BoolVar(&showProtocol, "P", false, protocolUsage)
	flag.BoolVar(&showHost, "host", false, hostUsage)
	flag.BoolVar(&showHost, "H", false, hostUsage)
	flag.BoolVar(&showPath, "path", false, pathUsage)
	flag.BoolVar(&showPath, "p", false, pathUsage)
	flag.BoolVar(&showDir, "directory", false, basenameUsage)
	flag.BoolVar(&showDir, "d", false, basenameUsage)
	flag.BoolVar(&showBase, "base", false, basenameUsage)
	flag.BoolVar(&showBase, "b", false, basenameUsage)
	flag.BoolVar(&showExtension, "extension", false, extensionUsage)
	flag.BoolVar(&showExtension, "e", false, extensionUsage)
}

func main() {
	appName := path.Base(os.Args[0])
	flag.Parse()
	args := flag.Args()

	// Configuration and command line interation
	cfg := cli.New(appName, strings.ToUpper(appName), datatools.Version)
	cfg.LicenseText = fmt.Sprintf(datatools.LicenseText, appName, datatools.Version)
	cfg.UsageText = fmt.Sprintf(usage, appName)
	cfg.DescriptionText = fmt.Sprintf(description, appName)
	cfg.OptionText = "OPTIONS"
	cfg.ExampleText = fmt.Sprintf(examples, appName, appName, appName, appName, appName, appName)

	if showHelp == true {
		if len(args) > 0 {
			fmt.Println(cfg.Help(args...))
		} else {
			fmt.Println(cfg.Usage())
		}
		os.Exit(0)
	}

	if showExamples == true {
		if len(args) > 0 {
			fmt.Println(cfg.Example(args...))
		} else {
			fmt.Println(cfg.ExampleText)
		}
		os.Exit(0)
	}

	if showLicense == true {
		fmt.Println(cfg.License())
		os.Exit(0)
	}

	if showVersion == true {
		fmt.Println(cfg.Version())
		os.Exit(0)
	}

	results := []string{}
	urlToParse := flag.Arg(0)
	if urlToParse == "" {
		fmt.Fprintf(os.Stderr, "Missing URL to parse")
		os.Exit(1)
	}
	u, err := url.Parse(urlToParse)
	if err != nil {
		fmt.Fprintf(os.Stderr, "%s", err)
		os.Exit(1)
	}

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
		fmt.Fprintf(os.Stdout, "%s%s%s%s%s",
			u.Scheme, useDelim, u.Host, useDelim, u.Path)
	} else {
		fmt.Fprint(os.Stdout, strings.Join(results, useDelim))
	}
}
