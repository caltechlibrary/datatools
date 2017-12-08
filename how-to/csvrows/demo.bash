#!/bin/bash

echo "Turn parameters into rows:"
csvrows -o 4rows.csv "First,Second,Third" "one,two,three"
csvrows "ein,zwei,drei" "1,2,3" >> 4rows.csv
cat 4rows.csv

#echo "Turn delimiter input parameters into rows"
#csvrows -d "|" "First,Second,Third|one,two,three" > 4rows.csv
#csvrows -delimiter "|" "ein,zwei,drei|1,2,3" >> 4rows.csv
#cat 4rows.csv

echo "Extract rows 1,3:"
cat 4rows.csv | csvrows -row 1,3 > result1.csv
csvrows -i 4rows.csv -row 1,3 > result2.csv
