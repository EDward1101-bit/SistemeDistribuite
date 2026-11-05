package main

import (
	"fmt"
	"sync"
)

func graphNode(id int, diameter int, in []<-chan int, out []chan<- int, wg *sync.WaitGroup, leaderChan chan<- int) {
	defer wg.Done()
	maxId := id

	for round := 0; round < diameter; round++ {
		// Transmisie catre vecini
		for _, ch := range out {
			ch <- maxId
		}

		// Receptie de la vecini
		for _, ch := range in {
			val := <-ch
			if val > maxId {
				maxId = val
			}
		}
	}

	if maxId == id {
		leaderChan <- id
	}
}

func main() {
	ids := []int{4, 1, 9, 6}
	n := len(ids)
	diameter := 3 
	
	
	edges := [][2]int{
		{0, 1}, {1, 2}, {2, 3}, {3, 0}, {1, 3},
	}

	inChans := make([][]chan int, n)
	outChans := make([][]chan int, n)

	for _, edge := range edges {
		u, v := edge[0], edge[1]
		ch := make(chan int, diameter)
		outChans[u] = append(outChans[u], ch)
		inChans[v] = append(inChans[v], ch)
	}

	var wg sync.WaitGroup
	wg.Add(n)
	leaderChan := make(chan int, n)

	for i := 0; i < n; i++ {
		var in []<-chan int
		for _, c := range inChans[i] { in = append(in, c) }
		
		var out []chan<- int
		for _, c := range outChans[i] { out = append(out, c) }
		
		go graphNode(ids[i], diameter, in, out, &wg, leaderChan)
	}

	wg.Wait()
	close(leaderChan)

	fmt.Println("ID-uri graf tare conex:", ids)
	for leader := range leaderChan {
		fmt.Println("Liderul ales are ID-ul:", leader)
	}
}