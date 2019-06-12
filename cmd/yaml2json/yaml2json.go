//
// yaml2json is a command line utility that converts an YAML
// to JSON.
//
// @Author R. S. Doiel
//
// Copyright (c) 2019, Caltech
// All rights not granted herein are expressly reserved by Caltech.
//
// Redistribution and use in source and binary forms, with or without modification, are permitted provided that the following conditions are met:
//
// 1. Redistributions of source code must retain the above copyright notice, this list of conditions and the following disclaimer.
//
// 2. Redistributions in binary form must reproduce the above copyright notice, this list of conditions and the following disclaimer in the documentation and/or other materials provided with the distribution.
//
// 3. Neither the name of the copyright holder nor the names of its contributors may be used to endorse or promote products derived from this software without specific prior written permission.
//
// THIS SOFTWARE IS PROVIDED BY THE COPYRIGHT HOLDERS AND CONTRIBUTORS "AS IS" AND ANY EXPRESS OR IMPLIED WARRANTIES, INCLUDING, BUT NOT LIMITED TO, THE IMPLIED WARRANTIES OF MERCHANTABILITY AND FITNESS FOR A PARTICULAR PURPOSE ARE DISCLAIMED. IN NO EVENT SHALL THE COPYRIGHT HOLDER OR CONTRIBUTORS BE LIABLE FOR ANY DIRECT, INDIRECT, INCIDENTAL, SPECIAL, EXEMPLARY, OR CONSEQUENTIAL DAMAGES (INCLUDING, BUT NOT LIMITED TO, PROCUREMENT OF SUBSTITUTE GOODS OR SERVICES; LOSS OF USE, DATA, OR PROFITS; OR BUSINESS INTERRUPTION) HOWEVER CAUSED AND ON ANY THEORY OF LIABILITY, WHETHER IN CONTRACT, STRICT LIABILITY, OR TORT (INCLUDING NEGLIGENCE OR OTHERWISE) ARISING IN ANY WAY OUT OF THE USE OF THIS SOFTWARE, EVEN IF ADVISED OF THE POSSIBILITY OF SUCH DAMAGE.
//
package main

import (
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"os"

	// CaltechLibrary packages
	"github.com/caltechlibrary/cli"
	"github.com/caltechlibrary/datatools"

	// 3rd Party packages
	"github.com/ghodss/yaml"
)

var (
	description = `
%s is a tool that converts YAML into JSON output.
`

	examples = `
These would get the file named "my.yaml" and save it as my.json

    %s my.yaml > my.json

    %s my.yaml my.json

	cat my.yaml | %s -i - > my.json
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

	// Application Options
	prettyPrint bool
)

type Object struct {
	Obj map[string]interface{} `yaml:",inline"`
}

func yaml2JSON(in io.Reader, out io.Writer, printPrint bool) error {
	var (
		src []byte
		err error
	)
	src, err = ioutil.ReadAll(in)
	if err != nil {
		return err
	}

	src, err = yaml.YAMLToJSON(src)
	if err != nil {
		return err
	}

	if prettyPrint == true {
		m := map[string]interface{}{}
		err = json.Unmarshal(src, &m)
		if err != nil {
			return err
		}
		src, err = json.MarshalIndent(m, "", "    ")
		if err != nil {
			return err
		}
	}
	fmt.Fprintf(out, "%s", src)
	return nil
}

func main() {
	app := cli.NewCli(datatools.Version)
	appName := app.AppName()

	// Add Help Docs
	app.AddHelp("license", []byte(fmt.Sprintf(datatools.LicenseText, appName, datatools.Version)))
	app.AddHelp("description", []byte(fmt.Sprintf(description, appName)))
	app.AddHelp("examples", []byte(fmt.Sprintf(examples, appName, appName, appName)))

	// Document non-option parameters
	app.SetParams("[YAML_FILENAME]", "[JSON_NAME]")

	// Standard Options
	app.BoolVar(&showHelp, "h,help", false, "display help")
	app.BoolVar(&showLicense, "l,license", false, "display license")
	app.BoolVar(&showVersion, "v,version", false, "display version")
	app.BoolVar(&showExamples, "examples", false, "display example(s)")
	//app.StringVar(&inputFName, "i,input", "", "input filename")
	app.StringVar(&outputFName, "o,output", "", "output filename")
	app.BoolVar(&quiet, "quiet", false, "suppress error messages")
	app.BoolVar(&generateMarkdown, "generate-markdown", false, "generate markdown documentation")
	app.BoolVar(&generateManPage, "generate-manpage", false, "generate man page")
	app.BoolVar(&newLine, "nl,newline", false, "if true add a trailing newline")

	// App Specific Options
	app.BoolVar(&prettyPrint, "p,pretty", false, "pretty print output")

	// Parse env and options
	app.Parse()
	args := app.Args()

	// Handle case of input/output filenames provided without -i, -o
	if len(args) > 0 {
		inputFName = args[0]
		if len(args) > 1 {
			outputFName = args[1]
		}
	}

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
	if generateManPage {
		app.GenerateManPage(app.Out)
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

	err = yaml2JSON(app.In, app.Out, prettyPrint)
	cli.ExitOnError(app.Eout, err, quiet)
	fmt.Fprintf(app.Out, "%s", eol)
}
