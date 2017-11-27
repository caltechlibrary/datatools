
# USAGE

## trimsuffix [OPTIONS] PREFIX_STRING [STRINGS_TO_CHECK] 

## SYNOPSIS

trimsuffix returns a version of the string(s) without the designated suffix string.

## OPTIONS

```
    -example    display example(s)
    -h    display help
    -help    display help
    -i    input filename
    -input    input filename
    -l    display license
    -license    display license
    -new-line    include a trailing newline in output
    -nl    include a trailing newline in output
    -o    output filename
    -output    output filename
    -v    display version
    -version    display version
```

## EXAMPLES

Trim the suffix "ly" from the words "unknowningly" and "unusually"

```shell
    trimsuffix "ly" "unknowningly" "unusually"
```

Should yield

```
    unknowning
    unusual
```

Trim the suffix "ly" from the the words "unknowningly" and "common"

```shell
    trimsuffix "ly" "unknowningly" "common"
```

Should yield

```
    unknowning
    common
```

trimsuffix v0.0.18
