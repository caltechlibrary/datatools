//
// Copyright (c) 2017, Caltech
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
	"flag"
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
	usage = `USAGE: %s [OPTIONS] [STRINGS_TO_JOIN]`

	description = `

SYNOPSIS

%s joins a JSON array or with options a new line delimited string into
a single string with output delimiter (default delimiter is an empty string).

`

	examples = `

EXAMPLES

Joining a JSON array into a single string delimited by a double pipe.

    %s -d '||' '["one", "two", "three"]'

This should yield

    one||two||three

Joining a newline delimited file into a single string

    cat <<EOF > %s -nl -d '||' "one||two||three"
    one
	two
	three
    EOF
This should yield

    one||two||three

`

	// Standard Options
	showHelp     bool
	showLicense  bool
	showVersion  bool
	showExamples bool
	inputFName   string
	outputFName  string

	// App Options
	delimiter string
	newLine   bool
)

func init() {
	// Standard options
	flag.BoolVar(&showHelp, "h", false, "display help")
	flag.BoolVar(&showHelp, "help", false, "display help")
	flag.BoolVar(&showLicense, "l", false, "display license")
	flag.BoolVar(&showLicense, "license", false, "display license")
	flag.BoolVar(&showVersion, "v", false, "display version")
	flag.BoolVar(&showVersion, "version", false, "display version")
	flag.BoolVar(&showExamples, "example", false, "display example(s)")
	flag.StringVar(&inputFName, "i", "", "input filename")
	flag.StringVar(&inputFName, "input", "", "input filename")
	flag.StringVar(&outputFName, "o", "", "output filename")
	flag.StringVar(&outputFName, "output", "", "output filename")

	// Application specific options
	flag.StringVar(&delimiter, "d", "", "set the output delimiting string value (default is empty string)")
	flag.StringVar(&delimiter, "delimiter", "", "set output delimiting string value (default is empty string)")
	flag.BoolVar(&newLine, "nl", false, "input as one substring per line rather than JSON")
	flag.BoolVar(&newLine, "newline", false, "input as one substring per line rather than JSON")
}

func main() {
	appName := path.Base(os.Args[0])
	flag.Parse()
	args := flag.Args()

	cfg := cli.New(appName, "", datatools.Version)
	cfg.UsageText = fmt.Sprintf(usage, appName)
	cfg.DescriptionText = fmt.Sprintf(description, appName)
	cfg.OptionText = "OPTIONS\n\n"
	cfg.ExampleText = fmt.Sprintf(examples, appName, appName)

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

	// Normalize the delimiter if \n or \t
	switch delimiter {
	case `\n`:
		delimiter = "\n"
	case `\t`:
		delimiter = "\t"
	}

	results := []string{}
	if inputFName != "" {
		src, err := ioutil.ReadAll(in)
		if err != nil {
			fmt.Fprintf(os.Stderr, "Can't read file, %s", err)
			os.Exit(1)
		}
		if newLine == true {
			results = strings.Split(string(src), "\n")
		} else {
			err := json.Unmarshal(src, &results)
			if err != nil {
				fmt.Fprintf(os.Stderr, "expected a JSON array of strings, %s\n", err)
				os.Exit(1)
			}
		}
	}

	for _, arg := range args {
		if strings.HasPrefix(arg, "[") && strings.HasSuffix(arg, "]") {
			parts := []string{}
			if err := json.Unmarshal([]byte(arg), &parts); err != nil {
				fmt.Fprintf(os.Stderr, "expected a JSON array of strings, %s\n", err)
				os.Exit(1)
			}
			for _, s := range parts {
				results = append(results, s)
			}
		} else {
			results = append(results, arg)
		}
	}
	fmt.Fprintf(out, "%s", strings.Join(results, delimiter))
}
