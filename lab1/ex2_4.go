package main

import (
	"errors"
	"fmt"
	"sync"
)

func nodeconnection(n int, adj [][]int) bool {
	reach := func(g [][]int) bool {
		seen := make([]bool, n)
		stack := []int{0}
		seen[0] = true
		count := 1
		for len(stack) > 0 {
			v := stack[len(stack)-1]
			stack = stack[:len(stack)-1]
			for _, w := range g[v] {
				if !seen[w] {
					seen[w] = true
					count++
					stack = append(stack, w)
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

func graphnode (id int, out []chan<- int, ins []<-chan int, round int, result *int, wg *sync.WaitGroup) {
	defer wg.Done()

	max := id
	for r:=0; r<round; r++ {
		for _, ch := range out {
			ch<-max
		}
		for _, in := range ins {
			if v := <-in; v > max {
				max = v
			}
		}
	}
	*result = max
}

func flooding(ids []int, adj [][]int)(int, error){
	n:= len(ids)
	if n==0 || len(adj) != n {
		return -1, errors.New("invalid input")
	}
	for _, neighbors := range adj {
		for _, v := range neighbors {
			if v < 0 || v >= n {
				return -1, errors.New("neighbor index out of range")
			}
		}
	}
	if !nodeconnection(n, adj) {
		return -1, errors.New("graph is not strongly connected")

	}
	
	maxIDs := append([]int(nil), ids...)
	for round := 0; round < n-1; round++ {
		next := append([]int(nil), maxIDs...)
		for u, neighbors := range adj {
			for _, v := range neighbors {
				if maxIDs[u] > next[v] {
					next[v] = maxIDs[u]
				}
			}
		}
		maxIDs = next
	}
	return maxIDs[0], nil
}

func main() {
	
}