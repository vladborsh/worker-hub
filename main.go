package main

import (
	"io"
	"log"
	"net/http"
)

func main() {
	http.HandleFunc("/job", ProcessJob)
	log.Fatal(http.ListenAndServe(":12345", nil))
}
