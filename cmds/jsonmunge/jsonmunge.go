//
// jsonmunge is a command line tool that takes a JSON document and
// a Go text/template rendering the result. Useful for
// reshaping a JSON document or transforming into a new format,
// or filter for specific content.
//
// @author R. S. Doiel, <rsdoiel@caltech.edu>
//
package main

import (
	"flag"
	"fmt"
	"io/ioutil"
	"os"
	"path"
	"text/template"

	// Caltech Library Packages
	"github.com/caltechlibrary/cli"
	"github.com/caltechlibrary/datatools"
	"github.com/caltechlibrary/dotpath"
	"github.com/caltechlibrary/tmplfn"
)

var (
	usage = `USAGE: %s [OPTIONS] TEMPLATE_FILENAME`

	description = `
SYSNOPSIS

%s is a command line tool that takes a JSON document and
one or more Go templates rendering the results. Useful for
reshaping a JSON document, transforming into a new format,
or filter for specific content.

+ TEMPLATE_FILENAME is the name of a Go text tempate file used to render
  the outbound JSON document
`

	examples = `
EXAMPLES

If person.json contained

   {"name": "Doe, Jane", "email":"jd@example.org", "age": 42}

and the template, name.tmpl, contained 

   {{- .name -}}

Getting just the name could be done with

    cat person.json | %s name.tmpl

This would yeild

    "Doe, Jane"
`

	// Basic Options
	showHelp    bool
	showLicense bool
	showVersion bool
	inputFName  string
	outputFName string

	// Application Specific Options
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
}

func main() {
	appName := path.Base(os.Args[0])
	flag.Parse()
	args := flag.Args()

	// Configuration and command line interation
	cfg := cli.New(appName, "DATATOOLS", fmt.Sprintf(datatools.LicenseText, appName, datatools.Version), datatools.Version)
	cfg.UsageText = fmt.Sprintf(usage, appName)
	cfg.DescriptionText = fmt.Sprintf(description, appName)
	cfg.ExampleText = fmt.Sprintf(examples, appName)

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
	// Read in and compile our templates
	tmpl, err := template.New(path.Base(args[0])).Funcs(tmplfn.AllFuncs()).ParseFiles(args...)
	if err != nil {
		fmt.Fprintf(os.Stderr, "%s\n", err)
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

	// READ in the JSON document
	buf, err := ioutil.ReadAll(in)
	if err != nil {
		fmt.Fprintf(os.Stderr, "%s\n", err)
		os.Exit(1)
	}
	// JSON Decode our document
	data, err := dotpath.JSONDecode(buf)
	if err != nil {
		fmt.Fprintf(os.Stderr, "%s\n", err)
		os.Exit(1)
	}

	// Execute template with data
	if err := tmpl.Execute(out, data); err != nil {
		fmt.Fprintf(os.Stderr, "%s\n", err)
		os.Exit(1)
	}
}
