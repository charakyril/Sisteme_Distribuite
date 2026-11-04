package main

import "fmt"

func main() {
	ch := make(chan string)
	done := make(chan struct{})

	//goroutine 
	go func() {
		msg := <-ch
		fmt.Println("goroutine received:", msg)
		close(done) // tell main we are finished, so it does not exit early
	}()

	// main sends the message
	ch <- "hello"

	<-done
}