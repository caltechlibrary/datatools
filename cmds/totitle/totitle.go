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
	"bufio"
	"flag"
	"fmt"
	"os"
	"path"
	"strings"

	// Caltech Library Packages
	"github.com/caltechlibrary/cli"
	"github.com/caltechlibrary/datatools"
)

var (
	usage = `USAGE: %s [OPTIONS] [STRINGS]`

	description = `

SYNOPSIS

%s turns a UTF-8 string(s) into title case (in English the same as upper case).
If the option -c or -capitalize is used it'll use a naive approach to
captialization rather than title case. If the option -ce or -capitalize-english
added rules English specific capitalization rules will be used.

`

	examples = `

EXAMPLE

Title case the string "the cat in the hat"

    %s "the cat in the hat"

This should yield

    THE CAT IN THE HAT

Usage -c or -capitalize option "the cat in the hat"

    %s -c "the cat in the hat"

should yeild

    "The Cat In The Hat"

Using -ce or -capitalize-english option "the cat in the hat"

    %s -ce "the cat in the hat"

should yeild

    "The Cat in the Hat"

`

	// Standard Options
	showHelp     bool
	showLicense  bool
	showVersion  bool
	showExamples bool
	inputFName   string
	outputFName  string
	quiet        bool

	// App Options
	newLine           bool
	capitalize        bool
	capitalizeEnglish bool
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
	flag.BoolVar(&quiet, "quiet", false, "suppress error messages")

	// Application specific options
	flag.BoolVar(&newLine, "nl", true, "output a newline")
	flag.BoolVar(&newLine, "newline", true, "output a newline")
	flag.BoolVar(&capitalize, "c", false, "capitalize words")
	flag.BoolVar(&capitalize, "capitalize", false, "capitalize words")
	flag.BoolVar(&capitalizeEnglish, "ce", false, "english capitalization of words")
	flag.BoolVar(&capitalizeEnglish, "capitalize-english", false, "english capitalization of words")
}

func main() {
	appName := path.Base(os.Args[0])
	flag.Parse()
	args := flag.Args()

	cfg := cli.New(appName, "", datatools.Version)
	cfg.UsageText = fmt.Sprintf(usage, appName)
	cfg.DescriptionText = fmt.Sprintf(description, appName)
	cfg.OptionText = "OPTIONS\n\n"
	cfg.ExampleText = fmt.Sprintf(examples, appName, appName, appName)

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
	cli.ExitOnError(os.Stderr, err, quiet)
	defer cli.CloseFile(inputFName, in)

	out, err := cli.Create(outputFName, os.Stdout)
	cli.ExitOnError(os.Stderr, err, quiet)
	defer cli.CloseFile(outputFName, out)

	nl := "\n"
	if newLine == false {
		nl = ""
	}

	if inputFName != "" {
		scanner := bufio.NewScanner(in)
		for scanner.Scan() {
			s := scanner.Text()
			if s != "" {
				switch {
				case capitalizeEnglish == true:
					fmt.Fprintf(out, "%s%s", datatools.EnglishTitle(s), nl)
				case capitalize == true:
					fmt.Fprintf(out, "%s%s", strings.Title(s), nl)
				default:
					fmt.Fprintf(out, "%s%s", strings.ToTitle(s), nl)
				}
			}
		}
		err = scanner.Err()
		cli.OnError(os.Stderr, err, quiet)
	}

	for _, s := range args {
		switch {
		case capitalizeEnglish == true:
			fmt.Fprintf(out, "%s%s", datatools.EnglishTitle(s), nl)
		case capitalize == true:
			fmt.Fprintf(out, "%s%s", strings.Title(s), nl)
		default:
			fmt.Fprintf(out, "%s%s", strings.ToTitle(s), nl)
		}
	}
}
