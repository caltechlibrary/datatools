#!/bin/bash

echo "Creating sheet 'My worksheet 1' in MyWorkbook.xlsx using options"
csv2xlsx -i data.csv MyWorkbook.xlsx 'My worksheet 1'
echo "Creating sheet 'My worksheet 2' in MyWorkbook.xlsx using pipes"
cat data.csv | csv2xlsx MyWorkbook.xlsx 'My worksheet 2'

