package main

import (
	"bufio"
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

+ changing cases
+ checking for prefixes, suffixes or substrings
+ trimming prefixes, suffixes or specific characters (i.e. cutsets)
+ locating, counting and replacing substrings
+ splitting a string into a JSON array of strings, joining JSON a string arrays into a string
+ formatting and padding strings and numbers
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
	nl           string
	englishTitle bool
	plainText    bool
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

func doToUpper(in io.Reader, out io.Writer, eout io.Writer, args []string) int {
	exitCode := 0
	if inputFName != "" {
		scanner := bufio.NewScanner(in)
		for scanner.Scan() {
			s := scanner.Text()
			fmt.Fprintln(out, strings.ToUpper(s))
		}
		if err := scanner.Err(); err != nil {
			onError(eout, err, quiet)
			exitCode = 1
		}
	}
	if len(args) > 0 {
		fmt.Fprintf(out, "%s%s", strings.ToUpper(strings.Join(args, " ")), nl)
	}
	return exitCode
}

func doToLower(in io.Reader, out io.Writer, eout io.Writer, args []string) int {
	exitCode := 0
	if inputFName != "" {
		scanner := bufio.NewScanner(in)
		for scanner.Scan() {
			s := scanner.Text()
			fmt.Fprintln(out, strings.ToLower(s))
		}
		if err := scanner.Err(); err != nil {
			onError(eout, err, quiet)
			exitCode = 1
		}
	}
	if len(args) > 0 {
		fmt.Fprintf(out, "%s%s", strings.ToLower(strings.Join(args, " ")), nl)
	}
	return exitCode
}

func doToTitle(in io.Reader, out io.Writer, eout io.Writer, args []string) int {
	// Handle file input
	if inputFName != "" {
		src, err := ioutil.ReadAll(in)
		exitOnError(eout, err, quiet)
		args = append(args, string(src))
	}
	// Now title our args
	for _, arg := range args {
		fmt.Fprintf(out, "%s%s", strings.ToTitle(arg), nl)
	}
	return 0
}

func doEnglishTitle(in io.Reader, out io.Writer, eout io.Writer, args []string) int {
	// Handle file input
	if inputFName != "" {
		src, err := ioutil.ReadAll(in)
		exitOnError(eout, err, quiet)
		args = append(args, string(src))
	}
	// Now title our args
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
		if err != nil {
			fmt.Fprintf(eout, "can't read %s, %s\n", inputFName, err)
			return 1
		}
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
		if err != nil {
			fmt.Fprintf(eout, "can't read %s, %s\n", inputFName, err)
			return 1
		}
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
		if err != nil {
			fmt.Fprintf(eout, "can't read %s, %s\n", inputFName, err)
			return 1
		}
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
		fmt.Fprintf(eout, "first parameter is the prefix%s", nl)
		return 1
	}
	prefix := args[0]
	args = args[1:]

	// Handle content coming from input
	exitCode := 0
	if inputFName != "" {
		scanner := bufio.NewScanner(in)
		for scanner.Scan() {
			s := scanner.Text()
			hasPrefix := strings.HasPrefix(s, prefix)
			fmt.Fprintln(out, "%T", hasPrefix)
			if hasPrefix == false {
				exitCode = 1
			}
		}
	}
	// Handle content common from args
	sep := ""
	for _, s := range args {
		hasPrefix := strings.HasPrefix(s, prefix)
		fmt.Fprintf(out, "%s%T", sep, hasPrefix)
		sep = " "
		if hasPrefix == false {
			exitCode = 1
		}
	}
	fmt.Fprintf(out, "%s", nl)
	return exitCode
}

func doTrimPrefix(in io.Reader, out io.Writer, eout io.Writer, args []string) int {
	// Validate parameters
	if len(args) < 1 {
		fmt.Fprintf(eout, "first parameter is the prefix%s", nl)
		return 1
	}
	prefix := args[0]
	args = args[1:]

	// Handle content coming from input
	exitCode := 0
	if inputFName != "" {
		scanner := bufio.NewScanner(in)
		for scanner.Scan() {
			s := scanner.Text()
			fmt.Fprintln(out, "%s", strings.TrimPrefix(s, prefix))
		}
	}
	// Handle content common from args
	sep := ""
	for _, s := range args {
		fmt.Fprintf(out, "%s%T", sep, strings.TrimPrefix(s, prefix))
		sep = " "
	}
	fmt.Fprintf(out, "%s", nl)
	return exitCode
}

func doHasSuffix(in io.Reader, out io.Writer, eout io.Writer, args []string) int {
	// Validate parameters
	if len(args) < 1 {
		fmt.Fprintf(eout, "first parameter is the suffix%s", nl)
		return 1
	}
	suffix := args[0]
	args = args[1:]

	// Handle content coming from input
	exitCode := 0
	if inputFName != "" {
		scanner := bufio.NewScanner(in)
		for scanner.Scan() {
			s := scanner.Text()
			hasSuffix := strings.HasSuffix(s, suffix)
			fmt.Fprintln(out, "%T", hasSuffix)
			if hasSuffix == false {
				exitCode = 1
			}
		}
	}
	// Handle content common from args
	sep := ""
	for _, s := range args {
		hasSuffix := strings.HasSuffix(s, suffix)
		fmt.Fprintf(out, "%s%T", sep, hasSuffix)
		sep = " "
		if hasSuffix == false {
			exitCode = 1
		}
	}
	fmt.Fprintf(out, "%s", nl)
	return exitCode
}

func doTrimSuffix(in io.Reader, out io.Writer, eout io.Writer, args []string) int {
	// Validate parameters
	if len(args) < 1 {
		fmt.Fprintf(eout, "first parameter is the suffix%s", nl)
		return 1
	}
	suffix := args[0]
	args = args[1:]

	// Handle content coming from input
	exitCode := 0
	if inputFName != "" {
		scanner := bufio.NewScanner(in)
		for scanner.Scan() {
			s := scanner.Text()
			fmt.Fprintln(out, "%s", strings.TrimSuffix(s, suffix))
		}
	}
	// Handle content common from args
	sep := ""
	for _, s := range args {
		fmt.Fprintf(out, "%s%T", sep, strings.TrimSuffix(s, suffix))
		sep = " "
	}
	fmt.Fprintf(out, "%s", nl)
	return exitCode
}

func doTrim(in io.Reader, out io.Writer, eout io.Writer, args []string) int {
	if len(args) < 1 {
		fmt.Fprintf(eout, "first parameter is the cutset%s", nl)
		return 1
	}
	cutset := args[0]
	args = args[1:]

	// Handle content coming from input
	exitCode := 0
	if inputFName != "" {
		scanner := bufio.NewScanner(in)
		for scanner.Scan() {
			s := scanner.Text()
			fmt.Fprintln(out, "%s", strings.Trim(s, cutset))
		}
	}
	// Handle content common from args
	sep := ""
	for _, s := range args {
		fmt.Fprintf(out, "%s%T", sep, strings.Trim(s, cutset))
		sep = " "
	}
	fmt.Fprintf(out, "%s", nl)
	return exitCode
}

func doTrimLeft(in io.Reader, out io.Writer, eout io.Writer, args []string) int {
	if len(args) < 1 {
		fmt.Fprintf(eout, "first parameter is the cutset%s", nl)
		return 1
	}
	cutset := args[0]
	args = args[1:]

	// Handle content coming from input
	exitCode := 0
	if inputFName != "" {
		scanner := bufio.NewScanner(in)
		for scanner.Scan() {
			s := scanner.Text()
			fmt.Fprintln(out, "%s", strings.TrimLeft(s, cutset))
		}
	}
	// Handle content common from args
	sep := ""
	for _, s := range args {
		fmt.Fprintf(out, "%s%T", sep, strings.TrimLeft(s, cutset))
		sep = " "
	}
	fmt.Fprintf(out, "%s", nl)
	return exitCode
}

func doTrimRight(in io.Reader, out io.Writer, eout io.Writer, args []string) int {
	if len(args) < 1 {
		fmt.Fprintf(eout, "first parameter is the cutset%s", nl)
		return 1
	}
	cutset := args[0]
	args = args[1:]

	// Handle content coming from input
	exitCode := 0
	if inputFName != "" {
		scanner := bufio.NewScanner(in)
		for scanner.Scan() {
			s := scanner.Text()
			fmt.Fprintln(out, "%s", strings.TrimRight(s, cutset))
		}
	}
	// Handle content common from args
	sep := ""
	for _, s := range args {
		fmt.Fprintf(out, "%s%T", sep, strings.TrimRight(s, cutset))
		sep = " "
	}
	fmt.Fprintf(out, "%s", nl)
	return exitCode
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
	app.BoolVar(&plainText, "t,text", false, "handle arrays as plain text")

	// Add verbs and functions
	app.AddAction("toupper", doToUpper, "converts a string(s) to upper case")
	app.AddAction("tolower", doToLower, "converts a string(s) to lower case")
	app.AddAction("totitle", doToTitle, "converts a string(s) to title case")
	app.AddAction("englishtitle", doEnglishTitle, "converts a string(s) to English style title case")
	app.AddAction("split", doSplit, "splits a string into a JSON array on a delimiter, first parameter is the delimiter")
	app.AddAction("splitn", doSplitN, "splits a string into an N length JSON array on delimiter, first parameter is the delimiter, second N")
	app.AddAction("join", doJoin, "join JSON array(s) of strings or join delimited input, first parameter is delimiter")
	app.AddAction("hasprefix", doHasPrefix, "output true if string(s) have prefix otherwise false, first parameter is prefix")
	app.AddAction("trimprefix", doTrimPrefix, "trims the prefix from a string(s), first parameter is prefix")
	app.AddAction("hassuffix", doHasSuffix, "output true if string(s) have suffix otherwise false, first parameter is suffix")
	app.AddAction("trimsuffix", doTrimSuffix, "trims the suffix from a string(s), first parameter is suffix")
	app.AddAction("trim", doTrim, "trims the cutset from beginning and end of string(s), first parameter is cutset")
	app.AddAction("trimleft", doTrimLeft, "left trim the cutset from a string(s), first parameter is cutset")
	app.AddAction("trimright", doTrimRight, "right trim the cutset from a string(s), first parameter is cutset")

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
