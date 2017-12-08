#!/bin/bash

if [ ! -d bin ]; then
    echo "Run make before running this script."
    exit
fi
ls -1 bin/ | while read ITEM; do
   D=$(basename "${ITEM}")
   "bin/${ITEM}" -generate-markdown-docs > "docs/${D}/index.md"
done
