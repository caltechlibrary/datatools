%mergepath(1) user manual | version 1.2.12 03b4ff7
% R. S. Doiel
% 2024-11-14

# NAME

mergepath

# SYNOPSIS

mergepath [OPTIONS] NEW_PATH_PARTS

# DESCRIPTION

mergepath can merge the new path parts with the existing path with
creating duplications.  It can also re-order existing path elements by
prefixing or appending to the existing path and removing the resulting
duplicate.

# OPTIONS

-help
: display help

-license
: display license

-version
: display version

-a, -append
: Append the directory to the path removing any duplication

-c, -clip
: Remove a directory from the path

-d, -directory
: The directory you want to add to the path.

-e, -envpath
: The path you want to merge with.

-nl, -newline
: if true add a trailing newline

-p, -prepend
: Prepend the directory to the path removing any duplication

-quiet
: suppress error messages


# EXAMPLES

This would put your home bin directory at the beginning of your path.

~~~
	export PATH=$(mergepath -p $HOME/bin)
~~~

mergepath 1.2.12

