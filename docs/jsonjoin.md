# USAGE

## jsonjoin [OPTIONS] JSON_FILE_1 JSON_FILE_2

## SYSNOPSIS

jsonjoin is a command line tool that takes two (or more) JSON object 
documents and combines into a new JSON object document based on 
the options chosen.

## OPTIONS

	-h	display help
	-help	display help
	-l	display license
	-license	display license
	-o	output filename
	-output	output filename
	-overwrite	copy all key/values from second object into the first
	-update	copy unique key/values from second object into the first
	-v	display version
	-version	display version

## EXAMPLES

### Joining two JSON objects (maps)

person.json containes

```json
   {"name": "Doe, Jane", "email":"jd@example.org", "age": 42}
```

profile.json containes

```json
   {"name": "Doe, Jane", "bio": "World renowned geophysist.",
   	"email": "jane.doe@example.edu"}
```

A simple join of person.json with profile.json

```shell
   jsonjoin person.json profile.json
```

would yeild

```json
   {
   	"person": {"name": "Doe, Jane", "email":"jd@example.org", "age": 42},
    "profile": {"name": "Doe, Jane", "bio": "World renowned geophysist.", 
				"email": "jane.doe@example.edu"}
	}
```

You can modify this behavor with -add or -merge. Both options are
order dependant (i.e. not guaranteed to be associative, A add B does
not necessarily equal B add A). 

+ -update will add unique key/values from the second object to the first object
+ -overwrite replace key/values in first object one with second objects'

Running

```shell
	jsonjoin -update person.json profile.json
```

would yield

```json
   { "name": "Doe, Jane", "email":"jd@example.org", "age": 42,
     "bio": "World renowned geophysist." }
```

Running

```shell
	jsonjoin -update profile.json person.json
```

would yield

```json
   	{ "name": "Doe, Jane",  "age": 42, 
		"bio": "World renowned geophysist.", 
		"email": "jane.doe@example.edu" }
```

Running 

```shell
	jsonjoin -overwrite person.json profile.json
```

would yield

```json
   	{ "name": "Doe, Jane", "email":"jane.doe@example.edu", "age": 42,
    	"bio": "World renowned geophysist." }
```


jsonjoin v0.0.12
