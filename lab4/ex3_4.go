package main

import (
	"errors"
	"fmt"
	"sync"
)

//after flooding phase.
type Result struct {
	Known    int     // how many votes the node collected
	Average  float64 //of all vates 
	Decision string
}

func Connected(n int, adj [][]int) bool {
	reach := func(g [][]int) bool {
		seen := make([]bool, n)
		stack := []int{0}
		seen[0] = true
		count := 1
		for len(stack) > 0 {
			u := stack[len(stack)-1]
			stack = stack[:len(stack)-1]
			for _, v := range g[u] {
				if !seen[v] {
					seen[v] = true
					count++
					stack = append(stack, v)
				}
			}
		}
		return count == n
	}
	rev := make([][]int, n)
	for u := range adj {
		for _, v := range adj[u] {
			rev[v] = append(rev[v], u)
		}
	}
	return reach(adj) && reach(rev)
}

func copyVotes(m map[int]int) map[int]int {
	c := make(map[int]int, len(m))
	for k, v := range m {
		c[k] = v
	}
	return c
}

func majorityNode(idx, vote int, outs []chan<- map[int]int, ins []<-chan map[int]int, rounds int, result *Result, wg *sync.WaitGroup) {
	defer wg.Done()

	known := map[int]int{idx: vote}
	for r := 0; r < rounds; r++ {
		snapshot := copyVotes(known) // only read it
		for _, o := range outs {
			o <- snapshot
		}
		for _, in := range ins {
			for k, v := range <-in {
				known[k] = v
			}
		}
	}

	sum := 0
	for _, v := range known {
		sum += v
	}
	count := len(known)

	res := Result{Known: count, Average: float64(sum) / float64(count)}
	switch {
	case 2*sum > count:
		res.Decision = "majority YES (1)"
	case 2*sum < count:
		res.Decision = "majority NO (0)"
	default:
		res.Decision = "TIE"
	}
	*result = res
}

// Flooding runson a strongly connected directed graph
func Flooding(votes []int, adj [][]int) ([]Result, error) {
	n := len(votes)
	if n == 0 || len(adj) != n {
		return nil, errors.New("invalid input")
	}
	for _, v := range votes {
		if v != 0 && v != 1 {
			return nil, errors.New("votes must be 0 or 1")
		}
	}
	for _, neighbours := range adj {
		for _, v := range neighbours {
			if v < 0 || v >= n {
				return nil, errors.New("neighbour index out of range")
			}
		}
	}
	if !Connected(n, adj) {
		return nil, errors.New("graph is not strongly connected")
	}

	outs := make([][]chan<- map[int]int, n)
	ins := make([][]<-chan map[int]int, n)
	for u := range adj {
		for _, v := range adj[u] {
			ch := make(chan map[int]int, n) // one message per round
			outs[u] = append(outs[u], ch)
			ins[v] = append(ins[v], ch)
		}
	}

	results := make([]Result, n)
	var wg sync.WaitGroup
	wg.Add(n)
	for i := 0; i < n; i++ {
		go majorityNode(i, votes[i], outs[i], ins[i], n-1, &results[i], &wg)
	}
	wg.Wait()
	return results, nil
}

func main() {
	votes := []int{1, 0, 1, 1, 0}
	adj := [][]int{
		{1},
		{2},
		{3, 0},
		{4, 1},
		{0},
	}

	results, err := Flooding(votes, adj)
	if err != nil {
		fmt.Println("error:", err)
		return
	}

	fmt.Println("votes:", votes)
	for i, r := range results {
		fmt.Printf("node %d: collected %d votes, average = %.2f -> %s\n", i, r.Known, r.Average, r.Decision)
	}
}

