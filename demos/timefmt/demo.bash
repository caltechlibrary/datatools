#!/bin/basj

echo ' running: timefmt -if "2006-01-02" -of "01/02/2006" "2017-12-02"'
echo 'expected: 12/02/2017'
echo "     got: $(timefmt -if "2006-01-02" -of "01/02/2006" "2017-12-02")"
echo ''

echo ' running: timefmt -input-format mysql -output-format RFC822  "2017-12-02 08:08:08"'
echo 'expected: 02 Dec 17 08:08 UTC'
echo "     got: $(timefmt -input-format mysql -output-format RFC822  "2017-12-02 08:08:08")"
echo ''
