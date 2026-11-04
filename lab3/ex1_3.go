package main

import (
	"fmt"
	"math/rand"
	"time"
)

// Stores the deviation from the real time
type Clock struct{ offset time.Duration }

func (c *Clock) Now() time.Time { return time.Now().Add(c.offset) }


func estimate(t1, t2, t3, t4 time.Time) (theta, delta time.Duration) {
	theta = (t2.Sub(t1) + t3.Sub(t4)) / 2
	delta = t4.Sub(t1) - t3.Sub(t2)
	return theta, delta
}

func measure(c_clock, s_clock *Clock) (theta, delta time.Duration) {
	time.Sleep(3*time.Millisecond + time.Duration(rand.Intn(10))*time.Millisecond)

	
	time.Sleep(5*time.Millisecond + time.Duration(rand.Intn(20))*time.Millisecond)
	
	time.Sleep(3*time.Millisecond + time.Duration(rand.Intn(10))*time.Millisecond)

	return estimate(t1, t2, t3, t4)
}

func abs(d time.Duration) time.Duration {
	if d < 0 {
		return -d
	}
	return d
}

func main() {
	const initialOffset = 500 * time.Millisecond // the client clock is 500 ms ahead
	server := &Clock{}                           // the server has the correct time

	// ---- (a) a single request/response ----
	fmt.Println("=== (a) single measurement ===")
	client := &Clock{offset: initialOffset}
	theta, delta := measure(client, server)
	fmt.Printf("theta = %v, delta = %v\n", theta, delta)

	before := abs(client.offset)
	client.offset += theta // correct the client clock
	after := abs(client.offset)
	fmt.Printf("client offset before: %v, after: %v\n", before, after)
	if after < before {
		fmt.Println("The offset decreased.")
	}

	// ---- (b) 10 repetitions with the same initial offset ----
	fmt.Println("\n=== (b) 10 measurements ===")
	const runs = 10
	thetas := make([]time.Duration, runs)
	deltas := make([]time.Duration, runs)
	for i := 0; i < runs; i++ {
		c := &Clock{offset: initialOffset} // same initial offset every time
		thetas[i], deltas[i] = measure(c, server)
		fmt.Printf("run %2d: theta = %v, delta = %v\n", i+1, thetas[i], deltas[i])
	}

	minTheta, minDelta := thetas[0], deltas[0]
	thetaAtMinDelta := thetas[0]
	for i := 1; i < runs; i++ {
		if thetas[i] < minTheta {
			minTheta = thetas[i]
		}
		if deltas[i] < minDelta {
			minDelta = deltas[i]
			thetaAtMinDelta = thetas[i] // the estimate with the smallest round trip is the most reliable
		}
	}
	fmt.Printf("\nminimum delta: %v\n", minDelta)
	fmt.Printf("minimum theta: %v\n", minTheta)
	fmt.Printf("theta from the run with minimum delta: %v\n", thetaAtMinDelta)

	// Comparison with (a): residual error after correcting the client clock
	fmt.Printf("\nresidual error in (a):                    %v\n", after)
	fmt.Printf("residual error using min-delta theta (b): %v\n", abs(initialOffset+thetaAtMinDelta))
}