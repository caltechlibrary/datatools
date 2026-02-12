%urldecode(1) user manual | version 1.3.5 f86e208
% R. S. Doiel
% 2026-02-12

# NAME

urldecode

# SYNOPSIS

urldecode [OPTIONS] [URL_ENCODED_STRING]

# DESCRIPTION

urldecode is a simple command line utility to URL decode content. By default
it reads from standard input and writes to standard out.  You can
also specifty the string to decode as a command line parameter.

You can provide the URL encoded string as a command line parameter or if none
present it will be read from standard input.

# OPTIONS

-help
: display help

-license
: display license

-version
: display version

-query
: use query escape (pluses for spaces)

-newline
: Append a trailing newline

# EXAMPLE

echo 'This%20is%20the%20string%20to%20encode%20&%20nothing%20else%0A' | urldecode

would yield (without the double quotes)

	"This is the string to encode & nothing else!" 


