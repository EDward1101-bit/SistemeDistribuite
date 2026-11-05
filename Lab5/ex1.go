package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

type Robot struct {
	X float64
	Y float64
}

func main() {
	var n int
	robots := make([]Robot, n)

	file, err := os.Open("ex1.txt")
	if err != nil {
		fmt.Println(err)
	}

	defer file.Close()

	scanner := bufio.NewScanner(file)

	scanner.Split(bufio.ScanWords)

	scanner.Scan()
	n, _ = strconv.Atoi(scanner.Text())
	for scanner.Scan() {
		var X, Y float64
		X, _ = strconv.ParseFloat(scanner.Text(), 64)
		scanner.Scan()
		Y, _ = strconv.ParseFloat(scanner.Text(), 64)
		robots = append(robots, Robot{X: X, Y: Y})
	}

	err = scanner.Err()
	if err != nil {
		fmt.Println("Eroare la citire: ", err)
	}

	for _, robot := range robots {
		fmt.Println(robot.X, robot.Y)
	}

}
