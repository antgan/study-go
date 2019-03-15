package main

import (
	"fmt"
	"sync"
)

type worker struct {
	id int
	c  chan int
	wg *sync.WaitGroup
}

func createWorker(id int, wg *sync.WaitGroup) worker {
	c := make(chan int)
	return worker{id, c, wg}
}

func (worker worker) doWork() {
	for v := range worker.c {
		fmt.Printf("Worker %d do work: %c\n", worker.id, v)
		worker.wg.Done()
	}
	//for {
	//	v, ok := <-worker.c
	//	if !ok {
	//		break
	//	}
	//	fmt.Printf("Worker %d do work: %c\n", worker.id, v)
	//	worker.wg.Done()
	//}
}

func main() {
	var wg sync.WaitGroup
	var workers [10]worker
	//init worker
	for i := range workers {
		workers[i] = createWorker(i, &wg)
	}
	//Do work
	wg.Add(10)
	for _, worker := range workers {
		go worker.doWork()
	}
	//Send data
	for i, worker := range workers {
		worker.c <- 'a' + i
	}

	wg.Wait()
}
