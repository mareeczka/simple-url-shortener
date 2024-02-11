package main

import (
	"database/sql"

	_ "github.com/mattn/go-sqlite3"
)

func findExistingUrl(db *sql.DB, url string) (string, error) { //Find URL in DB. Used by postHandler
	var existing string
	err := db.QueryRow("SELECT Id from url WHERE Url = ?;", url).Scan(&existing)
	if err == sql.ErrNoRows {
		return "", nil
	}
	return existing, err
}

func insertUrl(db *sql.DB, url string) (int64, error) { //Insert URL to DB. Used by postHandler
		stmt, err := db.Prepare("INSERT INTO url(Url) VALUES(?)")
		if err != nil {
		     return 0, err
	    	}
		defer stmt.Close()
	
		res, err := stmt.Exec(url)
		if err != nil {
			return 0, err
		}
	
		return res.LastInsertId()
}
