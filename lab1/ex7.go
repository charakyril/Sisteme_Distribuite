package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
)

func main() {
	file, err := os.Open("words.txt")
	if err != nil {
		fmt.Println(err)
		return
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	scanner.Split(bufio.ScanWords) // read word by word

	var words []string
	for scanner.Scan() {
		words = append(words, scanner.Text())
	}
	if err = scanner.Err(); err != nil {
		fmt.Println("Read error:", err)
		return
	}

	counts := make(map[string]int)
	for _, w := range words {
		counts[w]++
	}

	// Sort keys only to get a deterministic output order
	keys := make([]string, 0, len(counts))
	for w := range counts {
		keys = append(keys, w)
	}
	sort.Strings(keys)
	for _, w := range keys {
		fmt.Printf("%s: %d\n", w, counts[w])
	}
}