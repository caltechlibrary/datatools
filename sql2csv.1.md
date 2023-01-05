---
title: "sql2csv (1) user manual"
author: "R. S. Doiel"
pubDate: 2023-01-05
---

# NAME

sql2csv

# SYNOPSIS

sql2csv [OPTIONS] SQL_STATEMENT

sql2csv [OPTIONS] CONFIG_FILE SQL_STATEMENT

# DESCRIPTION

sql2csv takes a config file describing a SQL database connection
and output options needed and a SQL statement as the final parameter.
The output of the SQL query is rendered in CSV format to standard
out. sql2csv supports querying MySQL 8, Postgres and SQLite3 
databases.

The configuration file is a JSON document with the following 
key value pairs.

dsn_url
: (string) A data source name in URL form where the "protocol" element
identifies the database resource being accessed (e.g. "sqlite://",
"mysql://", "postgres://"). A data source name are rescribed
at <https://github.com/golang/go/wiki/SQLInterface>.

header_row
: (boolean) if true print a header row in the output, false for no 
header row output

delimiter
: (single character, default is ","), sets the field delimited used
in output. It can be set to "\t" for tab separated values.

use_crlf
: (boolean, default is false) if set to true to use "\r\n" as the
line terminator between rows of output. 

To connect with a database sql2csv relies on a data source name (DSN)
in URL format. In the URL form the URL's scheme indicates the type
of database you are connecting to (e.g. sqlite, mysql, postgres). The
rest of the DNS has the following form

~~~
[username[:password]@][protocol[(address)]]/dbname[?param1=value1&...&paramN=valueN]
~~~

For a simple database like SQLite3 a minimal DSN in url form 
for a database file "my_database.sqlite3" would look like

~~~
	sqlite://file:my_database.sqlite3
~~~

For MySQL you need to provide more information to connect (e.g. username,
password). In this example the username is "jane.doe", password is
"something_secret" the database is "my_database". (this example
assumes that MySQL 8 is running on localhost at the usual port).

~~~~
	mysql://jane.doe:something_secret@/my_database
~~~~

Postgres is similar to the MySQL connection string except the 
"scheme" is "postgres" instead of "mysql".


# OPTIONS

-help
: display help

-version
: display version

-license
: display license

A the following options will override a configuration.

-dsn
: use the data source name in URL form instead of a JSON 
configuration file

-header
: use a header row  if true, false skip the header row

-delimiter
: Set the delimiter to use, default is comma

-use-cdlf
: Force the line ending per row to carage return and
line feed if true, false use line feed

# EXAMPLE

Using the "dbcfg.json" configuration file, display ten rows
from table "mytable" in database indicated in "dbcfg.json".

  sql2csv dbcfg.json 'SELECT * FROM mytable LIMIT 10'

The CSV output is written standard out and can be redirected into
a file if desired.

  sql2csv dbcfg.json 'SELECT * FROM mytable LIMIT 10' \
      >ten-rows.csv

sql2csv 1.1.5
