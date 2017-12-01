
# string replace OLD NEW [STRING]

This command will return the replace the OLD value with NEW
for the STRING.

## Typical command line

NOTE: we'll use the option `-nl` to append a new line to the output
and make it easier to read.

```shell
    string -nl replace "friendly" "enemies" "The people were friendly"
```

Would return `The people were enemies`

```shell
    string replace "tomato" "ememies" "The people were friendly"
```

Would return `The people were friendly`

## Piping content

NOTE: To read content from standard input we use the `-i -` option.

```shell
    echo "The people were friendly" | string -nl -i - replace "friendly" "enemies"
```

Would return `The people were enemies`

```shell
    echo "The people were friendly" | string -nl -i - replace "apple" "enemies"
```

Would return `The people were friendly`

