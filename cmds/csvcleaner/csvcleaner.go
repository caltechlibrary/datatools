package main

import (
	"encoding/csv"
	"fmt"
	"io"
	"os"
	"strings"

	// Caltech Library Packages
	"github.com/caltechlibrary/cli"
	"github.com/caltechlibrary/datatools"
)

var (
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
	showHelp             bool
	showLicense          bool
	showVersion          bool
	showExamples         bool
	inputFName           string
	outputFName          string
	generateMarkdownDocs bool
	quiet                bool
	newLine              bool

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

func main() {
	app := cli.NewCli(datatools.Version)
	appName := app.AppName()

	// Add Help Docs
	app.AddHelp("license", []byte(fmt.Sprintf(datatools.LicenseText, appName, datatools.Version)))
	app.AddHelp("description", []byte(fmt.Sprintf(description, appName)))
	app.AddHelp("examples", []byte(fmt.Sprintf(examples, appName, appName, appName, appName)))

	// Standard options
	app.BoolVar(&showHelp, "h,help", false, "display help")
	app.BoolVar(&showLicense, "l,license", false, "display license")
	app.BoolVar(&showVersion, "v,version", false, "display version")
	app.BoolVar(&showExamples, "examples", false, "display example(s)")
	app.StringVar(&inputFName, "i,input", "", "input filename")
	app.StringVar(&outputFName, "o,output", "", "output filename")
	app.BoolVar(&generateMarkdownDocs, "generate-markdown-docs", false, "generation markdown documentation")
	app.BoolVar(&quiet, "quiet", false, "suppress error messages")
	app.BoolVar(&newLine, "nl,newline", true, "include trailing newline in output")

	// Application specific options
	app.IntVar(&fieldsPerRecord, "fields-per-row", 0, "set the number of columns to output right padding empty cells as needed")
	app.BoolVar(&lazyQuotes, "use-lazy-quoting", false, "If LazyQuotes is true, a quote may appear in an unquoted field and a non-doubled quote may appear in a quoted field.")
	app.BoolVar(&trimSpace, "trim-spaces", false, "If set to true leading and trailing white space in a field is ignored.")
	app.BoolVar(&trimLeadingSpace, "left-trim-spaces", false, "If set to true leading white space in a field is ignored.")
	app.BoolVar(&trimTrailingSpace, "right-trim-spaces", false, "If set to true trailing white space in a field is ignored.")
	app.BoolVar(&reuseRecord, "reuse", true, "if false then a new array is allocated for each row processed, if true the array gets reused")
	app.StringVar(&comma, "comma", "", "if set use this character in place of a comma for delimiting cells")
	app.StringVar(&rowComment, "comment-char", "", "if set, rows starting with this character will be ignored as comments")
	app.StringVar(&commaOut, "output-comma", "", "if set use this character in place of a comma for delimiting output cells")
	app.BoolVar(&useCRLF, "use-crlf", false, "if set use a charage return and line feed in output")
	app.BoolVar(&stopOnError, "stop-on-error", false, "exit on error, useful if you're trying to debug a problematic CSV file")

	app.BoolVar(&verbose, "verbose", false, "write verbose output to standard error")

	// Parse environment and options
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

	// Process options
	if generateMarkdownDocs {
		app.GenerateMarkdownDocs(app.Out)
		os.Exit(0)
	}
	if showHelp || showExamples {
		if len(args) > 0 {
			fmt.Fprintln(app.Out, app.Help(args...))
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

	nl := "\n"
	if newLine == false {
		nl = ""
	}

	// Loop through input CSV, apply options, write to output CSV
	if trimSpace == true {
		trimLeadingSpace = true
		trimTrailingSpace = true
	}

	// Setup our CSV reader with any cli options
	var rStr []rune

	r := csv.NewReader(app.In)
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

	w := csv.NewWriter(app.Out)
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
					cli.OnError(app.Eout, fmt.Errorf("row %d: expected %d, got %d cells\n", i, expectedCellCount, len(row)), quiet)
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
					cli.OnError(app.Eout, fmt.Errorf("%s", err), quiet)
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
			cli.OnError(app.Eout, fmt.Errorf("error writing row %d: %s", i, err), quiet)
			hasError = true
		}
		i++
		if verbose == true && (i%100) == 0 {
			cli.OnError(app.Eout, fmt.Errorf("Processed %d rows\n", i), quiet)
		}
		if hasError && stopOnError {
			os.Exit(1)
		}
	}
	// Finally we need to flush any remaining output...
	w.Flush()
	err = w.Error()
	cli.ExitOnError(app.Eout, err, quiet)
	if verbose == true {
		cli.OnError(app.Eout, fmt.Errorf("Processed %d rows\n", i), quiet)
	}
	fmt.Fprintf(app.Out, "%s", nl)
}
