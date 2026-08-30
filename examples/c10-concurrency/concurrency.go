package main

import (
	"fmt"
	"time"
)

func count(label string) {
	for i := range 5 {
		time.Sleep(50 * time.Millisecond) // pause so the swap between loops is visible
		fmt.Println(label, i)
	}
}

func square(number int, channel chan int) {
	channel <- number * number // send the result
}

func main() {
	fmt.Println("-- goroutines --")
	go count("goroutine") // no waiting, main moves on immediately
	count("main")         // normal call, keeps main alive while both loops run

	fmt.Println()
	fmt.Println("-- channels --")
	channel := make(chan int)
	go square(3, channel)
	go square(4, channel)

	first := <-channel // receive blocks until a send happens
	second := <-channel
	fmt.Println("results:", first, second)

	fmt.Println()
	fmt.Println("-- buffered --")
	buffered := make(chan int, 2)
	buffered <- 1 // sends up to capacity succeed with no receiver waiting
	buffered <- 2
	fmt.Println("buffered:", <-buffered, <-buffered)

	fmt.Println()
	fmt.Println("-- select --")
	go func() {
		time.Sleep(2 * time.Second)
		channel <- 25 // arrives too late to be received
	}()

	select {
	case value := <-channel:
		fmt.Println("received:", value)
	case <-time.After(1 * time.Second): // fires first, main waited only 1 second
		fmt.Println("timed out after 1 second")
	}
}
