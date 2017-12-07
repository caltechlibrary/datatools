#!/bin/bash

echo "Generating person.json"
cat <<EOF > person.json
{"name": "Doe, Jane", "email":"jane.doe@example.org", "age": 42}
EOF
cat person.json
echo ''

echo ' running: cat person.json | jsonrange -i -'
cat person.json | jsonrange -i - > result1.txt
cat <<EOF > expected1.txt
name
email
age
EOF
echo -n 'expected: '
cat expected1.txt
echo -n '     got: '
cat result1.txt
echo ''

echo ' running: jsonrange -i person.json -values'
jsonrange -i person.json -values | result2.txt
cat <<EOF > expected2.txt
"Doe, Jane"
"jane.doe@example.org"
42
EOF
echo -n 'expected: '
cat expected2.txt
echo -n '     got: '
cat result2.txt
echo ''

echo "Generating array1.json"
cat <<EOF > array1.json
["one", 2, {"label":"three","value":3}]
EOF

echo ' running: jsonrange -i array1.json'
jsonrange -i array1.json > result3.txt
cat <<EOF > expected3.txt
0
1
2
EOF
echo -n 'expected: '
cat expected3.txt
echo -n '     got: '
cat result3.txt
echo ''
cmp expected3.txt result3.txt

echo 'jsonrange -i array1.json -values'
jsonrange -i array1.json -values > result4.txt
cat <<EOF > expected4.txt
"one"
2
{"label":"three","value":3}
EOF
echo -n 'expected: '
cat expected4.txt
echo -n '     got: '
cat result4.txt
echo ''
cmp expected4.txt result4.txt

echo 'Generating array2.json'
echo '["one","two","three"]' > array2.json
cat array2.json
echo ''

echo ' running: jsonrange -i array2.json -length'
jsonrange -i array2.json -length > result5.txt
echo -n "3" > expected5.txt
echo -n 'expected: '
cat expected5.txt
echo -n '     got: '
cat result5.txt
echo ''
cmp expected5.txt result5.txt

echo ' running: jsonrange -i array2.json -last'
jsonrange -i array2.json -last > result6.txt
echo -n '2' > expected6.txt
echo -n 'expected: '
cat expected6.txt
echo -n '     got: '
cat result6.txt
echo ''
cmp expected6.txt result6.txt

echo ' running: jsonrange -i array2.json -values -last'
jsonrange -i array2.json -values -last > result7.txt
echo '"three"' > expected7.txt
echo -n 'expected: '
cat expected7.txt
echo -n '     got: '
cat result7.txt
echo ''
cmp expected7.txt result7.txt

echo 'Generating array3.json'
echo '[10,20,30,40,50]' > array3.json
echo ''

echo ' running: jsonrange -i array3.json -limit 2'
jsonrange -i array3.json -limit 2 > result8.txt
cat <<EOF > expected8.txt
1
2
EOF
echo -n 'expected: '
cat expected8.txt
echo -n '     got: '
cat result8.txt
echo ''
cmp expected8.txt result8.txt

echo ' running: jsonrange -i array3.json -values -limit 2'
jsonrange -i array3.json -values -limit 2 > result9.txt
cat <<EOF > expected9.txt
10
20
EOF
echo -n 'expected: '
cat expected9.txt
echo -n '     got: '
cat result9.txt
echo ''
cmp expected9.txt result9.txt

