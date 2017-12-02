
# string replacen OLD NEW N [STRING]

This command will return the replace the OLD value with NEW
for up to N times in STRING. If N is negative one then all
OLD values will be replaced with NEW.

## Typical command line

NOTE: we'll use the option `-nl` to append a new line to the output
and make it easier to read.

```shell
    string -nl replace "e" "@" 2 "The people were friendly"
```

Would return `The p@opl@ were friendly`

```shell
    string replace "e" "@" -1 "The people were friendly"
```

Would return `The p@opl@ w@r@ fri@ndly`

## Piping content

NOTE: To read content from standard input we use the `-i -` option.

```shell
    echo "The people were friendly" | string -nl -i - replace "e" "@" 2
```

Would return `The p@opl@ were friendly`

```shell
    echo "The people were friendly" | string -nl -i - replace "e" "@" -1
```

Would return `The p@opl@ w@r@ fri@ndly`

