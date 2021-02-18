//
// string is a command line utility to expose some of the Golang strings functions to the command line.
//
// @author R. S. Doiel, <rsdoiel@library.caltech.edu>
//
// Copyright (c) 2021, Caltech
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
	"io"
	"io/ioutil"
	"os"
	"strconv"
	"strings"

	// Caltech Library Packages
	"github.com/caltechlibrary/cli"
	"github.com/caltechlibrary/datatools"
)

var (
	description = `
string is a command line tool for transforming strings in common ways.

+ string length
+ changing cases
+ checking for prefixes, suffixes 
+ trimming prefixes, suffixes and cutsets (i.e. list of characters to cut)
+ position, counting and replacing substrings
+ splitting a string into a JSON array of strings, joining JSON a string arrays into a string
`

	examples = `
Convert text to upper case

	string toupper "one"

Convert text to lower case

	string tolower "ONE"

Captialize an English phrase

	string englishtitle "one more thing to know"

Split a space newline delimited list of words into a JSON array

	string -i wordlist.txt split "\n"

Join a JSON array of strings into a newline delimited list

	string join '\n' '["one","two","three","four","five"]'

`

	// Standard Options
	showHelp         bool
	showLicense      bool
	showVersion      bool
	showExamples     bool
	inputFName       string
	outputFName      string
	newLine          bool
	quiet            bool
	generateMarkdown bool
	generateManPage  bool
	eol              string

	// App Options
	delimiter       string
	outputDelimiter string
)

//
// Application functionality
//
func onError(eout io.Writer, err error, suppress bool) {
	if err != nil && suppress == false {
		fmt.Fprintln(eout, err)
	}
}

func exitOnError(eout io.Writer, err error, suppress bool) {
	if err != nil {
		onError(eout, err, suppress)
		os.Exit(1)
	}
}

func fnLength(in io.Reader, out io.Writer, eout io.Writer, args []string, flagSet *flag.FlagSet) int {
	err := flagSet.Parse(args)
	if err != nil {
		fmt.Fprintf(eout, "%s\n", err)
		return 1
	}
	args = flagSet.Args()

	if inputFName != "" {
		src, err := ioutil.ReadAll(in)
		exitOnError(eout, err, quiet)
		args = append(args, string(src))
	}
	for _, arg := range args {
		fmt.Fprintf(out, "%d%s", len(arg), eol)
	}
	return 0
}

func fnCount(in io.Reader, out io.Writer, eout io.Writer, args []string, flagSet *flag.FlagSet) int {
	err := flagSet.Parse(args)
	if err != nil {
		fmt.Fprintf(eout, "%s\n", err)
		return 1
	}
	args = flagSet.Args()
	if len(args) < 1 {
		fmt.Fprintf(eout, "first parameter is the sub string you're counting\n")
		return 1
	}
	target := args[0]
	args = args[1:]
	if inputFName != "" {
		src, err := ioutil.ReadAll(in)
		exitOnError(eout, err, quiet)
		args = append(args, string(src))
	}
	for _, arg := range args {
		fmt.Fprintf(out, "%d%s", strings.Count(arg, target), eol)
	}
	return 0
}

func fnToUpper(in io.Reader, out io.Writer, eout io.Writer, args []string, flagSet *flag.FlagSet) int {
	err := flagSet.Parse(args)
	if err != nil {
		fmt.Fprintf(eout, "%s\n", err)
		return 1
	}
	args = flagSet.Args()
	if inputFName != "" {
		src, err := ioutil.ReadAll(in)
		exitOnError(eout, err, quiet)
		args = append(args, string(src))
	}
	for _, arg := range args {
		fmt.Fprintf(out, "%s%s", strings.ToUpper(arg), eol)
	}
	return 0
}

func fnToLower(in io.Reader, out io.Writer, eout io.Writer, args []string, flagSet *flag.FlagSet) int {
	err := flagSet.Parse(args)
	if err != nil {
		fmt.Fprintf(eout, "%s\n", err)
		return 1
	}
	args = flagSet.Args()
	if inputFName != "" {
		src, err := ioutil.ReadAll(in)
		exitOnError(eout, err, quiet)
		args = append(args, string(src))
	}
	for _, arg := range args {
		fmt.Fprintf(out, "%s%s", strings.ToLower(arg), eol)
	}
	return 0
}

func fnToTitle(in io.Reader, out io.Writer, eout io.Writer, args []string, flagSet *flag.FlagSet) int {
	err := flagSet.Parse(args)
	if err != nil {
		fmt.Fprintf(eout, "%s\n", err)
		return 1
	}
	args = flagSet.Args()
	if inputFName != "" {
		src, err := ioutil.ReadAll(in)
		exitOnError(eout, err, quiet)
		args = append(args, string(src))
	}
	for _, arg := range args {
		fmt.Fprintf(out, "%s%s", strings.ToTitle(arg), eol)
	}
	return 0
}

func fnEnglishTitle(in io.Reader, out io.Writer, eout io.Writer, args []string, flagSet *flag.FlagSet) int {
	err := flagSet.Parse(args)
	if err != nil {
		fmt.Fprintf(eout, "%s\n", err)
		return 1
	}
	args = flagSet.Args()
	if inputFName != "" {
		src, err := ioutil.ReadAll(in)
		exitOnError(eout, err, quiet)
		args = append(args, string(src))
	}
	for _, arg := range args {
		fmt.Fprintf(out, "%s%s", datatools.EnglishTitle(arg), eol)
	}
	return 0
}

func fnSplit(in io.Reader, out io.Writer, eout io.Writer, args []string, flagSet *flag.FlagSet) int {
	err := flagSet.Parse(args)
	if err != nil {
		fmt.Fprintf(eout, "%s\n", err)
		return 1
	}
	args = flagSet.Args()
	if len(args) < 1 {
		fmt.Fprintln(eout, "first parameter is the delimiting string")
		return 1
	}
	delimiter := datatools.NormalizeDelimiter(args[0])
	args = args[1:]
	// Handle the case where out input is piped in or read from a file.
	if inputFName != "" {
		src, err := ioutil.ReadAll(in)
		exitOnError(eout, err, quiet)
		args = append(args, string(src))
	}
	// Now process the args
	for _, arg := range args {
		parts := strings.Split(arg, delimiter)
		// Now assemble our JSON array and display it
		src, err := json.Marshal(parts)
		exitOnError(eout, err, quiet)
		fmt.Fprintf(out, "%s%s", src, eol)
	}
	return 0
}

func fnSplitN(in io.Reader, out io.Writer, eout io.Writer, args []string, flagSet *flag.FlagSet) int {
	err := flagSet.Parse(args)
	if err != nil {
		fmt.Fprintf(eout, "%s\n", err)
		return 1
	}
	args = flagSet.Args()
	if len(args) < 2 {
		fmt.Fprintln(eout, "first parameter is the delimiting string, second is the count")
		return 1
	}
	delimiter := datatools.NormalizeDelimiter(args[0])
	// Now convert to cnt an integer
	cnt, err := strconv.Atoi(args[1])
	if err != nil {
		fmt.Fprintf(eout, "second parameter should be an integer, got %s, errror %s\n", args[1], err)
		return 1
	}
	args = args[2:]
	// Handle the case where out input is piped in or read from a file.
	if inputFName != "" {
		src, err := ioutil.ReadAll(in)
		exitOnError(eout, err, quiet)
		args = append(args, string(src))
	}
	// Handle the case of args being used for input
	for _, arg := range args {
		parts := strings.SplitN(arg, delimiter, cnt)

		// Now assemble our JSON array and display it
		src, err := json.Marshal(parts)
		exitOnError(eout, err, quiet)
		fmt.Fprintf(out, "%s%s", src, eol)
	}
	return 0
}

func fnJoin(in io.Reader, out io.Writer, eout io.Writer, args []string, flagSet *flag.FlagSet) int {
	err := flagSet.Parse(args)
	if err != nil {
		fmt.Fprintf(eout, "%s\n", err)
		return 1
	}
	args = flagSet.Args()
	if len(args) < 1 {
		fmt.Fprintln(eout, "first parameter is the delimiter to join with")
		return 1
	}
	delimiter := datatools.NormalizeDelimiter(args[0])
	args = args[1:]

	// Handle the case where out input is piped in or read from a file.
	if inputFName != "" {
		src, err := ioutil.ReadAll(in)
		exitOnError(eout, err, quiet)
		args = append(args, string(src))
	}
	// Handle the case of args being used for input
	for _, arg := range args {
		parts := []string{}
		// Now we've Unmarshal our object, join it
		err := json.Unmarshal([]byte(arg), &parts)
		exitOnError(eout, err, quiet)
		s := strings.Join(parts, delimiter)
		fmt.Fprintf(out, "%s%s", s, eol)
	}
	return 0
}

func fnHasPrefix(in io.Reader, out io.Writer, eout io.Writer, args []string, flagSet *flag.FlagSet) int {
	err := flagSet.Parse(args)
	if err != nil {
		fmt.Fprintf(eout, "%s\n", err)
		return 1
	}
	args = flagSet.Args()
	// Validate parameters
	if len(args) < 1 {
		fmt.Fprintf(eout, "first parameter is the prefix\n")
		return 1
	}
	prefix := args[0]
	args = args[1:]

	// Handle content coming from a file
	if inputFName != "" {
		src, err := ioutil.ReadAll(in)
		exitOnError(eout, err, quiet)
		args = append(args, string(src))
	}
	// Handle content common from args
	for _, s := range args {
		fmt.Fprintf(out, "%t%s", strings.HasPrefix(s, prefix), eol)
	}
	return 0
}

func fnTrimPrefix(in io.Reader, out io.Writer, eout io.Writer, args []string, flagSet *flag.FlagSet) int {
	err := flagSet.Parse(args)
	if err != nil {
		fmt.Fprintf(eout, "%s\n", err)
		return 1
	}
	args = flagSet.Args()
	// Validate parameters
	if len(args) < 1 {
		fmt.Fprintf(eout, "first parameter is the prefix\n")
		return 1
	}
	prefix := args[0]
	args = args[1:]

	// Handle content coming from input
	if inputFName != "" {
		src, err := ioutil.ReadAll(in)
		exitOnError(eout, err, quiet)
		args = append(args, string(src))
	}
	// Handle content common from args
	for _, arg := range args {
		fmt.Fprintf(out, "%s%s", strings.TrimPrefix(arg, prefix), eol)
	}
	return 0
}

func fnHasSuffix(in io.Reader, out io.Writer, eout io.Writer, args []string, flagSet *flag.FlagSet) int {
	err := flagSet.Parse(args)
	if err != nil {
		fmt.Fprintf(eout, "%s\n", err)
		return 1
	}
	args = flagSet.Args()
	// Validate parameters
	if len(args) < 1 {
		fmt.Fprintf(eout, "first parameter is the suffix\n")
		return 1
	}
	suffix := args[0]
	args = args[1:]

	// Handle content coming from input
	if inputFName != "" {
		src, err := ioutil.ReadAll(in)
		exitOnError(eout, err, quiet)
		args = append(args, string(src))
	}
	// Handle content common from args
	for _, arg := range args {
		fmt.Fprintf(out, "%t%s", strings.HasSuffix(arg, suffix), eol)
	}
	return 0
}

func fnTrimSuffix(in io.Reader, out io.Writer, eout io.Writer, args []string, flagSet *flag.FlagSet) int {
	err := flagSet.Parse(args)
	if err != nil {
		fmt.Fprintf(eout, "%s\n", err)
		return 1
	}
	args = flagSet.Args()
	// Validate parameters
	if len(args) < 1 {
		fmt.Fprintf(eout, "first parameter is the suffix\n")
		return 1
	}
	suffix := args[0]
	args = args[1:]

	// Handle content coming from input
	if inputFName != "" {
		src, err := ioutil.ReadAll(in)
		exitOnError(eout, err, quiet)
		args = append(args, string(src))
	}
	// Handle content common from args
	for _, arg := range args {
		fmt.Fprintf(out, "%s%s", strings.TrimSuffix(arg, suffix), eol)
	}
	return 0
}

func fnTrim(in io.Reader, out io.Writer, eout io.Writer, args []string, flagSet *flag.FlagSet) int {
	err := flagSet.Parse(args)
	if err != nil {
		fmt.Fprintf(eout, "%s\n", err)
		return 1
	}
	args = flagSet.Args()
	if len(args) < 1 {
		fmt.Fprintf(eout, "first parameter is the cutset\n")
		return 1
	}
	cutset := args[0]
	args = args[1:]

	// Handle content coming from input
	if inputFName != "" {
		src, err := ioutil.ReadAll(in)
		exitOnError(eout, err, quiet)
		args = append(args, string(src))
	}
	// Handle content common from args
	for _, arg := range args {
		fmt.Fprintf(out, "%s%s", strings.Trim(arg, cutset), eol)
	}
	return 0
}

func fnTrimLeft(in io.Reader, out io.Writer, eout io.Writer, args []string, flagSet *flag.FlagSet) int {
	err := flagSet.Parse(args)
	if err != nil {
		fmt.Fprintf(eout, "%s\n", err)
		return 1
	}
	args = flagSet.Args()
	if len(args) < 1 {
		fmt.Fprintf(eout, "first parameter is the cutset\n")
		return 1
	}
	cutset := args[0]
	args = args[1:]

	// Handle content coming from file
	if inputFName != "" {
		src, err := ioutil.ReadAll(in)
		exitOnError(eout, err, quiet)
		args = append(args, string(src))
	}
	// Handle content common from args
	for _, arg := range args {
		fmt.Fprintf(out, "%s%s", strings.TrimLeft(arg, cutset), eol)
	}
	return 0
}

func fnTrimRight(in io.Reader, out io.Writer, eout io.Writer, args []string, flagSet *flag.FlagSet) int {
	err := flagSet.Parse(args)
	if err != nil {
		fmt.Fprintf(eout, "%s\n", err)
		return 1
	}
	args = flagSet.Args()
	if len(args) < 1 {
		fmt.Fprintf(eout, "first parameter is the cutset\n")
		return 1
	}
	cutset := args[0]
	args = args[1:]

	// Handle content coming from file
	if inputFName != "" {
		src, err := ioutil.ReadAll(in)
		exitOnError(eout, err, quiet)
		args = append(args, string(src))
	}
	// Handle content common from args
	for _, arg := range args {
		fmt.Fprintf(out, "%s%s", strings.TrimRight(arg, cutset), eol)
	}
	return 0
}

func fnTrimSpace(in io.Reader, out io.Writer, eout io.Writer, args []string, flagSet *flag.FlagSet) int {
	err := flagSet.Parse(args)
	if err != nil {
		fmt.Fprintf(eout, "%s\n", err)
		return 1
	}
	args = flagSet.Args()
	// Handle content coming from file
	if inputFName != "" {
		src, err := ioutil.ReadAll(in)
		exitOnError(eout, err, quiet)
		args = append(args, string(src))
	}
	// Handle content common from args
	for _, arg := range args {
		fmt.Fprintf(out, "%s%s", strings.TrimSpace(arg), eol)
	}
	return 0
}

func fnContains(in io.Reader, out io.Writer, eout io.Writer, args []string, flagSet *flag.FlagSet) int {
	err := flagSet.Parse(args)
	if err != nil {
		fmt.Fprintf(eout, "%s\n", err)
		return 1
	}
	args = flagSet.Args()
	if len(args) < 1 {
		fmt.Fprintf(eout, "first parameter is the target string\n")
		return 1
	}
	target := args[0]
	args = args[1:]

	if inputFName != "" {
		src, err := ioutil.ReadAll(in)
		exitOnError(eout, err, quiet)
		args = append(args, string(src))
	}

	for _, arg := range args {
		fmt.Fprintf(out, "%t%s", strings.Contains(arg, target), eol)
	}
	return 0
}

func fnPosition(in io.Reader, out io.Writer, eout io.Writer, args []string, flagSet *flag.FlagSet) int {
	err := flagSet.Parse(args)
	if err != nil {
		fmt.Fprintf(eout, "%s\n", err)
		return 1
	}
	args = flagSet.Args()
	if len(args) < 1 {
		fmt.Fprintf(eout, "first parameter is the slice you're looking for")
		return 1
	}
	target := args[0]
	args = args[1:]
	if inputFName != "" {
		src, err := ioutil.ReadAll(in)
		exitOnError(eout, err, quiet)
		args = append(args, string(src))
	}
	for _, arg := range args {
		fmt.Fprintf(out, "%d%s", strings.Index(arg, target), eol)
	}
	return 0
}

func fnReplace(in io.Reader, out io.Writer, eout io.Writer, args []string, flagSet *flag.FlagSet) int {
	err := flagSet.Parse(args)
	if err != nil {
		fmt.Fprintf(eout, "%s\n", err)
		return 1
	}
	args = flagSet.Args()
	if len(args) < 2 {
		fmt.Fprintf(eout, "first parameter is the target, second the replacement string\n")
		return 1
	}
	target := args[0]
	replacement := args[1]
	args = args[2:]
	if inputFName != "" {
		src, err := ioutil.ReadAll(in)
		exitOnError(eout, err, quiet)
		args = append(args, string(src))
	}
	for _, arg := range args {
		fmt.Fprintf(out, "%s%s", strings.Replace(arg, target, replacement, -1), eol)
	}
	return 0
}

func fnReplacen(in io.Reader, out io.Writer, eout io.Writer, args []string, flagSet *flag.FlagSet) int {
	err := flagSet.Parse(args)
	if err != nil {
		fmt.Fprintf(eout, "%s\n", err)
		return 1
	}
	args = flagSet.Args()
	if len(args) < 3 {
		fmt.Fprintf(eout, "first parameter is the target, second the replacement string, third is the replacement count (must be positive integer)\n")
		return 1
	}
	target := args[0]
	replacement := args[1]
	cnt, err := strconv.Atoi(args[2])
	exitOnError(eout, err, quiet)
	if cnt < 0 {
		fmt.Fprintf(eout, "third parameter must be a positive integer\n")
		return 1
	}
	args = args[3:]
	if inputFName != "" {
		src, err := ioutil.ReadAll(in)
		exitOnError(eout, err, quiet)
		args = append(args, string(src))
	}
	for _, arg := range args {
		fmt.Fprintf(out, "%s%s", strings.Replace(arg, target, replacement, cnt), eol)
	}
	return 0
	exitOnError(eout, fmt.Errorf("doReplacen not implemented"), false)
	return 1
}

func fnPadLeft(in io.Reader, out io.Writer, eout io.Writer, args []string, flagSet *flag.FlagSet) int {
	err := flagSet.Parse(args)
	if err != nil {
		fmt.Fprintf(eout, "%s\n", err)
		return 1
	}
	args = flagSet.Args()
	if len(args) < 2 {
		fmt.Fprintf(eout, "first parameter is the padding, second the max width of the padded string\n")
		return 1
	}
	pad := args[0]
	maxWidth, err := strconv.Atoi(args[1])
	exitOnError(eout, err, quiet)
	args = args[2:]
	if inputFName != "" {
		src, err := ioutil.ReadAll(in)
		exitOnError(eout, err, quiet)
		args = append(args, string(src))
	}
	// NOTE: we want to the integer modulo of maxWidth and length of pad
	x := maxWidth / len(pad)
	pad = strings.Repeat(pad, x)
	for _, arg := range args {
		l := len(arg)
		if l >= maxWidth {
			fmt.Fprintf(out, "%s%s", arg, eol)
		} else {
			t := fmt.Sprintf("%s%s", pad, arg)
			fmt.Fprintf(out, "%s%s", t[maxWidth-l:], eol)
		}
	}
	return 0
}

func fnPadRight(in io.Reader, out io.Writer, eout io.Writer, args []string, flagSet *flag.FlagSet) int {
	err := flagSet.Parse(args)
	if err != nil {
		fmt.Fprintf(eout, "%s\n", err)
		return 1
	}
	args = flagSet.Args()
	if len(args) < 2 {
		fmt.Fprintf(eout, "first parameter is the padding, second the max width of the padded string\n")
		return 1
	}
	pad := args[0]
	maxWidth, err := strconv.Atoi(args[1])
	exitOnError(eout, err, quiet)
	args = args[2:]
	if inputFName != "" {
		src, err := ioutil.ReadAll(in)
		exitOnError(eout, err, quiet)
		args = append(args, string(src))
	}
	// NOTE: we want to the integer modulo of maxWidth and length of pad
	x := maxWidth / len(pad)
	pad = strings.Repeat(pad, x)
	for _, arg := range args {
		l := len(arg)
		if l >= maxWidth {
			fmt.Fprintf(out, "%s%s", arg, eol)
		} else {
			t := fmt.Sprintf("%s%s", arg, pad)
			fmt.Fprintf(out, "%s%s", t[0:maxWidth], eol)
		}
	}
	return 0
}

func fnSlice(in io.Reader, out io.Writer, eout io.Writer, args []string, flagSet *flag.FlagSet) int {
	err := flagSet.Parse(args)
	if err != nil {
		fmt.Fprintf(eout, "%s\n", err)
		return 1
	}
	args = flagSet.Args()
	if len(args) < 2 {
		fmt.Fprintf(eout, "first parameter is the start position (zero basedindex, inclusive), second is the end position (exclusive) of the substring\n")
		return 1
	}

	start, err := strconv.Atoi(args[0])
	exitOnError(eout, err, quiet)
	end, err := strconv.Atoi(args[1])
	exitOnError(eout, err, quiet)
	if start < 0 && end < 0 {
		fmt.Fprintf(eout, "start and end must be a positive integer\n")
		return 1
	}
	if end <= start {
		fmt.Fprintf(eout, "end is less than or equal to start\n")
		return 1
	}
	args = args[2:]
	if inputFName != "" {
		src, err := ioutil.ReadAll(in)
		exitOnError(eout, err, quiet)
		args = append(args, string(src))
	}

	for _, arg := range args {
		if start > len(arg) {
			fmt.Fprintf(eout, "start %d is past end of string %q\n", start, arg)
			return 1
		}
		fmt.Fprintf(out, "%s%s", arg[start:end], eol)
	}
	return 0
}

func main() {
	// Configuration and creation on or command line interface
	app := cli.NewCli(datatools.Version)

	// Add Help Docs
	app.AddHelp("description", []byte(description))
	app.AddHelp("examples", []byte(examples))

	// Standard Options
	app.BoolVar(&showHelp, "h,help", false, "display help")
	app.BoolVar(&showLicense, "l,license", false, "display license")
	app.BoolVar(&showVersion, "v,version", false, "display version")
	app.BoolVar(&showExamples, "e,examples", false, "display examples")
	app.StringVar(&inputFName, "i,input", "", "input file name")
	app.StringVar(&outputFName, "o,output", "", "output file name")
	app.BoolVar(&quiet, "quiet", false, "suppress error messages")
	app.BoolVar(&newLine, "nl,newline", false, "if true add a trailing newline")
	app.BoolVar(&generateMarkdown, "generate-markdown", false, "generate Markdown documentation")
	app.BoolVar(&generateManPage, "generate-manpage", false, "generate man page")

	// App Options
	app.StringVar(&delimiter, "d,delimiter", "", "set the delimiter")
	app.StringVar(&outputDelimiter, "do,output-delimiter", "", "set the output delimiter")

	// Add verbs and functions
	verb := app.NewVerb("toupper", "to upper case: [STRING]", fnToUpper)
	verb.SetParams("[STRING]")

	verb = app.NewVerb("tolower", "to lower case: [STRING]", fnToLower)
	verb.SetParams("[STRING]")

	verb = app.NewVerb("totitle", "to title case: [STRING]", fnToTitle)
	verb.SetParams("[STRING]")

	verb = app.NewVerb("englishtitle", "English style title case: [STRING]", fnEnglishTitle)
	verb.SetParams("[STRING]")

	verb = app.NewVerb("split", "split into a JSON array: DELIMITER [STRING]", fnSplit)
	verb.SetParams("DELIMITER", "[STRING]")

	verb = app.NewVerb("splitn", "split into an N length JSON array: DELIMITER N [STRING]", fnSplitN)
	verb.SetParams("DELIMITER", "N", "[STRING]")

	verb = app.NewVerb("join", "join JSON array into string: DELIMITER [JSON_ARRAY]", fnJoin)
	verb.SetParams("DELIMITER", "[JSON_ARRAY]")

	verb = app.NewVerb("hasprefix", "true/false on prefix: PREFIX [STRING]", fnHasPrefix)
	verb.SetParams("PREFIX", "[STRING]")

	verb = app.NewVerb("trimprefix", "trims prefix: PREFIX [STRING]", fnTrimPrefix)
	verb.SetParams("PREFIX", "[STRING]")

	verb = app.NewVerb("hassuffix", "true/false on suffix: SUFFIX [STRING]", fnHasSuffix)
	verb.SetParams("SUFFIX", "[STRING]")

	verb = app.NewVerb("trimsuffix", "trim suffix: SUFFIX [STRING]", fnTrimSuffix)
	verb.SetParams("SUFFIX", "[STRING]")

	verb = app.NewVerb("trim", "trim (beginning and end), CUTSET [STRING]", fnTrim)
	verb.SetParams("CURSET", "[STRING]")

	verb = app.NewVerb("trimleft", "left trim: CUTSET [STRING]", fnTrimLeft)
	verb.SetParams("CUTSET", "[STRING]")

	verb = app.NewVerb("trimright", "right trim: CUTSET [STRING]", fnTrimRight)
	verb.SetParams("CUTSET", "[STRING]")

	verb = app.NewVerb("trimspace", "trim leading and trailing spaces: [STRING]", fnTrimSpace)
	verb.SetParams("[STRING]")

	verb = app.NewVerb("count", "count substrings: SUBSTRING [STRING]", fnCount)
	verb.SetParams("SUBSTRING", "[STRING]")

	verb = app.NewVerb("contains", "has substrings: SUBSTRING [STRING]", fnContains)
	verb.SetParams("SUBSTRING", "[STRING]")

	verb = app.NewVerb("length", "length of string: [STRING]", fnLength)
	verb.SetParams("[STRING]")

	verb = app.NewVerb("position", "position of substring: SUBSTRING [STRING]", fnPosition)
	verb.SetParams("SUBSTRING", "[STRING]")

	verb = app.NewVerb("slice", "copy a substring: START END [STRING]", fnSlice)
	verb.SetParams("START", "END", "[STRING]")

	verb = app.NewVerb("replace", "replace: OLD NEW [STRING]", fnReplace)
	verb.SetParams("OLD", "NEW", "[STRING]")

	verb = app.NewVerb("replacen", "replace n times: OLD NEW N [STRING]", fnReplacen)
	verb.SetParams("OLD", "NEW", "N", "[STRING]")

	verb = app.NewVerb("padleft", "left pad: PADDING MAX_LENGTH [STRING]", fnPadLeft)
	verb.SetParams("PADDING", "MAX_LENGTH", "[STRING]")

	verb = app.NewVerb("padright", "right pad: PADDING MAX_LENGTH [STRING]", fnPadRight)
	verb.SetParams("PADDING", "MAX_LENGTH", "[STRING]")

	// We're ready to process args
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

	// Handle options
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
			fmt.Fprintf(app.Out, app.Help(args...))
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

	// Run the app!
	os.Exit(app.Run(args))
}
