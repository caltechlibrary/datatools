#
# Simple Makefile
#
PROJECT = datatools

VERSION = $(shell grep -m1 'Version = ' $(PROJECT).go | cut -d\`  -f 2)

BRANCH = $(shell git branch | grep '* ' | cut -d\  -f 2)

build: bin/csvcols bin/csvrows bin/csvfind bin/csvjoin bin/jsoncols bin/jsonrange bin/xlsx2json bin/xlsx2csv bin/csv2mdtable bin/csv2xlsx bin/csv2json bin/vcard2json bin/jsonjoin bin/jsonmunge bin/findfile bin/finddir bin/mergepath bin/reldate bin/range bin/timefmt bin/urlparse bin/splitstring


bin/csvcols: datatools.go cmds/csvcols/csvcols.go
	go build -o bin/csvcols cmds/csvcols/csvcols.go 

bin/csvrows: datatools.go cmds/csvrows/csvrows.go
	go build -o bin/csvrows cmds/csvrows/csvrows.go 

bin/csvjoin: datatools.go cmds/csvjoin/csvjoin.go
	go build -o bin/csvjoin cmds/csvjoin/csvjoin.go 

bin/jsoncols: datatools.go cmds/jsoncols/jsoncols.go
	go build -o bin/jsoncols cmds/jsoncols/jsoncols.go 

bin/jsonrange: datatools.go cmds/jsonrange/jsonrange.go
	go build -o bin/jsonrange cmds/jsonrange/jsonrange.go 

bin/xlsx2json: datatools.go cmds/xlsx2json/xlsx2json.go
	go build -o bin/xlsx2json cmds/xlsx2json/xlsx2json.go

bin/xlsx2csv: datatools.go cmds/xlsx2csv/xlsx2csv.go
	go build -o bin/xlsx2csv cmds/xlsx2csv/xlsx2csv.go

bin/csv2mdtable: datatools.go cmds/csv2mdtable/csv2mdtable.go
	go build -o bin/csv2mdtable cmds/csv2mdtable/csv2mdtable.go

bin/csv2xlsx: datatools.go cmds/csv2xlsx/csv2xlsx.go
	go build -o bin/csv2xlsx cmds/csv2xlsx/csv2xlsx.go

bin/csv2json: datatools.go cmds/csv2json/csv2json.go
	go build -o bin/csv2json cmds/csv2json/csv2json.go

bin/csvfind: datatools.go cmds/csvfind/csvfind.go
	go build -o bin/csvfind cmds/csvfind/csvfind.go

bin/vcard2json: datatools.go cmds/vcard2json/vcard2json.go
	go build -o bin/vcard2json cmds/vcard2json/vcard2json.go

bin/jsonmunge: datatools.go cmds/jsonmunge/jsonmunge.go
	go build -o bin/jsonmunge cmds/jsonmunge/jsonmunge.go

bin/jsonjoin: datatools.go cmds/jsonjoin/jsonjoin.go
	go build -o bin/jsonjoin cmds/jsonjoin/jsonjoin.go

bin/findfile: datatools.go cmds/findfile/findfile.go
	go build -o bin/findfile cmds/findfile/findfile.go 

bin/finddir: datatools.go cmds/finddir/finddir.go
	go build -o bin/finddir cmds/finddir/finddir.go 

bin/mergepath: datatools.go cmds/mergepath/mergepath.go
	go build -o bin/mergepath cmds/mergepath/mergepath.go 

bin/reldate: datatools.go cmds/reldate/reldate.go
	go build -o bin/reldate cmds/reldate/reldate.go 

bin/range: datatools.go cmds/range/range.go
	go build -o bin/range cmds/range/range.go 

bin/timefmt: datatools.go cmds/timefmt/timefmt.go
	go build -o bin/timefmt cmds/timefmt/timefmt.go 

bin/urlparse: datatools.go cmds/urlparse/urlparse.go
	go build -o bin/urlparse cmds/urlparse/urlparse.go 

bin/splitstring: datatools.go cmds/splitstring/splitstring.go
	go build -o bin/splitstring cmds/splitstring/splitstring.go

test:
	go test

website:
	./mk-website.bash

status:
	git status

save:
	if [ "$(msg)" != "" ]; then git commit -am "$(msg)"; else git commit -am "Quick Save"; fi
	git push origin $(BRANCH)

refresh:
	git fetch origin
	git pull origin $(BRANCH)

publish:
	./mk-website.bash
	./publish.bash

clean: 
	if [ -d bin ]; then /bin/rm -fR bin; fi
	if [ -d dist ]; then /bin/rm -fR dist; fi

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
	env GOBIN=$(GOPATH)/bin go install cmds/vcard2json/vcard2json.go
	env GOBIN=$(GOPATH)/bin go install cmds/xlsx2json/xlsx2json.go
	env GOBIN=$(GOPATH)/bin go install cmds/xlsx2csv/xlsx2csv.go
	env GOBIN=$(GOPATH)/bin go install cmds/splitstring/splitstring.go

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
	env  GOOS=linux GOARCH=amd64 go build -o dist/bin/vcard2json cmds/vcard2json/vcard2json.go
	env  GOOS=linux GOARCH=amd64 go build -o dist/bin/findfile cmds/findfile/findfile.go
	env  GOOS=linux GOARCH=amd64 go build -o dist/bin/finddir cmds/finddir/finddir.go
	env  GOOS=linux GOARCH=amd64 go build -o dist/bin/mergepath cmds/mergepath/mergepath.go
	env  GOOS=linux GOARCH=amd64 go build -o dist/bin/reldate cmds/reldate/reldate.go
	env  GOOS=linux GOARCH=amd64 go build -o dist/bin/range cmds/range/range.go
	env  GOOS=linux GOARCH=amd64 go build -o dist/bin/timefmt cmds/timefmt/timefmt.go
	env  GOOS=linux GOARCH=amd64 go build -o dist/bin/urlparse cmds/urlparse/urlparse.go
	env  GOOS=linux GOARCH=amd64 go build -o dist/bin/splitstring cmds/splitstring/splitstring.go
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
	env  GOOS=darwin GOARCH=amd64 go build -o dist/bin/vcard2json cmds/vcard2json/vcard2json.go
	env  GOOS=darwin GOARCH=amd64 go build -o dist/bin/jsonmunge cmds/jsonmunge/jsonmunge.go
	env  GOOS=darwin GOARCH=amd64 go build -o dist/bin/jsonjoin cmds/jsonjoin/jsonjoin.go
	env  GOOS=darwin GOARCH=amd64 go build -o dist/bin/findfile cmds/findfile/findfile.go
	env  GOOS=darwin GOARCH=amd64 go build -o dist/bin/finddir cmds/finddir/finddir.go
	env  GOOS=darwin GOARCH=amd64 go build -o dist/bin/mergepath cmds/mergepath/mergepath.go
	env  GOOS=darwin GOARCH=amd64 go build -o dist/bin/reldate cmds/reldate/reldate.go
	env  GOOS=darwin GOARCH=amd64 go build -o dist/bin/range cmds/range/range.go
	env  GOOS=darwin GOARCH=amd64 go build -o dist/bin/timefmt cmds/timefmt/timefmt.go
	env  GOOS=darwin GOARCH=amd64 go build -o dist/bin/urlparse cmds/urlparse/urlparse.go
	env  GOOS=darwin GOARCH=amd64 go build -o dist/bin/splitstring cmds/splitstring/splitstring.go
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
	env  GOOS=windows GOARCH=amd64 go build -o dist/bin/vcard2json.exe cmds/vcard2json/vcard2json.go
	env  GOOS=windows GOARCH=amd64 go build -o dist/bin/jsonmunge.exe cmds/jsonmunge/jsonmunge.go
	env  GOOS=windows GOARCH=amd64 go build -o dist/bin/jsonjoin.exe cmds/jsonjoin/jsonjoin.go
	env  GOOS=windows GOARCH=amd64 go build -o dist/bin/findfile.exe cmds/findfile/findfile.go
	env  GOOS=windows GOARCH=amd64 go build -o dist/bin/finddir.exe cmds/finddir/finddir.go
	env  GOOS=windows GOARCH=amd64 go build -o dist/bin/mergepath.exe cmds/mergepath/mergepath.go
	env  GOOS=windows GOARCH=amd64 go build -o dist/bin/reldate.exe cmds/reldate/reldate.go
	env  GOOS=windows GOARCH=amd64 go build -o dist/bin/range.exe cmds/range/range.go
	env  GOOS=windows GOARCH=amd64 go build -o dist/bin/timefmt.exe cmds/timefmt/timefmt.go
	env  GOOS=windows GOARCH=amd64 go build -o dist/bin/urlparse.exe cmds/urlparse/urlparse.go
	env  GOOS=windows GOARCH=amd64 go build -o dist/bin/splitstring.exe cmds/splitstring/splitstring.go
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
	env  GOOS=linux GOARCH=arm GOARM=7 go build -o dist/bin/vcard2json cmds/vcard2json/vcard2json.go
	env  GOOS=linux GOARCH=arm GOARM=7 go build -o dist/bin/jsonmunge cmds/jsonmunge/jsonmunge.go
	env  GOOS=linux GOARCH=arm GOARM=7 go build -o dist/bin/jsonjoin cmds/jsonjoin/jsonjoin.go
	env  GOOS=linux GOARCH=arm GOARM=7 go build -o dist/bin/findfile cmds/findfile/findfile.go
	env  GOOS=linux GOARCH=arm GOARM=7 go build -o dist/bin/finddir cmds/finddir/finddir.go
	env  GOOS=linux GOARCH=arm GOARM=7 go build -o dist/bin/mergepath cmds/mergepath/mergepath.go
	env  GOOS=linux GOARCH=arm GOARM=7 go build -o dist/bin/reldate cmds/reldate/reldate.go
	env  GOOS=linux GOARCH=arm GOARM=7 go build -o dist/bin/range cmds/range/range.go
	env  GOOS=linux GOARCH=arm GOARM=7 go build -o dist/bin/timefmt cmds/timefmt/timefmt.go
	env  GOOS=linux GOARCH=arm GOARM=7 go build -o dist/bin/urlparse cmds/urlparse/urlparse.go
	env  GOOS=linux GOARCH=arm GOARM=7 go build -o dist/bin/splitstring cmds/splitstring/splitstring.go
	cd dist && zip -r $(PROJECT)-$(VERSION)-raspbian-arm7.zip README.md LICENSE INSTALL.md bin/* docs/* how-to/* demos/*
	rm -fR dist/bin

distribute_docs:
	mkdir -p dist/docs
	mkdir -p dist/how-to
	cp -v README.md dist/
	cp -v LICENSE dist/
	cp -v INSTALL.md dist/
	cp -v docs/*.md dist/docs/
	if [ -f dist/docs/nav.md ]; then rm dist/docs/nav.md; fi
	if [ -f dist/docs/index.md ]; then rm dist/docs/index.md; fi
	cp -v how-to/*.md dist/how-to/
	if [ -f dist/how-to/nav.md ]; then rm dist/how-to/nav.md; fi
	if [ -f dist/how-to/index.md ]; then rm dist/how-to/index.md; fi
	cp -vR demos dist/
	./package-versions.bash > dist/package-versions.txt
	
release: distribute_docs dist/linux-amd64 dist/macosx-amd64 dist/windows-amd64 dist/raspbian-arm7

