//
// jsonrange iterates over an array or map returning either a JSON expression
// or map keep to stdout
//
// @Author R. S. Doiel, <rsdoiel@caltech.edu>
//
package main

import (
	"bytes"
	"encoding/json"
	"flag"
	"fmt"
	"os"
	"path"
	"strings"

	// 3rd Party Captech
	"github.com/rsdoiel/jid" // this is a fork of github.com/simeji/jid, adds Eval() and EvalString()

	// CaltechLibrary Packages
	"github.com/caltechlibrary/cli"
	"github.com/caltechlibrary/datatools"
)

var (
	usage = `USAGE: %s [OPTIONS] JSON_EXPRESSION `

	description = `
SYSNOPSIS

%s turns either the JSON expression that is a map or array into delimited
elements suitable for processing in a "for" style loop in Bash. If the
JSON expression is an array then the elements of the array are returned else
if the expression is a map/object then the keys or attribute names are turned.

+ EXPRESSION can be an empty string contains a JSON array or map.
`

	examples = `
EXAMPLES

Working with a map

    %s '{"name": "Doe, Jane", "email":"jane.doe@example.org", "age": 42}'

This would yield

    name
    email
    age

Working with an array

    %s '["one", 2, {"label":"three","value":3}]'

would yield

    one
    2
    {"label":"three","value":3}

Checking the length of a map or array

    %s -length '["one","two","three"]'

would yield

    3

Limitting the number of items returned

    %s -limit 2 '[1,2,3,4,5]'

would yield

    1
    2

Likewise you can have the JSON expression read from stnin

	echo '[1,2,3,4,5]' | %s -limit 2

would yield

    1
    2
`

	// Basic Options
	showHelp    bool
	showLicense bool
	showVersion bool
	inputFName  string
	outputFName string

	// Application Specific Options
	showLength bool
	delimiter  = "\n"
	limit      int
	dotPath    string
)

func srcKeys(inSrc string, limit int) ([]string, error) {
	data := map[string]interface{}{}
	if err := json.Unmarshal([]byte(inSrc), &data); err != nil {
		return nil, err
	}
	result := []string{}
	i := 0
	for keys := range data {
		result = append(result, keys)
		if limit > 0 && i == limit {
			return result, nil
		}
		i++
	}
	return result, nil
}

func srcVals(inSrc string, limit int) ([]string, error) {
	data := []interface{}{}
	if err := json.Unmarshal([]byte(inSrc), &data); err != nil {
		return nil, err
	}
	result := []string{}
	for i, val := range data {
		outSrc, err := json.Marshal(val)
		if err != nil {
			return nil, err
		}
		result = append(result, fmt.Sprintf("%s", outSrc))
		if limit != 0 && i == limit {
			return result, nil
		}
	}
	return result, nil
}

func init() {
	// Standard Options
	flag.BoolVar(&showHelp, "h", false, "display help")
	flag.BoolVar(&showLicense, "l", false, "display license")
	flag.BoolVar(&showVersion, "v", false, "display version")
	flag.StringVar(&inputFName, "i", "", "read JSON from file")
	flag.StringVar(&inputFName, "input", "", "read JSON from file")
	flag.StringVar(&outputFName, "o", "", "write to output file")
	flag.StringVar(&outputFName, "output", "", "write to output file")

	// Application Options
	flag.BoolVar(&showLength, "length", false, "return the number of keys or values")
	flag.StringVar(&delimiter, "d", "\n", "set delimiter for range output")
	flag.StringVar(&delimiter, "delimiter", "\n", "set delimiter for range output")
	flag.IntVar(&limit, "limit", 0, "limit the number of items output")
	flag.StringVar(&dotPath, "p", "", "range on given dot path")
	flag.StringVar(&dotPath, "dotpath", "", "range on given dot path")
}

func getLength(inSrc string) (int, error) {
	if strings.HasPrefix(inSrc, "{") {
		data := map[string]interface{}{}
		if err := json.Unmarshal([]byte(inSrc), &data); err != nil {
			return 0, err
		}
		return len(data), nil
	}
	data := []interface{}{}
	if err := json.Unmarshal([]byte(inSrc), &data); err != nil {
		return 0, err
	}
	return len(data), nil
}

func main() {
	appName := path.Base(os.Args[0])
	flag.Parse()

	args := flag.Args()

	// Configuration and command line interation
	cfg := cli.New(appName, "DATATOOLS", fmt.Sprintf(datatools.LicenseText, appName, datatools.Version), datatools.Version)
	cfg.UsageText = fmt.Sprintf(usage, appName)
	cfg.DescriptionText = fmt.Sprintf(description, appName)
	cfg.ExampleText = fmt.Sprintf(examples, appName, appName, appName, appName, appName)

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

	// OK, let's see what keys/values we're going to output...
	src := ""
	if len(args) == 0 {
		lines, err := cli.ReadLines(in)
		if err != nil {
			fmt.Fprintf(os.Stderr, "%s\n", err)
			os.Exit(1)
		}
		src = strings.Join(lines, "\n")
	} else {
		// If nothing is coming from stdin or a file then get the string from the command line
		if strings.HasPrefix(args[0], "'") == true {
			src = strings.Trim(args[0], "'")
		} else if strings.HasPrefix(args[0], "\"") == true {
			src = strings.Trim(args[0], "\"")
		} else {
			src = args[0]
		}
	}

	// If a dotPath is privided extract the desired field for range.
	if len(dotPath) > 0 {
		buf := bytes.NewBufferString(src)
		ea := &jid.EngineAttribute{
			DefaultQuery: ".",
			Monochrome:   true,
		}
		e, err := jid.NewEngine(buf, ea)
		if err != nil {
			fmt.Fprintf(os.Stderr, "%s\n", err)
			os.Exit(1)
		}
		result := e.EvalString(dotPath)

		if err := result.GetError(); err != nil {
			fmt.Fprintf(os.Stderr, "%s\n", err)
			os.Exit(1)
		}
		src = result.GetContent()
	}
	if len(src) == 0 {
		fmt.Println(cfg.Usage())
		os.Exit(1)
	}
	switch {
	case showLength:
		l, err := getLength(src)
		if err != nil {
			fmt.Fprintf(os.Stderr, "%s\n", err)
			os.Exit(1)
		}
		fmt.Fprintf(out, "%d", l)
	case strings.HasPrefix(src, "{"):
		elems, err := srcKeys(src, limit-1)
		if err != nil {
			fmt.Fprintf(os.Stderr, "%s\n", err)
			os.Exit(1)
		}
		fmt.Fprintln(out, strings.Join(elems, delimiter))
	case strings.HasPrefix(src, "["):
		elems, err := srcVals(src, limit-1)
		if err != nil {
			fmt.Fprintf(os.Stderr, "%s\n", err)
			os.Exit(1)
		}
		fmt.Fprintln(out, strings.Join(elems, delimiter))
	default:
		fmt.Fprintf(os.Stderr, "Cannot iterate over %q\n", src)
		os.Exit(1)
	}
}
