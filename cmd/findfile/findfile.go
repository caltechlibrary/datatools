//
// findfile - a simple directory tree walker that looks for files
// by name, basename or extension. Basically a unix "find" light to
// demonstrate walking the file system
//
// @author R. S. Doiel, <rsdoiel@caltech.edu>
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
	"io"
	"os"
	"path/filepath"
	"strings"
	"time"

	// Caltech Library Packages
	"github.com/caltechlibrary/cli"
	"github.com/caltechlibrary/datatools"
)

var (
	description = `
%s finds files based on matching prefix, suffix or contained text in base filename.
`

	examples = `
Search the current directory and subdirectories for Markdown files with extension of ".md".

	%s -s .md
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
	newLine          bool
	eol              string

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
	app := cli.NewCli(datatools.Version)
	appName := app.AppName()

	// Document non-option paramaters
	app.AddParams("[TARGET]", "[DIRECTORIES_TO_SEARCH]")

	// Add Help Docs
	app.AddHelp("license", []byte(fmt.Sprintf(datatools.LicenseText, appName, datatools.Version)))
	app.AddHelp("description", []byte(fmt.Sprintf(description, appName)))
	app.AddHelp("examples", []byte(fmt.Sprintf(examples, appName)))

	// Standard Options
	app.BoolVar(&showHelp, "h,help", false, "display this help message")
	app.BoolVar(&showLicense, "l,license", false, "display license information")
	app.BoolVar(&showVersion, "v,version", false, "display version message")
	app.BoolVar(&showExamples, "examples", false, "display example(s)")
	app.StringVar(&outputFName, "o,output", "", "output filename")
	app.BoolVar(&generateMarkdown, "generate-markdown", false, "generate markdown documentation")
	app.BoolVar(&generateManPage, "generate-manpage", false, "generate man page")
	app.BoolVar(&quiet, "quiet", false, "suppress error messages")
	app.BoolVar(&newLine, "nl,newline", false, "if true add a trailing newline")

	// App Specific Options
	app.BoolVar(&showModificationTime, "m,mod-time", false, "display file modification time before the path")
	app.BoolVar(&stopOnErrors, "error,stop-on-error", false, "Stop walk on file system errors (e.g. permissions)")
	app.BoolVar(&findPrefix, "p,prefix", false, "find file(s) based on basename prefix")
	app.BoolVar(&findContains, "c,contains", false, "find file(s) based on basename containing text")
	app.BoolVar(&findSuffix, "s,suffix", false, "find file(s) based on basename suffix")
	app.BoolVar(&outputFullPath, "f,full-path", false, "list full path for files found")
	app.IntVar(&optDepth, "d,depth", 0, "Limit depth of directories walked")
	pathSep = string(os.PathSeparator)

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

	// Process options
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

	if findPrefix == false && findSuffix == false && findContains == false {
		findAll = true
	}

	if findAll == false && len(args) == 0 {
		app.Usage(app.Eout)
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
		err = walkPath(app.Out, dir, target)
		cli.ExitOnError(app.Eout, err, quiet)
	}
	fmt.Fprintf(app.Out, "%s", eol)
}
