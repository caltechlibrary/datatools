
# string tolower [STRING]

This command will with return a lower case version of
string.

## Typical command line

```shell
    string tolower "THE PEOPLE WERE FREINDLY"
```

Would return `the people were freindly`

## Piping content

NOTE: To read content from standard input we use the `-i -` option.

```shell
    echo "THE PEOPLE WERE FREINDLY" | string -i - tolower 
```

Would return `the people were freindly`

