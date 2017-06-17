#!/bin/bash

if [ "$1" = "" ] && [ "$2" = "" ]; then
	echo "USAGE: $(basename "$0") CSV_FILENAME CSV_COL_NO"
	exit 1
fi
if [ "$1" = "" ]; then
	echo "Missing CSV FILE name"
	exit 1
fi

if [ "$2" = "" ]; then
	echo "Missing column number to match on"
	exit 1
fi

CSV_FILE="$1"
CSV_COL_NO="$2"

csvcols -i "$CSV_FILE" -col "$CSV_COL_NO" | sort -u | while read CELL; do
	if [ "$CELL" != "" ]; then
		csvfind -i "$CSV_FILE" -trim-spaces -col "$CSV_COL_NO"  "${CELL}"
	fi
done
