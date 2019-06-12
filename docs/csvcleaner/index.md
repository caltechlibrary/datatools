
# USAGE

	csvcleaner [OPTIONS]

## DESCRIPTION


csvcleaner normalizes a CSV file based on the options selected. It
helps to address issues like variable number of columns, leading/trailing
spaces in columns, and non-UTF-8 encoding issues.

By default input is expected from standard in and output is sent to 
standard out (errors to standard error). These can be modified by
appropriate options. The csv file is processed as a stream of rows so 
minimal memory is used to operate on the file. 


## OPTIONS

Below are a set of options available.

```
    -V, -verbose          write verbose output to standard error
    -comma                if set use this character in place of a comma for delimiting cells
    -comment-char         if set, rows starting with this character will be ignored as comments
    -examples             display example(s)
    -fields-per-row       set the number of columns to output right padding empty cells as needed
    -generate-manpage     generation man page
    -generate-markdown    generation markdown documentation
    -h, -help             display help
    -i, -input            input filename
    -l, -license          display license
    -left-trim            left trim spaces on CSV out
    -o, -output           output filename
    -output-comma         if set use this character in place of a comma for delimiting output cells
    -quiet                suppress error messages
    -reuse                if false then a new array is allocated for each row processed, if true the array gets reused
    -right-trim           right trim spaces on CSV out
    -stop-on-error        exit on error, useful if you're trying to debug a problematic CSV file
    -trim, -trim-spaces   trim spaces on CSV out
    -trim-leading-space   trim leading space from field(s) for CSV input
    -use-crlf             if set use a charage return and line feed in output
    -use-lazy-quotes      use lazy quotes for CSV input
    -v, -version          display version
```


## EXAMPLES


Normalizing a spread sheet's column count to 5 padding columns as needed per row.

    cat mysheet.csv | csvcleaner -field-per-row=5

Trim leading spaces from output.

    cat mysheet.csv | csvcleaner -left-trim

Trim trailing spaces from output.

    cat mysheet.csv | csvcleaner -right-trim

Trim leading and trailing spaces from output.

    cat mysheet.csv | csvcleaner -trim-space


csvcleaner v0.0.25
