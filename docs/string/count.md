
# string ount TARGET [STRING]

This command will count the number of times TARGET occurs in
the STRING.

## using command line paramters only

```shell
    string count "friend" "The people were freindly"
```

Would return `1`

```shell
    string count "tomato" "The people were freindly"
```

Would return `0`

## Piping in the strings to check

```shell
    echo "The people were friendly" | string count "friend"
```

Would return `1`

```shell
    echo "The people were friendly" | string count "apple"
```

Would return `0`

