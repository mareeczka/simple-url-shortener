package main

import (
	"encoding/json"
	"net/http"
	"regexp"
)

func sendJson(w http.ResponseWriter, s interface{}) { //Used by formJson
	jsonData, err := json.Marshal(s)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.Write(jsonData)
}

func formJson(w http.ResponseWriter, id string, url string) { //Forms JSON from entry. Used by postHandler
	entry := UrlResponse {
		Url: url,
		Id: id,
	}
	sendJson(w, entry)
}

func validUrl(u string) bool { //Is the URL valid? Used by postHandler.
	pattern := regexp.MustCompile(`^[a-zA-Z][\w+.-]*:(\/\/[^\/\s]+(\/[^?#\s]*)?)?(\?[^#\s]*)?(#.*)?$`)
	return pattern.MatchString(u)
}