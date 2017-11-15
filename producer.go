package main

import (
	"log"
	"net/http"
)

type Producer struct {
	Queue chan Job
}

func (h *Producer) Init() {
	h.Queue = make(chan Job, 100)
}

func (h *Producer) ProcessIncomingJob(w http.ResponseWriter, req *http.Request) {
	if r.Method == "POST" {
		// Now, we take the delay, and the person's name, and make a WorkRequest out of them.
		job := Job{Name: "name", Type: "standard"}

		// Push the work onto the queue.
		Queue <- job
		log.Println("Work request queued")
		return
	}
	return
}
