package entities

import (
	"database/sql"

	_ "github.com/go-sql-driver/go-sqlite3"
)

var mydb *sql.DB

const dbPath string = "../DateBase.db"

func init() {
	//https://stackoverflow.com/questions/45040319/unsupported-scan-storing-driver-value-type-uint8-into-type-time-time
	db, err := sql.Open("mysql", "root:root@tcp(127.0.0.1:3306)/test?charset=utf8&parseTime=true")
	if err != nil {
		panic(err)
	}
	mydb = db
}

// SQLExecer interface for supporting sql.DB and sql.Tx to do sql statement
type SQLExecer interface {
	Exec(query string, args ...interface{}) (sql.Result, error)
	Prepare(query string) (*sql.Stmt, error)
	Query(query string, args ...interface{}) (*sql.Rows, error)
	QueryRow(query string, args ...interface{}) *sql.Row
}

// DaoSource Data Access Object Source
type DataAccessObject struct {
	// if DB, each statement execute sql with random conn.
	// if Tx, all statements use the same conn as the Tx's connection
	SQLExecutor
}

func CheckErr(err error) {
	if err != nil {
		panic(err)
	}
}

func initTables() {
	_, err := mydb.Exec(`
        CREATE TABLE users(
            id INTEGER PRIMARY KEY,
            key TEXT,
            username TEXT,
            password TEXT,
            email TEXT,
            phone TEXT
        );
        CREATE TABLE meetings(
            id INTEGER PRIMARY KEY,
						owner TEXT,
            title TEXT,
            members TEXT,
            starttime TEXT,
            endtime TEXT
        );
    `)
	CheckErr(err)
}
