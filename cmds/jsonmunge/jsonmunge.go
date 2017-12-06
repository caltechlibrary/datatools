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
	description = `
%s is a command line tool that takes a JSON document and
one or more Go templates rendering the results. Useful for
reshaping a JSON document, transforming into a new format,
or filter for specific content.

+ TEMPLATE_FILENAME is the name of a Go text tempate file used to render
  the outbound JSON document
`

	examples = `
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
	showHelp             bool
	showLicense          bool
	showVersion          bool
	showExamples         bool
	inputFName           string
	outputFName          string
	generateMarkdownDocs bool
	quiet                bool
	newLine              bool
	eol                  string

	// Application Specific Options
	templateExpr string
)

func main() {
	app := cli.NewCli(datatools.Version)
	appName := app.AppName()

	// Add Help Docs
	app.AddHelp("license", []byte(fmt.Sprintf(datatools.LicenseText, appName, datatools.Version)))
	app.AddHelp("description", []byte(fmt.Sprintf(description, appName)))
	app.AddHelp("examples", []byte(fmt.Sprintf(examples, appName)))

	// Document non-option parameters
	app.AddParams("TEMPLATE_FILENAME")

	// Basic Options
	app.BoolVar(&showHelp, "h,help", false, "display help")
	app.BoolVar(&showLicense, "l,license", false, "display license")
	app.BoolVar(&showVersion, "v,version", false, "display version")
	app.BoolVar(&showExamples, "examples", false, "display example(s)")
	app.StringVar(&inputFName, "i,input", "", "input filename")
	app.StringVar(&outputFName, "o,output", "", "output filename")
	app.BoolVar(&generateMarkdownDocs, "generateMarkdownDocs", false, "generate markdown documentation")
	app.BoolVar(&quiet, "quiet", false, "suppress error messages")
	app.BoolVar(&newLine, "nl,newline", false, "if true add a trailing newline")

	// Application Specific Options
	app.StringVar(&templateExpr, "E,expression", "", "use template expression as template")

	// Parse env and options
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

	// Process Options
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
	if newLine {
		eol = "\n"
	}

	if len(args) == 0 && templateExpr == "" {
		cli.ExitOnError(app.Eout, fmt.Errorf("Need to provide at least one template name"), quiet)
	}

	var (
		tmpl *template.Template
	)
	if templateExpr != "" {
		// Read in and compile our templates expression
		tmpl, err = template.New("default").Funcs(tmplfn.AllFuncs()).Parse(templateExpr)
		cli.ExitOnError(app.Eout, err, quiet)
	} else {
		// Read in and compile our templates
		tmpl, err = template.New(path.Base(args[0])).Funcs(tmplfn.AllFuncs()).ParseFiles(args...)
		cli.ExitOnError(app.Eout, err, quiet)
	}

	// READ in the JSON document
	buf, err := ioutil.ReadAll(app.In)
	cli.ExitOnError(app.Eout, err, quiet)

	// JSON Decode our document
	data, err := dotpath.JSONDecode(buf)
	cli.ExitOnError(app.Eout, err, quiet)

	// Execute template with data
	err = tmpl.Execute(app.Out, data)
	cli.ExitOnError(app.Eout, err, quiet)
	fmt.Fprintf(app.Out, "%s", eol)
}
