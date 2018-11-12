package main

import (
	"net/http"
	"os"
)

func main() {
	http.HandleFunc("/dogPic.jpg", sendPic)
	http.ListenAndServe(":8080", nil)
}

func sendPic(w http.ResponseWriter, req *http.Request) {
	// find the file in the filesystem with that path
	f, err := os.Open("dogPic.jpg")
	if err != nil {
		http.Error(w, "file not found", 404)
	}
	// don't close the connection until we have copied the file back into the writer
	defer f.Close()

	fi, err := f.Stat()

	if err != nil {
		http.Error(w, "file not found", 404)
		return
	}

	// serve file f into the response writer, with a name and the last mod time from Stat() (for the eTag in the cache)
	http.ServeContent(w, req, f.Name(), fi.ModTime(), f)
}
