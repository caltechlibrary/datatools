
# demo jsonrange

Working with a map

```shell
    echo '{"name": "Doe, Jane", "email":"jane.doe@example.org", "age": 42}' \
       | jsonrange
```

This would yield

```
    name
    email
    age
```

Using the -values option on a map

```shell
    echo '{"name": "Doe, Jane", "email":"jane.doe@example.org", "age": 42}' \
      | jsonrange -values
```

This would yield

```
    "Doe, Jane"
    "jane.doe@example.org"
    42
```


Working with an array

```shell
    echo '["one", 2, {"label":"three","value":3}]' | jsonrange
```

would yield

```
    0
    1
    2
```

Using the -values option on the same array

```shell
    echo '["one", 2, {"label":"three","value":3}]' | jsonrange -values
```

would yield

```
    one
    2
    {"label":"three","value":3}
```

Checking the length of a map or array or number of keys in map

```shell
    echo '["one","two","three"]' | jsonrange -length
```

would yield

```
    3
```

Check for the index value of last element

```shell
    echo '["one","two","three"]' | jsonrange -last
```

would yield

```
    2
```

Limitting the number of items returned

```shell
    echo '[1,2,3,4,5]' | jsonrange -limit 2
```

would yield

```
    1
    2
```

