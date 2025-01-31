package datatools

import (
	"bytes"
	"database/sql"
	"encoding/csv"
	"fmt"
	"os"
	"path"
	"testing"

	// SQL drivers
	_ "github.com/glebarez/go-sqlite"
)

func setupTestData(workDir string, fName string) (string, error) {
	if _, err := os.Stat(workDir); os.IsNotExist(err) {
		os.MkdirAll(workDir, 0775)
	}
	if _, err := os.Stat(fName); err == nil {
		// Remove the stale test data if necessary
		os.Remove(fName)
	}
	dsn := fmt.Sprintf("file:%s", fName)
	db, err := sql.Open("sqlite", dsn)
	if err != nil {
		return dsn, err
	}
	defer db.Close()
	stmt := `CREATE TABLE data (
id INTEGER PRIMARY KEY,
val VARCHAR(256),
created DATETIME DEFAULT CURRENT_TIMESTAMP)`
	_, err = db.Exec(stmt)
	if err != nil {
		return dsn, err
	}

	stmt = `INSERT INTO data (id, val) VALUES (?, ?)`
	for i := 0; i < 10; i++ {
		val := fmt.Sprintf("v%06d", i)
		if _, err := db.Exec(stmt, i, val); err != nil {
			return dsn, err
		}
	}
	return fmt.Sprintf("sqlite://%s", dsn), nil
}

func TestSQLQueryToCSV(t *testing.T) {
	// Setup a test SQLite 3 database for conversion
	workDir, _ := os.Getwd()
	workDir = path.Join(workDir, "testout")
	fName := path.Join(workDir, "test.sqlite3")
	dsnURL, err := setupTestData(workDir, fName)
	if err != nil {
		t.Error(err)
		t.FailNow()
	}
	store, err := OpenSQLStore(dsnURL)
	if err != nil {
		t.Error(err)
		t.FailNow()
	}
	defer (func() {
		if err := store.Close(); err != nil {
			t.Error(err)
		}
	})()

	// Here's the test SQL query string.
	stmt := "SELECT id, val, created FROM data ORDER BY id"
	// Setup the CSV writer so we're writing to memory.
	buf := bytes.NewBuffer([]byte{})
	w := csv.NewWriter(buf)
	w.UseCRLF = false
	// Check the results written
	if err := store.QueryToCSV(w, stmt); err != nil {
		t.Error(err)
		t.FailNow()
	}
	// Write any buffered data to the underlying writer (standard output).
	w.Flush()
	if err := w.Error(); err != nil {
		t.Error(err)
		t.FailNow()
	}
}
