package main

import (
	"io"
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
	// copy the file into the writer, to send to the user
	io.Copy(w, f)
}
