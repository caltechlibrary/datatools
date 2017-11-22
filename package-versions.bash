#!/bin/bash

#
# Crawl the included Caltech Library package and display their versions in GOPATH
#
grep 'github.com/caltechlibrary/' *.go cmds/*/*.go | cut -d \" -f 2 | sort -u | while read PNAME; do 
    echo -n "$PNAME -- ";
    V=$(grep 'Version = `' "$GOPATH/src/$PNAME/$(basename $PNAME).go" | cut -d \` -f 2)
    if [ "$V" = "" ]; then
        echo "Unknown"
    else
        echo "$V"
    fi
done
