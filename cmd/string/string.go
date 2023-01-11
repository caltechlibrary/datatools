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
package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"io"
	"io/ioutil"
	"os"
	"path"
	"strconv"
	"strings"

	// Caltech Library Packages
	"github.com/caltechlibrary/datatools"
)

var (
	helpText = `---
title: "{app_name} (1) user manual"
author: "R. S. Doiel"
pubDate: 2023-01-09
---

# NAME

{app_name}

# SYNOPSIS

{app_name} [OPTIONS] [VERB] [VERB PARAMETERS...]

# DESCRIPTION

{app_name} is a command line tool for transforming strings in common ways.

- string length
- changing cases
- checking for prefixes, suffixes 
- trimming prefixes, suffixes and cutsets (i.e. list of characters to cut)
- position, counting and replacing substrings
- splitting a string into a JSON array of strings, joining JSON a string arrays into a string

VERB refers to the operation to performed on the supplied string(s).
VER PARAMETERS are thsose additional terms need to complete the process
provided by the VERB.

# OPTIONS

Options always come before the VERB.

-help
: display help

-license
:display license

-version
: display version

-d, -delimiter
: set the delimiter

-do, -output-delimiter
: set the output delimiter

-i, -input
: input file name

-nl, -newline
: if true add a trailing newline

-o, -output
: output file name

-quiet
: suppress error messages


## VERBS

contains
: has substrings: SUBSTRING [STRING] ` + "`" + `{app_name} contains SUBSTRING [STRING]` + "`" + `

count
: count substrings: SUBSTRING [STRING] ` + "`" + `{app_name} count SUBSTRING [STRING]` + "`" + `

englishtitle
: English style title case: [STRING] ` + "`" + `{app_name} englishtitle [STRING]` + "`" + `

hasprefix
: true/false on prefix: PREFIX [STRING] ` + "`" + `{app_name} hasprefix PREFIX [STRING]` + "`" + `

hassuffix
: true/false on suffix: SUFFIX [STRING] ` + "`" + `{app_name} hassuffix SUFFIX [STRING]` + "`" + `

join
: join JSON array into string: DELIMITER [JSON_ARRAY] ` + "`" + `{app_name} join DELIMITER [JSON_ARRAY]` + "`" + `

length
: length of string: [STRING] ` + "`" + `{app_name} length [STRING]` + "`" + `

padleft
: left pad PADDING MAX_LENGTH [STRING] ` + "`" + `{app_name} padleft PADDING MAX_LENGTH [STRING]` + "`" + `

padright
: right pad PADDING MAX_LENGTH [STRING] ` + "`" + `{app_name} padright PADDING MAX_LENGTH [STRING]` + "`" + `

position
: position of substring: SUBSTRING [STRING] ` + "`" + `{app_name} position SUBSTRING [STRING]` + "`" + `

replace
: replace: OLD NEW [STRING] ` + "`" + `{app_name} replace OLD NEW [STRING]` + "`" + `

replacen
: replace n times: OLD NEW N [STRING] ` + "`" + `{app_name} replacen OLD NEW N [STRING]` + "`" + `

slice
: copy a substring: START END [STRING] ` + "`" + `{app_name} slice START END [STRING]` + "`" + `

split
: split into a JSON array: DELIMITER [STRING] ` + "`" + `{app_name} split DELIMITER [STRING]` + "`" + `

splitn
: split into an N length JSON array: DELIMITER N [STRING] ` + "`" + `{app_name} splitn DELIMITER N [STRING]` + "`" + `

tolower
: to lower case: [STRING] ` + "`" + `{app_name} tolower [STRING]` + "`" + `

totitle
: to title case: [STRING] ` + "`" + `{app_name} totitle [STRING]` + "`" + `

toupper
: to upper case: [STRING] ` + "`" + `{app_name} toupper [STRING]` + "`" + `

trim
: trim (beginning and end), CUTSET [STRING] ` + "`" + `{app_name} trim CURSET [STRING]` + "`" + `

trimleft
: left trim CUTSET [STRING] ` + "`" + `{app_name} trimleft CUTSET [STRING]` + "`" + `

trimprefix
: trims prefix: PREFIX [STRING] ` + "`" + `{app_name} trimprefix PREFIX [STRING]` + "`" + `

trimright
: right trim: CUTSET [STRING] ` + "`" + `{app_name} trimright CUTSET [STRING]` + "`" + `

trimspace
: trim leading and trailing spaces: [STRING] ` + "`" + `{app_name} trimspace [STRING]` + "`" + `

trimsuffix
: trim suffix: SUFFIX [STRING] ` + "`" + `{app_name} trimsuffix SUFFIX [STRING]` + "`" + `

# EXAMPLES

Convert text to upper case

~~~
	{app_name} toupper "one"
~~~

Convert text to lower case

~~~
	{app_name} tolower "ONE"
~~~

Captialize an English phrase

~~~
	{app_name} englishtitle "one more thing to know"
~~~

Split a space newline delimited list of words into a JSON array

~~~
	{app_name} -i wordlist.txt split "\n"
~~~

Join a JSON array of strings into a newline delimited list

~~~
	{app_name} join '\n' '["one","two","three","four","five"]'
~~~

{app_name} {version}
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
	eol              string

	// App Options
	delimiter       string
	outputDelimiter string
)

func fmtTxt(src string, appName string, version string) string {
	return strings.ReplaceAll(strings.ReplaceAll(src, "{app_name}", appName), "{version}", version)
}

// Application functionality
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

func fnLength(in io.Reader, out io.Writer, eout io.Writer, args []string) int {
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

func fnCount(in io.Reader, out io.Writer, eout io.Writer, args []string) int {
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

func fnToUpper(in io.Reader, out io.Writer, eout io.Writer, args []string) int {
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

func fnToLower(in io.Reader, out io.Writer, eout io.Writer, args []string) int {
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

func fnToTitle(in io.Reader, out io.Writer, eout io.Writer, args []string) int {
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

func fnEnglishTitle(in io.Reader, out io.Writer, eout io.Writer, args []string) int {
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

func fnSplit(in io.Reader, out io.Writer, eout io.Writer, args []string) int {
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

func fnSplitN(in io.Reader, out io.Writer, eout io.Writer, args []string) int {
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

func fnJoin(in io.Reader, out io.Writer, eout io.Writer, args []string) int {
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

func fnHasPrefix(in io.Reader, out io.Writer, eout io.Writer, args []string) int {
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

func fnTrimPrefix(in io.Reader, out io.Writer, eout io.Writer, args []string) int {
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

func fnHasSuffix(in io.Reader, out io.Writer, eout io.Writer, args []string) int {
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

func fnTrimSuffix(in io.Reader, out io.Writer, eout io.Writer, args []string) int {
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

func fnTrim(in io.Reader, out io.Writer, eout io.Writer, args []string) int {
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

func fnTrimLeft(in io.Reader, out io.Writer, eout io.Writer, args []string) int {
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

func fnTrimRight(in io.Reader, out io.Writer, eout io.Writer, args []string) int {
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

func fnTrimSpace(in io.Reader, out io.Writer, eout io.Writer, args []string) int {
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

func fnContains(in io.Reader, out io.Writer, eout io.Writer, args []string) int {
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

func fnPosition(in io.Reader, out io.Writer, eout io.Writer, args []string) int {
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

func fnReplace(in io.Reader, out io.Writer, eout io.Writer, args []string) int {
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

func fnReplacen(in io.Reader, out io.Writer, eout io.Writer, args []string) int {
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

func fnPadLeft(in io.Reader, out io.Writer, eout io.Writer, args []string) int {
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

func fnPadRight(in io.Reader, out io.Writer, eout io.Writer, args []string) int {
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

func fnSlice(in io.Reader, out io.Writer, eout io.Writer, args []string) int {
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

func runApp(in io.Reader, out io.Writer, eout io.Writer, args []string, appName string, version string) int {
	if len(args) == 0 {
		fmt.Fprintf(eout, "no action provided, see %s -help", appName)
		return 1
	}

	verb := args[0]
	// Add verbs and functions
	ops := map[string]func(io.Reader, io.Writer, io.Writer, []string) int{
		"toupper":      fnToUpper,
		"tolower":      fnToLower,
		"totitle":      fnToTitle,
		"englishtitle": fnEnglishTitle,
		"split":        fnSplit,
		"splitn":       fnSplitN,
		"join":         fnJoin,
		"hasprefix":    fnHasPrefix,
		"trimprefix":   fnTrimPrefix,
		"hassuffix":    fnHasSuffix,
		"trimsuffix":   fnTrimSuffix,
		"trim":         fnTrim,
		"trimleft":     fnTrimLeft,
		"trimright":    fnTrimRight,
		"trimspace":    fnTrimSpace,
		"count":        fnCount,
		"contains":     fnContains,
		"length":       fnLength,
		"position":     fnPosition,
		"slice":        fnSlice,
		"replace":      fnReplace,
		"replacen":     fnReplacen,
		"padleft":      fnPadLeft,
		"padright":     fnPadRight,
	}
	if fn, ok := ops[verb]; ok {
		return fn(in, out, eout, args[1:])
	}
	fmt.Fprintf(eout, "%q is not supported", verb)
	return 1
}

func main() {
	appName := path.Base(os.Args[0])

	// Standard Options
	flag.BoolVar(&showHelp, "help", false, "display help")
	flag.BoolVar(&showLicense, "license", false, "display license")
	flag.BoolVar(&showVersion, "version", false, "display version")

	flag.StringVar(&inputFName, "i", "", "input file name")
	flag.StringVar(&inputFName, "input", "", "input file name")
	flag.StringVar(&outputFName, "o", "", "output file name")
	flag.StringVar(&outputFName, "output", "", "output file name")
	flag.BoolVar(&quiet, "quiet", false, "suppress error messages")
	flag.BoolVar(&newLine, "nl", false, "if true add a trailing newline")
	flag.BoolVar(&newLine, "newline", false, "if true add a trailing newline")

	// App Options
	flag.StringVar(&delimiter, "d", "", "set the delimiter")
	flag.StringVar(&delimiter, "delimiter", "", "set the delimiter")
	flag.StringVar(&outputDelimiter, "do", "", "set the output delimiter")
	flag.StringVar(&outputDelimiter, "output-delimiter", "", "set the output delimiter")

	// We're ready to process args
	flag.Parse()
	args := flag.Args()

	// Setup IO
	var err error

	in := os.Stdin
	out := os.Stdout
	eout := os.Stderr

	if inputFName != ""  && inputFName != "-" {
		in, err = os.Open(inputFName)
		if err != nil {
			fmt.Fprintln(eout, err)
			os.Exit(1)
		}
		defer in.Close()
	}

	if outputFName != "" && outputFName != "-" {
		out, err = os.Create(outputFName)
		if err != nil {
			fmt.Fprintln(eout, err)
			os.Exit(1)
		}
		defer out.Close()
	}

	// Handle options
	if showHelp {
		fmt.Fprintf(out, "%s\n", fmtTxt(helpText, appName, datatools.Version))
		os.Exit(0)
	}
	if showLicense {
		fmt.Fprintf(out, "%s\n", datatools.LicenseText)
		os.Exit(0)
	}
	if showVersion {
		fmt.Fprintf(out, "%s %s\n", appName, datatools.Version)
		os.Exit(0)
	}
	if newLine {
		eol = "\n"
	}

	// Run the app!
	os.Exit(runApp(in, out, eout, args, appName, datatools.Version))
}
