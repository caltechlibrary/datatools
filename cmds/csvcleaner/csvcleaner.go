package main

import (
	"encoding/csv"
	"flag"
	"fmt"
	"io"
	"os"
	"path"
	"strings"

	// Caltech Library Packages
	"github.com/caltechlibrary/cli"
	"github.com/caltechlibrary/datatools"
)

var (
	usage = `USAGE: %s [OPTIONS]`

	description = `
%s normalizes a CSV file based on the options selected. It
helps to address issues like variable number of columns, leading/trailing
spaces in columns, and non-UTF-8 encoding issues.

By default input is expected from standard in and output is sent to 
standard out (errors to standard error). These can be modified by
appropriate options. The csv file is processed as a stream of rows so 
minimal memory is used to operate on the file. 

`

	examples = `

EXAMPLES

Normalizing a spread sheet's column count to 5 padding columns as needed per row.

    cat mysheet.csv | %s -field-per-row=5

Trim leading spaces.

    cat mysheet.csv | %s -left-trim-spaces

Trim trailing spaces.

    cat mysheet.csv | %s -right-trim-spaces

Trim leading and trailing spaces

    cat mysheet.csv | %s -trim-spaces

`

	// Standard Options
	showHelp     bool
	showLicense  bool
	showVersion  bool
	showExamples bool
	inputFName   string
	outputFName  string

	// App Options
	comma             string
	rowComment        string
	fieldsPerRecord   int
	lazyQuotes        bool
	trailingComma     bool
	trimSpace         bool
	trimLeadingSpace  bool
	trimTrailingSpace bool
	reuseRecord       bool
	commaOut          string
	useCRLF           bool
	stopOnError       bool

	verbose bool
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
	flag.IntVar(&fieldsPerRecord, "fields-per-row", 0, "set the number of columns to output right padding empty cells as needed")
	flag.BoolVar(&lazyQuotes, "use-lazy-quoting", false, "If LazyQuotes is true, a quote may appear in an unquoted field and a non-doubled quote may appear in a quoted field.")
	flag.BoolVar(&trimSpace, "trim-spaces", false, "If set to true leading and trailing white space in a field is ignored.")
	flag.BoolVar(&trimLeadingSpace, "left-trim-spaces", false, "If set to true leading white space in a field is ignored.")
	flag.BoolVar(&trimTrailingSpace, "right-trim-spaces", false, "If set to true trailing white space in a field is ignored.")
	flag.BoolVar(&reuseRecord, "reuse", true, "if false then a new array is allocated for each row processed, if true the array gets reused")
	flag.StringVar(&comma, "comma", "", "if set use this character in place of a comma for delimiting cells")
	flag.StringVar(&rowComment, "comment-char", "", "if set, rows starting with this character will be ignored as comments")
	flag.StringVar(&commaOut, "output-comma", "", "if set use this character in place of a comma for delimiting output cells")
	flag.BoolVar(&useCRLF, "use-crlf", false, "if set use a charage return and line feed in output")
	flag.BoolVar(&stopOnError, "stop-on-error", false, "exit on error, useful if you're trying to debug a problematic CSV file")

	flag.BoolVar(&verbose, "verbose", false, "write verbose output to standard error")
}

func main() {
	appName := path.Base(os.Args[0])
	flag.Parse()
	args := flag.Args()

	cfg := cli.New(appName, "", datatools.Version)
	cfg.UsageText = fmt.Sprintf(usage, appName)
	cfg.DescriptionText = fmt.Sprintf(description, appName)
	cfg.OptionText = "OPTIONS"
	cfg.ExampleText = fmt.Sprintf(examples, appName, appName, appName, appName)

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

	// Loop through input CSV, apply options, write to output CSV
	if trimSpace == true {
		trimLeadingSpace = true
		trimTrailingSpace = true
	}

	// Setup our CSV reader with any cli options
	var rStr []rune

	r := csv.NewReader(in)
	if comma != "" {
		rStr = []rune(comma)
		if len(rStr) > 0 {
			r.Comma = rStr[0]
		}
	}
	if rowComment != "" {
		rStr = []rune(rowComment)
		if len(rStr) > 0 {
			r.Comment = rStr[0]
		}
	}
	r.FieldsPerRecord = fieldsPerRecord
	r.LazyQuotes = lazyQuotes
	r.TrimLeadingSpace = trimLeadingSpace
	r.ReuseRecord = reuseRecord

	w := csv.NewWriter(out)
	if commaOut != "" {
		rStr = []rune(commaOut)
		if len(rStr) > 0 {
			w.Comma = rStr[0]
		}
	}
	w.UseCRLF = useCRLF

	// i is so we can track row count as we process each streamed in row
	expectedCellCount := 0
	hasError := false
	i := 1
	for {
		row, err := r.Read()
		if err == io.EOF {
			break
		}
		if i == 1 {
			expectedCellCount = len(row)
		}
		if err != nil {
			serr := fmt.Sprintf("%s", err)
			if strings.HasSuffix(serr, "wrong number of fields in line") == true && fieldsPerRecord >= 0 {
				if verbose {
					fmt.Fprintf(os.Stderr, "row %d: expected %d, got %d cells\n", i, expectedCellCount, len(row))
				}
				// Trim trailing cells if needed
				if fieldsPerRecord > 0 && len(row) >= fieldsPerRecord {
					row = row[0:fieldsPerRecord]
				}
				// Append cells if needed
				for len(row) < expectedCellCount {
					row = append(row, "")
				}
			} else {
				hasError = true
				if verbose {
					fmt.Fprintf(os.Stderr, "%s\n", err)
				}
			}
		}
		if trimSpace {
			for i := range row {
				s := row[i]
				row[i] = strings.TrimSpace(s)
			}
		}
		if err := w.Write(row); err != nil {
			fmt.Fprintf(os.Stderr, "error writing row %d: %s\n", i, err)
		}
		i++
		if verbose == true && (i%100) == 0 {
			fmt.Fprintf(os.Stderr, "Processed %d rows\n", i)
		}
		if hasError && stopOnError {
			os.Exit(1)
		}
	}
	// Finally we need to flush any remaining output...
	w.Flush()
	if verbose {
		fmt.Fprintf(os.Stderr, "Processed %d rows\n", i)
	}
}
