package main

import "fmt"

type Punct struct {
	x, y int
}

func robot(nume string, trimite chan<- Punct, primeste <-chan Punct, gata chan<- bool) {
	for i := 1; i <= 3; i++ {
		p := Punct{x: i, y: i * 10}
		fmt.Println(nume, "a trimis pozitia:", p.x, p.y)
		trimite <- p
	}

	
	close(trimite)

	for p := range primeste {
		fmt.Println(nume, "a primit pozitia partenerului:", p.x, p.y)
	}

	fmt.Println(nume, "si-a terminat complet sarcina.")


	gata <- true
}

func main() {

	ch1 := make(chan Punct, 5) //de la r1 la r2
	ch2 := make(chan Punct, 5) //de la r2 la r1


	gata1 := make(chan bool)
	gata2 := make(chan bool)

	go robot("Robot 1", ch1, ch2, gata1)
	go robot("Robot 2", ch2, ch1, gata2)

	<-gata1
	<-gata2

	fmt.Println("Ambii roboti au raportat finalizarea sarcinilor!")
}