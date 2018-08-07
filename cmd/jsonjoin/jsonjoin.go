//
// jsonjoin is a command line tool that takes two JSON documents and combined
// them into one depending on the options
//
// @author R. S. Doiel, <rsdoiel@caltech.edu>
//
package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
	"path"
	"strings"

	// Caltech Library Packages
	"github.com/caltechlibrary/cli"
	"github.com/caltechlibrary/datatools"
)

var (
	description = `
%s is a command line tool that takes one (or more) JSON objects files
and joins them to a root JSON object read from standard input (or
file identified by -input option).  By default the resulting
joined JSON object is written to standard out.

The default behavior for %s is to create key/value pairs
based on the joined JSON document names and their contents.
This can be thought of as a branching behavior. Each additional
file becomes a branch and its key/value pairs become leafs.
The root JSON object is assumed to come from standard input
but can be designated by the -input option or created by the
-create option. Each additional file specified as a command line
argument is then treated as a new branch.

In addition to the branching behavior you can join JSON objects in a
flat manner.  The flat joining process can be ether non-destructive
adding new key/value pairs (-update option) or destructive
overwriting key/value pairs (-overwrite option).

Note: %s doesn't support a JSON array as the root JSON object.
`

	examples = `
Consider two JSON objects one in person.json and another
in profile.json.

person.json contains

   { "name": "Doe, Jane", "email":"jd@example.org", "age": 42 }

profile.json contains

   { "name": "Doe, Jane", "bio": "World renowned geophysist.",
     "email": "jane.doe@example.edu" }

A simple join of person.json with profile.json (note the
-create option)

   %s -create person.json profile.json

would yield and object like

   {
     "person":  { "name": "Doe, Jane", "email":"jd@example.org",
	 			"age": 42},
     "profile": { "name": "Doe, Jane", "bio": "World renowned geophysist.",
                  "email": "jane.doe@example.edu" }
   }

Likewise if you want to treat person.json as the root object and add
profile.json as a branch try

   cat person.json | %s profile.json

or

   %s -i person.json profile.json

this yields an object like

   {
     "name": "Doe, Jane", "email":"jd@example.org", "age": 42,
     "profile": { "name": "Doe, Jane", "bio": "World renowned geophysist.",
                  "email": "jane.doe@example.edu" }
   }

You can modify this behavor with -update or -overwrite. Both options are
order dependant (i.e. not associative, A update B does
not necessarily equal B update A).

+ -update will add unique key/values from the second object to the first object
+ -overwrite replace key/values in first object one with second objects'

Running

    %s -create -update person.json profile.json

would yield

   { "name": "Doe, Jane", "email":"jd@example.org", "age": 42,
     "bio": "World renowned geophysist." }

Running

    %s -create -update profile.json person.json

would yield

   { "name": "Doe, Jane",  "age": 42,
     "bio": "World renowned geophysist.",
     "email": "jane.doe@example.edu" }

Running

    %s -create -overwrite person.json profile.json

would yield

   { "name": "Doe, Jane", "email":"jane.doe@example.edu", "age": 42,
     "bio": "World renowned geophysist." }
`

	// Standard Options
	showHelp         bool
	showLicense      bool
	showVersion      bool
	showExamples     bool
	inputFName       string
	outputFName      string
	generateMarkdown bool
	quiet            bool
	newLine          bool
	eol              string

	// Application Specific Options
	update     bool
	overwrite  bool
	createRoot bool
)

func main() {
	app := cli.NewCli(datatools.Version)
	appName := app.AppName()

	// Add Help Docs
	app.AddHelp("license", []byte(fmt.Sprintf(datatools.LicenseText, appName, datatools.Version)))
	app.AddHelp("description", []byte(fmt.Sprintf(description, appName, appName, appName)))
	app.AddHelp("examples", []byte(fmt.Sprintf(examples, appName, appName, appName, appName, appName, appName)))

	// Document non-option parameter
	app.AddParams("JSON_FILE_1", "[JSON_FILE_2 ...]")

	// Standard Options
	app.BoolVar(&showHelp, "h,help", false, "display help")
	app.BoolVar(&showLicense, "l,license", false, "display license")
	app.BoolVar(&showVersion, "v,version", false, "display version")
	app.BoolVar(&showExamples, "examples", false, "display example(s)")
	app.StringVar(&inputFName, "i,input", "", "input filename (for root object)")
	app.StringVar(&outputFName, "o,output", "", "output filename")
	app.BoolVar(&generateMarkdown, "generate-markdown", false, "generate markdown docs")
	app.BoolVar(&quiet, "quiet", false, "suppress error messages")
	app.BoolVar(&newLine, "nl,newline", false, "if true add a trailing newline")

	// Application Specific Options
	app.BoolVar(&createRoot, "create", false, "create an empty root object, {}")
	app.BoolVar(&update, "update", false, "copy new key/values pairs into root object")
	app.BoolVar(&overwrite, "overwrite", false, "copy all key/values into root object")

	// Parse env amd options
	app.Parse()
	args := app.Args()

	// Setup IO
	var err error

	app.Eout = os.Stderr

	app.In, err = cli.Open(inputFName, os.Stdin)
	cli.ExitOnError(app.Eout, err, quiet)
	defer cli.CloseFile(inputFName, app.In)

	app.Out, err = cli.Create(outputFName, os.Stdout)
	cli.ExitOnError(app.Eout, err, quiet)
	defer cli.CloseFile(outputFName, app.Out)

	// Process options
	if generateMarkdown {
		app.GenerateMarkdown(app.Out)
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
	if newLine {
		eol = "\n"
	}

	// Make sure we have some JSON objects to join...
	if len(args) < 1 {
		cli.ExitOnError(app.Eout, fmt.Errorf("Missing JSON document(s) to join"), quiet)
	}

	outObject := map[string]interface{}{}
	newObject := map[string]interface{}{}

	// READ in the JSON document if present on standard in or specified with -i.
	if createRoot == false {
		buf, err := ioutil.ReadAll(app.In)
		cli.ExitOnError(app.Eout, err, quiet)
		err = json.Unmarshal(buf, &outObject)
		cli.ExitOnError(app.Eout, err, quiet)
	}

	for _, arg := range args {
		src, err := ioutil.ReadFile(arg)
		cli.ExitOnError(app.Eout, err, quiet)
		err = json.Unmarshal(src, &newObject)
		cli.ExitOnError(app.Eout, err, quiet)
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
			outObject[key] = newObject
		}
	}

	src, err := json.Marshal(outObject)
	cli.ExitOnError(app.Eout, err, quiet)
	fmt.Fprintf(app.Out, "%s%s", src, eol)
}
