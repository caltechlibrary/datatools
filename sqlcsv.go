package datatools

import (
	"database/sql"
	"encoding/csv"
	"fmt"
	"os"
	//"path"
	"strings"

	// Database specific drivers
	_ "github.com/glebarez/go-sqlite"
	_ "github.com/go-sql-driver/mysql"
	_ "github.com/lib/pq"
)

// SQLSrouce represents a wrapper SQL database drivers
// using a common struct.
type SQLStore struct {
	// Protocol holds the database type string, e.g. mysql, sqlite, pg
	Protocol string
	// Host name of service where to connect
	Host string
	// Port of service
	Port string
	// Database name you're going to query against
	Database string
	// User name for access a database service
	User string
	// Password for accessing a database service
	Password string

	// FIXME: need to have the CSV encoding options for writing
	// result of query
	WriteHeaderRow bool

	// workpath is the working directory to use when accessing SQLite3
	// related database paths
	workPath string

	// driverName is the database driver is the type of database we're accessing
	driverName string

	// the data source name
	dsn string

	// The db handle of the opened connection
	db *sql.DB
}

func dsnFixUp(driverName string, dsn string, workPath string) string {
	switch driverName {
	case "postgres":
		return fmt.Sprintf("%s://%s", driverName, dsn)
	case "sqlite":
		// NOTE: the db needs to be stored in the dataset directory
		// to keep the dataset easily movable.
		//dbName := path.Base(dsn)
		return dsn
		//path.Join(workPath, dbName)
	}
	return dsn
}

// OpenSQLStore opens a mysql, postgres or SQLite database
// based on a data source name expressed as a URL.
// The URL is formed by using the "protocol" to identify
// the service (e.g. "mysql://", "sqlite3://", "pg://")
// followed by a data source name per golang sql package
// documentation.
func OpenSQLStore(dsnURL string) (*SQLStore, error) {
	if !strings.Contains(dsnURL, "://") {
		return nil, fmt.Errorf("missing protocol in url scheme")
	}
	driverName, dsn, ok := strings.Cut(dsnURL, "://")
	if !ok {
		return nil, fmt.Errorf("could not parse DSN URI, got %q", dsnURL)
	}
	fmt.Printf("DEBUG driverName %q, dsn %q\n", driverName, dsn)
	var err error
	store := new(SQLStore)
	store.driverName = driverName
	store.workPath, err = os.Getwd()
	if err != nil {
		return nil, err
	}
	store.dsn = dsnFixUp(driverName, dsn, store.workPath)

	db, err := sql.Open(store.driverName, store.dsn)
	if err != nil {
		return nil, err
	}
	store.db = db
	store.WriteHeaderRow = true
	return store, nil
}

// Close the previously openned database resource
func (store *SQLStore) Close() error {
	if store.db != nil {
		return store.db.Close()
	}
	return nil
}

// QueryToCSV runs a SQL query statement and returns to the results
// CSV encoded via an io.Writer
func (store *SQLStore) QueryToCSV(out *csv.Writer, stmt string) error {
	fmt.Printf("DEBUG trying %q\n", stmt)
	rows, err := store.db.Query(stmt)
	if err != nil {
		fmt.Printf("DEBUG store -> %+v\n", store)
		fmt.Printf("DEBUG erorr from stmt %q -> %s\n", stmt, err)
		return err
	}
	defer rows.Close()
	fmt.Printf("DEBUG rows queries with %q\n", stmt)
	columns, err := rows.Columns()
	if err != nil {
		return err
	}
	fmt.Printf("DEBUG column names -> %s\n", strings.Join(columns, ", "))
	// Write out our header row is configurable
	if store.WriteHeaderRow {
		if err := out.Write(columns); err != nil {
			return err
		}
	}
	// Make an array of cells
	cells := make([]string, len(columns))
	for rows.Next() {
		// Retrieve the raw column data from the row
		vals := make([]string, len(columns))
		if err := rows.Scan(vals...); err != nil {
			return nil
		}
		fmt.Printf("DEBUG vals -> %+v\n", vals)
		// Convert values to strings for CSV write
		for i := 0; i < len(columns); i++ {
			switch vals[i].(type) {
			case string:
				cells[i] = vals[i].(string)
			default:
				cells[i] = fmt.Sprintf("%+v", vals[i])
			}
		}
		if err := out.Write(cells); err != nil {
			return err
		}
	}
	out.Flush()
	if err := rows.Err(); err != nil {
		return err
	}
	if err := out.Error(); err != nil {
		return err
	}
	return nil
}
