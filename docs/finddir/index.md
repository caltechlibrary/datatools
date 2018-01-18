
# USAGE

	finddir [OPTIONS] [TARGET] [DIRECTORIES_TO_SEARCH]

## SYNOPSIS


finddir finds directory based on matching prefix, suffix or contained text in base filename.


## OPTIONS

```
    -c, -contains             find file(s) based on basename containing text
    -d, -depth                Limit depth of directories walked
    -e, -error-stop           Stop walk on file system errors (e.g. permissions)
    -examples                 display example(s)
    -f, -full-path            list full path for files found
    -generate-markdown-docs   generate markdown documemtations
    -h, -help                 display this help message
    -l, -license              display license information
    -m, -mod-time             display file modification time before the path
    -nl, -newline             if true add a trailing newline
    -o, -output               output filename
    -p, -prefix               find file(s) based on basename prefix
    -quiet                    suppress error messages
    -s, -suffix               find file(s) based on basename suffix
    -v, -version              display version message
```


## EXAMPLES


Find all subdirectories starting with "img".

	finddir -p img


finddir v0.0.24-pre
