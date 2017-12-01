//
// string is a command line utility to expose some of the Golang strings functions to the command line.
//
// @author R. S. Doiel, <rsdoiel@library.caltech.edu>
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
`

	// Standard Options
	showHelp             bool
	showLicense          bool
	showVersion          bool
	showExamples         bool
	inputFName           string
	outputFName          string
	newLine              bool
	quiet                bool
	generateMarkdownDocs bool

	// App Options

	// Internal globals
	nl string
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

func doLength(in io.Reader, out io.Writer, eout io.Writer, args []string) int {
	if inputFName != "" {
		src, err := ioutil.ReadAll(in)
		exitOnError(eout, err, quiet)
		args = append(args, string(src))
	}
	for _, arg := range args {
		fmt.Fprintf(out, "%d%s", len(arg), nl)
	}
	return 0
}

func doCount(in io.Reader, out io.Writer, eout io.Writer, args []string) int {
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
		fmt.Fprintf(out, "%d%s", strings.Count(arg, target), nl)
	}
	return 0
}

func doToUpper(in io.Reader, out io.Writer, eout io.Writer, args []string) int {
	if inputFName != "" {
		src, err := ioutil.ReadAll(in)
		exitOnError(eout, err, quiet)
		args = append(args, string(src))
	}
	for _, arg := range args {
		fmt.Fprintf(out, "%s%s", strings.ToUpper(arg), nl)
	}
	return 0
}

func doToLower(in io.Reader, out io.Writer, eout io.Writer, args []string) int {
	if inputFName != "" {
		src, err := ioutil.ReadAll(in)
		exitOnError(eout, err, quiet)
		args = append(args, string(src))
	}
	for _, arg := range args {
		fmt.Fprintf(out, "%s%s", strings.ToLower(arg), nl)
	}
	return 0
}

func doToTitle(in io.Reader, out io.Writer, eout io.Writer, args []string) int {
	if inputFName != "" {
		src, err := ioutil.ReadAll(in)
		exitOnError(eout, err, quiet)
		args = append(args, string(src))
	}
	for _, arg := range args {
		fmt.Fprintf(out, "%s%s", strings.ToTitle(arg), nl)
	}
	return 0
}

func doEnglishTitle(in io.Reader, out io.Writer, eout io.Writer, args []string) int {
	if inputFName != "" {
		src, err := ioutil.ReadAll(in)
		exitOnError(eout, err, quiet)
		args = append(args, string(src))
	}
	for _, arg := range args {
		fmt.Fprintf(out, "%s%s", datatools.EnglishTitle(arg), nl)
	}
	return 0
}

func doSplit(in io.Reader, out io.Writer, eout io.Writer, args []string) int {
	if len(args) < 1 {
		fmt.Fprintln(eout, "first parameter is the delimiting string")
		return 1
	}
	delimiter := args[0]
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
		fmt.Fprintf(out, "%s%s", src, nl)
	}
	return 0
}

func doSplitN(in io.Reader, out io.Writer, eout io.Writer, args []string) int {
	if len(args) < 2 {
		fmt.Fprintln(eout, "first parameter is the delimiting string, second is the count")
		return 1
	}
	delimiter := args[0]
	countS := args[1]
	args = args[2:]
	// Now convert countS to cnt
	cnt, err := strconv.Atoi(countS)
	if err != nil {
		fmt.Fprintf(eout, "second parameter should be an integer, got %s, errror %s\n", countS, err)
		return 1
	}
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
		fmt.Fprintf(out, "%s%s", src, nl)
	}
	return 0
}

func doJoin(in io.Reader, out io.Writer, eout io.Writer, args []string) int {
	if len(args) < 1 {
		fmt.Fprintln(eout, "first parameter is the delimiter to join with")
		return 1
	}
	delimiter := args[0]
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
		fmt.Fprintf(out, "%s%s", s, nl)
	}
	return 0
}

func doHasPrefix(in io.Reader, out io.Writer, eout io.Writer, args []string) int {
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
		fmt.Fprintf(out, "%t%s", strings.HasPrefix(s, prefix), nl)
	}
	return 0
}

func doTrimPrefix(in io.Reader, out io.Writer, eout io.Writer, args []string) int {
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
		fmt.Fprintf(out, "%s%s", strings.TrimPrefix(arg, prefix), nl)
	}
	return 0
}

func doHasSuffix(in io.Reader, out io.Writer, eout io.Writer, args []string) int {
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
		fmt.Fprintf(out, "%t%s", strings.HasSuffix(arg, suffix), nl)
	}
	return 0
}

func doTrimSuffix(in io.Reader, out io.Writer, eout io.Writer, args []string) int {
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
		fmt.Fprintf(out, "%s%s", strings.TrimSuffix(arg, suffix), nl)
	}
	return 0
}

func doTrim(in io.Reader, out io.Writer, eout io.Writer, args []string) int {
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
		fmt.Fprintf(out, "%s%s", strings.Trim(arg, cutset), nl)
	}
	return 0
}

func doTrimLeft(in io.Reader, out io.Writer, eout io.Writer, args []string) int {
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
		fmt.Fprintf(out, "%s%s", strings.TrimLeft(arg, cutset), nl)
	}
	return 0
}

func doTrimRight(in io.Reader, out io.Writer, eout io.Writer, args []string) int {
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
		fmt.Fprintf(out, "%s%s", strings.TrimRight(arg, cutset), nl)
	}
	return 0
}

func doContains(in io.Reader, out io.Writer, eout io.Writer, args []string) int {
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
		fmt.Fprintf(out, "%t%s", strings.Contains(arg, target), nl)
	}
	return 0
}

func doPosition(in io.Reader, out io.Writer, eout io.Writer, args []string) int {
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
		fmt.Fprintf(out, "%d%s", strings.Index(arg, target), nl)
	}
	return 0
}

func doReplace(in io.Reader, out io.Writer, eout io.Writer, args []string) int {
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
		fmt.Fprintf(out, "%s%s", strings.Replace(arg, target, replacement, -1), nl)
	}
	return 0
}

func doReplacen(in io.Reader, out io.Writer, eout io.Writer, args []string) int {
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
		fmt.Fprintf(out, "%s%s", strings.Replace(arg, target, replacement, cnt), nl)
	}
	return 0
	exitOnError(eout, fmt.Errorf("doReplacen not implemented"), false)
	return 1
}

func doPadLeft(in io.Reader, out io.Writer, eout io.Writer, args []string) int {
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
			fmt.Fprintf(out, "%s%s", arg, nl)
		} else {
			t := fmt.Sprintf("%s%s", pad, arg)
			fmt.Fprintf(out, "%s%s", t[maxWidth-l:], nl)
		}
	}
	return 0
}

func doPadRight(in io.Reader, out io.Writer, eout io.Writer, args []string) int {
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
			fmt.Fprintf(out, "%s%s", arg, nl)
		} else {
			t := fmt.Sprintf("%s%s", arg, pad)
			fmt.Fprintf(out, "%s%s", t[0:maxWidth], nl)
		}
	}
	return 0
}

func doSlice(in io.Reader, out io.Writer, eout io.Writer, args []string) int {
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
		fmt.Fprintf(out, "%s%s", arg[start:end], nl)
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
	app.BoolVar(&newLine, "nl,newline", false, "output a trailing newline")
	app.BoolVar(&generateMarkdownDocs, "generate-markdown-docs", false, "output documentation in Markdown")

	// App Options

	// Add verbs and functions
	app.AddAction("toupper", doToUpper, "to upper case: [STRING]")
	app.AddAction("tolower", doToLower, "to lower case: [STRING]")
	app.AddAction("totitle", doToTitle, "to title case: [STRING]")
	app.AddAction("englishtitle", doEnglishTitle, "English style title case: [STRING]")
	app.AddAction("split", doSplit, "split into a JSON array: DELIMITER [STRING]")
	app.AddAction("splitn", doSplitN, "split into an N length JSON array: DELIMITER N [STRING]")
	app.AddAction("join", doJoin, "join JSON array into string: DELIMITER [STRING]")
	app.AddAction("hasprefix", doHasPrefix, "true/false on prefix: PREFIX [STRING]")
	app.AddAction("trimprefix", doTrimPrefix, "trims prefix: PREFIX [STRING]")
	app.AddAction("hassuffix", doHasSuffix, "true/false on suffix: SUFFIX [STRING]")
	app.AddAction("trimsuffix", doTrimSuffix, "trim suffix: SUFFIX [STRING]")
	app.AddAction("trim", doTrim, "trim (beginning and end), CUTSET [STRING]")
	app.AddAction("trimleft", doTrimLeft, "left trim: CUTSET [STRING]")
	app.AddAction("trimright", doTrimRight, "right trim: CUTSET [STRING]")
	app.AddAction("count", doCount, "count substrings: SUBSTRING [STRING]")
	app.AddAction("contains", doContains, "has substrings: SUBSTRING [STRING]")
	app.AddAction("length", doLength, "length of string: [STRING]")
	app.AddAction("position", doPosition, "position of substring: SUBSTRING [STRING]")
	app.AddAction("slice", doSlice, "copy a substring: START END [STRING]")
	app.AddAction("replace", doReplace, "replace: OLD NEW [STRING]")
	app.AddAction("replacen", doReplacen, "replace n times: OLD NEW N [STRING]")
	app.AddAction("padleft", doPadLeft, "left pad: PADDING MAX_LENGTH [STRING]")
	app.AddAction("padright", doPadRight, "right pad: PADDING MAX_LENGTH [STRING]")

	// We're ready to process args
	app.Parse()
	args := app.Args()

	// Setup IO
	var err error

	app.Eout = os.Stderr
	app.In, err = cli.Open(inputFName, os.Stdin)
	cli.ExitOnError(app.Eout, err, quiet)
	defer cli.CloseFile(inputFName, app.In)
	app.Out, err = cli.Create(inputFName, os.Stdout)
	cli.ExitOnError(app.Eout, err, quiet)
	defer cli.CloseFile(outputFName, app.Out)

	// Handle options
	if generateMarkdownDocs {
		app.GenerateMarkdownDocs(app.Out)
		os.Exit(0)
	}
	if showHelp || showExamples {
		if len(args) > 0 {
			fmt.Fprintf(app.Out, app.Help(args...))
		} else if showExamples {
			fmt.Fprintf(app.Out, app.Help("examples"))
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
		nl = "\n"
	}

	// Run the app!
	os.Exit(app.Run(args))
}
