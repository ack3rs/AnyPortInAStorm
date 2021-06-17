package main

import (
	"net/http"

	"github.com/acky666/AnyPortInAStorm/Client/Upload"
	l "github.com/acky666/ackyLog"
)

func main() {

	l.INFO("Starting Client - Please go to http://localhost:8080")
	http.HandleFunc("/upload", Upload.HandleUploadedFile)
	http.HandleFunc("/", Upload.HandleForm)
	http.ListenAndServe(":8080", nil)

}
