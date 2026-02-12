%finddir(1) user manual | version 1.3.5 effbad2
% R. S. Doiel
% 2026-02-12

# NAME

finddir

# SYNOPSIS

finddir [OPTIONS] [TARGET] [DIRECTORIES_TO_SEARCH]

# DESCRIPTION

finddir finds directory based on matching prefix, suffix o 
contained text in base filename.

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

-e, -error-stop
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

Find all subdirectories starting with "img".

~~~
	finddir -p img
~~~

finddir 1.3.5

