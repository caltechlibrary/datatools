
# string trimleft CUTSET [STRING]

This command will with trim all the characters in
CUTSET from left side of STRING.

## Typical command line

```shell
    string trimleft "-" "--people--"
```

Would return `people--`

## Piping content

NOTE: To read content from standard input we use the `-i -` option.

```shell
    echo "--people--" | string -i - trimleft "-"
```

Would return `people--`

