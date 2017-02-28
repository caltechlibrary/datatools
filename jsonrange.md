
# USAGE

    jsonrange [OPTIONS] JSON_EXPRESSION 

## SYSNOPSIS

jsonrange turns either the JSON expression that is a map or array into delimited
elements suitable for processing in a "for" style loop in Bash. If the
JSON expression is an array then the elements of the array are returned else
if the expression is a map/object then the keys or attribute names are turned.

+ EXPRESSION can be an empty string contains a JSON array or map.

## OPTIONS

```
	-d	set delimiter for range output
	-delimiter	set delimiter for range output
	-dotpath	range on given dot path
	-h	display help
	-i	read JSON from file
	-input	read JSON from file
	-l	display license
	-length	return the number of keys or values
	-last	return the index of the last element
	-limit	limit the number of items output
	-p	range on given dot path
	-v	display version
```

## EXAMPLES

Working with a map

```
    jsonrange '{"name": "Doe, Jane", "email":"jane.doe@example.org", "age": 42}'
```

This would yield

```
    name
    email
    age
```

Working with an array

```
    jsonrange '["one", 2, {"label":"three","value":3}]'
```

would yield

```
    one
    2
    {"label":"three","value":3}
```

Checking the length of a map or array

```
    jsonrange -length '["one","two","three"]'
```

would yield

```
    3
```

Checking the last element index of a an array

```
    jsonrange -last '["one","two","three"]'
```

would yield

```
    2
```


Limitting the number of items returned

```
    jsonrange -limit 2 '[1,2,3,4,5]'
```

would yield

```
    1
    2
```

