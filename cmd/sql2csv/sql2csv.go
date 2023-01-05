package main

import (
	"encoding/csv"
	"encoding/json"
	"flag"
	"fmt"
	"os"
	"path"
	"strings"

	// Caltech Library Packages
	"github.com/caltechlibrary/datatools"
)

const (
	helpText = `% {app_name} (1) user manual
% R. S. Doiel
% 2023-01-03

# NAME

{app_name}

# SYNOPSIS

{app_name} [OPTIONS] SQL_STATEMENT

{app_name} [OPTIONS] CONFIG_FILE SQL_STATEMENT

# DESCRIPTION

{app_name} takes a config file describing a SQL database connection
and output options needed and a SQL statement as the final parameter.
The output of the SQL query is rendered in CSV format to standard
out. {app_name} supports querying MySQL 8, Postgres and SQLite3 
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

To connect with a database {app_name} relies on a data source name (DSN)
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

  {app_name} dbcfg.json 'SELECT * FROM mytable LIMIT 10'

The CSV output is written standard out and can be redirected into
a file if desired.

  {app_name} dbcfg.json 'SELECT * FROM mytable LIMIT 10' \
      >ten-rows.csv

{app_name} {version}
`
)

func fmtTxt(src string, name string, version string) string {
	return strings.ReplaceAll(strings.ReplaceAll(src, `{app_name}`, name), `{version}`, version)
}

func main() {
	// Program configuration defaults
	sqlCfg := new(datatools.SQLCfg)
	sqlCfg.Delimiter = ","
	sqlCfg.WriteHeaderRow = true
	sqlCfg.UseCRLF = false

	// Running details
	appName := path.Base(os.Args[0])
	showHelp, showLicense, showVersion := false, false, false
	writeHeaderRow, useCRLF := sqlCfg.WriteHeaderRow, sqlCfg.UseCRLF
	dsn, delimiter := "", sqlCfg.Delimiter
	fName, stmt := "", ""

	// Handle options
	flag.BoolVar(&showHelp, "help", showHelp, "display help")
	flag.BoolVar(&showLicense, "license", showLicense, "display license")
	flag.BoolVar(&showVersion, "version", showVersion, "display version")
	flag.BoolVar(&writeHeaderRow, "header-row", writeHeaderRow, "write a header row if true")
	flag.BoolVar(&useCRLF, "use-crlf", useCRLF, "delimited rows with a carriage return and line feed")
	flag.StringVar(&dsn, "dsn", dsn, "connect using the data source name provided in URL form")
	flag.StringVar(&delimiter, "delimiter", delimiter, "set the delimiter, defaults to comma")
	flag.Parse()
	args := flag.Args()

	// Handle help and information options
	if showHelp {
		fmt.Fprintf(os.Stdout, "%s", fmtTxt(helpText, appName, datatools.Version))
		os.Exit(0)
	}
	if showVersion {
		fmt.Fprintf(os.Stdout, "%s %s\n", appName, datatools.Version)
		os.Exit(0)
	}
	if showLicense {
		fmt.Fprintf(os.Stdout, "%s %s\n%s\n", appName, datatools.Version, fmtTxt(datatools.LicenseText, appName, datatools.Version))
		os.Exit(0)
	}

	// Make sure the number of parameters makes sense
	switch len(args) {
	case 1:
		fName, stmt = "", args[0]
	case 2:
		fName, stmt = args[0], args[1]
	default:
		fmt.Fprintf(os.Stderr, "missing configuration and SQL query statement")
		os.Exit(1)
	}

	// Load configuration if provided
	if fName != "" {
		src, err := os.ReadFile(fName)
		if err != nil {
			fmt.Fprintf(os.Stderr, "%s\n", err)
			os.Exit(1)
		}
		if err := json.Unmarshal(src, &sqlCfg); err != nil {
			fmt.Fprintf(os.Stderr, "%s\n", err)
		}
	}

	// Apply options to configuration (overriding the config file is appropriate)
	sqlCfg.WriteHeaderRow = writeHeaderRow
	// Normalize escaped character
	switch delimiter {
	case "\\t":
		delimiter = "\t"
	}
	sqlCfg.Delimiter = delimiter
	if dsn != "" {
		sqlCfg.DSN = dsn
	}
	sqlCfg.UseCRLF = useCRLF

	// Setup CSV writer
	w := csv.NewWriter(os.Stdout)
	w.Comma = []rune(sqlCfg.Delimiter)[0]
	w.UseCRLF = sqlCfg.UseCRLF

	// Open the SQL store
	store, err := datatools.OpenSQLStore(sqlCfg.DSN)
	if err != nil {
		fmt.Fprintf(os.Stderr, "%s\n", err)
		os.Exit(1)
	}
	defer store.Close()

	store.WriteHeaderRow = sqlCfg.WriteHeaderRow
	if err := store.QueryToCSV(w, stmt); err != nil {
		w.Flush()
		fmt.Fprintf(os.Stderr, "%s\n", err)
		os.Exit(1)
	}
	w.Flush()
	if err := w.Error(); err != nil {
		fmt.Fprintf(os.Stderr, "%s\n", err)
		os.Exit(1)
	}
}
