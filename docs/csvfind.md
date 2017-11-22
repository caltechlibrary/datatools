
# USAGE

## csvfind [OPTIONS] TEXT_TO_MATCH

## SYNOPSIS

csvfind processes a CSV file as input returning rows that contain the column
with matched text. Columns are count from one instead of zero. Supports 
exact match as well as some Levenshtein matching.

## OPTIONS

	-allow-duplicates	allow duplicates when searching for matches
	-append-edit-distance	append column with edit distance found (useful for tuning levenshtein)
	-case-sensitive	perform a case sensitive match (default is false)
	-col	column to search for match in the CSV file
	-contains	use contains phrase for matching
	-d	set delimiter character
	-delete-cost	set the delete cost to use for levenshtein matching
	-delimiter	set delimiter character
	-h	display help
	-help	display help
	-i	input filename
	-input	input filename
	-insert-cost	set the insert cost to use for levenshtein matching
	-l	display license
	-levenshtein	use levenshtein matching
	-license	display license
	-max-edit-distance	set the edit distance thresh hold for match, default 0
	-o	output filename
	-output	output filename
	-skip-header-row	skip the header row
	-stop-words	use the colon delimited list of stop words
	-substitute-cost	set the substitution cost to use for levenshtein matching
	-trim-spaces	trim spaces around cell values before comparing
	-v	display version
	-version	display version

## EXAMPLES

Find the rows where the third column matches "The Red Book of Westmarch" exactly

```shell
    csvfind -i books.csv -col=2 "The Red Book of Westmarch"
```

Find the rows where the third column (colums numbered 0,1,2) matches approximately 
"The Red Book of Westmarch"

```shell
    csvfind -i books.csv -col=2 -levenshtein \
       -insert-cost=1 -delete-cost=1 -substitute-cost=3 \
       -max-edit-distance=50 -append-edit-distance \
       "The Red Book of Westmarch"
```

In this example we've appended the edit distance to see how close the matches are.

You can also search for phrases in columns.

```shell
    csvfind -i books.csv -col=2 -contains "Red Book"
```


csvfind v0.0.17
