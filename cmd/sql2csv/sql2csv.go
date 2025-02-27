package main

import (
	"encoding/csv"
	"encoding/json"
	"flag"
	"fmt"
	"os"
	"path"
	"runtime"

	// Caltech Library Packages
	"github.com/caltechlibrary/datatools"
)

const (
	helpText = `%{app_name}(1) user manual | version {version} {release_hash}
% R. S. Doiel
% {release_date}

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
at <https://go.dev/wiki/SQLInterface>. For the specificly supported
datatabase connection strings see
<https://pkg.go.dev/github.com/glebarez/go-sqlite>,
<https://pkg.go.dev/github.com/go-sql-driver/mysql#readme-dsn-data-source-name>
and <https://pkg.go.dev/github.com/lib/pq>


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

-use-crlf, -crlf
: Force the line ending per row to carage return and
line feed if true, false use line feed. Defaults to true
on Windows, false otherwise.

-sql FILENAME
: Read sql statement from a file instead of the command line.

# EXAMPLES

Using the "dbcfg.json" configuration file, display ten rows
from table "mytable" in database indicated in "dbcfg.json".

~~~sql
  {app_name} dbcfg.json 'SELECT * FROM mytable LIMIT 10'
~~~

The CSV output is written standard out and can be redirected into
a file if desired.

~~~shell
  {app_name} dbcfg.json 'SELECT * FROM mytable LIMIT 10' \
      >ten-rows.csv
~~~

Read SQL from a file and connect to Postgres without SSL you
can pass the `+"`"+`-sql`+"`"+` and `+"`"+`-dsn`+"`"+` options.

~~~shell
{app_name} \
  -dsn "postgres://${USER}@/${DB_NAME}?sslmode=disable" \
  -sql query.sql \
  >my_data.csv
~~~

`
)


func main() {
	// Program configuration defaults
	sqlCfg := new(datatools.SQLCfg)
	sqlCfg.Delimiter = ","
	sqlCfg.WriteHeaderRow = true
	sqlCfg.UseCRLF = (runtime.GOOS == "windows")

	// Running details
	appName := path.Base(os.Args[0])
	version := datatools.Version
	license := datatools.LicenseText
	releaseDate := datatools.ReleaseDate
	releaseHash := datatools.ReleaseHash

	showHelp, showLicense, showVersion := false, false, false
	writeHeaderRow, useCRLF := sqlCfg.WriteHeaderRow, sqlCfg.UseCRLF
	dsn, delimiter := "", ""
	fName, stmt := "", ""
	sqlFName := ""

	// Handle options
	flag.BoolVar(&showHelp, "help", showHelp, "display help")
	flag.BoolVar(&showLicense, "license", showLicense, "display license")
	flag.BoolVar(&showVersion, "version", showVersion, "display version")
	flag.BoolVar(&writeHeaderRow, "header-row", writeHeaderRow, "write a header row if true")
	flag.BoolVar(&useCRLF, "use-crlf", useCRLF, "delimited rows with a carriage return and line feed")
	flag.BoolVar(&useCRLF, "crlf", useCRLF, "delimited rows with a carriage return and line feed")
	flag.StringVar(&dsn, "dsn", dsn, "connect using the data source name provided in URL form")
	flag.StringVar(&delimiter, "delimiter", "", "set the delimiter, defaults to comma")
	flag.StringVar(&sqlFName, "sql", "", "read the SQL statement from a file, '-' will cause a read from standard input")
	flag.Parse()
	args := flag.Args()

	out := os.Stdout
	eout := os.Stderr

	// Handle help and information options
	if showHelp {
		fmt.Fprintf(out, "%s\n", datatools.FmtHelp(helpText, appName, version, releaseDate, releaseHash))
		os.Exit(0)
	}
	if showLicense {
		fmt.Fprintf(os.Stdout, "%s\n", license)
		os.Exit(0)
	}
	if showVersion {
		fmt.Fprintf(out, "datatools, %s %s %s\n", appName, version, releaseHash)
		os.Exit(0)
	}

	// Make sure the number of parameters makes sense
	switch len(args) {
	case 1:
		fName, stmt = "", args[0]
	case 2:
		fName, stmt = args[0], args[1]
	default:
		if sqlFName == "" {
			fmt.Fprintf(eout, "missing configuration and SQL query statement")
			os.Exit(1)
		}
	}

	// Read SQL from file if specified.
	if sqlFName != "" {
		src, err := os.ReadFile(sqlFName)
		if err != nil {
			fmt.Fprintf(eout, "%s\n", err)
			os.Exit(1)
		}
		stmt = fmt.Sprintf("%s", src)
	}

	// Load configuration if provided
	if fName != "" {
		src, err := os.ReadFile(fName)
		if err != nil {
			fmt.Fprintf(eout, "%s\n", err)
			os.Exit(1)
		}
		if err := json.Unmarshal(src, &sqlCfg); err != nil {
			fmt.Fprintf(eout, "%s\n", err)
		}
	}

	// Apply options to configuration (overriding the config file is appropriate)
	sqlCfg.WriteHeaderRow = writeHeaderRow
	// Normalize escaped character
	switch delimiter {
	case "\\t":
		delimiter = "\t"
	}

	if delimiter != "" {
		sqlCfg.Delimiter = delimiter
	}
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
		fmt.Fprintf(eout, "%s\n", err)
		os.Exit(1)
	}
	defer store.Close()

	store.WriteHeaderRow = sqlCfg.WriteHeaderRow
	if err := store.QueryToCSV(w, stmt); err != nil {
		w.Flush()
		fmt.Fprintf(eout, "%s\n", err)
		os.Exit(1)
	}
	w.Flush()
	if err := w.Error(); err != nil {
		fmt.Fprintf(eout, "%s\n", err)
		os.Exit(1)
	}
}
