package main

import (
	"fmt"
	"time"
)

func worker(id int, c chan int) {
	for v := range c {
		fmt.Printf("Worker %d receive %c\n", id, v)
	}
}

func createWorker(id int, buffer int) chan int {
	c := make(chan int, buffer)
	go worker(id, c)
	return c
}

func channelDemo() {
	var channels [10]chan int
	for i := 0; i < 10; i++ {
		channels[i] = createWorker(i, 0)
	}

	// Send data to channel
	for i := 0; i < 10; i++ {
		channels[i] <- 'a' + i
	}
	time.Sleep(time.Millisecond)
}

func bufferChannelDemo() {
	c := make(chan int, 3)
	go worker(0, c)

	c <- 'a'
	c <- 'b'
	c <- 'c'
	c <- 'd'

	time.Sleep(time.Millisecond)
}

func closeChannelDemo() {
	c := make(chan int, 3)
	go worker(1, c)
	c <- 'a'
	c <- 'b'
	c <- 'c'
	c <- 'd'
	close(c)
	time.Sleep(time.Millisecond)
}

func main() {
	channelDemo()
	bufferChannelDemo()
	closeChannelDemo()
}
