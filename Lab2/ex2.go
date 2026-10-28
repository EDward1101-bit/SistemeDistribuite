package main

import "fmt"
import "sync"

func getMessage(_ch chan string,  _wg *sync.WaitGroup){
	fmt.Println(<-_ch)
	_wg.Done()
}

func main(){
	ch := make(chan string, 1)
	ch <- "Tom Holland"

	wg := sync.WaitGroup{}

	wg.Add(1)
	go getMessage(ch, &wg)

	wg.Wait()
	
}