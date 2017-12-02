
# string hassuffix TARGET [STRING]

This command will return true if the string has
the suffix or false otherwise.

## Typical command line

```shell
    string hassuffix "ly" "freindly"
```

Would return `true`

```shell
    string hassuffix "ing" "freindly"
```

Would return `false`


## Piping content

NOTE: To read content from standard input we use the `-i -` option.

```shell
    echo "freindly" | string -i - hassuffix "ly" 
```

Would return `true`

```shell
    echo "freindly" | string -i - hassuffix "ing" 
```

Would return `false`


