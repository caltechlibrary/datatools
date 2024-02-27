// csv2tab converts a CSV file to tab separated values.
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

csv2tab is a simple conversion utility to convert from CSV to tab separated values.
csv2tab reads from standard input and writes to standard out.

# OPTIONS

-help
: display help

-license
: display license

-version
: display version


# EXAMPLES

If my.tab contained

~~~
    "name","email","age"
	"Doe, Jane","jane.doe@example.org",42
~~~

Concert this to a tab separated values

~~~
    csv2tab < my.csv 
~~~

This would yield

~~~
    name	email	age
	Doe, Jane	jane.doe@example.org	42
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

func main() {
	appName := path.Base(os.Args[0])
	version := datatools.Version
	license := datatools.LicenseText
	releaseDate := datatools.ReleaseDate
	releaseHash := datatools.ReleaseHash

	flag.BoolVar(&showHelp, "h", false, "display help")
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

	// Setup the CSV output
	r := csv.NewReader(in)
	r.Comma = ','
	r.Comment = '#'
	r.FieldsPerRecord = fieldsPerRecord
	r.LazyQuotes = lazyQuotes
	r.TrimLeadingSpace = trimLeadingSpace
	r.ReuseRecord = reuseRecord

	exitCode := 0
	w := csv.NewWriter(out)
	w.Comma = '\t'
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
