package main

import (
	"fmt"
	"sort"
	"strings"
)

func main() {
	employees := map[string]int{
		"Ana Popescu":      1,
		"Mihai Ionescu":    2,
		"Bogdan Radu":      1,
		"Elena Dumitrescu": 3,
		"Cristian Vasile":  2,
		"Ioana Enache":     1,
	}

	departments := make(map[int][]string)
	for name, id := range employees {
		departments[id] = append(departments[id], name)
	}

	ids := make([]int, 0, len(departments))
	for id := range departments {
		ids = append(ids, id)
	}
	sort.Ints(ids) // map order is random, so sort for a stable output

	for _, id := range ids {
		sort.Strings(departments[id])
		fmt.Printf("Department %d: [%s]\n", id, strings.Join(departments[id], ", "))
	}
}