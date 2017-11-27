//
// jsonjoin is a command line tool that takes two JSON documents and combined
// them into one depending on the options
//
// @author R. S. Doiel, <rsdoiel@caltech.edu>
//
package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"path"
	"strings"

	// Caltech Library Packages
	"github.com/caltechlibrary/cli"
	"github.com/caltechlibrary/datatools"
)

var (
	usage = `USAGE: %s [OPTIONS] JSON_FILE_1 [JSON_FILE_2 ...]`

	description = `

SYSNOPSIS

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

EXAMPLE

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

	// Basic Options
	showHelp     bool
	showLicense  bool
	showVersion  bool
	showExamples bool
	inputFName   string
	outputFName  string

	// Application Specific Options
	update     bool
	overwrite  bool
	createRoot bool
)

func init() {
	// Basic Options
	flag.BoolVar(&showHelp, "h", false, "display help")
	flag.BoolVar(&showHelp, "help", false, "display help")
	flag.BoolVar(&showLicense, "l", false, "display license")
	flag.BoolVar(&showLicense, "license", false, "display license")
	flag.BoolVar(&showVersion, "v", false, "display version")
	flag.BoolVar(&showVersion, "version", false, "display version")
	flag.BoolVar(&showExamples, "example", false, "display example(s)")
	flag.StringVar(&inputFName, "i", "", "input filename (for root object)")
	flag.StringVar(&inputFName, "input", "", "input filename (for root object)")
	flag.StringVar(&outputFName, "o", "", "output filename")
	flag.StringVar(&outputFName, "output", "", "output filename")

	// Application Specific Options
	flag.BoolVar(&createRoot, "create", false, "create an empty root object, {}")
	flag.BoolVar(&update, "update", false, "copy new key/values pairs into root object")
	flag.BoolVar(&overwrite, "overwrite", false, "copy all key/values into root object")
}

func main() {
	appName := path.Base(os.Args[0])
	flag.Parse()
	args := flag.Args()

	// Configuration and command line interation
	cfg := cli.New(appName, strings.ToUpper(appName), datatools.Version)
	cfg.LicenseText = fmt.Sprintf(datatools.LicenseText, appName, datatools.Version)
	cfg.UsageText = fmt.Sprintf(usage, appName)
	cfg.DescriptionText = fmt.Sprintf(description, appName, appName, appName)
	cfg.OptionText = "OPTIONS\n\n"
	cfg.ExampleText = fmt.Sprintf(examples, appName, appName, appName, appName, appName, appName)

	if showHelp == true {
		if len(args) > 0 {
			fmt.Println(cfg.Help(args...))
		} else {
			fmt.Println(cfg.Usage())
		}
		os.Exit(0)
	}

	if showExamples == true {
		if len(args) > 0 {
			fmt.Println(cfg.Example(args...))
		} else {
			fmt.Println(cfg.ExampleText)
		}
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

	if len(args) == 0 {
		fmt.Fprintf(os.Stderr, "Need to provide at least one template name\n")
		os.Exit(1)
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

	// Make sure we have some JSON objects to join...
	if len(args) < 1 {
		fmt.Fprintln(os.Stderr, cfg.Usage())
		fmt.Fprintln(os.Stderr, "Missing JSON document(s) to join")
		os.Exit(1)
	}

	outObject := map[string]interface{}{}
	newObject := map[string]interface{}{}

	// READ in the JSON document if present on standard in or specified with -i.
	if createRoot == false {
		buf, err := ioutil.ReadAll(in)
		if err != nil {
			fmt.Fprintf(os.Stderr, "%s\n", err)
			os.Exit(1)
		}
		if err := json.Unmarshal(buf, &outObject); err != nil {
			log.Fatal(err)
		}
	}

	for _, arg := range args {
		src, err := ioutil.ReadFile(arg)
		if err != nil {
			log.Fatal(err)
		}
		if err := json.Unmarshal(src, &newObject); err != nil {
			log.Fatal(err)
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
			outObject[key] = newObject
		}
	}

	src, err := json.Marshal(outObject)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Fprintf(out, "%s", src)
}
