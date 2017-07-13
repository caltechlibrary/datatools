#!/bin/bash

PROJECT="datatools"

function checkApp() {
	APP_NAME=$(which "$1")
	if [ "$APP_NAME" = "" ] && [ ! -f "./bin/$1" ]; then
		echo "Missing $APP_NAME"
		exit 1
	fi
}

function softwareCheck() {
	for APP_NAME in "$@"; do
		checkApp "$APP_NAME"
	done
}

function MakePage() {
	nav="$1"
	content="$2"
	html="$3"
	# Always use the latest compiled mkpage
	APP=$(which mkpage)

	echo "Rendering $html"
	$APP \
		"Nav=$nav" \
		"Content=$content" \
		page.tmpl >"$html"
	git add "$html"
}

echo "Checking necessary software is installed"
softwareCheck mkpage

echo "Generating website index.html"
MakePage nav.md README.md index.html
echo "Generating install.html"
MakePage nav.md INSTALL.md install.html
echo "Generating license.html"
MakePage nav.md "markdown:$(cat LICENSE)" license.html

# Generate docs/nav.md
cat <<EOF1 >docs/nav.md

+ [Home](/)
+ [up](../)
EOF1
INDEX_MENU=$(cat docs/nav.md)

echo "+ [Documentation](./)" >> docs/nav.md
echo "+ [How To ...](../how-to/)" >> docs/nav.md
git add docs/nav.md

# Generate docs/index.md
cat <<EOF2 >docs/index.md

# $PROJECT command help

EOF2

find cmds -maxdepth 1 -type d | while read DNAME; do
	FNAME="$(basename "$DNAME")"
	echo "+ [$FNAME](${FNAME}.html)"
done >>docs/index.md
git add docs/index.md

MakePage "docs/nav.md" "docs/index.md" "docs/index.html"


# Generate nav for How To section
echo "$INDEX_MENU" > how-to/nav.md
echo "+ [Documentation](../docs/)" >> how-to/nav.md
echo "+ [How To ...](./)" >> how-to/nav.md
git add how-to/nav.md

# Generate index page for How To section
cat <<EOF3 >how-to/index.md

# How To ...

EOF3

find ./how-to -type f | grep -E "\.md$" | while read FNAME; do
    TITLE="$(titleline -i "$FNAME")"
    FNAME="$(basename "$FNAME" ".md")"
    if [ "$FNAME" != "nav" ] && [ "$FNAME" != "index" ]; then
        echo "+ [${TITLE}](${FNAME}.html)"
    fi
done >>how-to/index.md
git add how-to/index.md

MakePage "markdown:${INDEX_MENU}" "how-to/index.md" "how-to/index.html"
git add how-to/index.html

# Generate the individual How To documents
find ./how-to -type f | grep -E "\.md$" | while read FNAME; do
    FNAME="$(basename "$FNAME" ".md")"
    MakePage "how-to/nav.md" "how-to/${FNAME}.md"  "how-to/${FNAME}.html"
done


# Generate the individual command docuumentation pages
for FNAME in csvcols csvfind csvjoin csv2json csv2mdtable csv2xlsx jsoncols jsonrange xlsx2json xlsx2csv vcard2json jsonmunge; do
	echo "Generating docs/$FNAME.html"
	MakePage docs/nav.md "docs/$FNAME.md" "docs/$FNAME.html"
done
