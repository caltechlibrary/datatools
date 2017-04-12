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
	if [ -f ./bin/mkpage ]; then
		APP="./bin/mkpage"
	fi

	echo "Rendering $html"
	$APP \
		"title=text:$PROJECT -- a small collection of file and shell utilities" \
		"nav=$nav" \
		"content=$content" \
		"sitebuilt=text:Updated $(date)" \
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

# Generate the individual command docuumentation pages
for FNAME in csvcols csvfind csvjoin csv2json csv2mdtable csv2xlsx text2terms jsoncols jsonrange xlsx2json xlsx2csv; do
	echo "Generating $FNAME.html"
	MakePage nav.md "$FNAME.md" "$FNAME.html"
done
