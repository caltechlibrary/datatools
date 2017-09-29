
# USAGE

## urlparse [OPTIONS] URL_TO_PARSE

## SYNOPSIS

urlparse can parse a URL and return the specific elements
requested (e.g. protocol, hostname, path, query string)

## OPTIONS

	-D	Set the output delimited for parsed display. (defaults to tab)
	-H	Display the hostname (and port if specified) found in URL.
	-P	Display the protocol of URL (defaults to http)
	-b	Display the base filename at the end of the path.
	-base	Display the base filename at the end of the path.
	-d	Display the base filename at the end of the path.
	-delimiter	Set the output delimited for parsed display. (defaults to tab)
	-directory	Display the base filename at the end of the path.
	-e	Display the filename extension (e.g. .html).
	-extension	Display the filename extension (e.g. .html).
	-h	display help
	-help	display help
	-host	Display the hostname (and port if specified) found in URL.
	-l	display license
	-license	display license
	-p	Display the path after the hostname.
	-path	Display the path after the hostname.
	-protocol	Display the protocol of URL (defaults to http)
	-v	display verison
	-version	display version

## EXAMPLE

With no options returns "http\texample.com\t/my/page.html"

```shell
    urlparse http://example.com/my/page.html
```

Get protocol. Returns "http".

```shell
    urlparse --protocol http://example.com/my/page.html
```

Get host or domain name.  Returns "example.com".

```shell
    urlparse --host http://example.com/my/page.html
```

Get path. Returns "/my/page.html".

```shell
    urlparse --path http://example.com/my/page.html
```

Get basename. Returns "page.html".

```shell
    urlparse --basename http://example.com/my/page.html
```

Get extension. Returns ".html".

```shell
    urlparse --extension http://example.com/my/page.html
```

Without options urlparse returns protocol, host and path
fields separated by a tab.


urlparse v0.0.14
