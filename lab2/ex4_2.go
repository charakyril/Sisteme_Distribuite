package main

import (
	"fmt"
	"time"
)

func worker(done chan<- struct{}) {
	fmt.Println("worker: processing...")
	time.Sleep(500 * time.Millisecond) // simulate work
	fmt.Println("worker: finished")
	done <- struct{}{} // notify main
}

func main() {
	done := make(chan struct{})
	go worker(done)

	<-done // blocks until the worker notifies us (no WaitGroup needed)
	fmt.Println("main: worker has finished")
}