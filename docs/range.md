
# USAGE

## range [OPTIONS] START_INTEGER END_INTEGER [INCREMENT_INTEGER]

## SYNOPSIS

range is a simple utility for shell scripts that emits a list of 
integers starting with the first command line argument and 
ending with the last integer command line argument. It is a 
subset of functionality found in the Unix seq command.

If the first argument is greater than the last then it counts 
down otherwise it counts up.

## OPTIONS

```
    -e    The ending integer.
    -end    The ending integer.
    -example    display example(s)
    -h    display help
    -help    display help
    -i    The non-zero integer increment value.
    -increment    The non-zero integer increment value.
    -l    display license
    -license    display license
    -r    Pick a range value from range
    -random    Pick a range value from range
    -s    The starting integer.
    -start    The starting integer.
    -v    display version
    -version    display version
```

## EXAMPLES

Create a range of integers one through five

```shell
    range 1 5
```

Yields 1 2 3 4 5

Create a range of integer negative two to six

```shell
    range -- -2 6
```

Yields -2 -1 0 1 2 3 4 5 6

Create a range of even integers two to ten

```shell
    range -increment=2 2 10
```

Yields 2 4 6 8 10

Create a descending range of integers ten down to one

```shell
    range 10 1
```

Yields 10 9 8 7 6 5 4 3 2 1


Pick a random integer between zero and ten

```shell
    range -r 0 10
```

Yields a random integer from 0 to 10

range v0.0.18
