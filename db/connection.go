package db

import (
	"github.com/gocraft/dbr"
	_ "github.com/mattn/go-sqlite3"
)

func OpenDB(dbPath string) (*dbr.Session, error) {
	if conn, err := dbr.Open("sqlite3", dbPath, nil); err != nil {
		return nil, err
	} else {
		return conn.NewSession(nil), nil
	}
}
