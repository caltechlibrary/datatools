
# USAGE

## csvjoin [OPTIONS] CSV1 CSV2 COL1 COL2

## SYNOPSIS

csvjoin outputs CSV content based on two CSV files with matching column values.
Each CSV input file has a designated column to match on. The values are
compared as strings. Columns are counted from one rather than zero.

## OPTIONS

	-allow-duplicates	allow duplicates when searching for matches
	-case-sensitive	make a case sensitive match (default is case insensitive)
	-col1	column to on join on in first CSV file
	-col2	column to on join on in second CSV file
	-contains	match columns based on csv1/col1 contained in csv2/col2
	-csv1	first CSV filename
	-csv2	second CSV filename
	-d	set delimiter character
	-delete-cost	deletion cost to use when calculating Levenshtein edit distance
	-delimiter	set delimiter character
	-h	display help
	-help	display help
	-in-memory	if true read both CSV files
	-insert-cost	insertion cost to use when calculating Levenshtein edit distance
	-l	display license
	-levenshtein	match columns using Levensthein edit distance
	-license	display license
	-max-edit-distance	maximum edit distance for match using Levenshtein distance
	-o	output filename
	-output	output filename
	-stop-words	a column delimited list of stop words to ingnore when matching
	-substitute-cost	substitution cost to use when calculating Levenshtein edit distance
	-trim-spaces	trim spaces around cell values before comparing
	-v	display version
	-verbose	output processing count to stderr
	-version	display version

## EXAMPLES

Simple usage of building a merged CSV file from data1.csv
and data2.csv where column 1 in data1.csv matches the value in
column 3 of data2.csv with the results being written to 
merged-data.csv..

```shell
    csvjoin -csv1=data1.csv -col1=2 \
       -csv2=data2.csv -col2=4 \
       -output=merged-data.csv
```

csvjoin v0.0.14
