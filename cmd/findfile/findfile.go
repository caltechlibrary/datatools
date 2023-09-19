//
// findfile - a simple directory tree walker that looks for files
// by name, basename or extension. Basically a unix "find" light to
// demonstrate walking the file system
//
// @author R. S. Doiel, <rsdoiel@caltech.edu>
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
//
package main

import (
	"flag"
	"fmt"
	"io"
	"os"
	"path"
	"path/filepath"
	"strings"
	"time"

	// Caltech Library Packages
	"github.com/caltechlibrary/datatools"
)

var (
	helpText = `---
title: "{app_name} (1) user manual"
author: "R. S. Doiel"
pubDate: 2023-01-06
---

# NAME

{app_name}

# SYNOPSIS

{app_name} [OPTIONS] [TARGET] [DIRECTORIES_TO_SEARCH]

# DESCRIPTION

{app_name} finds files based on matching prefix, suffix or contained text in base filename.

# OPTIONS

-help
: display this help message

-license
: display license information

-version
: display version message

-c, -contains
: find file(s) based on basename containing text

-d, -depth
: Limit depth of directories walked

-error, -stop-on-error
: Stop walk on file system errors (e.g. permissions)

-f, -full-path
: list full path for files found

-m, -mod-time
: display file modification time before the path

-nl, -newline
: if true add a trailing newline

-o, -output
: output filename

-p, -prefix
: find file(s) based on basename prefix

-quiet
: suppress error messages

-s, -suffix
: find file(s) based on basename suffix


# EXAMPLES

Search the current directory and subdirectories for Markdown files with extension of ".md".

~~~
	{app_name} -s .md
~~~

{app_name} {version}

`

	// Standard Options
	showHelp         bool
	showVersion      bool
	showLicense      bool
	showExamples     bool
	outputFName      string
	generateMarkdown bool
	generateManPage  bool
	quiet            bool

	// App Specific Options
	showModificationTime bool
	findPrefix           bool
	findContains         bool
	findSuffix           bool
	findAll              bool
	stopOnErrors         bool
	outputFullPath       bool
	optDepth             int
	pathSep              string
)

func fmtTxt(src string, appName string, version string) string {
	return strings.ReplaceAll(strings.ReplaceAll(src, "{app_name}", appName), "{version}", version)
}


func display(w io.Writer, docroot, p string, m time.Time) {
	var s string
	if outputFullPath == true {
		s, _ = filepath.Abs(p)
	} else {
		s, _ = filepath.Rel(docroot, p)
	}
	if showModificationTime == true {
		fmt.Fprintf(w, "%s %s\n", m.Format("2006-01-02 15:04:05 -0700"), s)
		return
	}
	fmt.Fprintf(w, "%s\n", s)
}

func walkPath(out io.Writer, docroot string, target string) error {
	return filepath.Walk(docroot, func(p string, info os.FileInfo, err error) error {
		if err != nil {
			if stopOnErrors == true {
				return fmt.Errorf("Can't read %s, %s", p, err)
			}
			return nil
		}
		// Limit depth of walk by directory
		if optDepth > 0 {
			d := filepath.Dir(p)
			currentDepth := strings.Count(d, pathSep) + 1
			if currentDepth > optDepth {
				return filepath.SkipDir
			}
		}
		s := filepath.Base(p)
		// If a regular file then apply rules for display
		if info.Mode().IsRegular() == true {
			switch {
			case findPrefix == true && strings.HasPrefix(s, target) == true:
				display(out, docroot, p, info.ModTime())
			case findSuffix == true && strings.HasSuffix(s, target) == true:
				display(out, docroot, p, info.ModTime())
			case findContains == true && strings.Contains(s, target) == true:
				display(out, docroot, p, info.ModTime())
			case strings.Compare(s, target) == 0:
				display(out, docroot, p, info.ModTime())
			case findAll == true:
				display(out, docroot, p, info.ModTime())
			}
		}
		return nil
	})
}

func main() {
	appName := path.Base(os.Args[0])

	// Standard Options
	flag.BoolVar(&showHelp, "help", false, "display this help message")
	flag.BoolVar(&showLicense, "license", false, "display license information")
	flag.BoolVar(&showVersion, "version", false, "display version message")

	flag.StringVar(&outputFName, "o", "", "output filename")
	flag.StringVar(&outputFName, "output", "", "output filename")
	flag.BoolVar(&quiet, "quiet", false, "suppress error messages")

	// App Specific Options
	flag.BoolVar(&showModificationTime, "m", false, "display file modification time before the path")
	flag.BoolVar(&showModificationTime, "mod-time", false, "display file modification time before the path")
	flag.BoolVar(&stopOnErrors, "error", false, "Stop walk on file system errors (e.g. permissions)")
	flag.BoolVar(&stopOnErrors, "stop-on-error", false, "Stop walk on file system errors (e.g. permissions)")
	flag.BoolVar(&findPrefix, "p", false, "find file(s) based on basename prefix")
	flag.BoolVar(&findPrefix, "prefix", false, "find file(s) based on basename prefix")
	flag.BoolVar(&findContains, "c", false, "find file(s) based on basename containing text")
	flag.BoolVar(&findContains, "contains", false, "find file(s) based on basename containing text")
	flag.BoolVar(&findSuffix, "s", false, "find file(s) based on basename suffix")
	flag.BoolVar(&findSuffix, "suffix", false, "find file(s) based on basename suffix")
	flag.BoolVar(&outputFullPath, "f", false, "list full path for files found")
	flag.BoolVar(&outputFullPath, "full-path", false, "list full path for files found")
	flag.IntVar(&optDepth, "d", 0, "Limit depth of directories walked")
	flag.IntVar(&optDepth, "depth", 0, "Limit depth of directories walked")
	pathSep = string(os.PathSeparator)

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

	// Process options
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

	if findPrefix == false && findSuffix == false && findContains == false {
		findAll = true
	}

	if findAll == false && len(args) == 0 {
		fmt.Fprintf(eout, "%s\n", fmtTxt(helpText, appName, datatools.Version))
		fmt.Fprintln(eout, "Missing required parameters")
		os.Exit(1)
	}

	// Shift the target filename so args holds the directories to search
	target := ""
	if findAll == false && len(args) > 0 {
		target, args = args[0], args[1:]
	}

	// Handle the case where currect work directory is assumed
	if len(args) == 0 {
		args = []string{"."}
	}

	// For each directory to search walk the path
	for _, dir := range args {
		err = walkPath(out, dir, target)
		if err != nil {
			fmt.Fprintln(eout, err)
			os.Exit(1)
		}
	}
}
