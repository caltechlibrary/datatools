// jsoncols is a command line tool for filter JSON data from standard in or specified files.
// It was inspired by [jq](https://github.com/stedolan/jq) and [jid](https://github.com/simeji/jid).
// It facilitates extract one or more columns of data from a JSON document.
//
// @author R. S. Doiel, <rsdoiel@caltech.edu>
package main

import (
	"bytes"
	"encoding/csv"
	"encoding/json"
	"flag"
	"fmt"
	"io/ioutil"
	"os"
	"path"

	// Caltech Library Packages
	"github.com/caltechlibrary/datatools"
	"github.com/caltechlibrary/dotpath"
)

var (
	helpText = `%{app_name}(1) user manual | version {version} {release_hash}
% R. S. Doiel
% {release_date}

# NAME

{app_name}

# SYNOPSIS

{app_name} [OPTIONS] [EXPRESSION] [INPUT_FILENAME] [OUTPUT_FILENAME]

# DESCRIPTION

{app_name} provides scripting flexibility for data extraction from JSON data
returning the results in columns.  This is helpful in flattening content
extracted from JSON blobs.  The default delimiter for each value
extracted is a comma. This can be overridden with an option.

- EXPRESSION can be an empty string or dot notation for an object's path
- INPUT_FILENAME is the filename to read or a dash "-" if you want to
  explicitly read from stdin
	- if not provided then {app_name} reads from stdin
- OUTPUT_FILENAME is the filename to write or a dash "-" if you want to
  explicitly write to stdout
	- if not provided then {app_name} write to stdout

# OPTIONS

-help
: display help

-license
: display license

-version
: display version

-csv
: output as CSV or other flat delimiter row

-d, -delimiter
: set the delimiter for multi-field csv output

-i, -input
: input filename

-nl, -newline
: if true add a trailing newline

-o, -output
: output filename

-p, -pretty
: pretty print JSON output

-quiet
: suppress error messages

-quote
: quote strings and JSON notation

-r, -repl
: run interactively


# EXAMPLES

If myblob.json contained

~~~
    {"name": "Doe, Jane", "email":"jane.doe@example.org", "age": 42}
~~~

Getting just the name could be done with

~~~
    {app_name} -i myblob.json .name
~~~

This would yield

~~~
    "Doe, Jane"
~~~

Flipping .name and .age into pipe delimited columns is as
easy as listing each field in the expression inside a
space delimited string.

~~~
    {app_name} -i myblob.json -d\|  .name .age
~~~

This would yield

~~~
    Doe, Jane|42
~~~

You can also pipe JSON data in.

~~~
    cat myblob.json | {app_name} .name .email .age
~~~

Would yield

~~~
   "Doe, Jane","jane.doe@xample.org",42
~~~

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
	prettyPrint      bool

	// Application Specific Options
	runInteractive bool
	csvOutput      bool
	delimiter      = ","
	expressions    []string
	quote          bool
)


func main() {
	var (
		src []byte
		err error
	)

	appName := path.Base(os.Args[0])
	version := datatools.Version
	license := datatools.LicenseText
	releaseDate := datatools.ReleaseDate
	releaseHash := datatools.ReleaseHash

	// Basic Options
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

	// Application Specific Options
	flag.BoolVar(&runInteractive, "r", false, "run interactively")
	flag.BoolVar(&runInteractive, "repl", false, "run interactively")
	flag.BoolVar(&csvOutput, "csv", false, "output as CSV or other flat delimiter row")
	flag.StringVar(&delimiter, "d", delimiter, "set the delimiter for multi-field csv output")
	flag.StringVar(&delimiter, "delimiter", delimiter, "set the delimiter for multi-field csv output")
	flag.BoolVar(&quote, "quote", true, "quote strings and JSON notation")
	flag.BoolVar(&prettyPrint, "p", false, "pretty print JSON output")
	flag.BoolVar(&prettyPrint, "pretty", false, "pretty print JSON output")

	// Parse Environment and Options
	flag.Parse()
	args := flag.Args()

	// Setup IO
	in := os.Stdin
	out := os.Stdout
	eout := os.Stderr

	if inputFName != "" && inputFName != "-" {
		in, err = os.Open(inputFName)
		if err != nil {
			fmt.Fprintf(eout, "input error %q, %q\n", inputFName, err)
			os.Exit(1)
		}
		defer in.Close()
	}

	if outputFName != "" && outputFName != "-" {
		out, err = os.Create(outputFName)
		if err != nil {
			fmt.Fprintf(eout, "output error %q, %q\n", outputFName, err)
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

	// Handle ordered args to get expressions for each column output.
	for _, arg := range args {
		if len(arg) == 0 {
			arg = "."
		}
		expressions = append(expressions, arg)
	}
	// Make sure we have a default expression to run.
	if len(expressions) == 0 {
		expressions = []string{"."}
	}

	// READ in the JSON document
	buf, err := ioutil.ReadAll(in)
	if err != nil {
		fmt.Fprintln(eout, err)
		os.Exit(1)
	}

	// JSON Decode our document
	decoder := json.NewDecoder(bytes.NewBuffer(buf))
	decoder.UseNumber()
	var data interface{}
	if err := decoder.Decode(&data); err != nil {
		fmt.Fprintln(eout, err)
		os.Exit(1)
	}

	if csvOutput == true {
		// For each dotpath expression return a result
		row := []string{}
		for _, qry := range expressions {
			result, err := dotpath.Eval(qry, data)
			if err != nil {
				fmt.Fprintln(eout, err)
				os.Exit(1)
			}
			switch result.(type) {
			case string:
				row = append(row, result.(string))
			case json.Number:
				row = append(row, result.(json.Number).String())
			default:
				if prettyPrint {
					src, err = datatools.JSONMarshalIndent(result, "", "    ")
				} else {
					src, err = datatools.JSONMarshal(result)
				}
				if err != nil {
					fmt.Fprintln(eout, err)
					os.Exit(1)
				}
				row = append(row, fmt.Sprintf("%s", src))
			}
		}

		// Setup the CSV output
		w := csv.NewWriter(out)
		if delimiter != "" {
			w.Comma = datatools.NormalizeDelimiterRune(delimiter)
		}
		err = w.Write(row)
		if err != nil {
			fmt.Fprintln(eout, err)
			os.Exit(1)
		}
		w.Flush()
		err = w.Error()
		if err != nil {
			fmt.Fprintln(eout, err)
			os.Exit(1)
		}
		os.Exit(0)
	}

	// Output JSON format (default)
	// For each dotpath expression return a result
	for i, qry := range expressions {
		if i > 0 {
			fmt.Fprintf(out, "%s", delimiter)
		}
		if qry == "." {
			fmt.Fprintf(out, "%s", buf)
		} else {
			result, err := dotpath.Eval(qry, data)
			if err != nil {
				fmt.Fprintln(eout, err)
				os.Exit(1)
			}
			switch result.(type) {
			case string:
				if quote == true {
					fmt.Fprintf(out, "%q", result)
				} else {
					fmt.Fprintf(out, "%s", result)
				}
			case json.Number:
				fmt.Fprintf(out, "%s", result.(json.Number).String())
			default:
				if prettyPrint {
					src, err = datatools.JSONMarshalIndent(result, "", "    ")
				} else {
					src, err = datatools.JSONMarshal(result)
				}
				if err != nil {
					fmt.Fprintln(eout, err)
					os.Exit(1)
				}
				/*
					if quote == true {
						fmt.Fprintf(out, "%q", src)
					} else {
						fmt.Fprintf(out, "%s", src)
					}
				*/
				fmt.Fprintf(out, "%s", src)
			}
		}
	}
	if newLine {
		fmt.Fprintln(out, "")
	}
}
