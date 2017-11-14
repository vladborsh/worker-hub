package main

import (
	"fmt"
)

type Worker struct {
	id           string
	workersQueue chan chan Job
	jobChannel   chan Job
	quitChannel  chan bool
}

func (w *Worker) Run() {
	workersQueue <- jobChannel
	for {
		select {
		case work := <-jobChannel:
			fmt.Println(job.Name)
		case <-quitChannel:
			fmt.Printf("Strop worker %d", w.id)
		}
	}
}
