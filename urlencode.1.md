%urlencode(1) user manual | version 1.3.5 a8f53a7
% R. S. Doiel
% 2026-02-12

# NAME

urlencode

# SYNOPSIS

urlencode [OPTIONS] [STRING]

# DESCRIPTION

urlencode is a simple command line utility to URL encode content. By default
it reads from standard input and writes to standard out.  You can
also specifty the string to encode as a command line parameter.

You can provide the string to encode as a command line parameter otherwise it
will be read from standard input.

# OPTIONS

-help
: display help

-license
: display license

-version
: display version

-query
: Use URL query style encoding (use plus for spaces)

-newline
: Append a trailing newline

# EXAMPLE

echo "This is the string to encode & nothing else!" | urlencode

would yield

    This%20is%20the%20string%20to%20encode%20&%20nothing%20else%0A


