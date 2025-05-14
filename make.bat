@echo off
REM This is a Windows 10 Batch file for building dataset command
REM from the command prompt.
REM
REM It requires: go version 1.16.6 or better and the cli for git installed
REM
go version
mkdir bin
echo "Getting ready to build the datatools in bin"
go build -o bin\codemeta2cff.exe cmd\codemeta2cff\codemeta2cff.exe
go build -o bin\csv2json.exe cmd\csv2json\csv2json.exe
go build -o bin\csv2jsonl.exe cmd\csv2jsonl\csv2jsonl.exe
go build -o bin\csv2mdtable.exe cmd\csv2mdtable\csv2mdtable.exe
go build -o bin\csv2tab.exe cmd\csv2tab\csv2tab.exe
go build -o bin\csv2xlsx.exe cmd\csv2xlsx\csv2xlsx.exe
go build -o bin\csvcleaner.exe cmd\csvcleaner\csvcleaner.exe
go build -o bin\csvcols.exe cmd\csvcols\csvcols.exe
go build -o bin\csvfind.exe cmd\csvfind\csvfind.exe
go build -o bin\csvjoin.exe cmd\csvjoin\csvjoin.exe
go build -o bin\csvrows.exe cmd\csvrows\csvrows.exe
go build -o bin\finddir.exe cmd\finddir\finddir.exe
go build -o bin\findfile.exe cmd\findfile\findfile.exe
go build -o bin\json2toml.exe cmd\json2toml\json2toml.exe
go build -o bin\json2yaml.exe cmd\json2yaml\json2yaml.exe
go build -o bin\jsoncols.exe cmd\jsoncols\jsoncols.exe
go build -o bin\jsonjoin.exe cmd\jsonjoin\jsonjoin.exe
go build -o bin\jsonmunge.exe cmd\jsonmunge\jsonmunge.exe
go build -o bin\jsonrange.exe cmd\jsonrange\jsonrange.exe
go build -o bin\jsonobjects2csv.exe cmd\jsonobjects2csv\jsonobjects2csv.exe
go build -o bin\mergepath.exe cmd\mergepath\mergepath.exe
go build -o bin\range.exe cmd\range\range.exe
go build -o bin\reldate.exe cmd\reldate\reldate.exe
go build -o bin\reltime.exe cmd\reltime\reltime.exe
go build -o bin\sql2csv.exe cmd\sql2csv\sql2csv.exe
go build -o bin\string.exe cmd\string\string.exe
go build -o bin\tab2csv.exe cmd\tab2csv\tab2csv.exe
go build -o bin\timefmt.exe cmd\timefmt\timefmt.exe
go build -o bin\toml2json.exe cmd\toml2json\toml2json.exe
go build -o bin\urlparse.exe cmd\urlparse\urlparse.exe
go build -o bin\xlsx2csv.exe cmd\xlsx2csv\xlsx2csv.exe
go build -o bin\xlsx2json.exe cmd\xlsx2json\xlsx2json.exe
go build -o bin\yaml2json.exe cmd\yaml2json\yaml2json.exe
echo "Checking compile should see version number of dataset"
bin\codemeta2cff.exe -version
bin\csv2json.exe -version
bin\csv2jsonl.exe -version
bin\csv2mdtable.exe -version
bin\csv2tab.exe -version
bin\csv2xlsx.exe -version
bin\csvcleaner.exe -version
bin\csvcols.exe -version
bin\csvfind.exe -version
bin\csvjoin.exe -version
bin\csvrows.exe -version
bin\finddir.exe -version
bin\findfile.exe -version
bin\json2toml.exe -version
bin\json2yaml.exe -version
bin\jsoncols.exe -version
bin\jsonjoin.exe -version
bin\jsonmunge.exe -version
bin\jsonrange.exe -version
bin\jsonobjects2csv.exe -version
bin\mergepath.exe -version
bin\range.exe -version
bin\reldate.exe -version
bin\reltime.exe -version
bin\sql2csv.exe -version
bin\string.exe -version
bin\tab2csv.exe -version
bin\timefmt.exe -version
bin\toml2json.exe -version
bin\urlparse.exe -version
bin\xlsx2csv.exe -version
bin\xlsx2json.exe -version
bin\yaml2json.exe -version
echo "If OK, you can now copy the dataset.exe to %USERPROFILE%\goin"
echo ""
echo "      copy bin\* %USERPROFILE%\AppData\go\bin"
""
echo "or someplace else in your %PATH%"
""
