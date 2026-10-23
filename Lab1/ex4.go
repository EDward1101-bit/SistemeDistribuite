package main

import "fmt"

func main() {
	v := []float32{15.0, 234.0, 5.0, 0.72, 98.1, 3.14, 0.045, 68.0, 120.0, 0.875}

	var index int = -1
	fmt.Scanf("%d", &index)

	if 0 <= index && index < len(v) {
		v = append(v[:index], v[index+1:]...)
	}

	fmt.Println(v)
}
