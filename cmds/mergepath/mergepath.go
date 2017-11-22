//
// mergepath.go - merge the path variable to avoid duplicates
//
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
	"os"
	"path"
	"strings"

	// CaltechLibrary packages
	"github.com/caltechlibrary/cli"
	"github.com/caltechlibrary/datatools"
)

var (
	usage = `USAGE: %s NEW_PATH_PARTS`

	description = `

SYNOPSIS

%s can merge the new path parts with the existing path with creating duplications.
It can also re-order existing path elements by prefixing or appending to the existing
path and removing the resulting duplicate.

`

	examples = `

EXAMPLE

	export PATH=$(%s -p $HOME/bin)

This would put your home bin directory at the beginning of your path.

`

	// Standard Options
	showHelp     bool
	showLicense  bool
	showVersion  bool
	showExamples bool

	// Application Specific Options

	envPath     string
	dir         string
	appendPath  = true
	prependPath = false
	clipPath    = false
)

func init() {
	const (
		pathUsage    = "The path you want to merge with."
		dirUsage     = "The directory you want to add to the path."
		appendUsage  = "Append the directory to the path removing any duplication"
		prependUsage = "Prepend the directory to the path removing any duplication"
		clipUsage    = "Remove a directory from the path"
		helpUsage    = "This help document."
	)
	flag.BoolVar(&showHelp, "h", false, "display help")
	flag.BoolVar(&showHelp, "help", false, "display help")
	flag.BoolVar(&showLicense, "l", false, "display license")
	flag.BoolVar(&showLicense, "license", false, "display license")
	flag.BoolVar(&showVersion, "v", false, "display version")
	flag.BoolVar(&showVersion, "version", false, "display version")
	flag.BoolVar(&showExamples, "example", false, "display example(s)")

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
	flag.Parse()
	args := flag.Args()

	// Configuration and command line interation
	cfg := cli.New(appName, strings.ToUpper(appName), datatools.Version)
	cfg.LicenseText = fmt.Sprintf(datatools.LicenseText, appName, datatools.Version)
	cfg.UsageText = fmt.Sprintf(usage, appName)
	cfg.DescriptionText = fmt.Sprintf(description, appName)
	cfg.OptionText = "OPTIONS"
	cfg.ExampleText = fmt.Sprintf(examples, appName)

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

	if flag.NArg() > 0 {
		dir = flag.Arg(0)
		if flag.NArg() == 2 {
			envPath = flag.Arg(1)
		}
	}

	if envPath == "$PATH" {
		envPath = os.Getenv("PATH")
	}
	if dir == "" {
		fmt.Fprintf(os.Stderr, "Missing directory to add to path")
		os.Exit(1)
	}
	if clipPath == true {
		fmt.Printf("%s", clip(envPath, dir))
		os.Exit(0)
	}
	if prependPath == true {
		appendPath = false
	}
	if strings.Contains(envPath, dir) {
		envPath = clip(envPath, dir)
	}
	if appendPath == true {
		fmt.Printf("%s:%s", envPath, dir)
	} else {
		fmt.Printf("%s:%s", dir, envPath)
	}
}
