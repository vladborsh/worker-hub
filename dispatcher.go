package main

type Dispatcher struct {
	WorkersQueue chan chan Job
}

func (d *Dispatcher) Init(numberWorkers int, prod Producer) {
	d.WorkersQueue = make(chan chan Job, numberWorkers)

	for i := 0; i < numberWorkers; i++ {
		worker := NewWorker(i, WorkersQueue)
		worker.Run()
	}

	go func() {
		for {
			select {
			case job := <-prod.Queue:
				worker := <-WorkersQueue
				worker.JobChannel <- job
			}
		}
	}()

}
