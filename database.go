package main

import (
	"database/sql"
	_ "github.com/mattn/go-sqlite3"
	"log"
)

type Repository struct {
	db *sql.DB
}

func RepoOpen(c *Main, reponame string) (Repository, error) {
	dbpath := getRepoDBPath(c, reponame)

	log.Printf("Connecting to repo '%s'", reponame)
	db, err := sql.Open("sqlite3", dbpath)

	if err != nil {
		log.Fatal(err)
		return Repository{}, err
	}

	return Repository{
		db: db,
	}, nil
}

func FindPackage(repo Repository, pkgName string) (Package, error) {
	db := repo.db
	stmt, err := db.Prepare("select * from packages where name=?")

	if err != nil {
		log.Println(err)
		return Package{}, err
	}

	defer stmt.Close()

	var name string
	var version string
	var build int
	err = stmt.QueryRow(pkgName).Scan(&name, &version, &build)

	if err != nil {
		log.Println(err)
		return Package{}, err
	}

	return Package{
		name:    name,
		version: version,
		build:   build,
	}, nil
}
