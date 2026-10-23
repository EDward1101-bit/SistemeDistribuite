package main

import "fmt"

func main() {
	var x int
	ok := true
	fmt.Scanf("%d", &x)

	if (x%2 == 0) || x <= 1 {
		ok = false
	} else {
		for i := 3; i*i <= x; i += 2 {
			if x%i == 0 {
				ok = false
				break
			}
		}

	}

	if ok || x == 2 {
		fmt.Println("prim")
	} else {
		fmt.Println("nu e prim")
	}
}
