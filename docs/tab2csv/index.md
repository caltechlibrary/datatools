
USAGE

tab2csv is a simple conversion utility to convert from tabs to quoted CSV.
tab2csv reads from standard input and writes to standard out.


If my.tab contained

    name	email	age
	Doe, Jane	jane.doe@example.org	42

Concert this to a CSV file format

    tab2csv < my.tab 

This would yield

    "name","email","age"
	"Doe, Jane","jane.doe@example.org",42

