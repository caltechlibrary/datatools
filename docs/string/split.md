
# string split DELIMITER [STRING]

This command will return split a string at the delimiter
and return it as a JSON array.

## Typical command line

```shell
    string split '|' 'one|two|three'
```

Would return `["one","two","three"]`

## Piping content

NOTE: To read content from standard input we use the `-i -` option.

```shell
    echo 'one|two|three' | string -i - split '|'
```

Would return `["one","two","three"]`

