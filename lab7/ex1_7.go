
package main

import (
	"fmt"
	"math/rand"
)

// Message represents a message sent between nodes during Round 1 or Round 2
type Message struct {
	SenderID int
	Round    int // 1 = current state, 2 = Tie value
	Value    int // 0 or 1
}

type Node struct {
	id          int
	x           int // current state x_i(t)
	s           int
	n           int // number of nodes in the system
	isByzantine bool
	inbox       chan Message
}

// Broadcasts the same value to all nodes (including itself)
func (node *Node) Bcast(nodes []*Node, round int, val int) {
	for _, peer := range nodes {
		peer.inbox <- Message{
			SenderID: node.id,
			Round:    round,
			Value:    val,
		}
	}
}

// In Round 1, a Byzantine node sends random values (0 or 1)
func (node *Node) MaliciousBcastState(nodes []*Node) {
	for _, peer := range nodes {
		peer.inbox <- Message{
			SenderID: node.id,
			Round:    1,
			Value:    rand.Intn(2),
		}
	}
}

// In Round 2, a Byzantine leader sends random Tie values
// to different nodes
func (node *Node) MaliciousBcastTie(nodes []*Node) {
	for _, peer := range nodes {
		peer.inbox <- Message{
			SenderID: node.id,
			Round:    2,
			Value:    rand.Intn(2),
		}
	}
}

// Computes the majority value and its frequency
func calcMajority(votes []int, n int) (int, int) {
	ones := 0

	for _, v := range votes {
		if v == 1 {
			ones++
		}
	}

	zeros := n - ones

	if ones >= (n/2)+1 {
		return 1, ones
	}

	return 0, zeros
}

func main() {
	var n, s int

	if _, err := fmt.Scan(&n, &s); err != nil {
		return
	}

	// Byzantine Agreement requires n > 4s
	if n <= 4*s {
		fmt.Println("Error: n must be greater than 4s")
		return
	}

	nodes := make([]*Node, n)

	// Read the initial value of each node
	for i := 0; i < n; i++ {
		var val int
		fmt.Scan(&val)

		nodes[i] = &Node{
			id:    i,
			x:     val,
			s:     s,
			n:     n,
			inbox: make(chan Message, 2*n),
		}
	}

	// Randomly choose s Byzantine nodes
	for _, idx := range rand.Perm(n)[:s] {
		nodes[idx].isByzantine = true
	}

	// Execute s + 2 phases (t = 0, 1, ..., s + 1)
	for t := 0; t <= s+1; t++ {

		// Round 1: every node broadcasts its current state
		for _, node := range nodes {
			if node.isByzantine {
				node.MaliciousBcastState(nodes)
			} else {
				node.Bcast(nodes, 1, node.x)
			}
		}

		// Each node receives n values and computes
		// Majority and multiplicity
		majorities := make([]int, n)
		mults := make([]int, n)

		for i, node := range nodes {
			votes := make([]int, 0, n)

			for j := 0; j < n; j++ {
				msg := <-node.inbox
				votes = append(votes, msg.Value)
			}

			majorities[i], mults[i] = calcMajority(votes, n)
		}

		// Round 2: phase leader (id == t) broadcasts Tie
		if nodes[t].isByzantine {
			nodes[t].MaliciousBcastTie(nodes)
		} else {
			nodes[t].Bcast(nodes, 2, majorities[t])
		}

		// Each node receives the Tie value and updates x
		for i, node := range nodes {
			tieMsg := <-node.inbox
			tie := tieMsg.Value

			if float64(mults[i]) > float64(n)/2.0+float64(s) {
				node.x = majorities[i]
			} else {
				node.x = tie
			}
		}
	}

	// Output the final value of a correct node
	for _, node := range nodes {
		if !node.isByzantine {
			fmt.Println(node.x)
			return
		}
	}
}
