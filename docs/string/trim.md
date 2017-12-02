
# string trim CUTSET [STRING]

This command will with trim all the characters in
CUTSET from the (beginning and end) of STRING.

## Typical command line

```shell
    string trim "-" "--people--"
```

Would return `people`

## Piping content

NOTE: To read content from standard input we use the `-i -` option.

```shell
    echo "--people--" | string -i - trim "-"
```

Would return `people`

