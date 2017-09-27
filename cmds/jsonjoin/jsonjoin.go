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
	usage = `USAGE: %s [OPTIONS] JSON_FILE_1 JSON_FILE_2`

	description = `
SYSNOPSIS

%s is a command line tool that takes two (or more) JSON object 
documents and combines into a new JSON object document based on 
the options chosen.
`

	examples = `
EXAMPLES

    Joining two JSON objects (maps)

person.json containes

   {"name": "Doe, Jane", "email":"jd@example.org", "age": 42}

profile.json containes

   {"name": "Doe, Jane", "bio": "World renowned geophysist.",
   	"email": "jane.doe@example.edu"}

A simple join of person.json with profile.json

   %s person.json profile.json

would yeild

   {
   	"person": {"name": "Doe, Jane", "email":"jd@example.org", "age": 42},
    "profile": {"name": "Doe, Jane", "bio": "World renowned geophysist.", 
				"email": "jane.doe@example.edu"}
	}

You can modify this behavor with -add or -merge. Both options are
order dependant (i.e. not guaranteed to be associative, A add B does
not necessarily equal B add A). 

+ -update will add unique key/values from the second object to the first object
+ -overwrite replace key/values in first object one with second objects'

Running

	%s -update person.json profile.json

would yield

   { "name": "Doe, Jane", "email":"jd@example.org", "age": 42,
     "bio": "World renowned geophysist." }

Running

	%s -update profile.json person.json

would yield
   
   	{ "name": "Doe, Jane",  "age": 42, 
		"bio": "World renowned geophysist.", 
		"email": "jane.doe@example.edu" }

Running 

	%s -overwrite person.json profile.json

would yield

   	{ "name": "Doe, Jane", "email":"jane.doe@example.edu", "age": 42,
    	"bio": "World renowned geophysist." }
`

	// Basic Options
	showHelp    bool
	showLicense bool
	showVersion bool
	outputFName string

	// Application Specific Options
	update    bool
	overwrite bool
)

func init() {
	// Basic Options
	flag.BoolVar(&showHelp, "h", false, "display help")
	flag.BoolVar(&showHelp, "help", false, "display help")
	flag.BoolVar(&showLicense, "l", false, "display license")
	flag.BoolVar(&showLicense, "license", false, "display license")
	flag.BoolVar(&showVersion, "v", false, "display version")
	flag.BoolVar(&showVersion, "version", false, "display version")
	flag.StringVar(&outputFName, "o", "", "output filename")
	flag.StringVar(&outputFName, "output", "", "output filename")

	// Application Specific Options
	flag.BoolVar(&update, "update", false, "copy unique key/values from second object into the first")
	flag.BoolVar(&overwrite, "overwrite", false, "copy all key/values from second object into the first")
}

func main() {
	appName := path.Base(os.Args[0])
	flag.Parse()
	args := flag.Args()

	// Configuration and command line interation
	cfg := cli.New(appName, "DATATOOLS", fmt.Sprintf(datatools.LicenseText, appName, datatools.Version), datatools.Version)
	cfg.UsageText = fmt.Sprintf(usage, appName)
	cfg.DescriptionText = fmt.Sprintf(description, appName)
	cfg.ExampleText = fmt.Sprintf(examples, appName, appName, appName, appName)

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

	if len(args) == 0 {
		fmt.Fprintf(os.Stderr, "Need to provide at least one template name\n")
		os.Exit(1)
	}

	out, err := cli.Create(outputFName, os.Stdout)
	if err != nil {
		fmt.Fprintf(os.Stderr, "%s\n", err)
		os.Exit(1)
	}
	defer cli.CloseFile(outputFName, out)

	if len(args) < 2 {
		fmt.Println(cfg.Usage())
		os.Exit(1)
	}

	outObject := map[string]interface{}{}
	newObject := map[string]interface{}{}

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
