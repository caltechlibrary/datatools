
# string totitle [STRING]

This command will with return a title case version of
string. This normally means it will be in upper case.

## Typical command line

```shell
    string totitle "the people were friendly"
```

Would return `THE PEOPLE WERE FRIENDLY`

## Piping content

NOTE: To read content from standard input we use the `-i -` option.

```shell
    echo "the people were friendly" | string -i - totitle 
```

Would return `THE PEOPLE WERE FRIENDLY`

