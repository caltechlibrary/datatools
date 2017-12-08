#!/bin/bash

GIT_REPO="https://github.com/caltechlibrary/datatools"

function write_nav() {
DNAME="$1"
finddir -depth 2 ${DNAME} | while read D; do
    if [[ "$D" != "." ]]; then
        echo "Writing ${DNAME}${D}/nav.md"
        RELPATH=$(reldocpath ${DNAME}"${D}" .)
        mkpage nav.tmpl relroot="text:${RELPATH}" \
        readme="text:${RELPATH}index.html" \
        docs="text:${RELPATH}docs/" \
        install="text:${RELPATH}INSTALL.html" \
        howto="text:${RELPATH}how-to/" \
        gitrepo="text:${GIT_REPO}" \
        >"${DNAME}${D}/nav.md"
    fi
done
}

# root nav
echo "Writing root nav.md"
mkpage nav.tmpl relroot="text:" \
        readme="text:index.html" \
        docs="text:docs/" \
        install="text:INSTALL.html" \
        howto="text:how-to/" \
        gitrepo="text:${GIT_REPO}" \
        >"nav.md"

# walk docs/ and generate needed nav.md
write_nav "docs/"

# walk how-to/ and generate needed nav.md
write_nav "how-to/"
