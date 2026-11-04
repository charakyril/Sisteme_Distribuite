package main

import "fmt"

func main() {
	ch := make(chan string)

	go func() {
		ch <- "hello"
		ch <- "world"
	}()

	// main receives two messages on the same channel
	first := <-ch
	second := <-ch
	fmt.Println("first message: ", first)
	fmt.Println("second message:", second)
}