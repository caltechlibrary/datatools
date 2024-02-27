// tabs2csv converts a tab delimited file to a CSV formatted file.
//
// @author R. S. Doiel, <rsdoiel@caltech.edu>
package main

import (
	"encoding/csv"
	"flag"
	"fmt"
	"io"
	"os"
	"path"
	"strings"

	// Caltech Library Packages
	"github.com/caltechlibrary/datatools"
)

var (
	helpText = `%{app_name}(1) user manual | version {version} {release_hash}
% R. S. Doiel
% {release_date}

# NAME

{app_name} 

# SYNOPSIS

{app_name} [OPTIONS]

# DESCRIPTION

{app_name} is a simple conversion utility to convert from tabs to quoted CSV.

{app_name} reads from standard input and writes to standard out.


# OPTIONS

-help
: display help

-license
: display license

-version
: display version

-fields-per-record
: (int) sets the number o fields expected in each row, -1 turns this off

-reuse-record
: re-uses the backing array on reader

-trim-leading-space
: trims leading space read

-use-lazy-quotes
: use lazy quoting for reader

# EXAMPLES

If my.tab contained

~~~
    name	email	age
	Doe, Jane	jane.doe@example.org	42
~~~

Concert this to a CSV file format

~~~
    {app_name} < my.tab 
~~~

This would yield

~~~
    "name","email","age"
	"Doe, Jane","jane.doe@example.org",42
~~~

{app_name} {version}

`

	// Standard Options
	showHelp    bool
	showLicense bool
	showVersion bool

	// CSV Reader Options
	lazyQuotes       bool
	trimLeadingSpace bool
	reuseRecord      bool
	fieldsPerRecord  int
)

func fmtTxt(src string, appName string, version string) string {
	return strings.ReplaceAll(strings.ReplaceAll(src, "{app_name}", appName), "{version}", version)
}

func main() {
	appName := path.Base(os.Args[0])

	flag.BoolVar(&showHelp, "help", false, "display help")
	flag.BoolVar(&showLicense, "license", false, "display license")
	flag.BoolVar(&showVersion, "version", false, "display version")

	// CSV Reader options
	flag.IntVar(&fieldsPerRecord, "fields-per-record", 0, "sets the number o fields expected in each row, -1 turns this off")
	flag.BoolVar(&lazyQuotes, "use-lazy-quotes", false, "use lazy quoting for reader")
	flag.BoolVar(&trimLeadingSpace, "trim-leading-space", false, "trims leading space read")
	flag.BoolVar(&reuseRecord, "reuse-record", false, "re-uses the backing array on reader")

	// Parse Environment and Options
	flag.Parse()

	in := os.Stdin
	out := os.Stdout
	eout := os.Stderr

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

	// Setup the CSV output
	r := csv.NewReader(in)
	r.Comma = '\t'
	r.Comment = '#'
	r.FieldsPerRecord = fieldsPerRecord
	r.LazyQuotes = lazyQuotes
	r.TrimLeadingSpace = trimLeadingSpace
	r.ReuseRecord = reuseRecord

	exitCode := 0
	w := csv.NewWriter(out)
	/*
		if delimiter != "" {
			w.Comma = datatools.NormalizeDelimiterRune(delimiter)
		}
	*/
	for {
		row, err := r.Read()
		if err == io.EOF {
			break
		}
		if err != nil {
			fmt.Fprintln(eout, err)
			exitCode = 1
		} else if err := w.Write(row); err != nil {
			fmt.Fprintln(eout, err)
			exitCode = 1
		}
	}
	w.Flush()
	if err := w.Error(); err != nil {
		fmt.Fprintln(eout, err)
		exitCode = 1
	}
	os.Exit(exitCode)
}
