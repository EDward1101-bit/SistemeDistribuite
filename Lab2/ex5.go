package main

import "fmt"

func send(ch chan<- string) {
	ch <- "Un mesaj"
}

func receive(in <-chan string, out chan<- string) {
	msg := <-in
	out <- msg
}

func main() {
	ch1 := make(chan string)
	ch2 := make(chan string)

	go send(ch1)
	go receive(ch1, ch2)

	mesaj := <-ch2
	fmt.Println(mesaj)
}