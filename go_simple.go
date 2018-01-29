package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	a := asChan(1, 3, 5, 7)
	b := asChan(2, 4, 6, 8)
	c := merge(a, b)
	for v := range c {
		fmt.Println(v)
	}
    /*
        this should print something like this:
            1
            2
            3
            4
            5
            6
            7
            8
        make sure that you close both channels and program should exit successfully at the end.
    */
}

func merge(a, b <-chan int) <-chan int {
	// this should take a and b and return a new channel which will send all values from both a and b
	c := make(chan int)

	go func() {
		defer close(c)
		for a != nil || b != nil {
			select {
				case v, ok := <-a:
					if !ok {
						a = nil
						continue
					}
					c <- v
				case v, ok := <-b:
					if !ok {
						b = nil
						continue
					}
					c <- v
			}
		}
	}()
	return c
}

func asChan(vs ...int) <-chan int {
	// this should reutrn a channel and send vs values randomly to it.
	c := make(chan int)
	go func() {
		for _, v := range vs {
			c <- v
			time.Sleep(time.Duration(rand.Intn(1000)))
		}
		close(c)
	}()
	return c
}