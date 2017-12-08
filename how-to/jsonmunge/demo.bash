#!/bin/bash

echo -n 'Generating person.json: '
cat <<EOF > person.json
{"name": "Doe, Jane", "email":"jd@example.org", "age": 42}
EOF
cat person.json
echo ''

echo -n 'Generating name.tmpl: '
cat <<EOF > name.tmpl
{{- printf "%q" .name -}}
EOF
cat name.tmpl
echo ''

echo ' running: cat person.json | jsonmunge name.tmpl'
cat person.json | jsonmunge name.tmpl > result1.json
echo 'expected: "Doe, Jane"'
echo -n '     got: '
cat result1.json
echo ''
