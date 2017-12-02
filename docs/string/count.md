
# string count TARGET [STRING]

This command will count the number of times TARGET occurs in
the STRING.

## Typical command line

```shell
    string count "friend" "The people were freindly"
```

Would return `1`

```shell
    string count "tomato" "The people were freindly"
```

Would return `0`

## Piping content

NOTE: To read content from standard input we use the `-i -` option.

```shell
    echo "The people were friendly" | string -i - count "friend"
```

Would return `1`

```shell
    echo "The people were friendly" | string -i - count "apple"
```

Would return `0`

