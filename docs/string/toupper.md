
# string toupper [STRING]

This command will with return a upper case version of
string.

## Typical command line

```shell
    string toupper "the people were freindly"
```

Would return `THE PEOPLE WERE FREINDLY`

## Piping content

NOTE: To read content from standard input we use the `-i -` option.

```shell
    echo "the people were freindly" | string -i - toupper 
```

Would return `THE PEOPLE WERE FREINDLY`

