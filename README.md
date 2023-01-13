
datatools
=========

_datatools_ is a rich collection of command line programs targetting
data conversion, cleanup and analysis directly from your favorite
POSIX shell. It has proven useful for data collaberations where
individual members of a project may prefer different toolsets in their
analysis (e.g. Julia, R, Python) but want to work from a common baseline.
It also has been used intensively for internal reporting from various
Caltech Library metadata sources.

The tools fall into three broad categories 

- data transformation and conversion
- shell scripting helpers
- "string", a tool providing the common string operations missing from shell

See [user manual](user-manual.md) for a complete list of the command line
programs. The data transformation tools include support for formats such as
Excel XML, csv, tab delimited files, json, yaml and toml.

Compiled versions of the datatools collection are provided for Linux
(amd64), Mac OS X (amd64), Windows 10 (amd64) and Raspbian (ARM7).
See https://github.com/caltechlibrary/datatools/releases.

Use "-help" option for a full list of options for each utility (e.g. `csv2json -help`).

Data transformation
-------------------

The tooling around transformation includes data conversion. These
include tools that work with CSV, tab delimited, JSON, TOML, YAML
and Excel XML.

There is also tooling to change data shapes using JSON as the
intermediate data format.

For the shell
-------------

Various utilities for simplifying work on the command line. 

+ [findfile](docs/findfile/) - find files based on prefix, suffix or contained string
+ [finddir](docs/finddir/) - find directories based on prefix, suffix or contained string
+ [mergepath](docs/mergepath/) - prefix, append, clip path variables
+ [range](docs/range/) - emit a range of integers (useful for numbered loops in Bash)
+ [reldate](docs/reldate/) - display a relative date in YYYY-MM-DD format
+ [reltime](docs/reltime/) - display a relative time in 24 hour notation, HH:MM:SS format
+ [timefmt](docs/timefmt/) - format a time value based on Golang's time format language
+ [urlparse](docs/urlparse/) - split a URL into parts

For strings
-----------

_datatools_ provides the [string](docs/string/) command for working with 
text strings (limited to memory available).  This is commonly needed when 
cleanup data for analysis. The _string_ command was created for when the 
old Unix standbys- grep, awk, sed, tr are unwieldly or inconvient. 
_string_ provides operations are common in most language like, trimming, 
spliting, and transforming letter case.  The _string_ command also makes 
it easy to join JSON string arrays into single a string using a delimiter 
or split a string into a JSON array based on a delimiter. The form of the 
command is `string [OPTIONS] [ACTION] [ARCTION_PARAMETERS...]`

```shell
    string toupper "one two three"
```

Would yield "ONE TWO THREE".

Some of the features included

+ change case (upper, lower, title, English title)
+ length, position and count of substrings
+ has prefix, suffix or contains
+ trim prefix, suffix and cutsets
+ split and join to/from JSON string arrays

See [string](docs/string/) for full details

Installation
------------

See [INSTALL.md](install.html) for details for installing pre-compiled 
versions of the programs.

