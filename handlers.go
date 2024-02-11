package main

import (
	"net/http"
	"strconv"
	"fmt"
	"database/sql"
)


func htmlHandler(w http.ResponseWriter, r *http.Request) {
	html := `<!DOCTYPE html>
	<html>
		<head>
			<title>URL shortener</title>
		</head>
		<body>
			<p>Usage: [url-service]/api/<id></p>
			<form action="submit", method="POST">
				<fieldset>
					<legend>URL Shortener</legend>
					<label for="url-input">URL:</label>
					<input id="url-input" type="text" name="url" placeholder="https://www.google.com"/>
					<input type="submit" value="POST URL"/>
				</fieldset>
			</form>
		</body>
	</html>
	`
	fmt.Fprintf(w, html)
}

func apiHandler(w http.ResponseWriter, r *http.Request, db *sql.DB) { 
	endpoint := r.URL.Path[len("/api/"):]
	_, err := strconv.ParseInt(endpoint, 10, 64)
	if err != nil {
		http.Error(w, "Err: NaN", http.StatusBadRequest)
		return
	}

	var url string
	err = db.QueryRow("Select Url from url WHERE Id = $1", endpoint).Scan(&url)

	if err != nil {
		http.Error(w, "Err: Error retrieving url", http.StatusInternalServerError)
		return
	}

	formJson(w, endpoint, url)
}

func postHandler(w http.ResponseWriter, r *http.Request, db *sql.DB) {
	r.ParseForm()
	url := r.Form.Get("url")

	if !validUrl(url) {
		http.Error(w, "Err: Invalid URL", http.StatusBadRequest)
		return
	}
	
	existing, err := findExistingUrl(db, url)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	if existing == "" {
		id, err := insertUrl(db, url)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		formJson(w, strconv.FormatInt(id, 10), url)
	} else {
		formJson(w, existing, url)
	}
}

