//
// jsoncols is a command line tool for filter JSON data from standard in or specified files.
// It was inspired by [jq](https://github.com/stedolan/jq) and [jid](https://github.com/simeji/jid).
// It facilitates extract one or more columns of data from a JSON document.
//
// @author R. S. Doiel, <rsdoiel@caltech.edu>
//
package main

import (
	"encoding/csv"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"

	// Caltech Library Packages
	"github.com/caltechlibrary/cli"
	"github.com/caltechlibrary/datatools"
	"github.com/caltechlibrary/dotpath"
)

var (
	description = `
%s provides scripting flexibility for data extraction from JSON data
returning the results in columns.  This is helpful in flattening content
extracted from JSON blobs.  The default delimiter for each value
extracted is a comma. This can be overridden with an option.

+ EXPRESSION can be an empty string or dot notation for an object's path
+ INPUT_FILENAME is the filename to read or a dash "-" if you want to
  explicitly read from stdin
	+ if not provided then %s reads from stdin
+ OUTPUT_FILENAME is the filename to write or a dash "-" if you want to
  explicitly write to stdout
	+ if not provided then %s write to stdout
`

	examples = `
If myblob.json contained

    {"name": "Doe, Jane", "email":"jane.doe@example.org", "age": 42}

Getting just the name could be done with

    %s -i myblob.json .name

This would yield

    "Doe, Jane"

Flipping .name and .age into pipe delimited columns is as
easy as listing each field in the expression inside a
space delimited string.

    %s -i myblob.json -d\|  .name .age

This would yield

    Doe, Jane|42

You can also pipe JSON data in.

    cat myblob.json | %s .name .email .age

Would yield

   "Doe, Jane","jane.doe@xample.org",42
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

	app := cli.NewCli(datatools.Version)
	appName := app.AppName()

	// Document non-option parameters
	app.AddParams("[EXPRESSION]", "[INPUT_FILENAME]", "[OUTPUT_FILENAME]")

	// Add Help Docs
	app.AddHelp("license", []byte(fmt.Sprintf(datatools.LicenseText, appName, datatools.Version)))
	app.AddHelp("description", []byte(fmt.Sprintf(description, appName, appName, appName)))
	app.AddHelp("examples", []byte(fmt.Sprintf(examples, appName, appName, appName)))

	// Basic Options
	app.BoolVar(&showHelp, "h,help", false, "display help")
	app.BoolVar(&showLicense, "l,license", false, "display license")
	app.BoolVar(&showVersion, "v,version", false, "display version")
	app.BoolVar(&showExamples, "examples", false, "display example(s)")
	app.StringVar(&inputFName, "i,input", "", "input filename")
	app.StringVar(&outputFName, "o,output", "", "output filename")
	app.BoolVar(&quiet, "quiet", false, "suppress error messages")
	app.BoolVar(&newLine, "nl,newline", false, "if true add a trailing newline")
	app.BoolVar(&generateMarkdown, "generate-markdown", false, "generate markdown documentation")
	app.BoolVar(&generateManPage, "generate-manpage", false, "generate man page")

	// Application Specific Options
	app.BoolVar(&runInteractive, "r", false, "run interactively")
	app.BoolVar(&runInteractive, "repl", false, "run interactively")
	app.BoolVar(&csvOutput, "csv", false, "output as CSV or other flat delimiter row")
	app.StringVar(&delimiter, "d,delimiter", delimiter, "set the delimiter for multi-field csv output")
	app.BoolVar(&quote, "quote", true, "quote strings and JSON notation")
	app.BoolVar(&prettyPrint, "p,pretty", false, "pretty print JSON output")

	// Parse Environment and Options
	app.Parse()
	args := app.Args()

	// Setup IO
	app.Eout = os.Stderr
	app.In, err = cli.Open(inputFName, os.Stdin)
	cli.ExitOnError(app.Eout, err, quiet)
	defer cli.CloseFile(inputFName, app.In)

	app.Out, err = cli.Create(outputFName, os.Stdout)
	cli.ExitOnError(app.Eout, err, quiet)
	defer cli.CloseFile(outputFName, app.Out)

	// Process Options
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
	buf, err := ioutil.ReadAll(app.In)
	cli.ExitOnError(app.Eout, err, quiet)

	// JSON Decode our document
	data, err := dotpath.JSONDecode(buf)
	cli.ExitOnError(app.Eout, err, quiet)

	if csvOutput == true {
		// For each dotpath expression return a result
		row := []string{}
		for _, qry := range expressions {
			result, err := dotpath.Eval(qry, data)
			cli.ExitOnError(app.Eout, err, quiet)
			switch result.(type) {
			case string:
				row = append(row, result.(string))
			case json.Number:
				row = append(row, result.(json.Number).String())
			default:
				if prettyPrint {
					src, err = json.MarshalIndent(result, "", "    ")
				} else {
					src, err = json.Marshal(result)
				}
				cli.ExitOnError(app.Eout, err, quiet)
				row = append(row, fmt.Sprintf("%s", src))
			}
		}

		// Setup the CSV output
		w := csv.NewWriter(app.Out)
		if delimiter != "" {
			w.Comma = datatools.NormalizeDelimiterRune(delimiter)
		}
		err = w.Write(row)
		cli.ExitOnError(app.Eout, err, quiet)
		w.Flush()
		err = w.Error()
		cli.ExitOnError(app.Eout, err, quiet)
		os.Exit(0)
	}

	// Output JSON format (default)
	// For each dotpath expression return a result
	for i, qry := range expressions {
		if i > 0 {
			fmt.Fprintf(app.Out, "%s", delimiter)
		}
		if qry == "." {
			fmt.Fprintf(app.Out, "%s", buf)
		} else {
			result, err := dotpath.Eval(qry, data)
			cli.ExitOnError(app.Eout, err, quiet)
			switch result.(type) {
			case string:
				if quote == true {
					fmt.Fprintf(app.Out, "%q", result)
				} else {
					fmt.Fprintf(app.Out, "%s", result)
				}
			case json.Number:
				fmt.Fprintf(app.Out, "%s", result.(json.Number).String())
			default:
				if prettyPrint {
					src, err = json.MarshalIndent(result, "", "    ")
				} else {
					src, err = json.Marshal(result)
				}
				cli.ExitOnError(app.Eout, err, quiet)
				/*
					if quote == true {
						fmt.Fprintf(app.Out, "%q", src)
					} else {
						fmt.Fprintf(app.Out, "%s", src)
					}
				*/
				fmt.Fprintf(app.Out, "%s", src)
			}
		}
	}
	if newLine {
		fmt.Fprintln(app.Out, "")
	}
}
