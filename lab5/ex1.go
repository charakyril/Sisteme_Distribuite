package main

import (
	"fmt"
	"math"
	"os"
)

const radius = 17.0       // communication radius
const max = 1000 // safety limit so the program cannot run forever
const eps = 0.0001       // two positions closer than this are "the same"


func distance(x1, y1, x2, y2 float64) float64 {
	dx := x1 - x2
	dy := y1 - y2
	return math.Sqrt(dx*dx + dy*dy)
}

func main() {

	file, err := os.Open("input.txt")
	if err != nil {
		fmt.Println("Cannot open input.txt")
		return
	}
	defer file.Close()

	var n int
	fmt.Fscan(file, &n)

	x := make([]float64, n)
	y := make([]float64, n)
	for i := 0; i < n; i++ {
		fmt.Fscan(file, &x[i], &y[i])
	}

	//Repeat Look-Compute-Move 
	for i := 1; i <= max; i++ {
		newX := make([]float64, n)
		newY := make([]float64, n)

		for j := 0; j < n; j++ {
			// Start with the robot's own position.
			sumX := x[j]
			sumY := y[j]
			count := 1

			// Look: add every neighbor (robot closer than the radius).
			for k := 0; k < n; k++ {
				if k != j && distance(x[j], y[j], x[k], y[k]) < radius {
					sumX += x[k]
					sumY += y[k]
					count++
				}
			}

			// Compute: the new position is the average.
			newX[j] = sumX / float64(count)
			newY[j] = sumY / float64(count)
		}

		// Move
		x = newX
		y = newY

		// Check the position
		together := true
		for i := 1; i < n; i++ {
			if distance(x[0], y[0], x[i], y[i]) > eps {
				together = false
			}
		}
		if together {
			fmt.Println("Rendezvous reached after", i, "iterations")
			fmt.Printf("Meeting point: (%.4f, %.4f)\n", x[0], y[0])
			return
		}
	}

	//the robots did not meet.
	fmt.Println("Rendezvous not reached. Final positions:")
	for i := 0; i < n; i++ {
		fmt.Printf("Robot %d: (%.4f, %.4f)\n", i, x[i], y[i])
	}
}
