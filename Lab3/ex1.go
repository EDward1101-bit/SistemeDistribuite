package main

import (
	"fmt"
	"math/rand"
	"time"
)

type Clock struct{ offset time.Duration }

func (c Clock) Now() time.Time { return time.Now().Add(c.offset) }

func estimate(t1, t2, t3, t4 time.Time) (theta, delta time.Duration) {
	delta = t4.Sub(t1) - t3.Sub(t2)
	theta = (t2.Sub(t1) + t3.Sub(t4)) / 2
	return theta, delta
}

func measure(c_clock, s_clock Clock) (theta, delta time.Duration) {
	t1 := c_clock.Now()
	time.Sleep(3*time.Millisecond + time.Duration(rand.Intn(10))*time.Millisecond)
	t2 := s_clock.Now()

	time.Sleep(5*time.Millisecond + time.Duration(rand.Intn(20))*time.Millisecond)
	t3 := s_clock.Now()

	time.Sleep(3*time.Millisecond + time.Duration(rand.Intn(10))*time.Millisecond)
	t4 := c_clock.Now()

	return estimate(t1, t2, t3, t4)
}

func main() {
	clientClock := Clock{offset: 300 * time.Millisecond}
	serverClock := Clock{offset: 0}

	// a)
	theta, delta := measure(clientClock, serverClock)
	clientClock.offset -= theta

	fmt.Println("a)")
	fmt.Println("Theta:", theta, "| Delta:", delta)
	fmt.Println("Offset corectat:", clientClock.offset)
	fmt.Println()

	// b)
	clientClock.offset = 300 * time.Millisecond // reset pentru test

	var minDelta time.Duration = -1
	var bestTheta time.Duration

	for i := 0; i < 10; i++ {
		t, d := measure(clientClock, serverClock)
		if minDelta == -1 || d < minDelta {
			minDelta = d
			bestTheta = t
		}
	}

	clientClock.offset -= bestTheta

	fmt.Println("b)")
	fmt.Println("Theta optim:", bestTheta, "| Delta minim:", minDelta)
	fmt.Println("Offset corectat:", clientClock.offset)
}
