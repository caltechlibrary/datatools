
# string trimprefix PREFIX [STRING]


This command will with trim all PREFIX from the STRING.

## Typical command line

```shell
    string trimprefix "--" "--people--"
```

Would return `people--`

## Piping content

NOTE: To read content from standard input we use the `-i -` option.

```shell
    echo "--people--" | string -i - trimprefix "--"
```

Would return `people--`

