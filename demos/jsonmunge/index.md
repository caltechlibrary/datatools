
# demo jsonmunge

If person.json contained

```json
   {"name": "Doe, Jane", "email":"jd@example.org", "age": 42}
```

and the template, name.tmpl, contained

```template
   {{- .name -}}
```

Getting just the name could be done with

```shell
    cat person.json | jsonmunge name.tmpl
```

This would yield

```
    "Doe, Jane"
```

