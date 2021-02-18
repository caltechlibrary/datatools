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
	bin/csvcleaner$(EXT) bin/string$(EXT) \
	bin/toml2json$(EXT) bin/json2toml$(EXT) \
	bin/yaml2json$(EXT) bin/json2yaml$(EXT) 


bin/csvcols$(EXT): datatools.go cmd/csvcols/csvcols.go
	go build -o bin/csvcols$(EXT) cmd/csvcols/csvcols.go 

bin/csvrows$(EXT): datatools.go cmd/csvrows/csvrows.go
	go build -o bin/csvrows$(EXT) cmd/csvrows/csvrows.go 

bin/csvjoin$(EXT): datatools.go cmd/csvjoin/csvjoin.go
	go build -o bin/csvjoin$(EXT) cmd/csvjoin/csvjoin.go 

bin/jsoncols$(EXT): datatools.go cmd/jsoncols/jsoncols.go
	go build -o bin/jsoncols$(EXT) cmd/jsoncols/jsoncols.go 

bin/jsonrange$(EXT): datatools.go cmd/jsonrange/jsonrange.go
	go build -o bin/jsonrange$(EXT) cmd/jsonrange/jsonrange.go 

bin/xlsx2json$(EXT): datatools.go cmd/xlsx2json/xlsx2json.go
	go build -o bin/xlsx2json$(EXT) cmd/xlsx2json/xlsx2json.go

bin/xlsx2csv$(EXT): datatools.go cmd/xlsx2csv/xlsx2csv.go
	go build -o bin/xlsx2csv$(EXT) cmd/xlsx2csv/xlsx2csv.go

bin/csv2mdtable$(EXT): datatools.go cmd/csv2mdtable/csv2mdtable.go
	go build -o bin/csv2mdtable$(EXT) cmd/csv2mdtable/csv2mdtable.go

bin/csv2xlsx$(EXT): datatools.go cmd/csv2xlsx/csv2xlsx.go
	go build -o bin/csv2xlsx$(EXT) cmd/csv2xlsx/csv2xlsx.go

bin/csv2json$(EXT): datatools.go cmd/csv2json/csv2json.go
	go build -o bin/csv2json$(EXT) cmd/csv2json/csv2json.go

bin/csvfind$(EXT): datatools.go cmd/csvfind/csvfind.go
	go build -o bin/csvfind$(EXT) cmd/csvfind/csvfind.go

bin/jsonmunge$(EXT): datatools.go cmd/jsonmunge/jsonmunge.go
	go build -o bin/jsonmunge$(EXT) cmd/jsonmunge/jsonmunge.go

bin/jsonjoin$(EXT): datatools.go cmd/jsonjoin/jsonjoin.go
	go build -o bin/jsonjoin$(EXT) cmd/jsonjoin/jsonjoin.go

bin/findfile$(EXT): datatools.go cmd/findfile/findfile.go
	go build -o bin/findfile$(EXT) cmd/findfile/findfile.go 

bin/finddir$(EXT): datatools.go cmd/finddir/finddir.go
	go build -o bin/finddir$(EXT) cmd/finddir/finddir.go 

bin/mergepath$(EXT): datatools.go cmd/mergepath/mergepath.go
	go build -o bin/mergepath$(EXT) cmd/mergepath/mergepath.go 

bin/reldate$(EXT): datatools.go cmd/reldate/reldate.go
	go build -o bin/reldate$(EXT) cmd/reldate/reldate.go 

bin/range$(EXT): datatools.go cmd/range/range.go
	go build -o bin/range$(EXT) cmd/range/range.go 

bin/timefmt$(EXT): datatools.go cmd/timefmt/timefmt.go
	go build -o bin/timefmt$(EXT) cmd/timefmt/timefmt.go 

bin/urlparse$(EXT): datatools.go cmd/urlparse/urlparse.go
	go build -o bin/urlparse$(EXT) cmd/urlparse/urlparse.go 

bin/csvcleaner$(EXT): datatools.go cmd/csvcleaner/csvcleaner.go
	go build -o bin/csvcleaner$(EXT) cmd/csvcleaner/csvcleaner.go

bin/string$(EXT): datatools.go cmd/string/string.go
	go build -o bin/string$(EXT) cmd/string/string.go

bin/toml2json$(EXT): datatools.go cmd/toml2json/toml2json.go
	go build -o bin/toml2json$(EXT) cmd/toml2json/toml2json.go

bin/json2toml$(EXT): datatools.go cmd/json2toml/json2toml.go
	go build -o bin/json2toml$(EXT) cmd/json2toml/json2toml.go

bin/yaml2json$(EXT): datatools.go cmd/yaml2json/yaml2json.go
	go build -o bin/yaml2json$(EXT) cmd/yaml2json/yaml2json.go

bin/json2yaml$(EXT): datatools.go cmd/json2yaml/json2yaml.go
	go build -o bin/json2yaml$(EXT) cmd/json2yaml/json2yaml.go

test: build
	go test
	bash test_cmd.bash

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
	if [ -d man ]; then rm -fR man; fi

man: build
	mkdir -p man/man1
	bin/csvcols -generate-manpage | nroff -Tutf8 -man > man/man1/csvcols.1
	bin/csvrows -generate-manpage | nroff -Tutf8 -man > man/man1/csvrows.1
	bin/csvfind -generate-manpage | nroff -Tutf8 -man > man/man1/csvfind.1
	bin/csvjoin -generate-manpage | nroff -Tutf8 -man > man/man1/csvjoin.1
	bin/csv2mdtable -generate-manpage | nroff -Tutf8 -man > man/man1/csv2mdtable.1
	bin/csv2xlsx -generate-manpage | nroff -Tutf8 -man > man/man1/csv2xlsx.1
	bin/csv2json -generate-manpage | nroff -Tutf8 -man > man/man1/csv2json.1
	bin/findfile -generate-manpage | nroff -Tutf8 -man > man/man1/findfile.1
	bin/finddir -generate-manpage | nroff -Tutf8 -man > man/man1/finddir.1
	bin/jsoncols -generate-manpage | nroff -Tutf8 -man > man/man1/jsoncols.1
	bin/jsonrange -generate-manpage | nroff -Tutf8 -man > man/man1/jsonrange.1
	bin/jsonjoin -generate-manpage | nroff -Tutf8 -man > man/man1/jsonjoin.1
	bin/jsonmunge -generate-manpage | nroff -Tutf8 -man > man/man1/jsonmunge.1
	bin/mergepath -generate-manpage | nroff -Tutf8 -man > man/man1/mergepath.1
	bin/reldate -generate-manpage | nroff -Tutf8 -man > man/man1/reldate.1
	bin/range -generate-manpage | nroff -Tutf8 -man > man/man1/range.1
	bin/timefmt -generate-manpage | nroff -Tutf8 -man > man/man1/timefmt.1
	bin/urlparse -generate-manpage | nroff -Tutf8 -man > man/man1/urlparse.1
	bin/xlsx2json -generate-manpage | nroff -Tutf8 -man > man/man1/xlsx2json.1
	bin/xlsx2csv -generate-manpage | nroff -Tutf8 -man > man/man1/xlsx2csv.1
	bin/csvcleaner -generate-manpage | nroff -Tutf8 -man > man/man1/csvcleaner.1
	bin/string -generate-manpage | nroff -Tutf8 -man > man/man1/string.1
	bin/toml2json -generate-manpage | nroff -Tutf8 -man > man/man1/toml2json.1
	bin/json2toml -generate-manpage | nroff -Tutf8 -man > man/man1/json2toml.1
	bin/yaml2json -generate-manpage | nroff -Tutf8 -man > man/man1/yaml2json.1
	bin/json2yaml -generate-manpage | nroff -Tutf8 -man > man/man1/json2yaml.1

install:
	env GOBIN=$(GOPATH)/bin go install cmd/csvcols/csvcols.go
	env GOBIN=$(GOPATH)/bin go install cmd/csvrows/csvrows.go
	env GOBIN=$(GOPATH)/bin go install cmd/csvfind/csvfind.go
	env GOBIN=$(GOPATH)/bin go install cmd/csvjoin/csvjoin.go
	env GOBIN=$(GOPATH)/bin go install cmd/csv2mdtable/csv2mdtable.go
	env GOBIN=$(GOPATH)/bin go install cmd/csv2xlsx/csv2xlsx.go
	env GOBIN=$(GOPATH)/bin go install cmd/csv2json/csv2json.go
	env GOBIN=$(GOPATH)/bin go install cmd/findfile/findfile.go
	env GOBIN=$(GOPATH)/bin go install cmd/finddir/finddir.go
	env GOBIN=$(GOPATH)/bin go install cmd/jsoncols/jsoncols.go
	env GOBIN=$(GOPATH)/bin go install cmd/jsonrange/jsonrange.go
	env GOBIN=$(GOPATH)/bin go install cmd/jsonmunge/jsonmunge.go
	env GOBIN=$(GOPATH)/bin go install cmd/jsonjoin/jsonjoin.go
	env GOBIN=$(GOPATH)/bin go install cmd/mergepath/mergepath.go
	env GOBIN=$(GOPATH)/bin go install cmd/reldate/reldate.go
	env GOBIN=$(GOPATH)/bin go install cmd/range/range.go
	env GOBIN=$(GOPATH)/bin go install cmd/timefmt/timefmt.go
	env GOBIN=$(GOPATH)/bin go install cmd/urlparse/urlparse.go
	env GOBIN=$(GOPATH)/bin go install cmd/xlsx2json/xlsx2json.go
	env GOBIN=$(GOPATH)/bin go install cmd/xlsx2csv/xlsx2csv.go
	env GOBIN=$(GOPATH)/bin go install cmd/csvcleaner/csvcleaner.go
	env GOBIN=$(GOPATH)/bin go install cmd/string/string.go
	env GOBIN=$(GOPATH)/bin go install cmd/toml2json/toml2json.go
	env GOBIN=$(GOPATH)/bin go install cmd/json2toml/json2toml.go
	env GOBIN=$(GOPATH)/bin go install cmd/yaml2json/yaml2json.go
	env GOBIN=$(GOPATH)/bin go install cmd/json2yaml/json2yaml.go

dist/linux-amd64:
	mkdir -p dist/bin
	env  GOOS=linux GOARCH=amd64 go build -o dist/bin/csvcols cmd/csvcols/csvcols.go
	env  GOOS=linux GOARCH=amd64 go build -o dist/bin/csvrows cmd/csvrows/csvrows.go
	env  GOOS=linux GOARCH=amd64 go build -o dist/bin/csvfind cmd/csvfind/csvfind.go
	env  GOOS=linux GOARCH=amd64 go build -o dist/bin/csvjoin cmd/csvjoin/csvjoin.go
	env  GOOS=linux GOARCH=amd64 go build -o dist/bin/jsoncols cmd/jsoncols/jsoncols.go
	env  GOOS=linux GOARCH=amd64 go build -o dist/bin/jsonrange cmd/jsonrange/jsonrange.go
	env  GOOS=linux GOARCH=amd64 go build -o dist/bin/xlsx2json cmd/xlsx2json/xlsx2json.go
	env  GOOS=linux GOARCH=amd64 go build -o dist/bin/xlsx2csv cmd/xlsx2csv/xlsx2csv.go
	env  GOOS=linux GOARCH=amd64 go build -o dist/bin/csv2mdtable cmd/csv2mdtable/csv2mdtable.go
	env  GOOS=linux GOARCH=amd64 go build -o dist/bin/csv2xlsx cmd/csv2xlsx/csv2xlsx.go
	env  GOOS=linux GOARCH=amd64 go build -o dist/bin/csv2json cmd/csv2json/csv2json.go
	env  GOOS=linux GOARCH=amd64 go build -o dist/bin/jsonmunge cmd/jsonmunge/jsonmunge.go
	env  GOOS=linux GOARCH=amd64 go build -o dist/bin/jsonjoin cmd/jsonjoin/jsonjoin.go
	env  GOOS=linux GOARCH=amd64 go build -o dist/bin/findfile cmd/findfile/findfile.go
	env  GOOS=linux GOARCH=amd64 go build -o dist/bin/finddir cmd/finddir/finddir.go
	env  GOOS=linux GOARCH=amd64 go build -o dist/bin/mergepath cmd/mergepath/mergepath.go
	env  GOOS=linux GOARCH=amd64 go build -o dist/bin/reldate cmd/reldate/reldate.go
	env  GOOS=linux GOARCH=amd64 go build -o dist/bin/range cmd/range/range.go
	env  GOOS=linux GOARCH=amd64 go build -o dist/bin/timefmt cmd/timefmt/timefmt.go
	env  GOOS=linux GOARCH=amd64 go build -o dist/bin/urlparse cmd/urlparse/urlparse.go
	env  GOOS=linux GOARCH=amd64 go build -o dist/bin/csvcleaner cmd/csvcleaner/csvcleaner.go
	env  GOOS=linux GOARCH=amd64 go build -o dist/bin/string cmd/string/string.go
	env  GOOS=linux GOARCH=amd64 go build -o dist/bin/toml2json cmd/toml2json/toml2json.go
	env  GOOS=linux GOARCH=amd64 go build -o dist/bin/json2toml cmd/json2toml/json2toml.go
	env  GOOS=linux GOARCH=amd64 go build -o dist/bin/yaml2json cmd/yaml2json/yaml2json.go
	env  GOOS=linux GOARCH=amd64 go build -o dist/bin/json2yaml cmd/json2yaml/json2yaml.go
	cd dist && zip -r $(PROJECT)-$(VERSION)-linux-amd64.zip README.md LICENSE INSTALL.md bin/* docs/* how-to/* demos/*
	rm -fR dist/bin


dist/macos-amd64:
	mkdir -p dist/bin
	env  GOOS=darwin GOARCH=amd64 go build -o dist/bin/csvcols cmd/csvcols/csvcols.go
	env  GOOS=darwin GOARCH=amd64 go build -o dist/bin/csvrows cmd/csvrows/csvrows.go
	env  GOOS=darwin GOARCH=amd64 go build -o dist/bin/csvfind cmd/csvfind/csvfind.go
	env  GOOS=darwin GOARCH=amd64 go build -o dist/bin/csvjoin cmd/csvjoin/csvjoin.go
	env  GOOS=darwin GOARCH=amd64 go build -o dist/bin/jsoncols cmd/jsoncols/jsoncols.go
	env  GOOS=darwin GOARCH=amd64 go build -o dist/bin/jsonrange cmd/jsonrange/jsonrange.go
	env  GOOS=darwin GOARCH=amd64 go build -o dist/bin/xlsx2json cmd/xlsx2json/xlsx2json.go
	env  GOOS=darwin GOARCH=amd64 go build -o dist/bin/xlsx2csv cmd/xlsx2csv/xlsx2csv.go
	env  GOOS=darwin GOARCH=amd64 go build -o dist/bin/csv2mdtable cmd/csv2mdtable/csv2mdtable.go
	env  GOOS=darwin GOARCH=amd64 go build -o dist/bin/csv2xlsx cmd/csv2xlsx/csv2xlsx.go
	env  GOOS=darwin GOARCH=amd64 go build -o dist/bin/csv2json cmd/csv2json/csv2json.go
	env  GOOS=darwin GOARCH=amd64 go build -o dist/bin/jsonmunge cmd/jsonmunge/jsonmunge.go
	env  GOOS=darwin GOARCH=amd64 go build -o dist/bin/jsonjoin cmd/jsonjoin/jsonjoin.go
	env  GOOS=darwin GOARCH=amd64 go build -o dist/bin/findfile cmd/findfile/findfile.go
	env  GOOS=darwin GOARCH=amd64 go build -o dist/bin/finddir cmd/finddir/finddir.go
	env  GOOS=darwin GOARCH=amd64 go build -o dist/bin/mergepath cmd/mergepath/mergepath.go
	env  GOOS=darwin GOARCH=amd64 go build -o dist/bin/reldate cmd/reldate/reldate.go
	env  GOOS=darwin GOARCH=amd64 go build -o dist/bin/range cmd/range/range.go
	env  GOOS=darwin GOARCH=amd64 go build -o dist/bin/timefmt cmd/timefmt/timefmt.go
	env  GOOS=darwin GOARCH=amd64 go build -o dist/bin/urlparse cmd/urlparse/urlparse.go
	env  GOOS=darwin GOARCH=amd64 go build -o dist/bin/csvcleaner cmd/csvcleaner/csvcleaner.go
	env  GOOS=darwin GOARCH=amd64 go build -o dist/bin/string cmd/string/string.go
	env  GOOS=darwin GOARCH=amd64 go build -o dist/bin/toml2json cmd/toml2json/toml2json.go
	env  GOOS=darwin GOARCH=amd64 go build -o dist/bin/json2toml cmd/json2toml/json2toml.go
	env  GOOS=darwin GOARCH=amd64 go build -o dist/bin/yaml2json cmd/yaml2json/yaml2json.go
	env  GOOS=darwin GOARCH=amd64 go build -o dist/bin/json2yaml cmd/json2yaml/json2yaml.go
	cd dist && zip -r $(PROJECT)-$(VERSION)-macos-amd64.zip README.md LICENSE INSTALL.md bin/* docs/* how-to/* demos/*
	rm -fR dist/bin
	

dist/macos-arm64:
	mkdir -p dist/bin
	env  GOOS=darwin GOARCH=arm64 go build -o dist/bin/csvcols cmd/csvcols/csvcols.go
	env  GOOS=darwin GOARCH=arm64 go build -o dist/bin/csvrows cmd/csvrows/csvrows.go
	env  GOOS=darwin GOARCH=arm64 go build -o dist/bin/csvfind cmd/csvfind/csvfind.go
	env  GOOS=darwin GOARCH=arm64 go build -o dist/bin/csvjoin cmd/csvjoin/csvjoin.go
	env  GOOS=darwin GOARCH=arm64 go build -o dist/bin/jsoncols cmd/jsoncols/jsoncols.go
	env  GOOS=darwin GOARCH=arm64 go build -o dist/bin/jsonrange cmd/jsonrange/jsonrange.go
	env  GOOS=darwin GOARCH=arm64 go build -o dist/bin/xlsx2json cmd/xlsx2json/xlsx2json.go
	env  GOOS=darwin GOARCH=arm64 go build -o dist/bin/xlsx2csv cmd/xlsx2csv/xlsx2csv.go
	env  GOOS=darwin GOARCH=arm64 go build -o dist/bin/csv2mdtable cmd/csv2mdtable/csv2mdtable.go
	env  GOOS=darwin GOARCH=arm64 go build -o dist/bin/csv2xlsx cmd/csv2xlsx/csv2xlsx.go
	env  GOOS=darwin GOARCH=arm64 go build -o dist/bin/csv2json cmd/csv2json/csv2json.go
	env  GOOS=darwin GOARCH=arm64 go build -o dist/bin/jsonmunge cmd/jsonmunge/jsonmunge.go
	env  GOOS=darwin GOARCH=arm64 go build -o dist/bin/jsonjoin cmd/jsonjoin/jsonjoin.go
	env  GOOS=darwin GOARCH=arm64 go build -o dist/bin/findfile cmd/findfile/findfile.go
	env  GOOS=darwin GOARCH=arm64 go build -o dist/bin/finddir cmd/finddir/finddir.go
	env  GOOS=darwin GOARCH=arm64 go build -o dist/bin/mergepath cmd/mergepath/mergepath.go
	env  GOOS=darwin GOARCH=arm64 go build -o dist/bin/reldate cmd/reldate/reldate.go
	env  GOOS=darwin GOARCH=arm64 go build -o dist/bin/range cmd/range/range.go
	env  GOOS=darwin GOARCH=arm64 go build -o dist/bin/timefmt cmd/timefmt/timefmt.go
	env  GOOS=darwin GOARCH=arm64 go build -o dist/bin/urlparse cmd/urlparse/urlparse.go
	env  GOOS=darwin GOARCH=arm64 go build -o dist/bin/csvcleaner cmd/csvcleaner/csvcleaner.go
	env  GOOS=darwin GOARCH=arm64 go build -o dist/bin/string cmd/string/string.go
	env  GOOS=darwin GOARCH=arm64 go build -o dist/bin/toml2json cmd/toml2json/toml2json.go
	env  GOOS=darwin GOARCH=arm64 go build -o dist/bin/json2toml cmd/json2toml/json2toml.go
	env  GOOS=darwin GOARCH=arm64 go build -o dist/bin/yaml2json cmd/yaml2json/yaml2json.go
	env  GOOS=darwin GOARCH=arm64 go build -o dist/bin/json2yaml cmd/json2yaml/json2yaml.go
	cd dist && zip -r $(PROJECT)-$(VERSION)-macos-arm64.zip README.md LICENSE INSTALL.md bin/* docs/* how-to/* demos/*
	rm -fR dist/bin
	

dist/windows-amd64:
	mkdir -p dist/bin
	env  GOOS=windows GOARCH=amd64 go build -o dist/bin/csvcols.exe cmd/csvcols/csvcols.go
	env  GOOS=windows GOARCH=amd64 go build -o dist/bin/csvrows.exe cmd/csvrows/csvrows.go
	env  GOOS=windows GOARCH=amd64 go build -o dist/bin/csvfind.exe cmd/csvfind/csvfind.go
	env  GOOS=windows GOARCH=amd64 go build -o dist/bin/csvjoin.exe cmd/csvjoin/csvjoin.go
	env  GOOS=windows GOARCH=amd64 go build -o dist/bin/jsoncols.exe cmd/jsoncols/jsoncols.go
	env  GOOS=windows GOARCH=amd64 go build -o dist/bin/jsonrange.exe cmd/jsonrange/jsonrange.go
	env  GOOS=windows GOARCH=amd64 go build -o dist/bin/xlsx2json.exe cmd/xlsx2json/xlsx2json.go
	env  GOOS=windows GOARCH=amd64 go build -o dist/bin/xlsx2csv.exe cmd/xlsx2csv/xlsx2csv.go
	env  GOOS=windows GOARCH=amd64 go build -o dist/bin/csv2mdtable.exe cmd/csv2mdtable/csv2mdtable.go
	env  GOOS=windows GOARCH=amd64 go build -o dist/bin/csv2xlsx.exe cmd/csv2xlsx/csv2xlsx.go
	env  GOOS=windows GOARCH=amd64 go build -o dist/bin/csv2json.exe cmd/csv2json/csv2json.go
	env  GOOS=windows GOARCH=amd64 go build -o dist/bin/jsonmunge.exe cmd/jsonmunge/jsonmunge.go
	env  GOOS=windows GOARCH=amd64 go build -o dist/bin/jsonjoin.exe cmd/jsonjoin/jsonjoin.go
	env  GOOS=windows GOARCH=amd64 go build -o dist/bin/findfile.exe cmd/findfile/findfile.go
	env  GOOS=windows GOARCH=amd64 go build -o dist/bin/finddir.exe cmd/finddir/finddir.go
	env  GOOS=windows GOARCH=amd64 go build -o dist/bin/mergepath.exe cmd/mergepath/mergepath.go
	env  GOOS=windows GOARCH=amd64 go build -o dist/bin/reldate.exe cmd/reldate/reldate.go
	env  GOOS=windows GOARCH=amd64 go build -o dist/bin/range.exe cmd/range/range.go
	env  GOOS=windows GOARCH=amd64 go build -o dist/bin/timefmt.exe cmd/timefmt/timefmt.go
	env  GOOS=windows GOARCH=amd64 go build -o dist/bin/urlparse.exe cmd/urlparse/urlparse.go
	env  GOOS=windows GOARCH=amd64 go build -o dist/bin/csvcleaner.exe cmd/csvcleaner/csvcleaner.go
	env  GOOS=windows GOARCH=amd64 go build -o dist/bin/string.exe cmd/string/string.go
	env  GOOS=windows GOARCH=amd64 go build -o dist/bin/toml2json.exe cmd/toml2json/toml2json.go
	env  GOOS=windows GOARCH=amd64 go build -o dist/bin/json2toml.exe cmd/json2toml/json2toml.go
	env  GOOS=windows GOARCH=amd64 go build -o dist/bin/yaml2json.exe cmd/yaml2json/yaml2json.go
	env  GOOS=windows GOARCH=amd64 go build -o dist/bin/json2yaml.exe cmd/json2yaml/json2yaml.go
	cd dist && zip -r $(PROJECT)-$(VERSION)-windows-amd64.zip README.md LICENSE INSTALL.md bin/* docs/* how-to/* demos/*
	rm -fR dist/bin




dist/raspbian-arm7:
	mkdir -p dist/bin
	env  GOOS=linux GOARCH=arm GOARM=7 go build -o dist/bin/csvcols cmd/csvcols/csvcols.go
	env  GOOS=linux GOARCH=arm GOARM=7 go build -o dist/bin/csvrows cmd/csvrows/csvrows.go
	env  GOOS=linux GOARCH=arm GOARM=7 go build -o dist/bin/csvfind cmd/csvfind/csvfind.go
	env  GOOS=linux GOARCH=arm GOARM=7 go build -o dist/bin/csvjoin cmd/csvjoin/csvjoin.go
	env  GOOS=linux GOARCH=arm GOARM=7 go build -o dist/bin/jsoncols cmd/jsoncols/jsoncols.go
	env  GOOS=linux GOARCH=arm GOARM=7 go build -o dist/bin/jsonrange cmd/jsonrange/jsonrange.go
	env  GOOS=linux GOARCH=arm GOARM=7 go build -o dist/bin/xlsx2json cmd/xlsx2json/xlsx2json.go
	env  GOOS=linux GOARCH=arm GOARM=7 go build -o dist/bin/xlsx2csv cmd/xlsx2csv/xlsx2csv.go
	env  GOOS=linux GOARCH=arm GOARM=7 go build -o dist/bin/csv2mdtable cmd/csv2mdtable/csv2mdtable.go
	env  GOOS=linux GOARCH=arm GOARM=7 go build -o dist/bin/csv2xlsx cmd/csv2xlsx/csv2xlsx.go
	env  GOOS=linux GOARCH=arm GOARM=7 go build -o dist/bin/csv2json cmd/csv2json/csv2json.go
	env  GOOS=linux GOARCH=arm GOARM=7 go build -o dist/bin/jsonmunge cmd/jsonmunge/jsonmunge.go
	env  GOOS=linux GOARCH=arm GOARM=7 go build -o dist/bin/jsonjoin cmd/jsonjoin/jsonjoin.go
	env  GOOS=linux GOARCH=arm GOARM=7 go build -o dist/bin/findfile cmd/findfile/findfile.go
	env  GOOS=linux GOARCH=arm GOARM=7 go build -o dist/bin/finddir cmd/finddir/finddir.go
	env  GOOS=linux GOARCH=arm GOARM=7 go build -o dist/bin/mergepath cmd/mergepath/mergepath.go
	env  GOOS=linux GOARCH=arm GOARM=7 go build -o dist/bin/reldate cmd/reldate/reldate.go
	env  GOOS=linux GOARCH=arm GOARM=7 go build -o dist/bin/range cmd/range/range.go
	env  GOOS=linux GOARCH=arm GOARM=7 go build -o dist/bin/timefmt cmd/timefmt/timefmt.go
	env  GOOS=linux GOARCH=arm GOARM=7 go build -o dist/bin/urlparse cmd/urlparse/urlparse.go
	env  GOOS=linux GOARCH=arm GOARM=7 go build -o dist/bin/csvcleaner cmd/csvcleaner/csvcleaner.go
	env  GOOS=linux GOARCH=arm GOARM=7 go build -o dist/bin/string cmd/string/string.go
	env  GOOS=linux GOARCH=arm GOARM=7 go build -o dist/bin/toml2json cmd/toml2json/toml2json.go
	env  GOOS=linux GOARCH=arm GOARM=7 go build -o dist/bin/json2toml cmd/json2toml/json2toml.go
	env  GOOS=linux GOARCH=arm GOARM=7 go build -o dist/bin/yaml2json cmd/yaml2json/yaml2json.go
	env  GOOS=linux GOARCH=arm GOARM=7 go build -o dist/bin/json2yaml cmd/json2yaml/json2yaml.go
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
	
release: distribute_docs dist/linux-amd64 dist/macos-amd64 dist/macos-arm64 dist/windows-amd64 dist/raspbian-arm7

