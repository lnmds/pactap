package main

import (
    "log"
    "database/sql"
    _ "github.com/mattn/go-sqlite3"
)

/* open a repo db file */
func repoDatabaseStart (dbpath string) *sql.DB {
    db, err := sql.Open("sqlite3", dbpath)

    if err != nil {
        log.Fatal(err)
        panic(err)
    }

    return db
}

func closeDb(db *sql.DB) {
    db.Close()
}

