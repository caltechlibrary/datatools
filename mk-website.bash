#!/bin/bash

function cleanUpHTML() {
    findfile -s ".html" . | while read P; do
        rm "$P"
    done
}

function FindNavMD() {
    DNAME="$1"
    if [ -f "${DNAME}/nav.md" ]; then
        echo "${DNAME}/nav.md"
        return
    fi
    DNAME=$(dirname "${DNAME}")
    FindNavMD "${DNAME}"
}

# Cleanup stale HTML files
cleanUpHTML

# Look through files and build new site
mkpage "nav=nav.md" "content=markdown:$(cat LICENSE)" page.tmpl > license.html
findfile -s ".md" . | while read P; do
    DNAME=$(dirname "$P")
    FNAME=$(basename "$P")
    case "$FNAME" in
        "INSTALL.md")
        HTML_NAME="${DNAME}/install.html"
        ;;
        "README.md")
        if [ ! -f "${DNAME}/index.md" ]; then
            HTML_NAME="${DNAME}/index.html"
        else
            HTML_NAME="${DNAME}/README.html"
        fi
        ;;
        *)
        HTML_NAME=$(echo "$P" | sed -E 's/.md$/.html/g')
        ;;
    esac
    if [ "${DNAME:0:4}" != "dist" ] && [ "${FNAME}" != "nav.md" ]; then
        NAV=$(FindNavMD "$DNAME")
        echo "Building $HTML_NAME from $DNAME/$FNAME and $NAV"
        mkpage "nav=$NAV" "content=${DNAME}/${FNAME}" page.tmpl >"${HTML_NAME}"
        git add "${HTML_NAME}"
    else
        echo "Skipping $P"
    fi
done


