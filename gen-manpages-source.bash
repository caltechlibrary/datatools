#!/bin/bash
#

PROGRAMS=$(ls -1 cmd/)
make build
for PROG in $PROGRAMS; do
	echo "Generating manpage source for $PROG"
	./bin/$PROG -help >"${PROG}.1.md"
	git add "${PROG}.1.md"
done
make man
git add man
