
USAGE: json2toml [OPTIONS] [JSON_FILENAME] [TOML_FILENAME]

DESCRIPTION

json2toml is a tool that converts JSON objects into TOML output.

OPTIONS

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


EXAMPLES

These would get the file named "my.json" and save it as my.toml

    json2toml my.json > my.toml

	json2toml my.json my.toml

	cat my.json | json2toml -i - > my.toml

json2toml 1.2.1
