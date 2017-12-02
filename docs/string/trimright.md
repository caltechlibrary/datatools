
# string trimright CUTSET [STRING]


This command will with trim all the characters in
CUTSET from right side of the STRING.

## Typical command line

```shell
    string trimright "-" "--people--"
```

Would return `--people`

## Piping content

NOTE: To read content from standard input we use the `-i -` option.

```shell
    echo "--people--" | string -i - trimright "-"
```

Would return `--people`

