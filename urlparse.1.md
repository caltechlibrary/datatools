%urlparse(1) user manual | version 1.3.5 f86e208
% R. S. Doiel
% 2026-02-12

# NAME

urlparse

# SYNOPSIS

urlparse [OPTIONS] URL_TO_PARSE

# DESCRIPTION

urlparse can parse a URL and return the specific elements
requested (e.g. protocol, hostname, path, query string)

# OPTIONS

-help
: display help

-license
: display license

-version
: display version

-H, -host
: Display the hostname (and port if specified) found in URL.

-P, -protocol
: Display the protocol of URL (defaults to http)

-base, -basename
: Display the base filename at the end of the path.

-d, -delimiter
: Set the output delimited for parsed display. (defaults to tab)

-dir, -dirname
: Display all but the last element of the path

-ext, -extname
: Display the filename extension (e.g. .html).

-i, -input
: input filename

-nl, -newline
: if true add a trailing newline

-o, -output
: output filename

-p, -path
: Display the path after the hostname.

-quiet
: suppress error messages


# EXAMPLES

With no options returns "http\texample.com\t/my/page.html"

~~~
    urlparse http://example.com/my/page.html
~~~

Get protocol. Returns "http".

~~~
    urlparse -protocol http://example.com/my/page.html
~~~

Get host or domain name.  Returns "example.com".

~~~
    urlparse -host http://example.com/my/page.html
~~~

Get path. Returns "/my/page.html".

~~~
    urlparse -path http://example.com/my/page.html
~~~

Get dirname. Returns "my"

~~~
    urlparse -dirname http://example.com/my/page.html
~~~

Get basename. Returns "page.html".

~~~
    urlparse -basename http://example.com/my/page.html
~~~

Get extension. Returns ".html".

~~~
    urlparse -extname http://example.com/my/page.html
~~~

Without options urlparse returns protocol, host and path
fields separated by a tab.

urlparse 1.3.5

