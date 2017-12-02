
# string join DELIMITER [JSON_ARRAY]

This command will return take a JSON array of strings
and return a single string joined with the delimiter.

## Typical command line

```shell
    string join '|' '["one","two","three"]'
```

Would return `one|two|three`

## Piping content

NOTE: To read content from standard input we use the `-i -` option.

```shell
    echo '["one","two","three"]' | string join '|'
```

Would return `one|two|three`



