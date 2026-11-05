package main

import (
	"fmt"
	"sync"
)

func avgNode(id int, initialValue float64, in []<-chan float64, out []chan<- float64, rounds int, wg *sync.WaitGroup, results chan<- float64) {
	defer wg.Done()
	x := initialValue

	for r := 0; r < rounds; r++ {
		for _, ch := range out {
			ch <- x
		}

		sum := x
		for _, ch := range in {
			sum += <-ch
		}

		// 1 / (d_i + 1) * (x_i + suma(vecini))
		x = sum / float64(len(in)+1)
	}

	results <- x
}

func main() {
	votes := []float64{1, 1, 0, 1, 0}
	n := len(votes)
	rounds := 30 

	inChans := make([][]chan float64, n)
	outChans := make([][]chan float64, n)

	// Legam nodurile bidirecțional
	for i := 0; i < n; i++ {
		next := (i + 1) % n
		prev := (i - 1 + n) % n

		chNext := make(chan float64, rounds)
		chPrev := make(chan float64, rounds)

		outChans[i] = append(outChans[i], chNext, chPrev)
		inChans[next] = append(inChans[next], chNext)
		inChans[prev] = append(inChans[prev], chPrev)
	}

	var wg sync.WaitGroup
	wg.Add(n)
	results := make(chan float64, n)

	for i := 0; i < n; i++ {
		var in []<-chan float64
		for _, c := range inChans[i] { in = append(in, c) }
		
		var out []chan<- float64
		for _, c := range outChans[i] { out = append(out, c) }
		
		go avgNode(i, votes[i], in, out, rounds, &wg, results)
	}

	wg.Wait()
	close(results)

	fmt.Println("Voturi initiale:", votes)
	var finalSum float64
	for res := range results {
		finalSum += res
	}
	
	avg := finalSum / float64(n)
	fmt.Printf("Media obtinuta la consens: %.3f\n", avg)
	
	if avg > 0.5 {
		fmt.Println("Rezultat: Majoritate PENTRU (1)")
	} else {
		fmt.Println("Rezultat: Majoritate CONTRA (0)")
	}
}