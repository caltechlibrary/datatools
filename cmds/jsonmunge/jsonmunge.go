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
	"strings"
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

This would yield

    "Doe, Jane"

`

	// Basic Options
	showHelp     bool
	showLicense  bool
	showVersion  bool
	showExamples bool
	inputFName   string
	outputFName  string
	quiet        bool

	// Application Specific Options
	templateExpr string
)

func init() {
	// Basic Options
	flag.BoolVar(&showHelp, "h", false, "display help")
	flag.BoolVar(&showHelp, "help", false, "display help")
	flag.BoolVar(&showLicense, "license", false, "display license")
	flag.BoolVar(&showVersion, "v", false, "display version")
	flag.BoolVar(&showVersion, "version", false, "display version")
	flag.BoolVar(&showExamples, "example", false, "display example(s)")
	flag.StringVar(&inputFName, "i", "", "input filename")
	flag.StringVar(&inputFName, "input", "", "input filename")
	flag.StringVar(&outputFName, "o", "", "output filename")
	flag.StringVar(&outputFName, "output", "", "output filename")
	flag.BoolVar(&quiet, "quiet", false, "suppress error messages")

	// Application Specific Options
	flag.StringVar(&templateExpr, "E", "", "use template expression as template")
	flag.StringVar(&templateExpr, "expression", "", "use template expression as template")
}

func main() {
	appName := path.Base(os.Args[0])
	flag.Parse()
	args := flag.Args()

	// Configuration and command line interation
	cfg := cli.New(appName, strings.ToUpper(appName), datatools.Version)
	cfg.LicenseText = fmt.Sprintf(datatools.LicenseText, appName, datatools.Version)
	cfg.UsageText = fmt.Sprintf(usage, appName)
	cfg.DescriptionText = fmt.Sprintf(description, appName)
	cfg.OptionText = "OPTIONS\n\n"
	cfg.ExampleText = fmt.Sprintf(examples, appName)

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

	if len(args) == 0 && templateExpr == "" {
		cli.ExitOnError(os.Stderr, fmt.Errorf("Need to provide at least one template name"), quiet)
	}

	var (
		tmpl *template.Template
		err  error
	)
	if templateExpr != "" {
		// Read in and compile our templates expression
		tmpl, err = template.New("default").Funcs(tmplfn.AllFuncs()).Parse(templateExpr)
		cli.ExitOnError(os.Stderr, err, quiet)
	} else {
		// Read in and compile our templates
		tmpl, err = template.New(path.Base(args[0])).Funcs(tmplfn.AllFuncs()).ParseFiles(args...)
		cli.ExitOnError(os.Stderr, err, quiet)
	}

	in, err := cli.Open(inputFName, os.Stdin)
	cli.ExitOnError(os.Stderr, err, quiet)
	defer cli.CloseFile(inputFName, in)

	out, err := cli.Create(outputFName, os.Stdout)
	cli.ExitOnError(os.Stderr, err, quiet)
	defer cli.CloseFile(outputFName, out)

	// READ in the JSON document
	buf, err := ioutil.ReadAll(in)
	cli.ExitOnError(os.Stderr, err, quiet)

	// JSON Decode our document
	data, err := dotpath.JSONDecode(buf)
	cli.ExitOnError(os.Stderr, err, quiet)

	// Execute template with data
	err = tmpl.Execute(out, data)
	cli.ExitOnError(os.Stderr, err, quiet)
}
