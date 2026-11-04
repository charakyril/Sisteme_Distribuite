package main 

import (
	"fmt"
	"sync"
)

type Msg struct {
    Value int
    Done bool // ο αρχηγός ανακοινώνει ότι έχει εκλεγεί
}

// εκτελεί τις εσωτερικές λειτουργίες του κόμβου
func node(id int, in <-chan Msg, out chan <- Msg, isLeader *bool, wg *sync.WaitGroup) {
	defer wg.Done()
	maxID := id
	out<- Msg{Value: id}
	for msg := range in {
		switch {
		case msg.Done:
			// ο αρχηγός ανακοινώνει ότι έχει εκλεγεί
			if msg.Value != id {
				out<-msg
			}
			return 
		case msg.Value == id:
			// ο κόμβος είναι ο αρχηγός
			*isLeader = true
			out<- Msg{Value: id, Done: true}
			return 
		case msg.Value > maxID:
			maxID = msg.Value
			out<- msg
		} 
	}
}
// επιστρέφει το ID του αρχηγού
func RingFlooding(ids []int) int {
	if len(ids) == 0 {
		return -1
	} 
	chans := make([]chan Msg, len(ids))
	for i := range chans {
		chans[i] = make(chan Msg, len(ids)+1)
	}
	leader := make([]bool, len(ids))
	var wg sync.WaitGroup
	wg.Add(len(ids))
	for i:=0; i<len(ids); i++ {
		go node(ids[i], chans[i], chans[(i+1)%len(ids)], &leader[i], &wg)
	}
	wg.Wait()
	for i, isLeader := range leader {
		if isLeader {
			return ids[i]
		}
	}
	return -1 
}


func main() {
	ids := []int{4, 17, 9, 10, 2, 20, 6, 11, 7}
	leader := RingFlooding(ids)
	fmt.Printf("The leader is: %d\n", leader)
}
