//
// jsonrange iterates over an array or map returning either a JSON expression
// or map keep to stdout
//
// @Author R. S. Doiel, <rsdoiel@caltech.edu>
//
package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
	"strings"

	// CaltechLibrary Packages
	"github.com/caltechlibrary/cli"
	"github.com/caltechlibrary/datatools"
	"github.com/caltechlibrary/dotpath"
)

var (
	description = `
%s returns returns a range of values based on the JSON structure being read and
options applied.  Without options the JSON structure is read from standard input
and writes a list of keys to standard out. Keys are either attribute names or for
arrays the index position (counting form zero).  If a DOT_PATH_EXPRESSION is included
on the command line then that is used to generate the results. Using options to 
can choose to read the JSON data structure from a file, write the output to a file
as well as display values instead of keys. a list of "keys" of an index or map in JSON.  

Using options it can also return a list of values.  The JSON object is read from standard in and the
resulting list is normally written to standard out. There are options to read or
write to files.  Additional parameters are assumed to be a dot path notation
select the parts of the JSON data structure you want from the range. 

DOT_PATH_EXPRESSION is a dot path stale expression indicating what you want range over.
E.g.

+ . would indicate the whole JSON data structure read is used to range over
+ .name would indicate to range over the value pointed at by the "name" attribute 
+ ["name"] would indicate to range over the value pointed at by the "name" attribute
+ [0] would indicate to range over the value held in the zero-th element of the array

The path can be chained together

+ .name.family would point to the value heald by the "name" attributes' "family" attribute.
`

	examples = `
Working with a map

    echo '{"name": "Doe, Jane", "email":"jane.doe@example.org", "age": 42}' \
       | %s

This would yield

    name
    email
    age

Using the -values option on a map

    echo '{"name": "Doe, Jane", "email":"jane.doe@example.org", "age": 42}' \
      | %s -values

This would yield

    "Doe, Jane"
    "jane.doe@example.org"
    42


Working with an array

    echo '["one", 2, {"label":"three","value":3}]' | %s

would yield

    0
    1
    2

Using the -values option on the same array

    echo '["one", 2, {"label":"three","value":3}]' | %s -values

would yield

    one
    2
    {"label":"three","value":3}

Checking the length of a map or array or number of keys in map

    echo '["one","two","three"]' | %s -length

would yield

    3

Check for the index value of last element

    echo '["one","two","three"]' | %s -last

would yield

    2

Limitting the number of items returned

    echo '[1,2,3,4,5]' | %s -limit 2

would yield

    1
    2
`

	// Basic Options
	showHelp             bool
	showLicense          bool
	showVersion          bool
	showExamples         bool
	inputFName           string
	outputFName          string
	generateMarkdownDocs bool
	quiet                bool

	// Application Specific Options
	showLength bool
	showLast   bool
	showValues bool
	delimiter  string
	limit      int
)

func mapKeys(data map[string]interface{}, limit int) ([]string, error) {
	result := []string{}
	i := 0
	for keys := range data {
		result = append(result, keys)
		if i == limit {
			return result, nil
		}
		i++
	}
	return result, nil
}

func arrayKeys(data []interface{}, limit int) ([]string, error) {
	result := []string{}
	for i := range data {
		if i == limit {
			return result, nil
		}
		result = append(result, fmt.Sprintf("%d", i))
	}
	return result, nil
}

func mapVals(data map[string]interface{}, limit int) ([]string, error) {
	result := []string{}
	i := 0
	for _, val := range data {
		if i == limit {
			return result, nil
		}
		outSrc, err := json.Marshal(val)
		if err != nil {
			return nil, err
		}
		result = append(result, fmt.Sprintf("%s", outSrc))
		i++
	}
	return result, nil
}

func arrayVals(data []interface{}, limit int) ([]string, error) {
	result := []string{}
	for i, val := range data {
		if i == limit {
			return result, nil
		}
		outSrc, err := json.Marshal(val)
		if err != nil {
			return nil, err
		}
		result = append(result, fmt.Sprintf("%s", outSrc))
	}
	return result, nil
}

func getLength(data interface{}) (int, error) {
	switch data.(type) {
	case map[string]interface{}:
		return len(data.(map[string]interface{})), nil
	case []interface{}:
		return len(data.([]interface{})), nil
	}
	return -1, fmt.Errorf("%T does not support length of range", data)
}

func srcKeys(data interface{}, limit int) ([]string, error) {
	switch data.(type) {
	case map[string]interface{}:
		return mapKeys(data.(map[string]interface{}), limit)
	case []interface{}:
		return arrayKeys(data.([]interface{}), limit)
	}
	return nil, fmt.Errorf("%T does not support for range, %s", data, data)
}

func srcVals(data interface{}, limit int) ([]string, error) {
	switch data.(type) {
	case map[string]interface{}:
		return mapVals(data.(map[string]interface{}), limit)
	case []interface{}:
		return arrayVals(data.([]interface{}), limit)
	}
	return nil, fmt.Errorf("%T does not support for range", data)
}

func main() {
	app := cli.NewCli(datatools.Version)
	appName := app.AppName()

	// Document non-option parameters
	app.AddParams("[DOT_PATH_EXPRESSION]")

	// Add Help Docs
	app.AddHelp("license", []byte(fmt.Sprintf(datatools.LicenseText, appName, datatools.Version)))
	app.AddHelp("description", []byte(fmt.Sprintf(description, appName)))
	app.AddHelp("examples", []byte(fmt.Sprintf(examples, appName, appName, appName, appName, appName, appName, appName)))

	// Standard Options
	app.BoolVar(&showHelp, "h,help", false, "display help")
	app.BoolVar(&showLicense, "l,license", false, "display license")
	app.BoolVar(&showVersion, "v,version", false, "display version")
	app.BoolVar(&showExamples, "examples", false, "display example(s)")
	app.StringVar(&inputFName, "i,input", "", "read JSON from file")
	app.StringVar(&outputFName, "o,output", "", "write to output file")
	app.BoolVar(&generateMarkdownDocs, "generateMarkdownDocs", false, "generate markdown docs")
	app.BoolVar(&quiet, "quiet", false, "suppress error messages")

	// Application Options
	app.BoolVar(&showLength, "length", false, "return the number of keys or values")
	app.BoolVar(&showLast, "last", false, "return the index of the last element in list (e.g. length - 1)")
	app.BoolVar(&showValues, "values", false, "return the values instead of the keys")
	app.StringVar(&delimiter, "d,delimiter", "", "set delimiter for range output")
	app.IntVar(&limit, "limit", 0, "limit the number of items output")

	// Parse options and environment
	app.Parse()
	args := app.Args()

	// Setup IO
	var err error

	app.Eout = os.Stderr

	app.In, err = cli.Open(inputFName, os.Stdin)
	cli.ExitOnError(os.Stderr, err, quiet)
	defer cli.CloseFile(inputFName, app.In)

	app.Out, err = cli.Create(outputFName, os.Stdout)
	cli.ExitOnError(os.Stderr, err, quiet)
	defer cli.CloseFile(outputFName, app.Out)

	// Process options
	if generateMarkdownDocs {
		app.GenerateMarkdownDocs(app.Out)
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

	// If no args then assume "." is desired
	if len(args) == 0 {
		args = []string{"."}
	}

	// Read in the complete JSON data structure
	buf, err := ioutil.ReadAll(app.In)
	cli.ExitOnError(app.Eout, err, quiet)

	if len(buf) == 0 {
		cli.ExitOnError(app.Eout, fmt.Errorf("no data"), quiet)
	}

	var (
		data interface{}
	)

	if len(delimiter) == 0 {
		delimiter = "\n"
	} else {
		delimiter = datatools.NormalizeDelimiter(delimiter)
	}

	for _, p := range args {
		if p == "." {
			data, err = dotpath.JSONDecode(buf)
		} else {
			data, err = dotpath.EvalJSON(p, buf)
		}
		cli.ExitOnError(app.Eout, err, quiet)

		switch {
		case showLength:
			l, err := getLength(data)
			cli.ExitOnError(app.Eout, err, quiet)
			fmt.Fprintf(app.Out, "%d", l)
		case showLast:
			l, err := getLength(data)
			cli.ExitOnError(app.Eout, err, quiet)
			fmt.Fprintf(app.Out, "%d", l-1)
		case showValues:
			elems, err := srcVals(data, limit-1)
			cli.ExitOnError(app.Eout, err, quiet)
			fmt.Fprintln(app.Out, strings.Join(elems, delimiter))
		default:
			elems, err := srcKeys(data, limit-1)
			cli.ExitOnError(app.Eout, err, quiet)
			fmt.Fprintln(app.Out, strings.Join(elems, delimiter))
		}
	}
}
