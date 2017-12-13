
# USAGE

	findfile [OPTIONS] [TARGET] [DIRECTORIES_TO_SEARCH]

## SYNOPSIS


findfile finds files based on matching prefix, suffix or contained text in base filename.


## OPTIONS

```
    -c, -contains             find file(s) based on basename containing text
    -d, -depth                Limit depth of directories walked
    -error, -stop-on-error    Stop walk on file system errors (e.g. permissions)
    -examples                 display example(s)
    -f, -full-path            list full path for files found
    -generate-markdown-docs   generate markdown documentation
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


Search the current directory and subdirectories for Markdown files with extension of ".md".

	findfile -s .md


findfile v0.0.22-pre
