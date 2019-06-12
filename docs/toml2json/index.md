
# USAGE

	toml2json [OPTIONS] [TOML_FILENAME] [JSON_NAME]

## DESCRIPTION


toml2json is a tool that converts TOML into JSON output.


## OPTIONS

Below are a set of options available.

```
    -examples            display example(s)
    -generate-manpage    generate man page
    -generate-markdown   generate markdown documentation
    -h, -help            display help
    -l, -license         display license
    -nl, -newline        if true add a trailing newline
    -o, -output          output filename
    -p, -pretty          pretty print output
    -quiet               suppress error messages
    -v, -version         display version
```


## EXAMPLES


These would get the file named "my.toml" and save it as my.json

    toml2json my.toml > my.json

    toml2json my.toml my.json

	cat my.toml | toml2json -i - > my.json


toml2json v0.0.24
