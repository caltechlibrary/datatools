#!/bin/bash

echo ' running: reldate -from=2014-08-01 3 days'
echo 'expected: 2014-08-04'
echo "     got: $(reldate -from=2014-08-01 3 days)"
echo ''

echo ' running: reldate --from=2014-08-03 3 days'
echo 'expected: 2014-08-06'
echo "     got: $(reldate --from=2014-08-03 3 days)"
echo ''

echo ' running: reldate --from=2014-08-03 -- -3 days'
echo 'expected: 2014-07-31'
echo "     got: $(reldate --from=2014-08-03 -- -3 days)"
echo ''

echo ' running: reldate --from=2015-02-10 Monday'
echo 'expected: 2015-02-09'
echo "     got: $(reldate --from=2015-02-10 Monday)"
echo ''
