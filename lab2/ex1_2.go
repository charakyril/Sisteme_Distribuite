package main

import "fmt"

func main() {
	ch := make(chan string)

	// A goroutine (not main) sends the message
	go func() {
		ch <- "hello"
	}()

	// main receives it; this blocks until the message arrives
	msg := <-ch
	fmt.Println("main received:", msg)
}