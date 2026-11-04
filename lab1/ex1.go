
// ελεγχος αν ειναι πρωτος αριθμος
package main

import "fmt"

func main() {
	var n int
	fmt.Print("Enter a natural number: ")
	fmt.Scan(&n)

	isPrime := n >= 2
	for d := 2; d*d <= n; d++ {
		if n%d == 0 {
			isPrime = false
			break
		}
	}

	if isPrime {
		fmt.Printf("%d is prime\n", n)
	} else {
		fmt.Printf("%d is not prime\n", n)
	}
}

