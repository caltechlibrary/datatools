#!/bin/bash
echo "Pipe csv1.csv to csv2mdtable rendering data1.md"
cat data1.csv | csv2mdtable > data1.md
cat data1.md

echo "Use option -i and -o to read data1.csv and write data1.md"
csv2mdtable -i data1.csv -o data1.md
cat data1.md
