
# string padright [STRING]

This command will with return a version of string
padding the left side with the characters provided up
to the a maximumn length of the new string.

## Typical command line

```shell
    string padright "-" 10 "people"
```

Would return `people------`

## Piping content

NOTE: To read content from standard input we use the `-i -` option.

```shell
    echo "people" | string -i - padright "-" 10
```

Would return `people------`

