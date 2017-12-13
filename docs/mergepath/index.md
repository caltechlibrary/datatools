
# USAGE

	mergepath [OPTIONS] NEW_PATH_PARTS

## SYNOPSIS


mergepath can merge the new path parts with the existing path with creating duplications.
It can also re-order existing path elements by prefixing or appending to the existing
path and removing the resulting duplicate.


## OPTIONS

```
    -a, -append               Append the directory to the path removing any duplication
    -c, -clip                 Remove a directory from the path
    -d, -directory            The directory you want to add to the path.
    -e, -envpath              The path you want to merge with.
    -examples                 display example(s)
    -generate-markdown-docs   generate markdown documentation
    -h, -help                 display help
    -l, -license              display license
    -nl, -newline             if true add a trailing newline
    -p, -prepend              Prepend the directory to the path removing any duplication
    -quiet                    suppress error messages
    -v, -version              display version
```


## EXAMPLES


This would put your home bin directory at the beginning of your path.

	export PATH=$(mergepath -p $HOME/bin)


mergepath v0.0.23-pre
