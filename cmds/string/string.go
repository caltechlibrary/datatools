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
		fmt.Fprintf(out, "%T%s", strings.HasPrefix(s, prefix), nl)
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
		fmt.Fprintf(out, "%s%T", strings.TrimRight(arg, cutset), nl)
	}
	return 0
}

func doContains(in io.Reader, out io.Writer, eout io.Writer, args []string) int {
	exitOnError(eout, fmt.Errorf("doContains not implemented"), false)
	return 1
}

func doPosition(in io.Reader, out io.Writer, eout io.Writer, args []string) int {
	exitOnError(eout, fmt.Errorf("doPosition not implemented"), false)
	return 1
}

func doReplace(in io.Reader, out io.Writer, eout io.Writer, args []string) int {
	exitOnError(eout, fmt.Errorf("doReplace not implemented"), false)
	return 1
}

func doReplacen(in io.Reader, out io.Writer, eout io.Writer, args []string) int {
	exitOnError(eout, fmt.Errorf("doReplacen not implemented"), false)
	return 1
}

func doPad(in io.Reader, out io.Writer, eout io.Writer, args []string) int {
	exitOnError(eout, fmt.Errorf("doPad not implemented"), false)
	return 1
}

func doPadLeft(in io.Reader, out io.Writer, eout io.Writer, args []string) int {
	exitOnError(eout, fmt.Errorf("doPadLeft not implemented"), false)
	return 1
}

func doPadRight(in io.Reader, out io.Writer, eout io.Writer, args []string) int {
	exitOnError(eout, fmt.Errorf("doPadRight not implemented"), false)
	return 1
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
	app.AddAction("toupper", doToUpper, "to upper case: [STRINGS]")
	app.AddAction("tolower", doToLower, "to lower case: [STRINGS]")
	app.AddAction("totitle", doToTitle, "to title case: [STRINGS]")
	app.AddAction("englishtitle", doEnglishTitle, "English style title case: [STRINGS]")
	app.AddAction("split", doSplit, "split into a JSON array: DELIMITER [STRINGS]")
	app.AddAction("splitn", doSplitN, "split into an N length JSON array: DELIMITER N [STRINGS]")
	app.AddAction("join", doJoin, "join JSON array into string: DELIMITER [STRINS]")
	app.AddAction("hasprefix", doHasPrefix, "true/false on prefix: PREFIX [STRINGS]")
	app.AddAction("trimprefix", doTrimPrefix, "trims prefix: PREFIX [STRINGS]")
	app.AddAction("hassuffix", doHasSuffix, "true/false on suffix: SUFFIX [STRINGS]")
	app.AddAction("trimsuffix", doTrimSuffix, "trim suffix: SUFFIX [STRINGS]")
	app.AddAction("trim", doTrim, "trim (beginning and end), CUTSET [STRINGS]")
	app.AddAction("trimleft", doTrimLeft, "left trim: CUTSET [STRINGS]")
	app.AddAction("trimright", doTrimRight, "right trim: CUTSET [STRINGS]")
	app.AddAction("count", doCount, "count substrings: SUBSTRING [STRINGS]")
	app.AddAction("contains", doContains, "has substrings: SUBSTRING [STRINGS]")
	app.AddAction("length", doLength, "length of string: [STRINGS]")
	app.AddAction("position", doPosition, "position of substring: SUBSTRING [STRINGS]")
	app.AddAction("replace", doReplace, "replace: TARGET REPLACEMENT [STRINGS]")
	app.AddAction("replacen", doReplace, "replace n times: TARGET REPLACEMENT COUNT [STRINGS]")
	app.AddAction("pad", doPad, "pad (beginning and end): PADDING MAX_LENGTH [STRINGS]")
	app.AddAction("padleft", doPadLeft, "left pad: PADDING MAX_LENGTH [STRINGS]")
	app.AddAction("padleft", doPadRight, "left pad: PADDING MAX_LENGTH [STRINGS]")

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
