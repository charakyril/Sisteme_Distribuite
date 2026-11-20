package main

import (
	"bufio"
	"fmt"
	"math/rand"
	"os"
	"sort"
	"strings"
)

func diametru(adj [][]int) int {
	n := len(adj)
	diam := 0
	for src := 0; src < n; src++ {
		dist := make([]int, n)
		for i := range dist {
			dist[i] = -1
		}
		dist[src] = 0
		queue := []int{src}
		for len(queue) > 0 {
			u := queue[0]
			queue = queue[1:]
			for _, v := range adj[u] {
				if dist[v] == -1 {
					dist[v] = dist[u] + 1
					queue = append(queue, v)
				}
			}
		}
		for _, d := range dist {
			if d == -1 {
				return -1
			}
			if d > diam {
				diam = d
			}
		}
	}
	return diam
}

// Info αντιπροσωπεύει ένα ζεύγος (id, τιμή) που κρατιέται στη μνήμη ενός κόμβου.
type Info struct {
	ID    int
	Value int
}

// Memory είναι ένα σύνολο από tokens (η ένωση γίνεται με εισαγωγή).
type Memory map[Info]struct{}

type Node struct {
	ID        int
	Neighbors []int
	M         Memory
}

func main() {
	// ---- Ανάγνωση εισόδου ----
	sc := bufio.NewScanner(os.Stdin)
	var lines []string
	for sc.Scan() {
		lines = append(lines, sc.Text())
	}

	var n, s int
	fmt.Sscan(lines[0], &n, &s)

	tokens := make([]int, n)
	for i, f := range strings.Fields(lines[1]) {
		fmt.Sscan(f, &tokens[i])
	}

	adj := make([][]int, n)
	for i := 0; i < n; i++ {
		for _, f := range strings.Fields(lines[2+i]) {
			var v int
			fmt.Sscan(f, &v)
			if v >= 0 && v < n { // αγνοούμε άκυρους γείτονες (π.χ. το "33")
				adj[i] = append(adj[i], v)
			}
		}
	}

	// ---- Διάμετρος και αριθμός γύρων ----
	D := diametru(adj)
	if D == -1 {
		fmt.Println("Graph not connected")
		return
	}
	rounds := (s + 1) * D

	// ---- Δημιουργία κόμβων: M_i(0) = {(i, v_i)} ----
	nodes := make([]Node, n)
	crashed := make([]bool, n)
	for i := 0; i < n; i++ {
		nodes[i] = Node{
			ID:        i,
			Neighbors: adj[i],
			M:         Memory{Info{i, tokens[i]}: {}},
		}
	}

	// rounds
	crashes := 0
	for r := 1; r <= rounds; r++ {
		// copy of memory in the beginnning
		snap := make([]Memory, n)
		for i := 0; i < n; i++ {
			snap[i] = Memory{}
			for info := range nodes[i].M {
				snap[i][info] = struct{}{}
			}
		}

		// random decision for crashing 
		crashing := -1
		if crashes < s && rand.Intn(2) == 0 {
			var alive []int
			for i := 0; i < n; i++ {
				if !crashed[i] {
					alive = append(alive, i)
				}
			}
			if len(alive) > 1 { // always stays one 
				crashing = alive[rand.Intn(len(alive))]
			}
		}

		//sent 
		for u := 0; u < n; u++ {
			if crashed[u] {
				continue
			}
			for _, v := range nodes[u].Neighbors {
				if u == crashing && rand.Intn(2) == 0 {
					continue 
				}
				if crashed[v] || v == crashing {
					continue
				}
				for info := range snap[u] {
					nodes[v].M[info] = struct{}{} // union
				}
			}
		}

		if crashing != -1 {
			crashed[crashing] = true
			crashes++
		}
	}

	correct := 0
	for i := 0; i < n; i++ {
		if !crashed[i] {
			correct = i
			break
		}
	}
	var res []Info
	for info := range nodes[correct].M {
		res = append(res, info)
	}
	sort.Slice(res, func(a, b int) bool { return res[a].ID < res[b].ID })

	fmt.Println("diameter", D)
	fmt.Println("loop before end ", rounds)
	for _, p := range res {
		fmt.Printf("(%d, %d) ", p.ID, p.Value)
	}
	fmt.Println()
}
