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
	"flag"
	"fmt"
	"io/ioutil"
	"os"
	"path"
	"strings"

	// Caltech Library Packages
	"github.com/caltechlibrary/cli"
	"github.com/caltechlibrary/datatools"
	"github.com/caltechlibrary/dotpath"
)

var (
	usage = `USAGE: %s [OPTIONS] [EXPRESSION] [INPUT_FILENAME] [OUTPUT_FILENAME]`

	description = `
SYSNOPSIS

%s provides scripting flexibility for data extraction from JSON data 
returning the results in columns.  This is helpful in flattening content 
extracted from JSON blobs.  The default delimiter for each value 
extracted is a comma. This can be overridden with an option.

+ EXPRESSION can be an empty stirng or dot notation for an object's path
+ INPUT_FILENAME is the filename to read or a dash "-" if you want to 
  explicity read from stdin
	+ if not provided then %s reads from stdin
+ OUTPUT_FILENAME is the filename to write or a dash "-" if you want to 
  explicity write to stdout
	+ if not provided then %s write to stdout
`

	examples = `
EXAMPLES

If myblob.json contained

{"name": "Doe, Jane", "email":"jane.doe@example.org", "age": 42}

Getting just the name could be done with

    %s -i myblob.json .name

This would yeild

    "Doe, Jane"

Flipping .name and .age into pipe delimited columns is as 
easy as listing each field in the expression inside a 
space delimited string.

    %s -i myblob.json -d\|  .name .age 

This would yeild

    "Doe, Jane"|42

You can also pipe JSON data in.

    cat myblob.json | %s .name .email .age

Would yield

   "Doe, Jane",jane.doe@xample.org,42
`

	// Basic Options
	showHelp    bool
	showLicense bool
	showVersion bool
	inputFName  string
	outputFName string

	// Application Specific Options
	monochrome     bool
	runInteractive bool
	csvOutput      bool
	delimiter      = ","
	expressions    []string
	permissive     bool
	quote          bool
)

func handleError(err error, exitCode int) {
	if permissive == false {
		fmt.Fprintf(os.Stderr, "%s\n", err)
	}
	if exitCode >= 0 {
		os.Exit(exitCode)
	}
}

func init() {
	// Basic Options
	flag.BoolVar(&showHelp, "h", false, "display help")
	flag.BoolVar(&showHelp, "help", false, "display help")
	flag.BoolVar(&showLicense, "l", false, "display license")
	flag.BoolVar(&showLicense, "license", false, "display license")
	flag.BoolVar(&showVersion, "v", false, "display version")
	flag.BoolVar(&showVersion, "version", false, "display version")
	flag.StringVar(&inputFName, "i", "", "input filename")
	flag.StringVar(&inputFName, "input", "", "input filename")
	flag.StringVar(&outputFName, "o", "", "output filename")
	flag.StringVar(&outputFName, "output", "", "output filename")

	// Application Specific Options
	flag.BoolVar(&monochrome, "m", false, "display output in monochrome")
	flag.BoolVar(&runInteractive, "r", false, "run interactively")
	flag.BoolVar(&runInteractive, "repl", false, "run interactively")
	flag.BoolVar(&csvOutput, "csv", false, "output as CSV or other flat delimiter row")
	flag.StringVar(&delimiter, "d", delimiter, "set the delimiter for multi-field csv output")
	flag.StringVar(&delimiter, "dimiter", delimiter, "set the delimiter for multi-field csv output")
	flag.BoolVar(&quote, "quote", false, "if dilimiter is found in column value add quotes for non-CSV output")
	flag.BoolVar(&permissive, "permissive", false, "suppress error messages")
	flag.BoolVar(&permissive, "quiet", false, "suppress error messages")
}

func main() {
	appName := path.Base(os.Args[0])
	flag.Parse()
	args := flag.Args()

	// Configuration and command line interation
	cfg := cli.New(appName, "DATATOOLS", fmt.Sprintf(datatools.LicenseText, appName, datatools.Version), datatools.Version)
	cfg.UsageText = fmt.Sprintf(usage, appName)
	cfg.DescriptionText = fmt.Sprintf(description, appName, appName, appName)
	cfg.ExampleText = fmt.Sprintf(examples, appName, appName, appName)

	//NOTE: Need to handle JSONQUERY_MONOCHROME setting
	monochrome = cfg.MergeEnvBool("monochrome", monochrome)

	if showHelp == true {
		fmt.Println(cfg.Usage())
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

	in, err := cli.Open(inputFName, os.Stdin)
	if err != nil {
		handleError(err, 1)
	}
	defer cli.CloseFile(inputFName, in)

	out, err := cli.Create(outputFName, os.Stdout)
	if err != nil {
		handleError(err, 1)
	}
	defer cli.CloseFile(outputFName, out)

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
		handleError(err, 1)
	}
	// JSON Decode our document
	data, err := dotpath.JSONDecode(buf)
	if err != nil {
		handleError(err, 1)
	}

	if csvOutput == true {
		// For each dotpath expression return a result
		row := []string{}
		for _, qry := range expressions {
			result, err := dotpath.Eval(qry, data)
			if err == nil {
				switch result.(type) {
				case string:
					row = append(row, result.(string))
				case json.Number:
					row = append(row, result.(json.Number).String())
				default:
					src, err := json.Marshal(result)
					if err != nil {
						handleError(err, 1)
					}
					row = append(row, fmt.Sprintf("%s", src))
				}
			} else {
				handleError(err, 1)
			}
		}

		// Setup the CSV output
		w := csv.NewWriter(out)
		if delimiter != "" {
			w.Comma = datatools.NormalizeDelimiterRune(delimiter)
		}
		if err := w.Write(row); err != nil {
			handleError(err, 1)
		}
		w.Flush()
		if err := w.Error(); err != nil {
			handleError(err, 1)
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
			if err == nil {
				switch result.(type) {
				case string:
					if quote == true && strings.Contains(result.(string), delimiter) == true {
						fmt.Fprintf(out, "%q", result)
					} else {
						fmt.Fprintf(out, "%s", result)
					}
				case json.Number:
					fmt.Fprintf(out, "%s", result.(json.Number).String())
				default:
					src, err := json.Marshal(result)
					if err != nil {
						handleError(err, 1)
					}
					if quote == true {
						fmt.Fprintf(out, "%q", src)
					} else {
						fmt.Fprintf(out, "%s", src)
					}
				}
			} else {
				handleError(err, 1)
			}
		}
	}
}
