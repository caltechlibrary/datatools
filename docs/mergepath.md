
# USAGE

## mergepath NEW_PATH_PARTS

## SYNOPSIS

mergepath can merge the new path parts with the existing path with creating duplications.
It can also re-order existing path elements by prefixing or appending to the existing
path and removing the resulting duplicate.

## OPTIONS

```
	-a	Append the directory to the path removing any duplication
	-append	Append the directory to the path removing any duplication
	-c	Remove a directory from the path
	-clip	Remove a directory from the path
	-d	The directory you want to add to the path.
	-directory	The directory you want to add to the path.
	-e	The path you want to merge with.
	-envpath	The path you want to merge with.
	-example	display example(s)
	-h	display help
	-help	display help
	-l	display license
	-license	display license
	-p	Prepend the directory to the path removing any duplication
	-prepend	Prepend the directory to the path removing any duplication
	-v	display version
	-version	display version
```

## EXAMPLE

This would put your home bin directory at the beginning of your path.

```
	export PATH=$(mergepath -p $HOME/bin)
```

mergepath v0.0.18
