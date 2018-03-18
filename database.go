package main

import (
    "log"
    "database/sql"
    _ "github.com/mattn/go-sqlite3"
)

type Repository struct {
    db *sql.DB
}

/* open a repo db file */
func RepoOpen(c *Main, reponame string) Repository {
    dbpath := getRepoDBPath(c, reponame)
    db, err := sql.Open("sqlite3", dbpath)

    if err != nil {
        log.Fatal(err)
        panic(err)
    }

    return Repository{
        db: db,
    }
}

func FindPackage(repo Repository, pkg_name string) (Package, error) {
    db := repo.db
    stmt, err := db.Prepare("select * from packages where name=?")

    if err != nil {
        log.Fatal(err)
        return Package{}, err
    }

    defer stmt.Close()

    var name string
    var version string
    var build int
    err = stmt.QueryRow(pkg_name).Scan(&name, &version, &build)

    if err != nil {
        log.Fatal(err)
        return Package{}, err
    }

    return Package{
        name: name,
        version: version,
        build: build,
    }, nil
}

func CloseDb(db *sql.DB) {
    db.Close()
}

