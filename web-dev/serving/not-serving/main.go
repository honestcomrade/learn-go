package main

import (
	"io"
	"net/http"
)

func main() {
	http.HandleFunc("/", dog)
	http.ListenAndServe(":8080", nil)
}

func dog(w http.ResponseWriter, req *http.Request) {
	w.Header().Set("content-Type", "text/html; charset=utf-8")
	// literaly pipe raw html directly to the input output string
	// we are not serving becuase we are actually pointing the browser to link on some other server
	io.WriteString(w, `
    <img src="https://upload.wikimedia.org/wikipedia/commons/d/d9/Collage_of_Nine_Dogs.jpg">
  `)
}
