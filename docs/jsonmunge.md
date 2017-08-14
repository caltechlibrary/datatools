
# jsonmunge

## USAGE

    jsonmunge [OPTIONS] TEMPLATE_FILENAME

## SYSNOPSIS

jsonmunge is a command line tool that takes a JSON document and
one or more Go templates rendering the results. Useful for
reshaping a JSON document, transforming into a new format,
or filter for specific content.

+ TEMPLATE_FILENAME is the name of a Go text tempate file used to render
  the outbound JSON document

## OPTIONS

```
	-h	display help
	-i	input filename
	-input	input filename
	-l	display license
	-o	output filename
	-output	output filename
	-v	display version
```

## EXAMPLES

If person.json contained

```json
   {"name": "Doe, Jane", "email":"jd@example.org", "age": 42}
```
and the template, name.tmpl, contained 

```template
   {{- .name -}}
```
Getting just the name could be done with

```shell
    cat person.json | jsonmunge name.tmpl
```
This would yeild

```shell
    "Doe, Jane"
```

jsonmunge v0.0.9
