
# USAGE

## joinstring [OPTIONS] [STRINGS_TO_JOIN]

## SYNOPSIS

joinstring joins a JSON array or with options a new line delimited string into
a single string with output delimiter (default delimiter is an empty string).

## OPTIONS

```
	-d	set the output delimiting string value (default is empty string)
	-delimiter	set output delimiting string value (default is empty string)
	-example	display example(s)
	-h	display help
	-help	display help
	-i	input filename
	-input	input filename
	-l	display license
	-license	display license
	-newline	input as one substring per line rather than JSON
	-nl	input as one substring per line rather than JSON
	-o	output filename
	-output	output filename
	-v	display version
	-version	display version
```

## EXAMPLES

Joining a JSON array into a single string delimited by a double pipe.

```
    joinstring -d '||' '["one", "two", "three"]'
```

This should yield

```
    one||two||three
```

Joining a newline delimited file into a single string

```
    cat <<EOF > joinstring -nl -d '||' "one||two||three"
    one
	two
	three
    EOF
```

This should yield

```
    one||two||three
```

joinstring v0.0.18
