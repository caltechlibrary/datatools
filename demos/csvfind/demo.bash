#!/bin/bash

echo "Look for 'The Red Book of Westmarch' in column 2:"
csvfind -i books.csv -o result1.csv \
    -col=2 "The Red Book of Westmarch"
echo ""
cat result1.csv

echo ""
echo "Look for 'The Red Book of Westmarch' in column 2 using fuzzy matching:"
csvfind -i books.csv -o result2.csv \
    -col=2 -levenshtein \
   -insert-cost=1 -delete-cost=1 -substitute-cost=3 \
   -max-edit-distance=50 -append-edit-distance \
   "The Red Book of Westmarch"
echo ""
cat result2.csv

echo ""
echo "Look for 'Red Book' in column 2:"
csvfind -i books.csv -o result3.csv \
    -col=2 -contains "Red Book"
echo ""
cat result3.csv
