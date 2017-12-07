
# demo jsoncols

If myblob.json contained

```shell
    {"name": "Doe, Jane", "email":"jane.doe@example.org", "age": 42}
```

Getting just the name could be done with

```shell
    jsoncols -i myblob.json .name
```

This would yield

```
    "Doe, Jane"
```

Flipping .name and .age into pipe delimited columns is as
easy as listing each field in the expression inside a
space delimited string.

```shell
    jsoncols -i myblob.json -d\|  .name .age
```

This would yield

```
    Doe, Jane|42
```

You can also pipe JSON data in.

```shell
    cat myblob.json | jsoncols .name .email .age
```

Would yield

```
   "Doe, Jane",jane.doe@xample.org,42
```

