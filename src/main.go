package main

import (
	"log"
	"net/http"
	"os"
	"database/sql"
)

func main() {
	var db *sql.DB
	var err error

	//Initialize db if it doesn't exist
	if _, err = os.Stat("url.db"); os.IsNotExist(err) {
		db, err = sql.Open("sqlite3", "url.db")
		if err != nil {
			log.Fatalf("Failed to open DB: %v\n", err)
		}

		createTableQuery := `CREATE TABLE IF NOT EXISTS url (
			Id INTEGER PRIMARY KEY,
			Url TEXT NOT NULL
		);`
		_, err = db.Exec(createTableQuery)
		if err != nil {
			log.Fatalf("Failed to create the table: %v\n", err)
		}
	} else {	
	//Open the existing db
		db, err = sql.Open("sqlite3", "url.db")
		if err != nil {
			log.Fatalf("Failed to open DB: %v\n", err)
		}
	}
	defer db.Close()

	router := http.NewServeMux()
	router.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		htmlHandler(w, r)
	})
	router.HandleFunc("/api/", func(w http.ResponseWriter, r *http.Request) {
		apiHandler(w, r, db)
	})
	router.HandleFunc("/submit", func(w http.ResponseWriter, r *http.Request) {
		postHandler(w, r, db)
	})
	http.ListenAndServe(":8080", router)
}
