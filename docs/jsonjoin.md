
# USAGE

## jsonjoin [OPTIONS] JSON_FILE_1 [JSON_FILE_2 ...]

## SYSNOPSIS

jsonjoin is a command line tool that takes one (or more) JSON objects files 
and joins them to a root JSON object read from standard input (or 
file identified by -input option).  By default the resulting
joined JSON object is written to standard out.

The default behavior for jsonjoin is to create key/value pairs
based on the joined JSON document names and their contents. 
This can be thought of as a branching behavior. Each additional 
file becomes a branch and its key/value pairs become leafs. 
The root JSON object is assumed to come from standard input
but can be designated by the -input option or created by the
-create option. Each additional file specified as a command line
argument is then treated as a new branch.

In addition to the branching behavior you can join JSON objects in a 
flat manner.  The flat joining process can be ether non-distructive 
adding new key/value pairs (-update option) or distructive 
overwriting key/value pairs (-overwrite option).

Note: jsonjoin doesn't support a JSON array as the root JSON object.

## OPTIONS

	-create	create an empty root object, {}
	-h	display help
	-help	display help
	-i	input filename (for root object)
	-input	input filename (for root object)
	-l	display license
	-license	display license
	-o	output filename
	-output	output filename
	-overwrite	copy all key/values into root object
	-update	copy new key/values pairs into root object
	-v	display version
	-version	display version

## EXAMPLES

Consider two JSON objects one in person.json and another 
in profile.json.

person.json containes

```json
   { "name": "Doe, Jane", "email":"jd@example.org", "age": 42 }
```

profile.json containes

```json
   { "name": "Doe, Jane", "bio": "World renowned geophysist.",
     "email": "jane.doe@example.edu" }
```

A simple join of person.json with profile.json (note the 
-create option)

```shell
   jsonjoin -create person.json profile.json
```

would yeild and object like

```json
   {
     "person":  { "name": "Doe, Jane", "email":"jd@example.org", 
	 			"age": 42},
     "profile": { "name": "Doe, Jane", "bio": "World renowned geophysist.", 
                  "email": "jane.doe@example.edu" }
   }
```

Likewise if you want to treat person.json as the root object and add
profile.json as a branch try

```shell
   cat person.json | jsonjoin profile.json
```

or

```shell
   jsonjoin -i person.json profile.json
```

this yields an object like

```json
   {
     "name": "Doe, Jane", "email":"jd@example.org", "age": 42,
     "profile": { "name": "Doe, Jane", "bio": "World renowned geophysist.", 
                  "email": "jane.doe@example.edu" }
   }
```

You can modify this behavor with -update or -overwrite. Both options are
order dependant (i.e. not associative, A update B does
not necessarily equal B update A). 

+ -update will add unique key/values from the second object to the first object
+ -overwrite replace key/values in first object one with second objects'

Running

```shell
    jsonjoin -create -update person.json profile.json
```

would yield

```json
   { "name": "Doe, Jane", "email":"jd@example.org", "age": 42,
     "bio": "World renowned geophysist." }
```

Running

```shell
    jsonjoin -create -update profile.json person.json
```

would yield

```json
   { "name": "Doe, Jane",  "age": 42, 
     "bio": "World renowned geophysist.", 
     "email": "jane.doe@example.edu" }
```

Running 

```shell
    jsonjoin -create -overwrite person.json profile.json
```

would yield

```json
   { "name": "Doe, Jane", "email":"jane.doe@example.edu", "age": 42,
     "bio": "World renowned geophysist." }
```


jsonjoin v0.0.17
