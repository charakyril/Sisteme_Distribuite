package main

import (
	"fmt"
	"math/rand"
	"time"
)

const (
	finishLine = 100 // meters
	stepSize   = 10  // meters per step
)

// robot runs in a straight line; each step takes a random amount of time.
// When it reaches the finish line it sends its name on the finish channel.
func robot(name string, finish chan<- string) {
	for pos := 0; pos < finishLine; pos += stepSize {
		time.Sleep(time.Duration(20+rand.Intn(60)) * time.Millisecond)
	}
	finish <- name
}

func main() {
	// Buffered (size 1) so the slower robot never blocks forever after the race
	finish1 := make(chan string, 1)
	finish2 := make(chan string, 1)

	fmt.Println("Observer: starting the race!")
	go robot("Robot 1", finish1)
	go robot("Robot 2", finish2)

	// select waits for whichever robot finishes first
	select {
	case winner := <-finish1:
		fmt.Println("Observer:", winner, "reached the finish line first!")
	case winner := <-finish2:
		fmt.Println("Observer:", winner, "reached the finish line first!")
	}
}