
# string slice DELIMITER [JSON_ARRAY]

This command will return slice a string starting
at START (inclusive) and finishing at END (exclusive).

## Typical command line

```shell
    string slice 3 8 'one|two|three'
```

Would return `|two|`

## Piping content

NOTE: To read content from standard input we use the `-i -` option.

```shell
    echo 'one|two|three' | string -i - slice 3 8
```

Would return `|two|`

