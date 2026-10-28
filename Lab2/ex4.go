package main 

import "fmt"

func f1(ok chan bool){
	fmt.Println("f1 s-a executat pana la final")
	ok <- true
}

func main(){
	ok := make(chan bool, 1)
	go f1(ok)

	<-ok

	fmt.Println("main s-a executat pana la final")
}