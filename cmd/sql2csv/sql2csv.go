package main

import (
	"fmt"
	"database/sql"
	"os"
	"path"

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

{app_name} DB_CONNECT_FILE DATABASE_NAME SQL_STATEMENT

# DESCRIPTION

{app_name} takes a config file describing a SQL database connection 
and runs the SQL statement provided as the final parameter. The
output is rendered in CSV format.

# OPTIONS

-help
: display help

-version
: display version

-license
: display license

# EXAMPLE

Using the ".my.cnf" configuration file, display ten rows
from table "mytable" in "mydata" database.

  {app_name} .my.cnf mydata 'SELECT * FROM mytable LIMIT 10'

The CSV output is written standard out and can be redirected into
a file if desired.

  {app_name} .my.cnf mydata 'SELECT * FROM mytable LIMIT 10' \
      >ten-rows.csv

{app_name} {version}
`

)

func fmtTxt(src string, name string, version string) string {
	return strings.ReplaceAll(strings.ReplaceAll(src, `{app_name}`, name), `{version}`, version)
}

func main() {

	appName := path.Base(os.Argv[0])
	showHelp, showLicense, showVersion := false, false, false

	flag.BoolVar(&showHelp, "help", showHelp, "display help")
	flag.BoolVar(&showLicense, "license", showLicense, "display license")
	flag.BoolVar(&showVersion, "version", showVersion, "display version")
	flag.Parse()
	args := flag.Args()

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

	if len(args) != 3 {
		fmt.Fprintln(os.Stderr, "expected DB_CONF_FILE DB_NAME SQL_QUERY")
		os.Exit(1)
	}
	// FIXME: Open DB connection using the config file and DB name
	// Create and execute a the SQL query provided on the command line
	// Loop through the returned rows and output each row CSV encoded
	//   First row should be used to build render header row

}
