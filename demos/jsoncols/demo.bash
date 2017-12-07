#!/bin/bash

cat <<EOF >myblob.json
{"name": "Doe, Jane", "email":"jane.doe@example.org", "age": 42}
EOF

echo ' running: jsoncols -i myblob.json .name'
jsoncols -i myblob.json .name > result1.txt
echo 'expected: "Doe, Jane"'
echo -n '     got: '
cat result1.txt
echo ''

echo ' running: jsoncols -i myblob.json -d "|"  .name .age'
jsoncols -i myblob.json -d "|"  .name .age > result2.txt
echo 'expected: "Doe, Jane"|42'
echo -n '     got: '
cat result2.txt
echo ''

echo ' running: cat myblob.json | jsoncols .name .email .age'
cat myblob.json | jsoncols .name .email .age > result3.txt
echo 'expected: "Doe, Jane","jane.doe@xample.org",42'
echo -n '     got: '
cat result3.txt
echo ''
