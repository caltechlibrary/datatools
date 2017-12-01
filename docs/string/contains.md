
# string contains TARGET [STRING]

This command will return "true" or "false" if strings contain the TARGET.
You can supply the STRING from standard input or as a command line argument.

## Typical command line

```shell
    string contains "friend" "The people were freindly"
```

Would return `true`

```shell
    string contains "tomato" "The people were freindly"
```

Would return `false`

## Piping content

NOTE: To read content from standard input we use the `-i -` option.

```shell
    echo "The people were friendly" | string -i - contains "friend"
```

Would return `true`

```shell
    echo "The people were friendly" | string -i - contains "apple"
```

Would return `false`

