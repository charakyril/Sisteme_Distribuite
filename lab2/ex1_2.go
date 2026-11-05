package main

import "fmt"

func main() {
	ch := make(chan string)

	// A goroutine sends the message
	go func() {
		ch <- "hello"
	}()

	// main receives it
	msg := <-ch
	fmt.Println("main received:", msg)
}
