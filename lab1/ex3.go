//άριοι αριθμόι ακολουθιας
package main

import "fmt"

func main() {
	numbers := []int{3, 8, 12, 7, 20, 5, 6}
	var evens []int
	for _, x := range numbers {
		if x%2 == 0 {
			evens = append(evens, x)
		}
	}
	fmt.Println(evens)
}