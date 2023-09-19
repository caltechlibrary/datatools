#
# Makefile for running pandoc on all Markdown docs ending in .md
#
PROJECT = datatools

MD_PAGES = $(shell ls -1 *.md)

HTML_PAGES = $(shell ls -1 *.md | sed -E 's/\.md/.html/g') 

build: about.md $(MD_PAGES) $(HTML_PAGES) how-to pagefind

about.md: codemeta.json
	echo "" | pandoc -metadata title="About $(PROJECT)" -metadata-file=codemeta.json -s --from markdown --to markdown --template=page.tmpl -o about.md
	git add about.md

$(HTML_PAGES): $(MD_PAGES) .FORCE
	pandoc --metadata title=$(basename $@) -s --to html5 $(basename $@).md -o $(basename $@).html \
	    --template=page.tmpl
	-if [ $@ = "README.html" ]; then mv README.html index.html; git add index.html; fi
	-if [ -f $(basename $@).html ]; then git add $(basename $@).html; fi

how-to:
	cd how-to && make -f website.mak

pagefind: .FORCE
	pagefind --verbose --exclude-selectors="nav,header,footer" --bundle-dir ./pagefind --source .
	git add pagefind

clean:
	if ls -1 *.html >/dev/null; then rm *.html; fi
	if ls -1 how-to/*.html>/dev/null; then rm how-to/*.html; fi

.FORCE:
