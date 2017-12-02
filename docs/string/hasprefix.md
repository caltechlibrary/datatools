
# string hasprefix TARGET [STRING]

This command will return true if the string has
the prefix or false otherwise.

## Typical command line

```shell
    string hasprefix "frei" "freindly"
```

Would return `true`

```shell
    string hasprefix "ing" "freindly"
```

Would return `false`


## Piping content

NOTE: To read content from standard input we use the `-i -` option.

```shell
    echo "freindly" | string -i - hasprefix "frei" 
```

Would return `true`

```shell
    echo "freindly" | string -i - hasprefix "ing" 
```

Would return `false`


