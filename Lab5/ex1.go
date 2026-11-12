package main

import (
	"fmt"
	"math"
	"sync"
)

const (
	R      = 16.0 
	EPS    = 1e-6
	MAX_IT = 1000
)

type Point struct {
	X, Y float64
}

func dist(a, b Point) float64 {
	dx := a.X - b.X
	dy := a.Y - b.Y
	return math.Sqrt(dx*dx + dy*dy)
}


func buildGraph(pos []Point) [][]int {
	n := len(pos)
	g := make([][]int, n)

	for i := 0; i < n; i++ {
		for j := i + 1; j < n; j++ {
			if dist(pos[i], pos[j]) < R {
				g[i] = append(g[i], j)
				g[j] = append(g[j], i)
			}
		}
	}

	return g
}

func robot(id int, pos Point, vecini []int, inbox []chan Point, newPos []Point, wg *sync.WaitGroup) {
	defer wg.Done()

	// Look
	for _, v := range vecini {
		inbox[v] <- pos
	}

	sum := pos
	for range vecini {
		p := <-inbox[id]
		sum.X += p.X
		sum.Y += p.Y
	}

	// Compute
	k := float64(len(vecini) + 1)
	next := Point{sum.X / k, sum.Y / k}

	// Move
	newPos[id] = next
}

// pozitiile aproape identice devin un singur robot
func compress(pos []Point) []Point {
	res := []Point{}

	for _, p := range pos {
		found := false
		for _, q := range res {
			if dist(p, q) < EPS {
				found = true
				break
			}
		}
		if !found {
			res = append(res, p)
		}
	}

	return res
}

func main() {
	var n int
	fmt.Scan(&n)

	pos := make([]Point, n)
	for i := range pos {
		fmt.Scan(&pos[i].X, &pos[i].Y)
	}

	it := 0
	for len(pos) > 1 && it < MAX_IT {
		it++
		g := buildGraph(pos)

		// fiecare robot primeste exact cate un mesaj de la fiecare vecin
		inbox := make([]chan Point, len(pos))
		for i := range inbox {
			inbox[i] = make(chan Point, len(g[i]))
		}

		newPos := make([]Point, len(pos))
		var wg sync.WaitGroup

		for i := range pos {
			wg.Add(1)
			go robot(i, pos[i], g[i], inbox, newPos, &wg)
		}
		wg.Wait()

		moved := false
		for i := range pos {
			if dist(pos[i], newPos[i]) > EPS {
				moved = true
				break
			}
		}

		pos = compress(newPos)
		fmt.Printf("Iteratia %d: %d roboti activi\n", it, len(pos))

		// clustere izolate, nu se mai misca nimic
		if !moved {
			break
		}
	}

	if len(pos) == 1 {
		fmt.Printf("Rendezvous in (%.4f, %.4f) dupa %d iteratii\n", pos[0].X, pos[0].Y, it)
	} else {
		fmt.Printf("Nu s-a ajuns la rendezvous, au ramas %d clustere:\n", len(pos))
		for _, p := range pos {
			fmt.Printf("(%.4f, %.4f)\n", p.X, p.Y)
		}
	}
}
