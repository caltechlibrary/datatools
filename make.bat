@echo off
REM This is a Windows 10 Batch file for building dataset command
REM from the command prompt.
REM
REM It requires: go version 1.16.6 or better and the cli for git installed
REM
go version
mkdir bin
echo off
echo Getting ready to build the datatools in bin
echo on
go build -o bin\codemeta2cff.exe cmd\codemeta2cff\codemeta2cff.go
go build -o bin\csv2json.exe cmd\csv2json\csv2json.go
go build -o bin\csv2mdtable.exe cmd\csv2mdtable\csv2mdtable.go
go build -o bin\csv2tab.exe cmd\csv2tab\csv2tab.go
go build -o bin\csv2xlsx.exe cmd\csv2xlsx\csv2xlsx.go
go build -o bin\csvcleaner.exe cmd\csvcleaner\csvcleaner.go
go build -o bin\csvcols.exe cmd\csvcols\csvcols.go
go build -o bin\csvfind.exe cmd\csvfind\csvfind.go
go build -o bin\csvjoin.exe cmd\csvjoin\csvjoin.go
go build -o bin\csvrows.exe cmd\csvrows\csvrows.go
go build -o bin\finddir.exe cmd\finddir\finddir.go
go build -o bin\findfile.exe cmd\findfile\findfile.go
go build -o bin\json2toml.exe cmd\json2toml\json2toml.go
go build -o bin\json2yaml.exe cmd\json2yaml\json2yaml.go
go build -o bin\jsoncols.exe cmd\jsoncols\jsoncols.go
go build -o bin\jsonjoin.exe cmd\jsonjoin\jsonjoin.go
go build -o bin\jsonmunge.exe cmd\jsonmunge\jsonmunge.go
go build -o bin\jsonrange.exe cmd\jsonrange\jsonrange.go
go build -o bin\mergepath.exe cmd\mergepath\mergepath.go
go build -o bin\range.exe cmd\range\range.go
go build -o bin\reldate.exe cmd\reldate\reldate.go
go build -o bin\reltime.exe cmd\reltime\reltime.go
go build -o bin\sql2csv.exe cmd\sql2csv\sql2csv.go
go build -o bin\string.exe cmd\string\string.go
go build -o bin\tab2csv.exe cmd\tab2csv\tab2csv.go
go build -o bin\timefmt.exe cmd\timefmt\timefmt.go
go build -o bin\toml2json.exe cmd\toml2json\toml2json.go
go build -o bin\urlparse.exe cmd\urlparse\urlparse.go
go build -o bin\xlsx2csv.exe cmd\xlsx2csv\xlsx2csv.go
go build -o bin\xlsx2json.exe cmd\xlsx2json\xlsx2json.go
go build -o bin\yaml2json.exe cmd\yaml2json\yaml2json.go
echo  off
echo Checking compile should see version number of dataset
echo on
bin\codemeta2cff.exe -version
bin\csv2json.exe -version
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
echo off
echo If OK, you can now copy the dataset.exe to %USERPROFILE%\go\bin in
echo off   
echo       copy bin\* %USERPROFILE%\AppData\go\bin
echo off
echo or someplace else in your PATH
echo off   
echo      %PATH%
echo on
