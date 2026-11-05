package main

import (
	"fmt"
	"math/rand"
	"sort"
	"sync"
	"time"
)

type LamportClock struct{ counter int }

// increments the counter
func (l *LamportClock) Tick() { l.counter++ }

// Send increments the counter
func (l *LamportClock) Send() { l.counter++ }

// Receive sets the counter to max(local, received) + 1
func (l *LamportClock) Receive(ts int) {
	if ts > l.counter {
		l.counter = ts
	}
	l.counter++
}

type Message struct{ sender, ts int }

type Event struct {
	process          int
	internal_counter int
	kind             string // "internal", "send", "recv"
}

func less(a, b Event) bool {
	if a.internal_counter != b.internal_counter {
		return a.internal_counter < b.internal_counter
	}
	return a.process < b.process
}

func process(id int, next chan<- Message, inbox <-chan Message, events chan<- Event, rounds int, wg *sync.WaitGroup) {
	defer wg.Done()
	clock := &LamportClock{}

	for r := 0; r < rounds; r++ {
		internal := rand.Intn(3)
		for k := 0; k < internal; k++ {
			clock.Tick()
			events <- Event{id, clock.counter, "internal"}
		}

		clock.Send()
		next <- Message{sender: id, ts: clock.counter}
		events <- Event{id, clock.counter, "send"}

		// Receive from the previous process in the ring
		msg := <-inbox
		clock.Receive(msg.ts)
		events <- Event{id, clock.counter, "recv"}

		time.Sleep(time.Duration(rand.Intn(20)) * time.Millisecond)
	}
}

func main() {
	const n = 3
	const rounds = 3

	// inboxes[i] is the inbox of process P(i+1)
	inboxes := make([]chan Message, n)
	for i := range inboxes {
		inboxes[i] = make(chan Message, rounds)
	}

	events := make(chan Event)
	var wg sync.WaitGroup

	// P1, P2, P3, P1
	for i := 0; i < n; i++ {
		wg.Add(1)
		go process(i+1, inboxes[(i+1)%n], inboxes[i], events, rounds, &wg)
	}

	go func() {
		wg.Wait()
		close(events)
	}()

	// αποθηκευση event
	var all []Event
	for e := range events {
		all = append(all, e)
	}

	// ταξινόμηση
	sort.Slice(all, func(i, j int) bool { return less(all[i], all[j]) })

	for _, e := range all {
		fmt.Printf("clock=%2d  P%d  %s\n", e.internal_counter, e.process, e.kind)
	}
}
