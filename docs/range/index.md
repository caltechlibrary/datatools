
# USAGE

	range [OPTIONS] START_INTEGER END_INTEGER [INCREMENT_INTEGER]

## DESCRIPTION


range is a simple utility for shell scripts that emits a list of 
integers starting with the first command line argument and 
ending with the last integer command line argument. It is a 
subset of functionality found in the Unix seq command.

If the first argument is greater than the last then it counts 
down otherwise it counts up.


## OPTIONS

Below are a set of options available.

```
    -e, -end             The ending integer.
    -examples            display example(s)
    -generate-manpage    generate man page
    -generate-markdown   generate markdown documentation
    -h, -help            display help
    -inc, -increment     The non-zero integer increment value.
    -l, -license         display license
    -nl, -newline        if true add a trailing newline
    -quiet               suppress error messages
    -random              Pick a range value from range
    -s, -start           The starting integer.
    -v, -version         display version
```


## EXAMPLES


Create a range of integers one through five

	range 1 5

Yields 1 2 3 4 5

Create a range of integer negative two to six

	range -- -2 6

Yields -2 -1 0 1 2 3 4 5 6

Create a range of even integers two to ten

	range -increment=2 2 10

Yields 2 4 6 8 10

Create a descending range of integers ten down to one

	range 10 1

Yields 10 9 8 7 6 5 4 3 2 1


Pick a random integer between zero and ten

	range -random 0 10

Yields a random integer from 0 to 10


range v0.0.25
