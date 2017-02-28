#
# Simple Makefile
#
PROJECT = datatools

VERSION = $(shell grep -m1 'Version = ' $(PROJECT).go | cut -d\"  -f 2)

BRANCH = $(shell git branch | grep '* ' | cut -d\  -f 2)

build: bin/csvcols bin/csvjoin bin/jsoncols bin/jsonrange bin/xlsx2json bin/xlsx2csv bin/csv2mdtable bin/csv2xlsx

bin/csvcols: datatools.go cmds/csvcols/csvcols.go
	go build -o bin/csvcols cmds/csvcols/csvcols.go 

bin/csvjoin: datatools.go cmds/csvjoin/csvjoin.go
	go build -o bin/csvjoin cmds/csvjoin/csvjoin.go 

bin/jsoncols: datatools.go dotpath/dotpath.go cmds/jsoncols/jsoncols.go
	go build -o bin/jsoncols cmds/jsoncols/jsoncols.go 

bin/jsonrange: datatools.go dotpath/dotpath.go cmds/jsonrange/jsonrange.go
	go build -o bin/jsonrange cmds/jsonrange/jsonrange.go 

bin/xlsx2json: datatools.go cmds/xlsx2json/xlsx2json.go
	go build -o bin/xlsx2json cmds/xlsx2json/xlsx2json.go

bin/xlsx2csv: datatools.go cmds/xlsx2csv/xlsx2csv.go
	go build -o bin/xlsx2csv cmds/xlsx2csv/xlsx2csv.go

bin/csv2mdtable: datatools.go cmds/csv2mdtable/csv2mdtable.go
	go build -o bin/csv2mdtable cmds/csv2mdtable/csv2mdtable.go

bin/csv2xlsx: datatools.go cmds/csv2xlsx/csv2xlsx.go
	go build -o bin/csv2xlsx cmds/csv2xlsx/csv2xlsx.go

test:
	cd dotpath && go test

website:
	./mk-website.bash

status:
	git status

save:
	git commit -am "Quick Save"
	git push origin $(BRANCH)

refresh:
	git fetch origin
	git pull origin $(BRANCH)

publish:
	./mk-website.bash
	./publish.bash

clean: 
	if [ -f index.html ]; then /bin/rm *.html;fi
	if [ -d bin ]; then /bin/rm -fR bin; fi
	if [ -d dist ]; then /bin/rm -fR dist; fi
	if [ -f $(PROJECT)-$(VERSION)-release.zip ]; then rm -f $(PROJECT)-$(VERSION)-release.zip; fi

install:
	env GOBIN=$(HOME)/bin go install cmds/csvcols/csvcols.go
	env GOBIN=$(HOME)/bin go install cmds/csvjoin/csvjoin.go
	env GOBIN=$(HOME)/bin go install cmds/jsoncols/jsoncols.go
	env GOBIN=$(HOME)/bin go install cmds/jsonrange/jsonrange.go
	env GOBIN=$(HOME)/bin go install cmds/xlsx2json/xlsx2json.go
	env GOBIN=$(HOME)/bin go install cmds/xlsx2csv/xlsx2csv.go
	env GOBIN=$(HOME)/bin go install cmds/csv2mdtable/csv2mdtable.go
	env GOBIN=$(HOME)/bin go install cmds/csv2xlsx/csv2xlsx.go

dist/linux-amd64:
	env  GOOS=linux GOARCH=amd64 go build -o dist/linux-amd64/csvcols cmds/csvcols/csvcols.go
	env  GOOS=linux GOARCH=amd64 go build -o dist/linux-amd64/csvjoin cmds/csvjoin/csvjoin.go
	env  GOOS=linux GOARCH=amd64 go build -o dist/linux-amd64/jsoncols cmds/jsoncols/jsoncols.go
	env  GOOS=linux GOARCH=amd64 go build -o dist/linux-amd64/jsonrange cmds/jsonrange/jsonrange.go
	env  GOOS=linux GOARCH=amd64 go build -o dist/linux-amd64/xlsx2json cmds/xlsx2json/xlsx2json.go
	env  GOOS=linux GOARCH=amd64 go build -o dist/linux-amd64/xlsx2csv cmds/xlsx2csv/xlsx2csv.go
	env  GOOS=linux GOARCH=amd64 go build -o dist/linux-amd64/csv2mdtable cmds/csv2mdtable/csv2mdtable.go
	env  GOOS=linux GOARCH=amd64 go build -o dist/linux-amd64/csv2xlsx cmds/csv2xlsx/csv2xlsx.go

dist/macosx-amd64:
	env  GOOS=darwin GOARCH=amd64 go build -o dist/macosx-amd64/csvcols cmds/csvcols/csvcols.go
	env  GOOS=darwin GOARCH=amd64 go build -o dist/macosx-amd64/csvjoin cmds/csvjoin/csvjoin.go
	env  GOOS=darwin GOARCH=amd64 go build -o dist/macosx-amd64/jsoncols cmds/jsoncols/jsoncols.go
	env  GOOS=darwin GOARCH=amd64 go build -o dist/macosx-amd64/jsonrange cmds/jsonrange/jsonrange.go
	env  GOOS=darwin GOARCH=amd64 go build -o dist/macosx-amd64/xlsx2json cmds/xlsx2json/xlsx2json.go
	env  GOOS=darwin GOARCH=amd64 go build -o dist/macosx-amd64/xlsx2csv cmds/xlsx2csv/xlsx2csv.go
	env  GOOS=darwin GOARCH=amd64 go build -o dist/macosx-amd64/csv2mdtable cmds/csv2mdtable/csv2mdtable.go
	env  GOOS=darwin GOARCH=amd64 go build -o dist/macosx-amd64/csv2xlsx cmds/csv2xlsx/csv2xlsx.go

dist/windows-amd64:
	env  GOOS=windows GOARCH=amd64 go build -o dist/windows-amd64/csvcols.exe cmds/csvcols/csvcols.go
	env  GOOS=windows GOARCH=amd64 go build -o dist/windows-amd64/csvjoin.exe cmds/csvjoin/csvjoin.go
	env  GOOS=windows GOARCH=amd64 go build -o dist/windows-amd64/jsoncols.exe cmds/jsoncols/jsoncols.go
	env  GOOS=windows GOARCH=amd64 go build -o dist/windows-amd64/jsonrange.exe cmds/jsonrange/jsonrange.go
	env  GOOS=windows GOARCH=amd64 go build -o dist/windows-amd64/xlsx2json.exe cmds/xlsx2json/xlsx2json.go
	env  GOOS=windows GOARCH=amd64 go build -o dist/windows-amd64/xlsx2csv.exe cmds/xlsx2csv/xlsx2csv.go
	env  GOOS=windows GOARCH=amd64 go build -o dist/windows-amd64/csv2mdtable.exe cmds/csv2mdtable/csv2mdtable.go
	env  GOOS=windows GOARCH=amd64 go build -o dist/windows-amd64/csv2xlsx.exe cmds/csv2xlsx/csv2xlsx.go

dist/raspbian-arm7:
	env  GOOS=linux GOARCH=arm GOARM=7 go build -o dist/raspbian-arm7/csvcols cmds/csvcols/csvcols.go
	env  GOOS=linux GOARCH=arm GOARM=7 go build -o dist/raspbian-arm7/csvjoin cmds/csvjoin/csvjoin.go
	env  GOOS=linux GOARCH=arm GOARM=7 go build -o dist/raspbian-arm7/jsoncols cmds/jsoncols/jsoncols.go
	env  GOOS=linux GOARCH=arm GOARM=7 go build -o dist/raspbian-arm7/jsonrange cmds/jsonrange/jsonrange.go
	env  GOOS=linux GOARCH=arm GOARM=7 go build -o dist/raspbian-arm7/xlsx2json cmds/xlsx2json/xlsx2json.go
	env  GOOS=linux GOARCH=arm GOARM=7 go build -o dist/raspbian-arm7/xlsx2csv cmds/xlsx2csv/xlsx2csv.go
	env  GOOS=linux GOARCH=arm GOARM=7 go build -o dist/raspbian-arm7/csv2mdtable cmds/csv2mdtable/csv2mdtable.go
	env  GOOS=linux GOARCH=arm GOARM=7 go build -o dist/raspbian-arm7/csv2xlsx cmds/csv2xlsx/csv2xlsx.go

release: dist/linux-amd64 dist/macosx-amd64 dist/windows-amd64 dist/raspbian-arm7
	mkdir -p dist
	cp -v README.md dist/
	cp -v LICENSE dist/
	cp -v INSTALL.md dist/
	cp -v csvcols.md dist/
	cp -v csvjoin.md dist/
	cp -v jsoncols.md dist/
	cp -v jsonrange.md dist/
	cp -v xlsx2json.md dist/
	cp -v xlsx2csv.md dist/
	cp -v csv2mdtable.md dist/
	cp -v csv2xlsx.md dist/
	zip -r $(PROJECT)-$(VERSION)-release.zip dist/

