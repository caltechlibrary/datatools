//
// jsonpath is a command line tool for filter JSON data from standard in or specified files.
// It was inspired by [jq](https://github.com/stedolan/jq) and [jid](https://github.com/simeji/jid).
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

	// Caltech Library Packages
	"github.com/caltechlibrary/cli"
	"github.com/caltechlibrary/datatools"
	"github.com/caltechlibrary/datatools/dotpath"
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
	delimiter      = ","
	expressions    []string
)

func init() {
	// Basic Options
	flag.BoolVar(&showHelp, "h", false, "display help")
	flag.BoolVar(&showLicense, "l", false, "display license")
	flag.BoolVar(&showVersion, "v", false, "display version")
	flag.StringVar(&inputFName, "i", "", "input filename")
	flag.StringVar(&inputFName, "input", "", "input filename")
	flag.StringVar(&outputFName, "o", "", "output filename")
	flag.StringVar(&outputFName, "output", "", "output filename")

	// Application Specific Options
	flag.BoolVar(&monochrome, "m", false, "display output in monochrome")
	flag.BoolVar(&runInteractive, "r", false, "run interactively")
	flag.BoolVar(&runInteractive, "repl", false, "run interactively")
	flag.StringVar(&delimiter, "d", delimiter, "set the delimiter for multi-field output")
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
		fmt.Fprintf(os.Stderr, "%s\n", err)
		os.Exit(1)
	}
	defer cli.CloseFile(inputFName, in)

	out, err := cli.Create(outputFName, os.Stdout)
	if err != nil {
		fmt.Fprintf(os.Stderr, "%s\n", err)
		os.Exit(1)
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
		fmt.Fprintf(os.Stderr, "%s\n", err)
		os.Exit(1)
	}
	var data interface{}
	err = json.Unmarshal(buf, data)

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
				fmt.Fprintf(os.Stderr, "%s\n", err)
				os.Exit(1)
			}
			switch result.(type) {
			case string:
				fmt.Fprintf(out, "%s", result)
			case float64:
				fmt.Fprintf(out, "%f", result)
			default:
				src, err := json.Marshal(result)
				if err != nil {
					fmt.Fprintf(os.Stderr, "%s\n", err)
					os.Exit(1)
				}
				fmt.Fprintf(out, "%s", src)
			}
		}
	}
}
