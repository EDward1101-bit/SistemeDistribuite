package main

import (
	"fmt"
	"sync"
)

func getMessage(_ch chan string, _wg *sync.WaitGroup) {
	_ch <- "LingLong DingDong"
	_wg.Done()
}

func main() {
	ch := make(chan string, 1)

	wg := sync.WaitGroup{}

	wg.Add(1)
	go getMessage(ch, &wg)

	wg.Wait()
	fmt.Println(<-ch)
}
