package main

func ProcessJob(w http.ResponseWriter, req *http.Request) {
	io.WriteString(w, "hello, world!\n")
}
