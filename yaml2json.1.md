
USAGE: yaml2json [OPTIONS] [YAML_FILENAME] [JSON_NAME]

DESCRIPTION

yaml2json is a tool that converts YAML into JSON output.

OPTIONS

    -examples           display example(s)
    -generate-manpage   generate man page
    -generate-markdown  generate markdown documentation
    -h, -help           display help
    -l, -license        display license
    -nl, -newline       if true add a trailing newline
    -o, -output         output filename
    -p, -pretty         pretty print output
    -quiet              suppress error messages
    -v, -version        display version


EXAMPLES

These would get the file named "my.yaml" and save it as my.json

    yaml2json my.yaml > my.json

    yaml2json my.yaml my.json

	cat my.yaml | yaml2json -i - > my.json

yaml2json 1.2.1
