#!/bin/bash

echo ' running: range 1 5'
echo 'expected: 1 2 3 4 5'
echo "     got: $(range 1 5)"
echo ''

echo ' running: range -- -2 6'
echo 'expected: -2 -1 0 1 2 3 4 5 6'
echo "     got: $(range -- -2 6)"
echo ''

echo ' running: range -increment=2 2 10'
echo 'expected: 2 4 6 8 10'
echo "     got: $(range -increment=2 2 10)"
echo ''

echo ' running: range 10 1'
echo 'expected: 10 9 8 7 6 5 4 3 2 1'
echo "     got: $(range 10 1)"
echo ''

echo ' running: range -r 0 10'
echo 'expected random number betwee (inclusive) 0 and 10'
I=$(range -random 0 10)
if [[ "$I" -lt "0" || "$I" -gt "10" ]]; then
    echo "     got: $I (error, out of range)"
else
    echo "     got: $I"
fi
echo ''
