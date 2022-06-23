#
# Simple Makefile
#
PROJECT = datatools

PROGRAMS = $(shell ls -1 cmd/)

PACKAGE = $(shell ls -1 *.go)

VERSION = $(shell grep '"version":' codemeta.json | cut -d\"  -f 4)

BRANCH = $(shell git branch | grep '* ' | cut -d\  -f 2)

CODEMETA2CFF = $(shell which codemeta2cff)

OS = $(shell uname)

#PREFIX = /usr/local/bin
PREFIX = $(HOME)

ifneq ($(prefix),)
        PREFIX = $(prefix)
endif

EXT = 
ifeq ($(OS), Windows)
        EXT = .exe
endif

build: version.go $(PROGRAMS)

version.go: .FORCE
	@echo "package $(PROJECT)" >version.go
	@echo '' >>version.go
	@echo 'const Version = "$(VERSION)"' >>version.go
	@echo '' >>version.go
	@git add version.go
	@if [ -f bin/codemeta ]; then ./bin/codemeta; fi
	$(CODEMETA2CFF)

$(PROGRAMS): $(PACKAGE)
	@mkdir -p bin
	go build -o bin/$@$(EXT) cmd/$@/*.go

test: $(PACKAGE)
	go test
#	cd reldate && go test
#	cd timefmt && go test
	cd codemeta && go test
	bash test_cmd.bash

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
	@if [ -d man ]; then rm -fR man; fi

# NOTE: macOS causes problems if you copy a binary versus move it.
install: build
	@echo "Installing programs in $(PREFIX)/bin"
	@for FNAME in $(PROGRAMS); do if [ -f ./bin/$$FNAME ]; then mv -v ./bin/$$FNAME $(PREFIX)/bin/$$FNAME; fi; done
	@echo ""
	@echo "Make sure $(PREFIX)/bin is in your PATH"

uninstall: .FORCE
	@echo "Removing programs in $(PREFIX)/bin"
	@for FNAME in $(PROGRAMS); do if [ -f $(PREFIX)/bin/$$FNAME ]; then rm -v $(PREFIX)/bin/$$FNAME; fi; done


dist/linux-amd64: $(PROGRAMS)
	@mkdir -p dist/bin
	@for FNAME in $(PROGRAMS); do env  GOOS=linux GOARCH=amd64 go build -o dist/bin/$$FNAME cmd/$$FNAME/*.go; done
	@cd dist && zip -r $(PROJECT)-v$(VERSION)-linux-amd64.zip LICENSE codemeta.json CITATION.cff *.md bin/* docs/* how-to/* demos/*
	@rm -fR dist/bin


dist/macos-amd64: $(PROGRAMS)
	@mkdir -p dist/bin
	@for FNAME in $(PROGRAMS); do env GOOS=darwin GOARCH=amd64 go build -o dist/bin/$$FNAME cmd/$$FNAME/*.go; done
	@cd dist && zip -r $(PROJECT)-v$(VERSION)-macos-amd64.zip LICENSE codemeta.json CITATION.cff *.md bin/* docs/* how-to/* demos/*
	@rm -fR dist/bin
	

dist/macos-arm64: $(PROGRAMS)
	@mkdir -p dist/bin
	@for FNAME in $(PROGRAMS); do env GOOS=darwin GOARCH=arm64 go build -o dist/bin/$$FNAME cmd/$$FNAME/*.go; done
	@cd dist && zip -r $(PROJECT)-v$(VERSION)-macos-arm64.zip LICENSE codemeta.json CITATION.cff *.md bin/* docs/* how-to/* demos/*
	@rm -fR dist/bin
	

dist/windows-amd64: $(PROGRAMS)
	@mkdir -p dist/bin
	@for FNAME in $(PROGRAMS); do env GOOS=windows GOARCH=amd64 go build -o dist/bin/$$FNAME.exe cmd/$$FNAME/*.go; done
	@cd dist && zip -r $(PROJECT)-v$(VERSION)-windows-amd64.zip LICENSE codemeta.json CITATION.cff *.md bin/* docs/* how-to/* demos/*
	@rm -fR dist/bin


dist/raspbian-arm7: $(PROGRAMS)
	@mkdir -p dist/bin
	@for FNAME in $(PROGRAMS); do env GOOS=linux GOARCH=arm GOARM=7 go build -o dist/bin/$$FNAME cmd/$$FNAME/*.go; done
	@cd dist && zip -r $(PROJECT)-v$(VERSION)-raspberry_pi_os-arm7.zip LICENSE codemeta.json CITATION.cff *.md bin/* docs/* how-to/* demos/*
	@rm -fR dist/bin

dist/datatools_$(VERSION)_amd64.snap:
	@mkdir -p dist/
	snapcraft
	@mv datatools_$(VERSION)_amd64.snap dist/
	@chmod 664 dist/datatools_$(VERSION)_amd64.snap

distribute_docs:
	@mkdir -p dist/
	@cp -v codemeta.json dist/
	@cp -v CITATION.cff dist/
	@cp -v README.md dist/
	@cp -v LICENSE dist/
	@cp -v INSTALL.md dist/
	@cp -vR docs dist/
	@cp -vR how-to dist/
	
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

release: build gen_batfiles distribute_docs dist/linux-amd64 dist/macos-amd64 dist/macos-arm64 dist/windows-amd64 dist/raspbian-arm7 snap


.FORCE:
