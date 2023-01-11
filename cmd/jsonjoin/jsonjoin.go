//
// jsonjoin is a command line tool that takes two JSON documents and combined
// them into one depending on the options
//
// @author R. S. Doiel, <rsdoiel@caltech.edu>
//
package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"io/ioutil"
	"os"
	"path"
	"strings"

	// Caltech Library Packages
	"github.com/caltechlibrary/datatools"
)

var (
	helpText = `---
title: "json2toml (1) user manual"
author: "R. S. Doiel"
pubDate: 2013-01-06
---

# NAME

json2toml 

# SYNOPSIS

json2toml [OPTIONS] [JSON_FILENAME] [TOML_FILENAME]

# DESCRIPTION

json2toml is a tool that converts JSON objects into TOML output.

# OPTIONS

-help
: display help

-license
: display license

-version:
display version

-nl, -newline
: if true add a trailing newline

-o, -output
: output filename

-p, -pretty
: pretty print output

-quiet
: suppress error messages


# EXAMPLES

These would get the file named "my.json" and save it as my.toml

~~~
    json2toml my.json > my.toml

	json2toml my.json my.toml

	cat my.json | json2toml -i - > my.toml
~~~

json2toml 1.2.1


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

	// Application Specific Options
	update     bool
	overwrite  bool
	createRoot bool
)

func fmtTxt(src string, appName string, version string) string {
	return strings.ReplaceAll(strings.ReplaceAll(src, "{app_name}", appName), "{version}", version)
}

func main() {
	appName := path.Base(os.Args[0])

	// Standard Options
	flag.BoolVar(&showHelp, "help", false, "display help")
	flag.BoolVar(&showLicense, "license", false, "display license")
	flag.BoolVar(&showVersion, "version", false, "display version")

	flag.StringVar(&inputFName, "i", "", "input filename (for root object)")
	flag.StringVar(&inputFName, "input", "", "input filename (for root object)")
	flag.StringVar(&outputFName, "o", "", "output filename")
	flag.StringVar(&outputFName, "output", "", "output filename")
	flag.BoolVar(&quiet, "quiet", false, "suppress error messages")
	flag.BoolVar(&newLine, "nl", false, "if true add a trailing newline")
	flag.BoolVar(&newLine, "newline", false, "if true add a trailing newline")

	// Application Specific Options
	flag.BoolVar(&createRoot, "create", false, "create an empty root object, {}")
	flag.BoolVar(&update, "update", false, "copy new key/values pairs into root object")
	flag.BoolVar(&overwrite, "overwrite", false, "copy all key/values into root object")

	// Parse env amd options
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
			fmt.Fprintln(eout,err)
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
	if newLine {
		eol = "\n"
	}

	// Make sure we have some JSON objects to join...
	if len(args) < 1 {
		if err != nil {
			fmt.Fprintln(eout, err)
			os.Exit(1)
		}
	}

	outObject := map[string]interface{}{}
	newObject := map[string]interface{}{}

	// READ in the JSON document if present on standard in or specified with -i.
	if createRoot == false {
		buf, err := ioutil.ReadAll(in)
		if err != nil {
			fmt.Fprintln(eout, err)
			os.Exit(1)
		}
		err = json.Unmarshal(buf, &outObject)
		if err != nil {
			fmt.Fprintln(eout, err)
			os.Exit(1)
		}
	}

	for _, arg := range args {
		src, err := ioutil.ReadFile(arg)
		if err != nil {
			fmt.Fprintln(eout, err)
			os.Exit(1)
		}
		err = json.Unmarshal(src, &newObject)
		if err != nil {
			fmt.Fprintln(eout, err)
			os.Exit(1)
		}
		switch {
		case update == true:
			for k, v := range newObject {
				if _, ok := outObject[k]; ok != true {
					outObject[k] = v
				}
			}
		case overwrite == true:
			for k, v := range newObject {
				outObject[k] = v
			}
		default:
			key := strings.TrimSuffix(path.Base(arg), ".json")
			outObject[key] = newObject
		}
	}

	src, err := json.Marshal(outObject)
	if err != nil {
		fmt.Fprintln(eout, err)
		os.Exit(1)
	}
	fmt.Fprintf(out, "%s%s", src, eol)
}
