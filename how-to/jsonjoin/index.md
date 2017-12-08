
# Using jsonjoin

Consider two JSON objects one in person.json and another
in profile.json.

person.json contains

```shell
   { "name": "Doe, Jane", "email":"jd@example.org", "age": 42 }
```

profile.json contains

```json
   { "name": "Doe, Jane", "bio": "World renowned geophysist.",
     "email": "jane.doe@example.edu" }
```

A simple join of person.json with profile.json (note the
-create option)

```shell
   jsonjoin -create person.json profile.json
```

would yield and object like

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

