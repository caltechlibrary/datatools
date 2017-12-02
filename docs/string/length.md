
# string length [STRING]

This command will return the length of STRING.

## Typical command line

```shell
    string length "friend"
```

Would return `6`

```shell
    string length "plum"
```

Would return `4`

## Piping content

NOTE: To read content from standard input we use the `-i -` option.

```shell
    echo -n "friend" | string -i - length
```

Would return `6`

```shell
    echo -n "plum" | string -i - length
```

Would return `4`

