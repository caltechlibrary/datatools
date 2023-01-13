#
# Makefile for running pandoc on all Markdown docs ending in .md
#
PROJECT = datatools, how to

MD_PAGES = $(shell ls -1 *.md)

HTML_PAGES = $(shell ls -1 *.md | sed -E 's/\.md/.html/g') 

build: $(MD_PAGES) $(HTML_PAGES) .FORCE

$(HTML_PAGES): $(MD_PAGES) .FORCE
	pandoc --metadata title=$(basename $@) -s --to html5 $(basename $@).md -o $(basename $@).html \
	    --template=page.tmpl
	git add $(basename $@).html

clean:
	if ls -1 *.html >/dev/null; then rm *.html; fi

.FORCE:
