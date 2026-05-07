
# generated with CMTools 1.3.5 b09862d

#
# Makefile for running pandoc on all Markdown docs ending in .md
#
PROJECT = datatools

PANDOC = $(shell which pandoc)

MD_PAGES = $(shell ls -1 *.md)

HTML_PAGES = $(shell ls -1 *.md | sed -E 's/\.md/\.html/g')

build: $(HTML_PAGES) $(MD_PAGES) # pagefind

$(HTML_PAGES): $(MD_PAGES) .FORCE
	if [ -f $(PANDOC) ]; then $(PANDOC) --metadata title=$(basename $@) -s --to html5 $(basename $@).md -o $(basename $@).html \
		--lua-filter=links-to-html.lua \
		--lua-filter=add-col-scope.lua \
	    --template=page.tmpl; fi
	@if [ $@ = "README.html" ]; then mv README.html index.html; fi

pagefind: .FORCE
	# NOTE: I am not including most of the archive in PageFind index since it doesn't make sense in this case.
	pagefind --verbose --glob="{*.html,docs/*.html}" --force-language en-US --exclude-selectors="nav,header,footer" --output-path ./pagefind --site .
	git add pagefind

clean:
	@rm *.html

.FORCE:
