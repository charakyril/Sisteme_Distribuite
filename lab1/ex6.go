//φθίνουσα σειρά κατα ονόματα 
package main

import (
	"fmt"
	"sort"
)

func main() {
	heights := map[string]int{
		"Andrei": 190,
		"Maria":  165,
		"Ion":    182,
		"Elena":  172,
		"George": 178,
		"Ana":    165,
	}

	names := make([]string, 0, len(heights))
	for n := range heights {
		names = append(names, n)
	}
	// Sort names by height, descending
	sort.Slice(names, func(i, j int) bool {
		return heights[names[i]] > heights[names[j]]
	})

	for _, n := range names {
		fmt.Println(n, heights[n])
	}
}