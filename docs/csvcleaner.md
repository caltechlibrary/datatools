
# USAGE

## csvcleaner [OPTIONS]

## SYNOPSIS

csvcleaner normalizes a CSV file based on the options selected. It
helps to address issues like variable number of columns, leading/trailing
spaces in columns, and non-UTF-8 encoding issues.

By default input is expected from standard in and output is sent to 
standard out (errors to standard error). These can be modified by
appropriate options. The csv file is processed as a stream of rows so 
minimal memory is used to operate on the file. 

## OPTIONS

```
	-comma	if set use this character in place of a comma for delimiting cells
	-comment-char	if set, rows starting with this character will be ignored as comments
	-example	display example(s)
	-fields-per-row	set the number of columns to output right padding empty cells as needed
	-h	display help
	-help	display help
	-i	input filename
	-input	input filename
	-l	display license
	-left-trim-spaces	If set to true leading white space in a field is ignored.
	-license	display license
	-o	output filename
	-output	output filename
	-output-comma	if set use this character in place of a comma for delimiting output cells
	-reuse	if false then a new array is allocated for each row processed, if true the array gets reused
	-right-trim-spaces	If set to true trailing white space in a field is ignored.
	-stop-on-error	exit on error, useful if you're trying to debug a problematic CSV file
	-trim-spaces	If set to true leading and trailing white space in a field is ignored.
	-use-crlf	if set use a charage return and line feed in output
	-use-lazy-quoting	If LazyQuotes is true, a quote may appear in an unquoted field and a non-doubled quote may appear in a quoted field.
	-v	display version
	-verbose	write verbose output to standard error
	-version	display version
```

## EXAMPLES

Normalizing a spread sheet's column count to 5 padding columns as needed per row.

```shell
    cat mysheet.csv | csvcleaner -field-per-row=5
```

Trim leading spaces.

```shell
    cat mysheet.csv | csvcleaner -left-trim-spaces
```

Trim trailing spaces.

```shell
    cat mysheet.csv | csvcleaner -right-trim-spaces
```

Trim leading and trailing spaces

```shell
    cat mysheet.csv | csvcleaner -trim-spaces
```

csvcleaner v0.0.18
