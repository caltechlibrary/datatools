#!/bin/bash

cat <<EOF > expected1.csv
Number,Value
one,1
two,2
three,3
EOF

xlsx2csv MyWorkbook.xlsx "My worksheet 1" > sheet1.csv
cmp sheet1.csv expected1.csv

echo -n "Count the sheets in MyWorkbook.xlsx: "
xlsx2csv -nl -count MyWorkbook.xlsx
echo ''

echo -n "List the sheets in MyWorkbook.xlsx: "
xlsx2csv -sheets MyWorkbook.xlsx
echo ''

xlsx2csv -N MyWorkbook.xlsx | while read SHEET_NAME; do
    CSV_NAME="${SHEET_NAME// /-}.csv"
    xlsx2csv -o "${CSV_NAME}" MyWorkbook.xlsx "${SHEET_NAME}" 
done
