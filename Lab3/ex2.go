package main

import (
	"fmt"
	"math/rand"
	"sort"
	"sync"
	"time"
)

type LamportClock struct {
	counter int
}

func (l *LamportClock) Tick() {
	l.counter++
}

func (l *LamportClock) Send() {
	l.Tick()
}

func (l *LamportClock) Receive(ts int) {
	if ts > l.counter {
		l.counter = ts
	}
	l.Tick()
}

type Message struct {
	sender int
	ts     int
}

type Event struct {
	process          int
	internal_counter int
	kind             string
}

func less(a, b Event) bool {
	if a.internal_counter == b.internal_counter {
		return a.process < b.process
	}
	return a.internal_counter < b.internal_counter
}

func runProcess(id int, in <-chan Message, out chan<- Message, eventsChan chan<- Event, wg *sync.WaitGroup) {
	defer wg.Done()
	clock := LamportClock{counter: 0}

	for step := 0; step < 3; step++ {
		time.Sleep(time.Duration(rand.Intn(100)) * time.Millisecond)

		for k := 0; k < rand.Intn(3); k++ {
			clock.Tick()
			eventsChan <- Event{process: id, internal_counter: clock.counter, kind: "internal"}
		}

		if id == 1 && step == 0 {
			clock.Send()
			eventsChan <- Event{process: id, internal_counter: clock.counter, kind: "send"}
			out <- Message{sender: id, ts: clock.counter}
		}

		select {
		case msg := <-in:
			clock.Receive(msg.ts)
			eventsChan <- Event{process: id, internal_counter: clock.counter, kind: "recv"}

			clock.Send()
			eventsChan <- Event{process: id, internal_counter: clock.counter, kind: "send"}
			out <- Message{sender: id, ts: clock.counter}
		default:
		}
	}
}

func main() {
	ch12 := make(chan Message, 10)
	ch23 := make(chan Message, 10)
	ch31 := make(chan Message, 10)

	eventsChan := make(chan Event, 100)
	var wg sync.WaitGroup
	wg.Add(3)

	go runProcess(1, ch31, ch12, eventsChan, &wg)
	go runProcess(2, ch12, ch23, eventsChan, &wg)
	go runProcess(3, ch23, ch31, eventsChan, &wg)

	wg.Wait()
	close(eventsChan)

	var allEvents []Event
	for ev := range eventsChan {
		allEvents = append(allEvents, ev)
	}

	sort.Slice(allEvents, func(i, j int) bool {
		return less(allEvents[i], allEvents[j])
	})

	for _, ev := range allEvents {
		fmt.Printf("P%d - %s (ceas: %d)\n", ev.process, ev.kind, ev.internal_counter)
	}
}
