
# USAGE

## finddir [OPTIONS] [TARGET] [DIRECTORIES_TO_SEARCH]

## SYNOPSIS

finddir finds directory based on matching prefix, suffix or contained text in base filename.

## OPTIONS

	-c	find file(s) based on basename containing text
	-contains	find file(s) based on basename containing text
	-d	Limit depth of directories walked
	-depth	Limit depth of directories walked
	-e	Stop walk on file system errors (e.g. permissions)
	-error-stop	Stop walk on file system errors (e.g. permissions)
	-f	list full path for files found
	-full-path	list full path for files found
	-h	display this help message
	-help	display this help message
	-l	display license information
	-license	display license information
	-m	display file modification time before the path
	-mod-time	display file modification time before the path
	-p	find file(s) based on basename prefix
	-prefix	find file(s) based on basename prefix
	-s	find file(s) based on basename suffix
	-suffix	find file(s) based on basename suffix
	-v	display version message
	-version	display version message

## EXAMPLE

```
	finddir -p img
```

Find all subdirectories starting with "img". 


finddir v0.0.15
