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
                license="text:${RELPATH}license.html" \
                docs="text:${RELPATH}docs/" \
                install="text:${RELPATH}install.html" \
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
        license="text:license.html" \
        docs="text:docs/" \
        install="text:install.html" \
        howto="text:how-to/" \
        gitrepo="text:${GIT_REPO}" \
        >"nav.md"

# generate docs nav
mkpage nav.tmpl relroot="text:" \
        readme="text:../index.html" \
        license="text:../license.html" \
        docs="text:./" \
        install="text:../install.html" \
        howto="text:../how-to/" \
        gitrepo="text:${GIT_REPO}" \
        >"docs/nav.md"

# generate how-to nav
mkpage nav.tmpl relroot="text:" \
        readme="text:../index.html" \
        license="text:../license.html" \
        docs="text:../docs/" \
        install="text:../install.html" \
        howto="text:./" \
        gitrepo="text:${GIT_REPO}" \
        >"how-to/nav.md"

# walk docs/ and generate needed nav.md
write_nav "docs/"

# walk how-to/ and generate needed nav.md
write_nav "how-to/"
