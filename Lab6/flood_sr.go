package main

import (
	"bufio"
	"fmt"
	"math/rand"
	"os"
	"sort"
	"strconv"
	"strings"
)

func diametru(adj [][]int) int {
	n := len(adj)
	diam := 0
	for src := 0; src < n; src++ {
		// BFS din nodul src
		dist := make([]int, n)
		for i := range dist {
			dist[i] = -1
		}
		dist[src] = 0
		queue := []int{src}

		for len(queue) > 0 {
			u := queue[0]
			queue = queue[1:]

			for _, v := range adj[u] {
				if dist[v] == -1 {
					dist[v] = dist[u] + 1
					queue = append(queue, v)
				}
			}
		}

		for _, d := range dist {
			if d == -1 {
				return -1 // avem un nod la care nu putem ajunge din src
			}
			if d > diam {
				diam = d
			}
		}
	}
	return diam
}

// Info reprezinta o pereche (id, valoare) retinuta in memoria unui nod.
type Info struct {
	ID    int
	Value int
}

// Memory este o multime de token-uri (reuniunea se face prin inserare).
type Memory map[Info]struct{}

type Node struct {
	ID        int
	Neighbors []int
	M         Memory
}

func readLineInts(scanner *bufio.Scanner) []int {
	if !scanner.Scan() {
		return nil
	}
	fields := strings.Fields(scanner.Text())
	res := make([]int, len(fields))
	for i, f := range fields {
		val, _ := strconv.Atoi(f)
		res[i] = val
	}
	return res
}

func main() {
	scanner := bufio.NewScanner(os.Stdin)

	header := readLineInts(scanner)
	if len(header) < 2 {
		return
	}
	n, s := header[0], header[1]

	tokens := readLineInts(scanner)
	adj := make([][]int, n)
	for i := 0; i < n; i++ {
		adj[i] = readLineInts(scanner)
	}

	diam := diametru(adj)
	if diam == -1 {
		fmt.Println("Graful nu este tare conex")
		return
	}

	nodes := make([]Node, n)
	alive := make([]bool, n)
	for i := 0; i < n; i++ {
		mem := make(Memory)
		mem[Info{ID: i, Value: tokens[i]}] = struct{}{}
		nodes[i] = Node{
			ID:        i,
			Neighbors: adj[i],
			M:         mem,
		}
		alive[i] = true
	}

	crashes := 0
	t := 0

	for {
		if crashes < s && rand.Intn(2) == 0 {
			var candidates []int
			for i := 0; i < n; i++ {
				if alive[i] {
					candidates = append(candidates, i)
				}
			}
			if len(candidates) > 1 {
				victim := candidates[rand.Intn(len(candidates))]
				alive[victim] = false
				crashes++
			}
		}

		nextM := make([]Memory, n)
		for i := 0; i < n; i++ {
			if !alive[i] {
				continue
			}
			nextM[i] = make(Memory)
			for info := range nodes[i].M {
				nextM[i][info] = struct{}{}
			}
		}

		for i := 0; i < n; i++ {
			if !alive[i] {
				continue
			}
			for _, neighbor := range nodes[i].Neighbors {
				if !alive[neighbor] {
					continue
				}
				for info := range nodes[i].M {
					nextM[neighbor][info] = struct{}{}
				}
			}
		}

		for i := 0; i < n; i++ {
			if alive[i] {
				nodes[i].M = nextM[i]
			}
		}

		if t > (s+1)*diam {
			break
		}
		t++
	}

	var finalMem Memory
	for i := 0; i < n; i++ {
		if alive[i] {
			finalMem = nodes[i].M
			break
		}
	}

	result := make([]Info, 0, len(finalMem))
	for info := range finalMem {
		result = append(result, info)
	}
	sort.Slice(result, func(i, j int) bool {
		return result[i].ID < result[j].ID
	})

	// fmt.Println("Diametru:", diam)
	// fmt.Println("Iteratia:", t)
	
	for i, info := range result {
		if i > 0 {
			fmt.Print(" ")
		}
		fmt.Printf("(%d, %d)", info.ID, info.Value)
	}
	fmt.Println()
}
