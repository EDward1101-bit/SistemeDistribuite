package main

import (
	"fmt"
	"sync"
)

type Msg struct {
	Value int
	Done  bool // liderul anunta ca a fost ales
}

func node(id int, in <-chan Msg, out chan<- Msg, isLeader *bool, wg *sync.WaitGroup, leaderChan chan<- int) {
	defer wg.Done()


	out <- Msg{Value: id, Done: false}

	for msg := range in {
		if msg.Done {
			if msg.Value != id {
				out <- msg
			}
			return
		}

		if msg.Value > id {
			out <- msg
		} else if msg.Value == id {
			// Daca a facut o tura completa, e cel mai mare
			*isLeader = true
			leaderChan <- id
			out <- Msg{Value: id, Done: true}
			return
		}

	}
}

func RingFlooding(ids []int) int {
	n := len(ids)
	chans := make([]chan Msg, n)
	for i := 0; i < n; i++ {
		chans[i] = make(chan Msg, n) 
	}

	var wg sync.WaitGroup
	wg.Add(n)
	leaderChan := make(chan int, 1)
	isLeaderFlags := make([]bool, n)

	
	for i := 0; i < n; i++ {
		go node(ids[i], chans[i], chans[(i+1)%n], &isLeaderFlags[i], &wg, leaderChan)
	}

	wg.Wait()
	return <-leaderChan
}

func main() {
	ids := []int{14, 25, 9, 31, 7}
	fmt.Println("ID-uri retea inel:", ids)
	
	leader := RingFlooding(ids)
	fmt.Println("Liderul ales are ID-ul:", leader)
}