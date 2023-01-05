package datatools

import (
	"database/sql"
	"encoding/csv"
	"fmt"
	"os"
	"strings"

	// 3rd Party packages
	//sql "github.com/jmoiron/sqlx"

	// Database specific drivers
	_ "github.com/glebarez/go-sqlite"
	_ "github.com/go-sql-driver/mysql"
	_ "github.com/lib/pq"
)

// SQLCfg holds the information for connecting to
// a SQLStore and options for the CSV output.
type SQLCfg struct {
	DSN            string `json:"dsn_url,omitempty"`
	WriteHeaderRow bool   `json:"header_row,omitempty"`
	Delimiter      string `json:"delimiter,omitempty"`
	UseCRLF        bool   `json:"use_crlf,omitempty"`
}

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

	// WriteHeaderRow tracks desired behavior about generating
	// a header row in the CSV encoded output. NOTE: using OpenSQLStore()
	// sets this value to true.
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
	rows, err := store.db.Query(stmt)
	if err != nil {
		return err
	}
	defer rows.Close()
	columns, err := rows.Columns()
	if err != nil {
		return err
	}
	columnTypes, err := rows.ColumnTypes()
	if err != nil {
		return err
	}
	// Write out our header row is configurable
	if store.WriteHeaderRow {
		if err := out.Write(columns); err != nil {
			return err
		}
	}
	// Make an array of cells
	cells := make([]string, len(columnTypes))
	vals := make([]interface{}, len(columnTypes))
	for i := 0; i < len(columnTypes); i++ {
		vals[i] = new(interface{})
	}
	for rows.Next() {
		// Retrieve the raw column data from the row
		err := rows.Scan(vals[:]...)
		if err != nil {
			return nil
		}
		for i := 0; i < len(columnTypes); i++ {
			val := *vals[i].(*interface{})
			if val == nil {
				// FIXME: this should be configurable as 'NULL' or empty
				// string.
				cells[i] = "NULL"
			} else {
				switch val.(type) {
				case []byte:
					s := fmt.Sprintf("%s", val.([]byte))
					cells[i] = s
				case string:
					s := val.(string)
					cells[i] = s
				case bool:
					x := val.(bool)
					cells[i] = fmt.Sprintf("%T", x)
				case float32:
					x := val.(float32)
					cells[i] = fmt.Sprintf("%f", x)
				case float64:
					x := val.(float64)
					cells[i] = fmt.Sprintf("%f", x)
				case int:
					x := val.(int)
					cells[i] = fmt.Sprintf("%d", x)
				case int8:
					x := val.(int8)
					cells[i] = fmt.Sprintf("%d", x)
				case int16:
					x := val.(int16)
					cells[i] = fmt.Sprintf("%d", x)
				case int32:
					x := val.(int32)
					cells[i] = fmt.Sprintf("%d", x)
				case int64:
					x := val.(int64)
					cells[i] = fmt.Sprintf("%d", x)
				default:
					cells[i] = fmt.Sprintf("%+v", val)
				}
			}
		}
		if err := out.Write(cells); err != nil {
			return err
		}
	}
	out.Flush()
	if err := out.Error(); err != nil {
		return err
	}
	if err := rows.Err(); err != nil {
		return err
	}
	return nil
}
