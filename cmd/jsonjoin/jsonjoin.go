// jsonjoin is a command line tool that takes two JSON documents and combined
// them into one depending on the options
//
// @author R. S. Doiel, <rsdoiel@caltech.edu>
package main

import (
	"flag"
	"fmt"
	"io/ioutil"
	"os"
	"path"
	"strings"

	// Caltech Library Packages
	"github.com/caltechlibrary/datatools"
)

var (
	helpText = `%{app_name}(1) irdmtools user manual | version {version} {release_hash}
% R. S. Doiel
% {release_date}

# NAME

{app_name} 

# SYNOPSIS

{app_name} [OPTIONS] JSON_FILENAME [JSON_FILENAME ...]

# DESCRIPTION

{app_name} joins one or more JSON objects. By default the
objects are each assigned to an attribute corresponding with their
filenames minus the ".json" extension. If the object is read from
standard input then "_" is used as it's attribute name.

If you use the update or overwrite options you will create a merged
object. The update option keeps the attribute value first encountered
and overwrite takes the last attribute value encountered.

# OPTIONS

-help
: display help

-license
: display license

-version:
display version

-nl, -newline
: if true add a trailing newline

-o, -output
: output filename

-p, -pretty
: pretty print output

-quiet
: suppress error messages

-create
: Create a root object placing each joined objects under their own attribute

-update
: update first object with the second object, ignore existing attributes

-overwrite
: update first object with the second object, overwriting existing attributes

# EXAMPLES

This is an example of take "my1.json" and "my2.json"
render "my.json"

~~~
    jsonjoin my1.json my2.json >my.json
~~~

my.json would have two attributes, "my1" and "my2" each
with their complete attributes.

Using the update option you can merge my1.json with any additional attribute
values found in m2.json.

~~~
    jsonjoin -update my1.json my2.json >my.json
~~~

Using the overwrite option you can merge my1.json with my2.json accepted
as replacement values.

~~~
    jsonjoin -overwrite my1.json my2.json >my.json
~~~






`

	// Standard Options
	showHelp         bool
	showLicense      bool
	showVersion      bool
	showExamples     bool
	pretty           bool
	outputFName      string
	generateMarkdown bool
	generateManPage  bool
	quiet            bool
	newLine          bool
	eol              string

	// Application Specific Options
	update     bool
	overwrite  bool
	createRoot bool
)

func main() {
	appName := path.Base(os.Args[0])

	// Standard Options
	flag.BoolVar(&showHelp, "help", false, "display help")
	flag.BoolVar(&showLicense, "license", false, "display license")
	flag.BoolVar(&showVersion, "version", false, "display version")

	flag.StringVar(&outputFName, "o", "", "output filename")
	flag.StringVar(&outputFName, "output", "", "output filename")
	flag.BoolVar(&quiet, "quiet", false, "suppress error messages")
	flag.BoolVar(&newLine, "nl", false, "if true add a trailing newline")
	flag.BoolVar(&newLine, "newline", false, "if true add a trailing newline")
	flag.BoolVar(&pretty, "p", false, "pretty print json")
	flag.BoolVar(&pretty, "pretty", false, "pretty print json")

	// Application Specific Options
	flag.BoolVar(&createRoot, "create", false, "for each object joined each under their own attribute.")
	flag.BoolVar(&update, "update", false, "copy new key/values pairs into root object")
	flag.BoolVar(&overwrite, "overwrite", false, "copy all key/values into root object")

	// Parse env amd options
	flag.Parse()
	args := flag.Args()

	// Setup IO
	var err error

	in := os.Stdin
	out := os.Stdout
	eout := os.Stderr

	// Process options
	if showHelp {
		fmt.Fprintf(out, "%s\n", datatools.FmtHelp(helpText, appName, datatools.Version, datatools.ReleaseDate, datatools.ReleaseHash))
		os.Exit(0)
	}
	if showLicense {
		fmt.Fprintf(out, "%s\n", datatools.LicenseText)
		os.Exit(0)
	}
	if showVersion {
		fmt.Fprintf(out, "%s %s %s\n", appName, datatools.Version, datatools.ReleaseHash)
		os.Exit(0)
	}
	if newLine {
		eol = "\n"
	}

	// Make sure we have some JSON objects to join...
	if len(args) < 1 {
		if err != nil {
			fmt.Fprintln(eout, err)
			os.Exit(1)
		}
	}

	outObject := map[string]interface{}{}
	newObject := map[string]interface{}{}

	for _, arg := range args {
		var src []byte
		if arg != "" && arg != "-" {
			src, err = ioutil.ReadFile(arg)
			if err != nil {
				fmt.Fprintln(eout, err)
				os.Exit(1)
			}
		} else {
			src, err = ioutil.ReadAll(in)
		}
		err = datatools.JSONUnmarshal(src, &newObject)
		if err != nil {
			fmt.Fprintln(eout, err)
			os.Exit(1)
		}
		switch {
		case update == true:
			for k, v := range newObject {
				if _, ok := outObject[k]; ok != true {
					outObject[k] = v
				}
			}
		case overwrite == true:
			for k, v := range newObject {
				outObject[k] = v
			}
		default:
			key := strings.TrimSuffix(path.Base(arg), ".json")
			if key == "" {
				key = "_"
			}
			outObject[key] = newObject
		}
	}

	var (
		src []byte
	)

	if pretty {
		src, err = datatools.JSONMarshalIndent(outObject, "", "    ")
	} else {
		src, err = datatools.JSONMarshal(outObject)
	}
	if err != nil {
		fmt.Fprintln(eout, err)
		os.Exit(1)
	}
	fmt.Fprintf(out, "%s%s", src, eol)
}
