// urlparse - a URL Parser library for use in Bash scripts.
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

{app_name} [OPTIONS] URL_TO_PARSE

# DESCRIPTION

{app_name} can parse a URL and return the specific elements
requested (e.g. protocol, hostname, path, query string)

# OPTIONS

-help
: display help

-license
: display license

-version
: display version

-H, -host
: Display the hostname (and port if specified) found in URL.

-P, -protocol
: Display the protocol of URL (defaults to http)

-base, -basename
: Display the base filename at the end of the path.

-d, -delimiter
: Set the output delimited for parsed display. (defaults to tab)

-dir, -dirname
: Display all but the last element of the path

-ext, -extname
: Display the filename extension (e.g. .html).

-i, -input
: input filename

-nl, -newline
: if true add a trailing newline

-o, -output
: output filename

-p, -path
: Display the path after the hostname.

-quiet
: suppress error messages


# EXAMPLES

With no options returns "http\texample.com\t/my/page.html"

~~~
    {app_name} http://example.com/my/page.html
~~~

Get protocol. Returns "http".

~~~
    {app_name} -protocol http://example.com/my/page.html
~~~

Get host or domain name.  Returns "example.com".

~~~
    {app_name} -host http://example.com/my/page.html
~~~

Get path. Returns "/my/page.html".

~~~
    {app_name} -path http://example.com/my/page.html
~~~

Get dirname. Returns "my"

~~~
    {app_name} -dirname http://example.com/my/page.html
~~~

Get basename. Returns "page.html".

~~~
    {app_name} -basename http://example.com/my/page.html
~~~

Get extension. Returns ".html".

~~~
    {app_name} -extname http://example.com/my/page.html
~~~

Without options {app_name} returns protocol, host and path
fields separated by a tab.

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

	appName := path.Base(os.Args[0])
	version := datatools.Version
	license := datatools.LicenseText
	releaseDate := datatools.ReleaseDate
	releaseHash := datatools.ReleaseHash

	// Standard Options
	flag.BoolVar(&showHelp, "help", false, "display help")
	flag.BoolVar(&showLicense, "license", false, "display license")
	flag.BoolVar(&showVersion, "version", false, "display version")

	flag.StringVar(&inputFName, "i", "", "input filename")
	flag.StringVar(&inputFName, "input", "", "input filename")
	flag.StringVar(&outputFName, "o", "", "output filename")
	flag.StringVar(&outputFName, "output", "", "output filename")

	flag.BoolVar(&quiet, "quiet", false, "suppress error messages")
	flag.BoolVar(&newLine, "nl", false, "if true add a trailing newline")
	flag.BoolVar(&newLine, "newline", false, "if true add a trailing newline")

	// App Specific Options
	flag.StringVar(&delimiter, "d", delimiter, delimiterUsage)
	flag.StringVar(&delimiter, "delimiter", delimiter, delimiterUsage)
	flag.BoolVar(&showProtocol, "P", false, protocolUsage)
	flag.BoolVar(&showProtocol, "protocol", false, protocolUsage)
	flag.BoolVar(&showHost, "H", false, hostUsage)
	flag.BoolVar(&showHost, "host", false, hostUsage)
	flag.BoolVar(&showPath, "p", false, pathUsage)
	flag.BoolVar(&showPath, "path", false, pathUsage)
	flag.BoolVar(&showDir, "dir", false, dirnameUsage)
	flag.BoolVar(&showDir, "dirname", false, dirnameUsage)
	flag.BoolVar(&showBase, "base", false, basenameUsage)
	flag.BoolVar(&showBase, "basename", false, basenameUsage)
	flag.BoolVar(&showExtension, "ext", false, extnameUsage)
	flag.BoolVar(&showExtension, "extname", false, extnameUsage)

	// Parse env and options
	flag.Parse()
	args := flag.Args()

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

	// Process Options
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

	argc := len(args)
	results := []string{}
	if argc == 0 {
		fmt.Fprintln(eout, "Missing URL to parse")
		os.Exit(1)
	}
	urlToParse := args[0]
	u, err := url.Parse(urlToParse)
	if err != nil {
		fmt.Fprintln(eout, err)
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
		fmt.Fprintf(out, "%s%s%s%s%s%s",
			u.Scheme, useDelim, u.Host, useDelim, u.Path, eol)
	} else {
		fmt.Fprintf(out, "%s%s", strings.Join(results, useDelim), eol)
	}
}
