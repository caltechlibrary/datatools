// Generates a time in HH:MM:SS format based on a relative time
// description (e.g. -13h12m11s for minus 13 hours, 12 minutes and 11 seconds)
//
// @Author R. S. Doiel, <rsdoiel@caltech.edu>
//
// Copyright (c) 2022, Caltech
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
	"io"
	"os"
	"path"
	"strings"
	"time"

	// Local package
	"github.com/caltechlibrary/datatools"
)

const (
	// 24 hour notation
	timeFmt = "15:04:05"
)

var (
	helpText = `%{app_name}(1) user manual | version {version} {release_hash}
% R. S. Doiel
% {release_date}

# NAME

{app_name}

# SYNOPSIS

{app_name} RELATIVE_TIME_STRING

# DESCRIPTION

{app_name} provides a relative time string in the "HH:MM:SS" in 24 hour format.

The notation for the relative time string is based on Go's time duration string. From
https://golang.google.cn/pkg/time/#ParseDuration, 

> A duration string is a possibly signed sequence of decimal numbers, each with 
> optional fraction and a unit suffix, such as "300ms", "-1.5h" or "2h45m".
> Valid time units are "ns", "us" (or "µs"), "ms", "s", "m", "h".


# EXAMPLES

Get the dime ninety minutes in the past.

~~~shell
    {app_name} -- -90m
~~~

Get the time 24 hours ago

~~~shell
    {app_name} -- -24h
~~~

Get the time 16 hours, 23 minutes and 4 seconds in the future.

~~~shell
	{app_name} 16h23m4s
~~~

`

	licenseText = `{app_name} {version}

Copyright (c) 2022, Caltech
All rights not granted herein are expressly reserved by Caltech.

Redistribution and use in source and binary forms, with or without modification, are permitted provided that the following conditions are met:

1. Redistributions of source code must retain the above copyright notice, this list of conditions and the following disclaimer.

2. Redistributions in binary form must reproduce the above copyright notice, this list of conditions and the following disclaimer in the documentation and/or other materials provided with the distribution.

3. Neither the name of the copyright holder nor the names of its contributors may be used to endorse or promote products derived from this software without specific prior written permission.

THIS SOFTWARE IS PROVIDED BY THE COPYRIGHT HOLDERS AND CONTRIBUTORS "AS IS" AND ANY EXPRESS OR IMPLIED WARRANTIES, INCLUDING, BUT NOT LIMITED TO, THE IMPLIED WARRANTIES OF MERCHANTABILITY AND FITNESS FOR A PARTICULAR PURPOSE ARE DISCLAIMED. IN NO EVENT SHALL THE COPYRIGHT HOLDER OR CONTRIBUTORS BE LIABLE FOR ANY DIRECT, INDIRECT, INCIDENTAL, SPECIAL, EXEMPLARY, OR CONSEQUENTIAL DAMAGES (INCLUDING, BUT NOT LIMITED TO, PROCUREMENT OF SUBSTITUTE GOODS OR SERVICES; LOSS OF USE, DATA, OR PROFITS; OR BUSINESS INTERRUPTION) HOWEVER CAUSED AND ON ANY THEORY OF LIABILITY, WHETHER IN CONTRACT, STRICT LIABILITY, OR TORT (INCLUDING NEGLIGENCE OR OTHERWISE) ARISING IN ANY WAY OUT OF THE USE OF THIS SOFTWARE, EVEN IF ADVISED OF THE POSSIBILITY OF SUCH DAMAGE.
`

	// Standard Options
	showHelp    bool
	showVersion bool
	showLicense bool
	outputFName string
	newLine     bool
	eol         string
)

func assertOk(eout io.Writer, e error, failMsg string) {
	if e != nil {
		fmt.Fprintf(eout, " %s, %s", failMsg, e)
		os.Exit(1)
	}
}

func fmtText(txt string, appName string, version string) string {
	return strings.ReplaceAll(strings.ReplaceAll(txt, "{app_name}", appName), "{version}", version)
}

func main() {
	// Setup IO
	var (
		useTime   string
		startTime time.Time
		tm        time.Time
	)

	appName := path.Base(os.Args[0])

	// Standard Options
	flag.BoolVar(&showHelp, "help", false, "display help")
	flag.BoolVar(&showLicense, "license", false, "display license")
	flag.BoolVar(&showVersion, "version", false, "display version")
	flag.BoolVar(&newLine, "nl", true, "if true add a trailing newline")

	// App option
	flag.StringVar(&useTime, "time", "", "use time, formatted as HH:MM:SS format (24 hours)")

	// Parse env and options
	flag.Parse()
	args := flag.Args()

	out := os.Stdout
	eout := os.Stderr

	// Process Options
	if showHelp {
		fmt.Fprintln(out, fmtText(helpText, appName, datatools.Version))
		os.Exit(0)
	}
	if showLicense {
		fmt.Fprintln(out, fmtText(licenseText, appName, datatools.Version))
		os.Exit(0)
	}
	if showVersion {
		fmt.Fprintln(out, datatools.Version)
		os.Exit(0)
	}
	if newLine {
		eol = "\n"
	}

	if len(args) == 0 {
		fmt.Fprintln(eout, "Must include a duration to add/subtract relative time, e.g. -10h13m8s")
	}
	duration := strings.TrimSpace(args[0])

	offset, err := time.ParseDuration(duration)
	assertOk(eout, err, "Expected a time duration")

	if useTime == "" {
		startTime = time.Now()
	} else {
		startTime, err = time.Parse(timeFmt, useTime)
		assertOk(eout, err, "Did not understand time string.")
	}
	tm = startTime.Add(offset)

	fmt.Fprintf(out, "%s%s", tm.Format(timeFmt), eol)
}
