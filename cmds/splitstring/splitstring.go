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
	usage = `USAGE: %s [OPTIONS] [STRING_TO_SPLIT]`

	description = `

%s splits a string based on a delimiting string provided. The default
delimiter is a space. You can specify a delimiting string via
the -d or --delimiter option.  %s will split the string provided
as a command line argument but can read split string(s) recieved on
stdin in with the -i or --input option. By default the split
strings are render as a JSON array but with the option -nl or
--newline you can render each split string one per line.

`

	examples = `

EXAMPLES

Splitting a string that is double pipe delimited rendering
one sub string per line.

    %s -nl -d '||' "one||two||three"

This should yield

    one
	two
	three

Splitting a string that is double pipe delimited rendering JSON

    %s -d '||' "one||two||three"

This should yield

   ["one","two","three"]

`

	// Standard Options
	showHelp     bool
	showLicense  bool
	showVersion  bool
	showExamples bool
	inputFName   string
	outputFName  string
	quiet        bool
	newLine      bool

	// App Options
	delimiter string
	plainText bool
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
	flag.BoolVar(&newLine, "no-newline", false, "exclude trailing newline in output")
	flag.BoolVar(&newLine, "nl", true, "include trailing newline in output")
	flag.BoolVar(&newLine, "newline", true, "include trailing newline in output")

	// Application specific options
	flag.StringVar(&delimiter, "d", " ", "set the delimiting string value")
	flag.StringVar(&delimiter, "delimiter", " ", "set the delimiting string value")
	flag.BoolVar(&plainText, "text", true, "output as plain text, one value per line rather than JSON")
	flag.BoolVar(&plainText, "plain-text", true, "output as plain text, one value per line rather than JSON")
}

func main() {
	appName := path.Base(os.Args[0])
	flag.Parse()
	args := flag.Args()

	cfg := cli.New(appName, "", datatools.Version)
	cfg.UsageText = fmt.Sprintf(usage, appName)
	cfg.DescriptionText = fmt.Sprintf(description, appName, appName)
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
	cli.ExitOnError(os.Stderr, err, quiet)
	defer cli.CloseFile(inputFName, in)

	out, err := cli.Create(outputFName, os.Stdout)
	cli.ExitOnError(os.Stderr, err, quiet)
	defer cli.CloseFile(outputFName, out)

	// Normalize the delimiter if \n or \t
	switch delimiter {
	case `\n`:
		delimiter = "\n"
	case `\t`:
		delimiter = "\t"
	}

	if inputFName != "" {
		src, err := ioutil.ReadAll(in)
		cli.ExitOnError(os.Stderr, err, quiet)
		args = strings.Split(string(src), "\n")
	}

	results := []string{}
	for _, s := range args {
		parts := strings.Split(s, delimiter)
		for _, p := range parts {
			results = append(results, p)
		}
	}

	nl := "\n"
	if newLine == false {
		nl = ""
	}

	if plainText == true {
		fmt.Fprintf(out, "%s%s", strings.Join(results, "\n"), nl)
	} else {
		src, err := json.Marshal(results)
		cli.ExitOnError(os.Stderr, err, quiet)
		fmt.Fprintf(out, "%s%s", src, nl)
	}
}
