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
	flag.StringVar(&delimiter, "d", " ", "set the delimiting string value")
	flag.StringVar(&delimiter, "delimiter", " ", "set the delimiting string value")
	flag.BoolVar(&newLine, "nl", false, "output as one substring per line rather than JSON")
	flag.BoolVar(&newLine, "newline", false, "output as one substring per line rather than JSON")
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

	if inputFName != "" {
		src, err := ioutil.ReadAll(in)
		if err != nil {
			fmt.Fprintf(os.Stderr, "Can't read file, %s", err)
			os.Exit(1)
		}
		args = strings.Split(string(src), "\n")
	}

	results := []string{}
	for _, s := range args {
		parts := strings.Split(s, delimiter)
		for _, p := range parts {
			results = append(results, p)
		}
	}

	if newLine == true {
		fmt.Fprintf(out, "%s\n", strings.Join(results, "\n"))
	} else {
		src, err := json.Marshal(results)
		if err != nil {
			fmt.Fprintf(os.Stderr, "Can't marshal %+v", err)
			os.Exit(1)
		}
		fmt.Fprintf(out, "%s\n", src)
	}
}
