//
// tabs2csv converts a tab delimited file to a CSV formatted file.
//
// @author R. S. Doiel, <rsdoiel@caltech.edu>
//
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
	description = `
USAGE

%s is a simple conversion utility to convert from tabs to quoted CSV.
%s reads from standard input and writes to standard out.

`

	examples = `
If my.tab contained

    name	email	age
	Doe, Jane	jane.doe@example.org	42

Concert this to a CSV file format

    %s < my.tab 

This would yield

    "name","email","age"
	"Doe, Jane","jane.doe@example.org",42

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

	if showHelp {
		fmt.Fprintf(os.Stdout, description, appName, appName)
		//FIXME: need to forse this to standard out ...
		flag.PrintDefaults()
		fmt.Fprintf(os.Stdout, examples, appName)
		os.Exit(0)
	}
	if showLicense {
		fmt.Fprintln(os.Stdout, datatools.LicenseText, appName, datatools.Version)
		os.Exit(0)
	}
	if showVersion {
		fmt.Fprintln(os.Stdout, datatools.Version)
		os.Exit(0)
	}

	// Setup the CSV output
	r := csv.NewReader(os.Stdin)
	r.Comma = '\t'
	r.Comment = '#'
	r.FieldsPerRecord = fieldsPerRecord
	r.LazyQuotes = lazyQuotes
	r.TrimLeadingSpace = trimLeadingSpace
	r.ReuseRecord = reuseRecord

	exitCode := 0
	w := csv.NewWriter(os.Stdout)
	/*
		if delimiter != "" {
			w.Comma = datatools.NormalizeDelimiterRune(delimiter)
		}
	*/
	for {
		row, err := r.Read()
		if err == io.EOF {
			break
		} else if err != nil {
			fmt.Fprintln(os.Stderr, err)
			exitCode = 1
		} else {
			if err := w.Write(row); err != nil {
				fmt.Fprintln(os.Stderr, err)
				exitCode = 1
			}
			w.Flush()
			if err := w.Error(); err != nil {
				fmt.Fprintln(os.Stderr, err)
				exitCode = 1
			}
		}
	}
	os.Exit(exitCode)
}
