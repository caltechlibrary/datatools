
# USAGE

## jsonmunge [OPTIONS] TEMPLATE_FILENAME

## SYSNOPSIS

jsonmunge is a command line tool that takes a JSON document and
one or more Go templates rendering the results. Useful for
reshaping a JSON document, transforming into a new format,
or filter for specific content.

+ TEMPLATE_FILENAME is the name of a Go text tempate file used to render
  the outbound JSON document

## OPTIONS

```
	-example	display example(s)
	-h	display help
	-help	display help
	-i	input filename
	-input	input filename
	-license	display license
	-o	output filename
	-output	output filename
	-v	display version
	-version	display version
```

## EXAMPLES

If person.json contained

```json
   {"name": "Doe, Jane", "email":"jd@example.org", "age": 42}
```

and the template, name.tmpl, contained

```
   {{- .name -}}
```

Getting just the name could be done with

```shell
    cat person.json | jsonmunge name.tmpl
```

This would yield

```
    "Doe, Jane"
```

jsonmunge v0.0.18
