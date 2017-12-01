
# string position TARGET [STRING]

This command will return the position (where zero is the first position)
of TARGET in STRING. If the TARGET is not found a negative 1 will be returned.

## Typical command line

NOTE: we'll use the option `-nl` to append a new line to the output
and make it easier to read.

```shell
    string -nl position "friend" "The people were friendly"
```

Would return `16`

```shell
    string position "tomato" "The people were friendly"
```

Would return `-1`

## Piping content

NOTE: To read content from standard input we use the `-i -` option.

```shell
    echo "The people were friendly" | string -nl -i - position "friend"
```

Would return `16`

```shell
    echo "The people were friendly" | string -nl -i - position "apple"
```

Would return `-1`

