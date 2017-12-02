
# string trimsuffix SUFFIX [STRING]


This command will with trim all the SUFFIX from the STRING.

## Typical command line

```shell
    string trimsuffix "--" "--people--"
```

Would return `--people`

## Piping content

NOTE: To read content from standard input we use the `-i -` option.

```shell
    echo "--people--" | string -i - trimsuffix "--"
```

Would return `--people`

