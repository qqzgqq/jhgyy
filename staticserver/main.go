package main

import (
	"net/http"
)

func main() {

	//static file handler.
	http.Handle("/staticfile/", http.StripPrefix("/staticfile/", http.FileServer(http.Dir("./staticfile"))))

	//Listen on port 8080
	http.ListenAndServe(":8080", nil)
}
