package main

import (
	"fmt"
	"time"
)

func worker(done chan<- struct{}) {
	fmt.Println("worker: processing...")
	time.Sleep(500 * time.Millisecond) 
	fmt.Println("worker: finished")
	done <- struct{}{} 
}

func main() {
	done := make(chan struct{})
	go worker(done)

	<-done 
	fmt.Println("main: worker has finished")
}
