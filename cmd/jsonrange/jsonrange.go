//
// jsonrange iterates over an array or map returning either a JSON expression
// or map keep to stdout
//
// @Author R. S. Doiel, <rsdoiel@caltech.edu>
//
package main

import (
	"flag"
	"fmt"
	"io/ioutil"
	"os"
	"path"
	"strings"

	// CaltechLibrary Packages
	"github.com/caltechlibrary/datatools"
	"github.com/caltechlibrary/dotpath"
)

var (
	helpText =`---
title: "{app_name} (1) user manual"
author: "R. S. Doiel"
pubDate: 2023-01-09
---

# NAME

{app_name} 

# SYNOPSIS

{app_name} [OPTIONS] [DOT_PATH_EXPRESSION]

# DESCRIPTION

{app_name} returns returns a range of values based on the JSON structure
being read and options applied.  Without options the JSON structure is
read from standard input and writes a list of keys to standard out. Keys
are either attribute names or for arrays the index position (counting
form zero).  If a DOT_PATH_EXPRESSION is included on the command line then
that is used to generate the results. Using options to can choose to read
the JSON data structure from a file, write the output to a file as well
as display values instead of keys. a list of "keys" of an index or map in
JSON.  

Using options it can also return a list of values.  The JSON object is
read from standard in and the resulting list is normally written to
standard out. There are options to read or write to files.  Additional
parameters are assumed to be a dot path notation select the parts of the
JSON data structure you want from the range. 

DOT_PATH_EXPRESSION is a dot path stale expression indicating what you
want range over.  E.g.

- . would indicate the whole JSON data structure read is used to range over
- .name would indicate to range over the value pointed at by the "name" attribute 
- ["name"] would indicate to range over the value pointed at by the "name" attribute
- [0] would indicate to range over the value held in the zero-th element of the array

The path can be chained together

- .name.family would point to the value heald by the "name" attributes' "family" attribute.

# OPTIONS

-help
: display help

-license
: display license

-version
: display version

-d, -delimiter
: set delimiter for range output

-i, -input
: read JSON from file

-last
: return the index of the last element in list (e.g. length - 1)

-length
: return the number of keys or values

-limit
: limit the number of items output

-nl, -newline
: if true add a trailing newline

-o, -output
: write to output file

-quiet
: suppress error messages

-values
: return the values instead of the keys


# EXAMPLES

Working with a map

~~~
    echo '{"name": "Doe, Jane", "email":"jane.doe@example.org", "age": 42}' \
       | {app_name}
~~~

This would yield

~~~
    name
    email
    age
~~~

Using the -values option on a map

~~~
    echo '{"name": "Doe, Jane", "email":"jane.doe@example.org", "age": 42}' \
      | {app_name} -values
~~~

This would yield

~~~
    "Doe, Jane"
    "jane.doe@example.org"
    42
~~~


Working with an array

~~~
    echo '["one", 2, {"label":"three","value":3}]' | {app_name}
~~~

would yield

~~~
    0
    1
    2
~~~

Using the -values option on the same array

~~~
    echo '["one", 2, {"label":"three","value":3}]' | {app_name} -values
~~~

would yield

~~~
    one
    2
    {"label":"three","value":3}
~~~

Checking the length of a map or array or number of keys in map

~~~
    echo '["one","two","three"]' | {app_name} -length
~~~

would yield

~~~
    3
~~~

Check for the index of last element

~~~
    echo '["one","two","three"]' | {app_name} -last
~~~

would yield

~~~
    2
~~~

Check for the index value of last element

~~~
    echo '["one","two","three"]' | {app_name} -values -last
~~~

would yield

~~~
    "three"
~~~

Limitting the number of items returned

~~~
    echo '[10,20,30,40,50]' | %!s(MISSING) -limit 2
~~~

would yield

~~~
    1
    2
~~~

Limitting the number of values returned

~~~
    echo '[10,20,30,40,50]' | %!s(MISSING) -values -limit 2
~~~

would yield

~~~
    10
    20
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
	eol              string

	// Application Specific Options
	showLength bool
	showLast   bool
	showValues bool
	delimiter  string
	limit      int
)

func fmtTxt(src string, appName string, version string) string {
	return strings.ReplaceAll(strings.ReplaceAll(src, "{app_name}", appName), "{version}", version)
}

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
		outSrc, err := datatools.JSONMarshal(val)
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
		outSrc, err := datatools.JSONMarshal(val)
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
	appName := path.Base(os.Args[0])

	// Standard Options
	flag.BoolVar(&showHelp, "help", false, "display help")
	flag.BoolVar(&showLicense, "license", false, "display license")
	flag.BoolVar(&showVersion, "version", false, "display version")

	flag.StringVar(&inputFName, "i", "", "read JSON from file")
	flag.StringVar(&inputFName, "input", "", "read JSON from file")
	flag.StringVar(&outputFName, "o", "", "write to output file")
	flag.StringVar(&outputFName, "output", "", "write to output file")
	

	flag.BoolVar(&quiet, "quiet", false, "suppress error messages")
	flag.BoolVar(&newLine, "nl", false, "if true add a trailing newline")
	flag.BoolVar(&newLine, "newline", false, "if true add a trailing newline")

	// Application Options
	flag.BoolVar(&showLength, "length", false, "return the number of keys or values")
	flag.BoolVar(&showLast, "last", false, "return the index of the last element in list (e.g. length - 1)")
	flag.BoolVar(&showValues, "values", false, "return the values instead of the keys")
	flag.StringVar(&delimiter, "d", "", "set delimiter for range output")
	flag.StringVar(&delimiter, "delimiter", "", "set delimiter for range output")
	flag.IntVar(&limit, "limit", -1, "limit the number of items output")

	// Parse options and environment
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
			fmt.Fprintln(eout, err)
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

	// If no args then assume "." is desired
	if len(args) == 0 {
		args = []string{"."}
	}

	// Read in the complete JSON data structure
	buf, err := ioutil.ReadAll(in)
	if err != nil {
		fmt.Fprintln(eout, err)
		os.Exit(1)
	}

	if len(buf) == 0 {
		if err != nil {
			fmt.Fprintln(eout, err)
			os.Exit(1)
		}
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
		if err != nil {
			fmt.Fprintln(eout, err)
			os.Exit(1)
		}

		switch {
		case showLength:
			l, err := getLength(data)
			if err != nil {
				fmt.Fprintln(eout, err)
				os.Exit(1)
			}
			fmt.Fprintf(out, "%d", l)
		case showLast:
			l, err := getLength(data)
			if err != nil {
				fmt.Fprintln(eout, err)
				os.Exit(1)
			}
			if showValues {
				elems, err := srcVals(data, limit)
				if err != nil {
					fmt.Fprintln(eout, err)
					os.Exit(1)
				}
				l := len(elems)
				fmt.Fprintf(out, "%s", elems[l-1])
			} else {
				fmt.Fprintf(out, "%d", l-1)
			}
		case showValues:
			elems, err := srcVals(data, limit)
			if err != nil {
				fmt.Fprintln(eout, err)
				os.Exit(1)
			}
			fmt.Fprintln(out, strings.Join(elems, delimiter))
		default:
			elems, err := srcKeys(data, limit)
			if err != nil {
				fmt.Fprintln(eout, err)
				os.Exit(1)
			}
			fmt.Fprintln(out, strings.Join(elems, delimiter))
		}
	}
	fmt.Fprintf(out, "%s", eol)
}
