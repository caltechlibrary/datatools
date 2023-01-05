#
# Simple Makefile
#
PROJECT = datatools

PROGRAMS = codemeta2cff csv2json csv2mdtable csv2tab csv2xlsx csvcleaner csvcols csvfind csvjoin csvrows finddir findfile json2toml json2yaml jsoncols jsonjoin jsonmunge jsonrange mergepath range reldate reltime sql2csv string tab2csv timefmt toml2json urlparse xlsx2csv xlsx2json yaml2json

MAN_PAGES = codemeta2cff.1 sql2csv.1

PACKAGE = $(shell ls -1 *.go)

VERSION = $(shell grep '"version":' codemeta.json | cut -d\"  -f 4)

BRANCH = $(shell git branch | grep '* ' | cut -d\  -f 2)

PANDOC = $(shell which pandoc)

OS = $(shell uname)

#PREFIX = /usr/local
PREFIX = $(HOME)

ifneq ($(prefix),)
        PREFIX = $(prefix)
endif

EXT =
ifeq ($(OS), Windows)
        EXT = .exe
endif

build: version.go $(PROGRAMS) CITATION.cff about.md

version.go: .FORCE
	@echo "package $(PROJECT)" >version.go
	@echo '' >>version.go
	@echo 'const Version = "$(VERSION)"' >>version.go
	@echo '' >>version.go
	@git add version.go

about.md: codemeta.json .FORCE
	cat codemeta.json | sed -E   's/"@context"/"at__context"/g;s/"@type"/"at__type"/g;s/"@id"/"at__id"/g' >_codemeta.json
	if [ -f $(PANDOC) ]; then echo "" | $(PANDOC) --metadata title="About $(PROJECT)" --metadata-file=_codemeta.json --template=codemeta-md.tmpl >about.md; fi

CITATION.cff: codemeta.json .FORCE
	cat codemeta.json | sed -E   's/"@context"/"at__context"/g;s/"@type"/"at__type"/g;s/"@id"/"at__id"/g' >_codemeta.json
	if [ -f $(PANDOC) ]; then echo "" | $(PANDOC) --metadata title="Cite $(PROJECT)" --metadata-file=_codemeta.json --template=codemeta-cff.tmpl >CITATION.cff; fi

$(PROGRAMS): $(PACKAGE)
	@mkdir -p bin
	go build -o bin/$@$(EXT) cmd/$@/*.go

test: $(PACKAGE)
	go test
#	cd reldate && go test
#	cd timefmt && go test
	cd codemeta && go test
	bash test_cmd.bash
	
$(MAN_PAGES): .FORCE
	mkdir -p man/man1
	pandoc $@.md --from markdown --to man -s >man/man1/$@

man: $(MAN_PAGES)

website:
	bash gen-nav.bash
	bash mk-website.bash

status:
	git status

save:
	@if [ "$(msg)" != "" ]; then git commit -am "$(msg)"; else git commit -am "Quick Save"; fi
	git push origin $(BRANCH)

refresh:
	git fetch origin
	git pull origin $(BRANCH)

publish:
	bash gen-nav.bash
	bash mk-website.bash
	bash publish.bash

clean:
	@if [ -f version.go ]; then rm version.go; fi
	@if [ -d bin ]; then rm -fR bin; fi
	@if [ -d dist ]; then rm -fR dist; fi
	#@if [ -d man ]; then rm -fR man; fi

# NOTE: macOS causes problems if you copy a binary versus move it.
install: build
	@echo "Installing programs in $(PREFIX)/bin"
	@for FNAME in $(PROGRAMS); do if [ -f ./bin/$$FNAME ]; then mv -v ./bin/$$FNAME $(PREFIX)/bin/$$FNAME; fi; done
	@echo ""
	@echo "Make sure $(PREFIX)/bin is in your PATH"
	@echo "Installing man pages in $(PREFIX)/man/man1"
	@mkdir -p $(PREFIX)/man/man1
	@for FNAME in $(MAN_PAGES); do cp -v man/man1/$$FNAME $(PREFIX)/man/man1/; done
	@echo "Make sure $(PREFIX)/man is in your MANPATH"

uninstall: .FORCE
	@echo "Removing programs in $(PREFIX)/bin"
	@for FNAME in $(PROGRAMS); do if [ -f $(PREFIX)/bin/$$FNAME ]; then rm -v $(PREFIX)/bin/$$FNAME; fi; done
	@echo "Removing man pages in $(PREFIX)/man"
	@for FNAME in $(MAN_PAGES); do if [ -f $(PREFIX)/man/man1/$$FNAME ]; then rm -v $(PREFIX)/man/man1/$$FNAME; fi; done



dist/linux-amd64: $(PROGRAMS)
	@mkdir -p dist/bin
	@for FNAME in $(PROGRAMS); do env  GOOS=linux GOARCH=amd64 go build -o dist/bin/$$FNAME cmd/$$FNAME/*.go; done
	@cd dist && zip -r $(PROJECT)-v$(VERSION)-linux-amd64.zip LICENSE codemeta.json CITATION.cff *.md bin/* docs/* how-to/* man/*
	@rm -fR dist/bin


dist/macos-amd64: $(PROGRAMS)
	@mkdir -p dist/bin
	@for FNAME in $(PROGRAMS); do env GOOS=darwin GOARCH=amd64 go build -o dist/bin/$$FNAME cmd/$$FNAME/*.go; done
	@cd dist && zip -r $(PROJECT)-v$(VERSION)-macos-amd64.zip LICENSE codemeta.json CITATION.cff *.md bin/* docs/* how-to/* man/*
	@rm -fR dist/bin


dist/macos-arm64: $(PROGRAMS)
	@mkdir -p dist/bin
	@for FNAME in $(PROGRAMS); do env GOOS=darwin GOARCH=arm64 go build -o dist/bin/$$FNAME cmd/$$FNAME/*.go; done
	@cd dist && zip -r $(PROJECT)-v$(VERSION)-macos-arm64.zip LICENSE codemeta.json CITATION.cff *.md bin/* docs/* how-to/* man/*
	@rm -fR dist/bin


dist/windows-amd64: $(PROGRAMS)
	@mkdir -p dist/bin
	@for FNAME in $(PROGRAMS); do env GOOS=windows GOARCH=amd64 go build -o dist/bin/$$FNAME.exe cmd/$$FNAME/*.go; done
	@cd dist && zip -r $(PROJECT)-v$(VERSION)-windows-amd64.zip LICENSE codemeta.json CITATION.cff *.md bin/* docs/* how-to/* man/*
	@rm -fR dist/bin

dist/windows-arm64: $(PROGRAMS)
	@mkdir -p dist/bin
	@for FNAME in $(PROGRAMS); do env GOOS=windows GOARCH=arm64 go build -o dist/bin/$$FNAME.exe cmd/$$FNAME/*.go; done
	@cd dist && zip -r $(PROJECT)-v$(VERSION)-windows-arm64.zip LICENSE codemeta.json CITATION.cff *.md bin/* docs/* how-to/* man/*
	@rm -fR dist/bin


dist/raspbian-arm7: $(PROGRAMS)
	@mkdir -p dist/bin
	@for FNAME in $(PROGRAMS); do env GOOS=linux GOARCH=arm GOARM=7 go build -o dist/bin/$$FNAME cmd/$$FNAME/*.go; done
	@cd dist && zip -r $(PROJECT)-v$(VERSION)-raspberry_pi_os-arm7.zip LICENSE codemeta.json CITATION.cff *.md bin/* docs/* how-to/* man/*
	@rm -fR dist/bin

#dist/datatools_$(VERSION)_amd64.snap:
#	@mkdir -p dist/
#	snapcraft
#	@mv datatools_$(VERSION)_amd64.snap dist/
#	@chmod 664 dist/datatools_$(VERSION)_amd64.snap

distribute_docs:
	@mkdir -p dist/
	@cp -v codemeta.json dist/
	@cp -v CITATION.cff dist/
	@cp -v README.md dist/
	@cp -v LICENSE dist/
	@cp -v INSTALL.md dist/
	@cp -vR docs dist/
	@cp -vR how-to dist/
	@cp -vR man dist/

gen_batfiles: .FORCE
	@echo '@echo off' >make.bat
	@echo 'REM This is a Windows 10 Batch file for building dataset command' >>make.bat
	@echo 'REM from the command prompt.' >>make.bat
	@echo 'REM' >>make.bat
	@echo 'REM It requires: go version 1.16.6 or better and the cli for git installed' >>make.bat
	@echo 'REM' >>make.bat
	@echo 'go version' >>make.bat
	@echo 'mkdir bin' >>make.bat
	@echo 'echo "Getting ready to build the datatools in bin"' >>make.bat
	@for FNAME in $(PROGRAMS); do echo "go build -o bin/$${FNAME}.exe cmd/$${FNAME}/$${FNAME}.exe" | sed -E 's/\//\\/g' >> make.bat; done
	@echo 'echo "Checking compile should see version number of dataset"' >>make.bat
	@for FNAME in $(PROGRAMS); do echo "bin/$${FNAME}.exe -version" | sed -E 's/\//\\/g' >> make.bat; done
	@echo 'echo "If OK, you can now copy the dataset.exe to %USERPROFILE%\go\bin"' >>make.bat
	@echo 'echo ""' >>make.bat
	@echo 'echo "      copy bin/* %USERPROFILE%/AppData/go/bin"' | sed -E 's/\//\\/g' >>make.bat
	@echo '""' >>make.bat
	@echo 'echo "or someplace else in your %PATH%"' >>make.bat
	@echo '""' >>make.bat
	@git add make.bat

snap: dist/datatools_$(VERSION)_amd64.snap

release: clean build man gen_batfiles distribute_docs dist/linux-amd64 dist/macos-amd64 dist/macos-arm64 dist/windows-amd64 dist/windows-arm64 dist/raspbian-arm7


.FORCE:
