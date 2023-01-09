---
title: "{app_name} (1) user manual"
author: "R. S. Doiel"
pubDate: 2023-01-09
---

# NAME

{app_name}

# SYNOPSIS

{app_name} [OPTIONS] URL_TO_PARSE

# DESCRIPTION

{app_name} can parse a URL and return the specific elements
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
    {app_name} http://example.com/my/page.html
~~~

Get protocol. Returns "http".

~~~
    {app_name} -protocol http://example.com/my/page.html
~~~

Get host or domain name.  Returns "example.com".

~~~
    {app_name} -host http://example.com/my/page.html
~~~

Get path. Returns "/my/page.html".

~~~
    {app_name} -path http://example.com/my/page.html
~~~

Get dirname. Returns "my"

~~~
    {app_name} -dirname http://example.com/my/page.html
~~~

Get basename. Returns "page.html".

~~~
    {app_name} -basename http://example.com/my/page.html
~~~

Get extension. Returns ".html".

~~~
    {app_name} -extname http://example.com/my/page.html
~~~

Without options {app_name} returns protocol, host and path
fields separated by a tab.

{app_name} {version}
