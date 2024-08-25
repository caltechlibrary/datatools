%findfile(1) user manual | version 1.2.9 89f7b4d
% R. S. Doiel
% 2024-08-25

# NAME

findfile

# SYNOPSIS

findfile [OPTIONS] [TARGET] [DIRECTORIES_TO_SEARCH]

# DESCRIPTION

findfile finds files based on matching prefix, suffix or contained text in base filename.

# OPTIONS

-help
: display this help message

-license
: display license information

-version
: display version message

-c, -contains
: find file(s) based on basename containing text

-d, -depth
: Limit depth of directories walked

-error, -stop-on-error
: Stop walk on file system errors (e.g. permissions)

-f, -full-path
: list full path for files found

-m, -mod-time
: display file modification time before the path

-nl, -newline
: if true add a trailing newline

-o, -output
: output filename

-p, -prefix
: find file(s) based on basename prefix

-quiet
: suppress error messages

-s, -suffix
: find file(s) based on basename suffix


# EXAMPLES

Search the current directory and subdirectories for Markdown files with extension of ".md".

~~~
	findfile -s .md
~~~

findfile 1.2.9


