
# USAGE

	timefmt [OPTIONS] TIME_STRING_TO_CONVERT

## DESCRIPTION


timefmt formats the current date or INPUT_DATE based on the output format
provided in options. The default input and  output format is RFC3339. 
Formats are specified based on Golang's time package including the
common constants (e.g. RFC822, RFC1123). 

For details see https://golang.org/pkg/time/#Time.Format.

One additional time layout provided by timefmt 
 
+ mysql "2006-01-02 15:04:05 -0700"


## OPTIONS

Below are a set of options available.

```
    -examples            display example(s)
    -generate-manpage    generate man page
    -generate-markdown   generate markdown documentation
    -h, -help            display help
    -if, -input-format   Set format for input
    -l, -license         display license
    -nl, -newline        if true add a trailing newline
    -o, -output          output filename
    -of, -output-format  Set format for output
    -quiet               suppress error messages
    -utc                 timestamps in UTC
    -v, -version         display version
```


## EXAMPLES


Format the date July, 7, 2016 in YYYY-MM-DD format

    timefmt -if "2006-01-02" -of "01/02/2006" "2017-12-02"

Yields "12/02/2017"

Format the MySQL date/time of 8:08am, July 2, 2016

    timefmt -input-format mysql -output-format RFC822  "2017-12-02 08:08:08"

Yields "02 Dec 17 08:08 UTC"


timefmt v0.0.25
