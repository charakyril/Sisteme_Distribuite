package main

import (
	"fmt"
	"math/rand"
	"time"
)

type Point struct {
	X, Y int
}

func robot(name string, route []Point, out chan<- Point, in <-chan Point, report chan<- string) {
	received := make(chan int)

	go func() {
		count := 0
		for p := range in {
			fmt.Printf("%s received position %v\n", name, p)
			count++
		}
		received <- count
	}()

	for _, p := range route {
		out <- p
		time.Sleep(time.Duration(20+rand.Intn(60)) * time.Millisecond)
	}
	fmt.Printf("%s finished its route and stopped\n", name)
	close(out) 

	count := <-received
	report <- fmt.Sprintf("%s is done (received %d positions)", name, count)
}

func main() {
	aToB := make(chan Point)
	bToA := make(chan Point)
	report := make(chan string)

	routeA := []Point{{0, 0}, {1, 0}, {2, 1}, {3, 1}, {4, 2}}
	routeB := []Point{{5, 5}, {4, 5}, {3, 4}, {2, 4}, {1, 3}, {0, 3}, {0, 2}}

	fmt.Println("Observer: starting the robots")
	go robot("Robot A", routeA, aToB, bToA, report)
	go robot("Robot B", routeB, bToA, aToB, report)

	fmt.Println("Observer:", <-report)
	fmt.Println("Observer:", <-report)
	fmt.Println("Observer: both robots completed their tasks")
}
