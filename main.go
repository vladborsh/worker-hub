package main

import (
	"io"
	"log"
	"net/http"
)

func main() {
	prod := Producer{}
	disp := Dispatcher{}
	prod.Init()
	disp.Init(3, prod)

	http.HandleFunc("/job", prod.ProcessIncomingJob)
	log.Fatal(http.ListenAndServe(":8080", nil))
}
