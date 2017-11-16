package main

import (
	"fmt"
)

type Worker struct {
	Id           string
	WorkersQueue chan chan Job
	JobChannel   chan Job
	QuitChannel  chan bool
}

func NewWorker(id int, workersQueue chan chan Job) Worker {
	return Worker{
		Id:           id,
		WorkersQueue: workersQueue,
		JobChannel:   make(chan Job),
		QuitChannel:  make(chan bool)}
}

func (w *Worker) Run() {
	go func() {
		for {
			WorkersQueue <- JobChannel
			select {
			case job := <-JobChannel:
				fmt.Println(job.Name)
			case <-quitChannel:
				fmt.Printf("Strop worker %d", w.Id)
			}
		}
	}()
}

func (w *Worker) Stop() {
	w.QuitChannel < true
}
