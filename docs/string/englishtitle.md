
# string englishtitle [STRING]

This command will with return a string in title format
using English capitialization conventions.

## Typical command line

```shell
    string englishtitle "THE PEOPLE WERE FREINDLY"
```

Would return `The People Were Freindly`

## Piping content

NOTE: To read content from standard input we use the `-i -` option.

```shell
    echo "THE PEOPLE WERE FREINDLY" | string -i - englishtitle 
```

Would return `The People Were Freindly`

