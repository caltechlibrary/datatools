#!/bin/bash
#

PROGRAMS="codemeta2cff sql2csv csv2json csv2mdtable csv2tab csv2xlsx csvcleaner"

make build
for PROG in $PROGRAMS; do
	echo "Generating manpage source for $PROG"
	./bin/$PROG -help >"${PROG}.1.md"
done
make man
