
# string contains TARGET [STRING]

This command will return "true" or "false" if strings contain the TARGET.
You can supply the STRING from standard input or as a command line argument.

## using command line paramters only

```shell
    string contains "friend" "The people were freindly"
```

Would return `true`

```shell
    string contains "tomator" "The people were freindly"
```

Would return `false`

## Piping in the strings to check

```shell
    echo "The people were friendly" | string contains "friend"
```

Would return `true`

```shell
    echo "The people were friendly" | string contains "apple"
```

Would return `false`

