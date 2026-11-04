package main

import "fmt"

// send puts a message on the channel.
func send(ch chan<- string) {
	ch <- "hello"
}

func receive(ch <-chan string) string {
	return <-ch
}

func main() {
	ch := make(chan string)
	result := make(chan string)

	go send(ch)

	go func() {
		result <- receive(ch)
	}()

	fmt.Println("main printed:", <-result)
}