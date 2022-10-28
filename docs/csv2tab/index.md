
USAGE

csv2tab is a simple conversion utility to convert from CSV to tab separated values.
csv2tab reads from standard input and writes to standard out.


If my.tab contained

    "name","email","age"
	"Doe, Jane","jane.doe@example.org",42

Concert this to a tab separated values

    csv2tab < my.csv 

This would yield

    name	email	age
	Doe, Jane	jane.doe@example.org	42

