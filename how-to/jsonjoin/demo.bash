#!/bin/bash

cat <<EOF > person.json
{ "name": "Doe, Jane", "email":"jd@example.org", "age": 42 }
EOF
echo -n "This is person.json: "
cat person.json
echo ''

cat <<EOF > profile.json
{ "name": "Doe, Jane", "bio": "World renowned geophysist.", "email": "jane.doe@example.edu" }
EOF
echo -n "This is profile.json: "
cat profile.json
echo ''

echo ' running: jsonjoin -create person.json profile.json'
jsonjoin -create person.json profile.json > result1.json
echo -n 'expected: '
cat expected1.json
echo -n '     got: '
cat result1.json
echo ''


echo ' running: cat person.json | jsonjoin profile.json'
cat person.json | jsonjoin -i -  profile.json > result2.json
echo -n 'expected: '
cat expected2.json
echo -n '     got: '
cat result2.json
echo ''

echo ' running: jsonjoin -i person.json profile.json'
jsonjoin -i person.json profile.json > result3.json
echo -n 'expected: '
cat expected3.json
echo -n '     got: '
cat result3.json
echo ''

echo ' running: jsonjoin -create -update person.json profile.json'
jsonjoin -create -update person.json profile.json > result4.json
echo -n 'expected: '
cat expected4.json
echo -n '     got: '
cat result4.json
echo ''

echo ' running: jsonjoin -create -update profile.json person.json'
jsonjoin -create -update profile.json person.json > result5.json
echo -n 'expected: '
cat expected5.json
echo -n '     got: '
cat result5.json
echo ''

echo ' running: jsonjoin -create -overwrite person.json profile.json'
jsonjoin -create -overwrite person.json profile.json > result6.json
echo -n 'expected: '
cat expected6.json
echo -n '     got: '
cat result6.json
echo ''
