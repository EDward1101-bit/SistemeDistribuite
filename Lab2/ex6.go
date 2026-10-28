package main

import (
	"fmt"
	"math/rand"
	"time"
)

func runRobot(name string, finish chan<- string) {
	time.Sleep(time.Duration(100+rand.Intn(420)) * time.Millisecond)
	finish <- name
}

func main() {
	finish1 := make(chan string)
	finish2 := make(chan string)

	go runRobot("R1", finish1)
	go runRobot("R2", finish2)

	select {
	case winner := <-finish1:
		fmt.Printf("Câștigătorul este: %s!\n", winner)
	case winner := <-finish2:
		fmt.Printf("Câștigătorul este: %s!\n", winner)
	}
}
