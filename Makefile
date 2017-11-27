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


build$(EXT): bin/csvcols$(EXT) bin/csvrows$(EXT) bin/csvfind$(EXT) bin/csvjoin$(EXT) bin/jsoncols$(EXT) bin/jsonrange$(EXT) bin/xlsx2json$(EXT) bin/xlsx2csv$(EXT) bin/csv2mdtable$(EXT) bin/csv2xlsx$(EXT) bin/csv2json$(EXT) bin/vcard2json$(EXT) bin/jsonjoin$(EXT) bin/jsonmunge$(EXT) bin/findfile$(EXT) bin/finddir$(EXT) bin/mergepath$(EXT) bin/reldate$(EXT) bin/range$(EXT) bin/timefmt$(EXT) bin/urlparse$(EXT) bin/splitstring$(EXT) bin/joinstring$(EXT) bin/hasprefix$(EXT) bin/hassuffix$(EXT) bin/trimprefix$(EXT) bin/trimsuffix$(EXT) bin/tolower$(EXT) bin/toupper$(EXT) bin/totitle$(EXT) bin/csvcleaner$(EXT) 


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

bin/vcard2json$(EXT): datatools.go cmds/vcard2json/vcard2json.go
	go build -o bin/vcard2json$(EXT) cmds/vcard2json/vcard2json.go

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

bin/splitstring$(EXT): datatools.go cmds/splitstring/splitstring.go
	go build -o bin/splitstring$(EXT) cmds/splitstring/splitstring.go

bin/joinstring$(EXT): datatools.go cmds/joinstring/joinstring.go
	go build -o bin/joinstring$(EXT) cmds/joinstring/joinstring.go

bin/hasprefix$(EXT): datatools.go cmds/hasprefix/hasprefix.go
	go build -o bin/hasprefix$(EXT) cmds/hasprefix/hasprefix.go

bin/hassuffix$(EXT): datatools.go cmds/hassuffix/hassuffix.go
	go build -o bin/hassuffix$(EXT) cmds/hassuffix/hassuffix.go

bin/trimprefix$(EXT): datatools.go cmds/trimprefix/trimprefix.go
	go build -o bin/trimprefix$(EXT) cmds/trimprefix/trimprefix.go

bin/trimsuffix$(EXT): datatools.go cmds/trimsuffix/trimsuffix.go
	go build -o bin/trimsuffix$(EXT) cmds/trimsuffix/trimsuffix.go

bin/tolower$(EXT): datatools.go cmds/tolower/tolower.go
	go build -o bin/tolower$(EXT) cmds/tolower/tolower.go

bin/toupper$(EXT): datatools.go cmds/toupper/toupper.go
	go build -o bin/toupper$(EXT) cmds/toupper/toupper.go

bin/totitle$(EXT): datatools.go cmds/totitle/totitle.go
	go build -o bin/totitle$(EXT) cmds/totitle/totitle.go

bin/csvcleaner$(EXT): datatools.go cmds/csvcleaner/csvcleaner.go
	go build -o bin/csvcleaner$(EXT) cmds/csvcleaner/csvcleaner.go

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
	env GOBIN=$(GOPATH)/bin go install cmds/joinstring/joinstring.go
	env GOBIN=$(GOPATH)/bin go install cmds/hasprefix/hasprefix.go
	env GOBIN=$(GOPATH)/bin go install cmds/hassuffix/hassuffix.go
	env GOBIN=$(GOPATH)/bin go install cmds/trimprefix/trimprefix.go
	env GOBIN=$(GOPATH)/bin go install cmds/trimsuffix/trimsuffix.go
	env GOBIN=$(GOPATH)/bin go install cmds/tolower/tolower.go
	env GOBIN=$(GOPATH)/bin go install cmds/toupper/toupper.go
	env GOBIN=$(GOPATH)/bin go install cmds/totitle/totitle.go
	env GOBIN=$(GOPATH)/bin go install cmds/csvcleaner/csvcleaner.go

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
	env  GOOS=linux GOARCH=amd64 go build -o dist/bin/joinstring cmds/joinstring/joinstring.go
	env  GOOS=linux GOARCH=amd64 go build -o dist/bin/hasprefix cmds/hasprefix/hasprefix.go
	env  GOOS=linux GOARCH=amd64 go build -o dist/bin/hassuffix cmds/hassuffix/hassuffix.go
	env  GOOS=linux GOARCH=amd64 go build -o dist/bin/trimprefix cmds/trimprefix/trimprefix.go
	env  GOOS=linux GOARCH=amd64 go build -o dist/bin/trimsuffix cmds/trimsuffix/trimsuffix.go
	env  GOOS=linux GOARCH=amd64 go build -o dist/bin/tolower cmds/tolower/tolower.go
	env  GOOS=linux GOARCH=amd64 go build -o dist/bin/toupper cmds/toupper/toupper.go
	env  GOOS=linux GOARCH=amd64 go build -o dist/bin/totitle cmds/totitle/totitle.go
	env  GOOS=linux GOARCH=amd64 go build -o dist/bin/csvcleaner cmds/csvcleaner/csvcleaner.go
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
	env  GOOS=darwin GOARCH=amd64 go build -o dist/bin/joinstring cmds/joinstring/joinstring.go
	env  GOOS=darwin GOARCH=amd64 go build -o dist/bin/hasprefix cmds/hasprefix/hasprefix.go
	env  GOOS=darwin GOARCH=amd64 go build -o dist/bin/hassuffix cmds/hassuffix/hassuffix.go
	env  GOOS=darwin GOARCH=amd64 go build -o dist/bin/trimprefix cmds/trimprefix/trimprefix.go
	env  GOOS=darwin GOARCH=amd64 go build -o dist/bin/trimsuffix cmds/trimsuffix/trimsuffix.go
	env  GOOS=darwin GOARCH=amd64 go build -o dist/bin/tolower cmds/tolower/tolower.go
	env  GOOS=darwin GOARCH=amd64 go build -o dist/bin/toupper cmds/toupper/toupper.go
	env  GOOS=darwin GOARCH=amd64 go build -o dist/bin/totitle cmds/totitle/totitle.go
	env  GOOS=darwin GOARCH=amd64 go build -o dist/bin/csvcleaner cmds/csvcleaner/csvcleaner.go
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
	env  GOOS=windows GOARCH=amd64 go build -o dist/bin/joinstring.exe cmds/joinstring/joinstring.go
	env  GOOS=windows GOARCH=amd64 go build -o dist/bin/hasprefix.exe cmds/hasprefix/hasprefix.go
	env  GOOS=windows GOARCH=amd64 go build -o dist/bin/hassuffix.exe cmds/hassuffix/hassuffix.go
	env  GOOS=windows GOARCH=amd64 go build -o dist/bin/trimprefix.exe cmds/trimprefix/trimprefix.go
	env  GOOS=windows GOARCH=amd64 go build -o dist/bin/trimsuffix.exe cmds/trimsuffix/trimsuffix.go
	env  GOOS=windows GOARCH=amd64 go build -o dist/bin/tolower.exe cmds/tolower/tolower.go
	env  GOOS=windows GOARCH=amd64 go build -o dist/bin/toupper.exe cmds/toupper/toupper.go
	env  GOOS=windows GOARCH=amd64 go build -o dist/bin/totitle.exe cmds/totitle/totitle.go
	env  GOOS=windows GOARCH=amd64 go build -o dist/bin/csvcleaner.exe cmds/csvcleaner/csvcleaner.go
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
	env  GOOS=linux GOARCH=arm GOARM=7 go build -o dist/bin/joinstring cmds/joinstring/joinstring.go
	env  GOOS=linux GOARCH=arm GOARM=7 go build -o dist/bin/hasprefix cmds/hasprefix/hasprefix.go
	env  GOOS=linux GOARCH=arm GOARM=7 go build -o dist/bin/hassuffix cmds/hassuffix/hassuffix.go
	env  GOOS=linux GOARCH=arm GOARM=7 go build -o dist/bin/trimprefix cmds/trimprefix/trimprefix.go
	env  GOOS=linux GOARCH=arm GOARM=7 go build -o dist/bin/trimsuffix cmds/trimsuffix/trimsuffix.go
	env  GOOS=linux GOARCH=arm GOARM=7 go build -o dist/bin/tolower cmds/tolower/tolower.go
	env  GOOS=linux GOARCH=arm GOARM=7 go build -o dist/bin/toupper cmds/toupper/toupper.go
	env  GOOS=linux GOARCH=arm GOARM=7 go build -o dist/bin/totitle cmds/totitle/totitle.go
	env  GOOS=linux GOARCH=arm GOARM=7 go build -o dist/bin/csvcleaner cmds/csvcleaner/csvcleaner.go
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

