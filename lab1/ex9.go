package main

import "fmt"

func main() {
	arr := [...]int{2, 7, 11, 8, 4, 9}
	var s int
	fmt.Print("s = ")
	fmt.Scan(&s)

	// value -> indices where it has appeared so far
	seen := make(map[int][]int)
	for j, x := range arr {
		for _, i := range seen[s-x] { // has the complement appeared already?
			fmt.Println(i, j)
		}
		seen[x] = append(seen[x], j)
	}
}