#
# Simple Makefile
#
PROJECT = datatools

VERSION = $(shell grep -m1 'Version = ' $(PROJECT).go | cut -d\`  -f 2)

BRANCH = $(shell git branch | grep '* ' | cut -d\  -f 2)

OS = $(shell uname)

EXT = 
ifeq ($(OS), Windows)
        EXT = .exe
endif


build$(EXT): bin/csvcols$(EXT) bin/csvrows$(EXT) bin/csvfind$(EXT) bin/csvjoin$(EXT) \
	bin/jsoncols$(EXT) bin/jsonrange$(EXT) bin/xlsx2json$(EXT) bin/xlsx2csv$(EXT) \
	bin/csv2mdtable$(EXT) bin/csv2xlsx$(EXT) bin/csv2json$(EXT) bin/jsonjoin$(EXT) \
	bin/jsonmunge$(EXT) bin/findfile$(EXT) bin/finddir$(EXT) bin/mergepath$(EXT) \
	bin/reldate$(EXT) bin/range$(EXT) bin/timefmt$(EXT) bin/urlparse$(EXT) \
	bin/csvcleaner$(EXT) bin/string$(EXT) 


bin/csvcols$(EXT): datatools.go cmds/csvcols/csvcols.go
	go build -o bin/csvcols$(EXT) cmds/csvcols/csvcols.go 

bin/csvrows$(EXT): datatools.go cmds/csvrows/csvrows.go
	go build -o bin/csvrows$(EXT) cmds/csvrows/csvrows.go 

bin/csvjoin$(EXT): datatools.go cmds/csvjoin/csvjoin.go
	go build -o bin/csvjoin$(EXT) cmds/csvjoin/csvjoin.go 

bin/jsoncols$(EXT): datatools.go cmds/jsoncols/jsoncols.go
	go build -o bin/jsoncols$(EXT) cmds/jsoncols/jsoncols.go 

bin/jsonrange$(EXT): datatools.go cmds/jsonrange/jsonrange.go
	go build -o bin/jsonrange$(EXT) cmds/jsonrange/jsonrange.go 

bin/xlsx2json$(EXT): datatools.go cmds/xlsx2json/xlsx2json.go
	go build -o bin/xlsx2json$(EXT) cmds/xlsx2json/xlsx2json.go

bin/xlsx2csv$(EXT): datatools.go cmds/xlsx2csv/xlsx2csv.go
	go build -o bin/xlsx2csv$(EXT) cmds/xlsx2csv/xlsx2csv.go

bin/csv2mdtable$(EXT): datatools.go cmds/csv2mdtable/csv2mdtable.go
	go build -o bin/csv2mdtable$(EXT) cmds/csv2mdtable/csv2mdtable.go

bin/csv2xlsx$(EXT): datatools.go cmds/csv2xlsx/csv2xlsx.go
	go build -o bin/csv2xlsx$(EXT) cmds/csv2xlsx/csv2xlsx.go

bin/csv2json$(EXT): datatools.go cmds/csv2json/csv2json.go
	go build -o bin/csv2json$(EXT) cmds/csv2json/csv2json.go

bin/csvfind$(EXT): datatools.go cmds/csvfind/csvfind.go
	go build -o bin/csvfind$(EXT) cmds/csvfind/csvfind.go

bin/jsonmunge$(EXT): datatools.go cmds/jsonmunge/jsonmunge.go
	go build -o bin/jsonmunge$(EXT) cmds/jsonmunge/jsonmunge.go

bin/jsonjoin$(EXT): datatools.go cmds/jsonjoin/jsonjoin.go
	go build -o bin/jsonjoin$(EXT) cmds/jsonjoin/jsonjoin.go

bin/findfile$(EXT): datatools.go cmds/findfile/findfile.go
	go build -o bin/findfile$(EXT) cmds/findfile/findfile.go 

bin/finddir$(EXT): datatools.go cmds/finddir/finddir.go
	go build -o bin/finddir$(EXT) cmds/finddir/finddir.go 

bin/mergepath$(EXT): datatools.go cmds/mergepath/mergepath.go
	go build -o bin/mergepath$(EXT) cmds/mergepath/mergepath.go 

bin/reldate$(EXT): datatools.go cmds/reldate/reldate.go
	go build -o bin/reldate$(EXT) cmds/reldate/reldate.go 

bin/range$(EXT): datatools.go cmds/range/range.go
	go build -o bin/range$(EXT) cmds/range/range.go 

bin/timefmt$(EXT): datatools.go cmds/timefmt/timefmt.go
	go build -o bin/timefmt$(EXT) cmds/timefmt/timefmt.go 

bin/urlparse$(EXT): datatools.go cmds/urlparse/urlparse.go
	go build -o bin/urlparse$(EXT) cmds/urlparse/urlparse.go 

bin/csvcleaner$(EXT): datatools.go cmds/csvcleaner/csvcleaner.go
	go build -o bin/csvcleaner$(EXT) cmds/csvcleaner/csvcleaner.go

bin/string$(EXT): datatools.go cmds/string/string.go
	go build -o bin/string$(EXT) cmds/string/string.go

test: build
	go test
	bash test_cmds.bash

website:
	bash gen-nav.bash
	bash mk-website.bash

status:
	git status

save:
	if [ "$(msg)" != "" ]; then git commit -am "$(msg)"; else git commit -am "Quick Save"; fi
	git push origin $(BRANCH)

refresh:
	git fetch origin
	git pull origin $(BRANCH)

publish:
	bash gen-nav.bash
	bash mk-website.bash
	bash publish.bash

clean: 
	if [ -d bin ]; then rm -fR bin; fi
	if [ -d dist ]; then rm -fR dist; fi

install:
	env GOBIN=$(GOPATH)/bin go install cmds/csvcols/csvcols.go
	env GOBIN=$(GOPATH)/bin go install cmds/csvrows/csvrows.go
	env GOBIN=$(GOPATH)/bin go install cmds/csvfind/csvfind.go
	env GOBIN=$(GOPATH)/bin go install cmds/csvjoin/csvjoin.go
	env GOBIN=$(GOPATH)/bin go install cmds/csv2mdtable/csv2mdtable.go
	env GOBIN=$(GOPATH)/bin go install cmds/csv2xlsx/csv2xlsx.go
	env GOBIN=$(GOPATH)/bin go install cmds/csv2json/csv2json.go
	env GOBIN=$(GOPATH)/bin go install cmds/findfile/findfile.go
	env GOBIN=$(GOPATH)/bin go install cmds/finddir/finddir.go
	env GOBIN=$(GOPATH)/bin go install cmds/jsoncols/jsoncols.go
	env GOBIN=$(GOPATH)/bin go install cmds/jsonrange/jsonrange.go
	env GOBIN=$(GOPATH)/bin go install cmds/jsonmunge/jsonmunge.go
	env GOBIN=$(GOPATH)/bin go install cmds/jsonjoin/jsonjoin.go
	env GOBIN=$(GOPATH)/bin go install cmds/mergepath/mergepath.go
	env GOBIN=$(GOPATH)/bin go install cmds/reldate/reldate.go
	env GOBIN=$(GOPATH)/bin go install cmds/range/range.go
	env GOBIN=$(GOPATH)/bin go install cmds/timefmt/timefmt.go
	env GOBIN=$(GOPATH)/bin go install cmds/urlparse/urlparse.go
	env GOBIN=$(GOPATH)/bin go install cmds/xlsx2json/xlsx2json.go
	env GOBIN=$(GOPATH)/bin go install cmds/xlsx2csv/xlsx2csv.go
	env GOBIN=$(GOPATH)/bin go install cmds/csvcleaner/csvcleaner.go
	env GOBIN=$(GOPATH)/bin go install cmds/string/string.go

dist/linux-amd64:
	mkdir -p dist/bin
	env  GOOS=linux GOARCH=amd64 go build -o dist/bin/csvcols cmds/csvcols/csvcols.go
	env  GOOS=linux GOARCH=amd64 go build -o dist/bin/csvrows cmds/csvrows/csvrows.go
	env  GOOS=linux GOARCH=amd64 go build -o dist/bin/csvfind cmds/csvfind/csvfind.go
	env  GOOS=linux GOARCH=amd64 go build -o dist/bin/csvjoin cmds/csvjoin/csvjoin.go
	env  GOOS=linux GOARCH=amd64 go build -o dist/bin/jsoncols cmds/jsoncols/jsoncols.go
	env  GOOS=linux GOARCH=amd64 go build -o dist/bin/jsonrange cmds/jsonrange/jsonrange.go
	env  GOOS=linux GOARCH=amd64 go build -o dist/bin/xlsx2json cmds/xlsx2json/xlsx2json.go
	env  GOOS=linux GOARCH=amd64 go build -o dist/bin/xlsx2csv cmds/xlsx2csv/xlsx2csv.go
	env  GOOS=linux GOARCH=amd64 go build -o dist/bin/csv2mdtable cmds/csv2mdtable/csv2mdtable.go
	env  GOOS=linux GOARCH=amd64 go build -o dist/bin/csv2xlsx cmds/csv2xlsx/csv2xlsx.go
	env  GOOS=linux GOARCH=amd64 go build -o dist/bin/csv2json cmds/csv2json/csv2json.go
	env  GOOS=linux GOARCH=amd64 go build -o dist/bin/jsonmunge cmds/jsonmunge/jsonmunge.go
	env  GOOS=linux GOARCH=amd64 go build -o dist/bin/jsonjoin cmds/jsonjoin/jsonjoin.go
	env  GOOS=linux GOARCH=amd64 go build -o dist/bin/findfile cmds/findfile/findfile.go
	env  GOOS=linux GOARCH=amd64 go build -o dist/bin/finddir cmds/finddir/finddir.go
	env  GOOS=linux GOARCH=amd64 go build -o dist/bin/mergepath cmds/mergepath/mergepath.go
	env  GOOS=linux GOARCH=amd64 go build -o dist/bin/reldate cmds/reldate/reldate.go
	env  GOOS=linux GOARCH=amd64 go build -o dist/bin/range cmds/range/range.go
	env  GOOS=linux GOARCH=amd64 go build -o dist/bin/timefmt cmds/timefmt/timefmt.go
	env  GOOS=linux GOARCH=amd64 go build -o dist/bin/urlparse cmds/urlparse/urlparse.go
	env  GOOS=linux GOARCH=amd64 go build -o dist/bin/csvcleaner cmds/csvcleaner/csvcleaner.go
	env  GOOS=linux GOARCH=amd64 go build -o dist/bin/string cmds/string/string.go
	cd dist && zip -r $(PROJECT)-$(VERSION)-linux-amd64.zip README.md LICENSE INSTALL.md bin/* docs/* how-to/* demos/*
	rm -fR dist/bin


dist/macosx-amd64:
	mkdir -p dist/bin
	env  GOOS=darwin GOARCH=amd64 go build -o dist/bin/csvcols cmds/csvcols/csvcols.go
	env  GOOS=darwin GOARCH=amd64 go build -o dist/bin/csvrows cmds/csvrows/csvrows.go
	env  GOOS=darwin GOARCH=amd64 go build -o dist/bin/csvfind cmds/csvfind/csvfind.go
	env  GOOS=darwin GOARCH=amd64 go build -o dist/bin/csvjoin cmds/csvjoin/csvjoin.go
	env  GOOS=darwin GOARCH=amd64 go build -o dist/bin/jsoncols cmds/jsoncols/jsoncols.go
	env  GOOS=darwin GOARCH=amd64 go build -o dist/bin/jsonrange cmds/jsonrange/jsonrange.go
	env  GOOS=darwin GOARCH=amd64 go build -o dist/bin/xlsx2json cmds/xlsx2json/xlsx2json.go
	env  GOOS=darwin GOARCH=amd64 go build -o dist/bin/xlsx2csv cmds/xlsx2csv/xlsx2csv.go
	env  GOOS=darwin GOARCH=amd64 go build -o dist/bin/csv2mdtable cmds/csv2mdtable/csv2mdtable.go
	env  GOOS=darwin GOARCH=amd64 go build -o dist/bin/csv2xlsx cmds/csv2xlsx/csv2xlsx.go
	env  GOOS=darwin GOARCH=amd64 go build -o dist/bin/csv2json cmds/csv2json/csv2json.go
	env  GOOS=darwin GOARCH=amd64 go build -o dist/bin/jsonmunge cmds/jsonmunge/jsonmunge.go
	env  GOOS=darwin GOARCH=amd64 go build -o dist/bin/jsonjoin cmds/jsonjoin/jsonjoin.go
	env  GOOS=darwin GOARCH=amd64 go build -o dist/bin/findfile cmds/findfile/findfile.go
	env  GOOS=darwin GOARCH=amd64 go build -o dist/bin/finddir cmds/finddir/finddir.go
	env  GOOS=darwin GOARCH=amd64 go build -o dist/bin/mergepath cmds/mergepath/mergepath.go
	env  GOOS=darwin GOARCH=amd64 go build -o dist/bin/reldate cmds/reldate/reldate.go
	env  GOOS=darwin GOARCH=amd64 go build -o dist/bin/range cmds/range/range.go
	env  GOOS=darwin GOARCH=amd64 go build -o dist/bin/timefmt cmds/timefmt/timefmt.go
	env  GOOS=darwin GOARCH=amd64 go build -o dist/bin/urlparse cmds/urlparse/urlparse.go
	env  GOOS=darwin GOARCH=amd64 go build -o dist/bin/csvcleaner cmds/csvcleaner/csvcleaner.go
	env  GOOS=darwin GOARCH=amd64 go build -o dist/bin/string cmds/string/string.go
	cd dist && zip -r $(PROJECT)-$(VERSION)-macosx-amd64.zip README.md LICENSE INSTALL.md bin/* docs/* how-to/* demos/*
	rm -fR dist/bin
	


dist/windows-amd64:
	mkdir -p dist/bin
	env  GOOS=windows GOARCH=amd64 go build -o dist/bin/csvcols.exe cmds/csvcols/csvcols.go
	env  GOOS=windows GOARCH=amd64 go build -o dist/bin/csvrows.exe cmds/csvrows/csvrows.go
	env  GOOS=windows GOARCH=amd64 go build -o dist/bin/csvfind.exe cmds/csvfind/csvfind.go
	env  GOOS=windows GOARCH=amd64 go build -o dist/bin/csvjoin.exe cmds/csvjoin/csvjoin.go
	env  GOOS=windows GOARCH=amd64 go build -o dist/bin/jsoncols.exe cmds/jsoncols/jsoncols.go
	env  GOOS=windows GOARCH=amd64 go build -o dist/bin/jsonrange.exe cmds/jsonrange/jsonrange.go
	env  GOOS=windows GOARCH=amd64 go build -o dist/bin/xlsx2json.exe cmds/xlsx2json/xlsx2json.go
	env  GOOS=windows GOARCH=amd64 go build -o dist/bin/xlsx2csv.exe cmds/xlsx2csv/xlsx2csv.go
	env  GOOS=windows GOARCH=amd64 go build -o dist/bin/csv2mdtable.exe cmds/csv2mdtable/csv2mdtable.go
	env  GOOS=windows GOARCH=amd64 go build -o dist/bin/csv2xlsx.exe cmds/csv2xlsx/csv2xlsx.go
	env  GOOS=windows GOARCH=amd64 go build -o dist/bin/csv2json.exe cmds/csv2json/csv2json.go
	env  GOOS=windows GOARCH=amd64 go build -o dist/bin/jsonmunge.exe cmds/jsonmunge/jsonmunge.go
	env  GOOS=windows GOARCH=amd64 go build -o dist/bin/jsonjoin.exe cmds/jsonjoin/jsonjoin.go
	env  GOOS=windows GOARCH=amd64 go build -o dist/bin/findfile.exe cmds/findfile/findfile.go
	env  GOOS=windows GOARCH=amd64 go build -o dist/bin/finddir.exe cmds/finddir/finddir.go
	env  GOOS=windows GOARCH=amd64 go build -o dist/bin/mergepath.exe cmds/mergepath/mergepath.go
	env  GOOS=windows GOARCH=amd64 go build -o dist/bin/reldate.exe cmds/reldate/reldate.go
	env  GOOS=windows GOARCH=amd64 go build -o dist/bin/range.exe cmds/range/range.go
	env  GOOS=windows GOARCH=amd64 go build -o dist/bin/timefmt.exe cmds/timefmt/timefmt.go
	env  GOOS=windows GOARCH=amd64 go build -o dist/bin/urlparse.exe cmds/urlparse/urlparse.go
	env  GOOS=windows GOARCH=amd64 go build -o dist/bin/csvcleaner.exe cmds/csvcleaner/csvcleaner.go
	env  GOOS=windows GOARCH=amd64 go build -o dist/bin/string.exe cmds/string/string.go
	cd dist && zip -r $(PROJECT)-$(VERSION)-windows-amd64.zip README.md LICENSE INSTALL.md bin/* docs/* how-to/* demos/*
	rm -fR dist/bin




dist/raspbian-arm7:
	mkdir -p dist/bin
	env  GOOS=linux GOARCH=arm GOARM=7 go build -o dist/bin/csvcols cmds/csvcols/csvcols.go
	env  GOOS=linux GOARCH=arm GOARM=7 go build -o dist/bin/csvrows cmds/csvrows/csvrows.go
	env  GOOS=linux GOARCH=arm GOARM=7 go build -o dist/bin/csvfind cmds/csvfind/csvfind.go
	env  GOOS=linux GOARCH=arm GOARM=7 go build -o dist/bin/csvjoin cmds/csvjoin/csvjoin.go
	env  GOOS=linux GOARCH=arm GOARM=7 go build -o dist/bin/jsoncols cmds/jsoncols/jsoncols.go
	env  GOOS=linux GOARCH=arm GOARM=7 go build -o dist/bin/jsonrange cmds/jsonrange/jsonrange.go
	env  GOOS=linux GOARCH=arm GOARM=7 go build -o dist/bin/xlsx2json cmds/xlsx2json/xlsx2json.go
	env  GOOS=linux GOARCH=arm GOARM=7 go build -o dist/bin/xlsx2csv cmds/xlsx2csv/xlsx2csv.go
	env  GOOS=linux GOARCH=arm GOARM=7 go build -o dist/bin/csv2mdtable cmds/csv2mdtable/csv2mdtable.go
	env  GOOS=linux GOARCH=arm GOARM=7 go build -o dist/bin/csv2xlsx cmds/csv2xlsx/csv2xlsx.go
	env  GOOS=linux GOARCH=arm GOARM=7 go build -o dist/bin/csv2json cmds/csv2json/csv2json.go
	env  GOOS=linux GOARCH=arm GOARM=7 go build -o dist/bin/jsonmunge cmds/jsonmunge/jsonmunge.go
	env  GOOS=linux GOARCH=arm GOARM=7 go build -o dist/bin/jsonjoin cmds/jsonjoin/jsonjoin.go
	env  GOOS=linux GOARCH=arm GOARM=7 go build -o dist/bin/findfile cmds/findfile/findfile.go
	env  GOOS=linux GOARCH=arm GOARM=7 go build -o dist/bin/finddir cmds/finddir/finddir.go
	env  GOOS=linux GOARCH=arm GOARM=7 go build -o dist/bin/mergepath cmds/mergepath/mergepath.go
	env  GOOS=linux GOARCH=arm GOARM=7 go build -o dist/bin/reldate cmds/reldate/reldate.go
	env  GOOS=linux GOARCH=arm GOARM=7 go build -o dist/bin/range cmds/range/range.go
	env  GOOS=linux GOARCH=arm GOARM=7 go build -o dist/bin/timefmt cmds/timefmt/timefmt.go
	env  GOOS=linux GOARCH=arm GOARM=7 go build -o dist/bin/urlparse cmds/urlparse/urlparse.go
	env  GOOS=linux GOARCH=arm GOARM=7 go build -o dist/bin/csvcleaner cmds/csvcleaner/csvcleaner.go
	env  GOOS=linux GOARCH=arm GOARM=7 go build -o dist/bin/string cmds/string/string.go
	cd dist && zip -r $(PROJECT)-$(VERSION)-raspbian-arm7.zip README.md LICENSE INSTALL.md bin/* docs/* how-to/* demos/*
	rm -fR dist/bin

distribute_docs:
	mkdir -p dist/
	mkdir -p dist/
	cp -v README.md dist/
	cp -v LICENSE dist/
	cp -v INSTALL.md dist/
	cp -vR docs dist/
	cp -vR how-to dist/
	./package-versions.bash > dist/package-versions.txt
	
release: distribute_docs dist/linux-amd64 dist/macosx-amd64 dist/windows-amd64 dist/raspbian-arm7

