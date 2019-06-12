
# USAGE

	json2yaml [OPTIONS] [JSON_FILENAME] [YAML_FILENAME]

## DESCRIPTION


json2yaml is a tool that converts JSON objects into YAML output.


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
    -quiet               suppress error messages
    -v, -version         display version
```


## EXAMPLES


These would get the file named "my.json" and save it as my.yaml

    json2yaml my.json > my.yaml

	%!s(MISSING) my.json my.taml

	cat my.json | %!s(MISSING) -i - > my.taml



json2yaml v0.0.24
