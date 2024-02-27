%csvfind(1) user manual | version 1.2.6 {release_hash}
% R. S. Doiel
% {release_date}

# NAME

csvfind

# SYNOPSIS

csvfind [OPTIONS] TEXT_TO_MATCH

# DESCRIPTION

csvfind processes a CSV file as input returning rows that contain
the column with matched text. Columns are counted from one instead of
zero. Supports exact match as well as some Levenshtein matching.

# OPTIONS

-help
: display help

-license
: display license

-version
: display version


-allow-duplicates
: allow duplicates when searching for matches

-append-edit-distance
: append column with edit distance found (useful for tuning levenshtein)

-case-sensitive
: perform a case sensitive match (default is false)

-col, -cols
: column to search for match in the CSV file

-contains
: use contains phrase for matching

-d, -delimiter
: set delimiter character

-delete-cost
: set the delete cost to use for levenshtein matching

-i, -input
: input filename

-insert-cost
: set the insert cost to use for levenshtein matching

-levenshtein
: use levenshtein matching

-max-edit-distance
: set the edit distance thresh hold for match, default 0

-nl, -newline
: include trailing newline from output

-o, -output
: output filename

-quiet
: suppress error messages

-skip-header-row
: skip the header row

-stop-words
: use the colon delimited list of stop words

-substitute-cost
: set the substitution cost to use for levenshtein matching

-trim-leading-space
: trim leadings space in field(s) for CSV input

-trimspace, -trimspaces
: trim spaces around cell values before comparing

-use-lazy-quotes
: use lazy quotes on CSV input


# EXAMPLES

Find the rows where the third column matches "The Red Book of Westmarch"
exactly

~~~
    csvfind -i books.csv -col=2 "The Red Book of Westmarch"
~~~

Find the rows where the third column (colums numbered 1,2,3) matches
approximately "The Red Book of Westmarch"

~~~
    csvfind -i books.csv -col=2 -levenshtein \
       -insert-cost=1 -delete-cost=1 -substitute-cost=3 \
       -max-edit-distance=50 -append-edit-distance \
       "The Red Book of Westmarch"
~~~

In this example we've appended the edit distance to see how close the
matches are.

You can also search for phrases in columns.

~~~
    csvfind -i books.csv -col=2 -contains "Red Book"
~~~

csvfind 1.2.6

