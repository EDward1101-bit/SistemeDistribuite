package main

import "fmt"

func produce(ch chan string) {
    ch <- "Primul mesaj"
    ch <- "Al doilea mesaj"
}

func main() {
    ch := make(chan string, 2)
    go produce(ch)

    fmt.Println(<-ch, <-ch)
}