package main

import (
	"net/http"
)

func main() {
	http.HandleFunc("/dogPic.jpg", sendPic)
	http.ListenAndServe(":8080", nil)
}

func sendPic(w http.ResponseWriter, req *http.Request) {
	// serve file with name directly into the writer
	http.ServeFile(w, req, "dogPic.jpg")
}
