#!/bin/bash

#
# Crawl the included Caltech Library package and display their versions in GOPATH
#
function crawl_code() {
	findfile -s ".go" . | while read FNAME; do
        grep "${1}" "${FNAME}" | cut -d \" -f 2 | while read PNAME; do
		V=$(grep 'Version = `' "${GOPATH}/src/${PNAME}/$(basename "${PNAME}").go" | cut -d \` -f 2)
		if [ "$V" != "" ]; then
		    echo "$PNAME -- $V"
		fi
        done
	done | sort -u
}

echo "The following package versions were used in this release"
crawl_code "github.com/caltechlibrary/"
crawl_code "github.com/rsdoiel/"
