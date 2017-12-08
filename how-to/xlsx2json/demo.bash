#!/bin/bash

echo '[["Number","Value"],["one","1"],["two","2"],["three","3"]]' > expected1.json
xlsx2json -nl MyWorkbook.xlsx "My worksheet 1" > sheet1.json
echo -n "Check sheet1.json against expected1.json: " 
R=$(cmp expected1.json sheet1.json)
if [ "$R" != "" ]; then
    echo "$R"
else
    echo "OK"
fi
echo ''

echo -n "Count the sheets in MyWorkbook.xlsx: "
xlsx2json -count MyWorkbook.xlsx
echo ''

echo -n "List the sheets in MyWorkbook.xlsx: "
xlsx2json -sheets MyWorkbook.xlsx
echo ''

xlsx2json -N MyWorkbook.xlsx | while read SHEET_NAME; do
    JSON_NAME="${SHEET_NAME// /-}.json"
    xlsx2json -nl -o "${JSON_NAME}" MyWorkbook.xlsx "$SHEET_NAME"
    echo -n "Check ${JSON_NAME} against expected1.json: "
    R=$(cmp expected1.json "${JSON_NAME}")
    if [ "$R" != "" ]; then
        echo "$R"
    else
        echo "OK"
    fi
done
