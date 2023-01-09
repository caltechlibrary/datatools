//
// jsonmunge is a command line tool that takes a JSON document and
// a Go text/template rendering the result. Useful for
// reshaping a JSON document or transforming into a new format,
// or filter for specific content.
//
// @author R. S. Doiel, <rsdoiel@caltech.edu>
//
package main

import (
	"flag"
	"fmt"
	"io/ioutil"
	"os"
	"path"
	"strings"
	"text/template"

	// Caltech Library Packages
	"github.com/caltechlibrary/datatools"
	"github.com/caltechlibrary/dotpath"
	"github.com/caltechlibrary/tmplfn"
)

var (
	helpText = `---
title: "{app_name}"
author: "R. S. Doiel"
pubDate: 2023-01-09
---

# NAME

{app_name} 

# SYNOPSIS

{app_name} [OPTIONS] TEMPLATE_FILENAME

# DESCRIPTION

{app_name} is a command line tool that takes a JSON document and
one or more Go templates rendering the results. Useful for
reshaping a JSON document, transforming into a new format,
or filter for specific content.

- TEMPLATE_FILENAME is the name of a Go text tempate file used to render
  the outbound JSON document

# OPTIONS

-help
: display help

-license
: display license

-version
: display version

-E, -expression
: use template expression as template

-i, -input
: input filename

-nl, -newline
: if true add a trailing newline

-o, -output
: output filename

-quiet
: suppress error messages


# EXAMPLES

If person.json contained

~~~
   {"name": "Doe, Jane", "email":"jd@example.org", "age": 42}
~~~

and the template, name.tmpl, contained

~~~
   {{- .name -}}
~~~

Getting just the name could be done with

~~~
    cat person.json | {app_name} name.tmpl
~~~

This would yield

~~~
    "Doe, Jane"
~~~

{app_name} {version}
`

	// Basic Options
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

	// Application Specific Options
	templateExpr string
)

func fmtTxt(src string, appName string, version string) string {
	return strings.ReplaceAll(strings.ReplaceAll(src, "{app_name}", appName), "{version}", version)
}

func main() {
	appName := path.Base(os.Args[0])

	// Basic Options
	flag.BoolVar(&showHelp, "help", false, "display help")
	flag.BoolVar(&showLicense, "license", false, "display license")
	flag.BoolVar(&showVersion, "version", false, "display version")
	flag.BoolVar(&showExamples, "examples", false, "display example(s)")

	flag.StringVar(&inputFName, "i", "", "input filename")
	flag.StringVar(&inputFName, "input", "", "input filename")
	flag.StringVar(&outputFName, "o", "", "output filename")
	flag.StringVar(&outputFName, "output", "", "output filename")
	flag.BoolVar(&quiet, "quiet", false, "suppress error messages")
	flag.BoolVar(&newLine, "nl", false, "if true add a trailing newline")
	flag.BoolVar(&newLine, "newline", false, "if true add a trailing newline")

	// Application Specific Options
	flag.StringVar(&templateExpr, "E", "", "use template expression as template")
	flag.StringVar(&templateExpr, "expression", "", "use template expression as template")

	// Parse env and options
	flag.Parse()
	args := flag.Args()

	// Setup IO
	var err error

	in := os.Stdin
	out := os.Stdout
	eout := os.Stderr

	if inputFName != "" {
		in, err = os.Open(inputFName)
		if err != nil {
			fmt.Fprintln(eout, err)
			os.Exit(1)
		}
		defer in.Close()
	}

	if outputFName != "" {
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

	if len(args) == 0 && templateExpr == "" {
		if err != nil {
			fmt.Fprintln(eout, err)
			os.Exit(1)
		}
	}

	var (
		tmpl *template.Template
	)
	if templateExpr != "" {
		// Read in and compile our templates expression
		tmpl, err = template.New("default").Funcs(tmplfn.AllFuncs()).Parse(templateExpr)
		if err != nil {
			fmt.Fprintln(eout, err)
			os.Exit(1)
		}
	} else {
		// Read in and compile our templates
		tmpl, err = template.New(path.Base(args[0])).Funcs(tmplfn.AllFuncs()).ParseFiles(args...)
		if err != nil {
			fmt.Fprintln(eout, err)
			os.Exit(1)
		}
	}

	// READ in the JSON document
	buf, err := ioutil.ReadAll(in)
	if err != nil {
		fmt.Fprintln(eout, err)
		os.Exit(1)
	}

	// JSON Decode our document
	data, err := dotpath.JSONDecode(buf)
	if err != nil {
		fmt.Fprintln(eout, err)
		os.Exit(1)
	}

	// Execute template with data
	err = tmpl.Execute(out, data)
	if err != nil {
		fmt.Fprintln(eout, err)
		os.Exit(1)
	}
	if newLine {
		fmt.Fprintln(out, "")
	}
}
