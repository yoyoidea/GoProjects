package main

import (
	"fmt"
	"time"
)

func main()  {
	// channel
	chanOwner := func() <- chan int {
		resultsStream := make(chan int, 5)
		go func() {
			defer close(resultsStream)
			for i := 0; i <= 6; i++ {
				resultsStream <- i
			}
		}()
		return resultsStream
	}
	resultsStream := chanOwner()
	for result := range resultsStream{
		fmt.Printf("Received: %d\n", result)
	}
	fmt.Println("Done")

	// select, example 1:
	start := time.Now()
	c := make(chan interface{})
	go func() {
		time.Sleep(1*time.Second)
		close(c)
	}()

	fmt.Println("Blocking on read....")

	select {
	case <-c:
		fmt.Printf("Unblocked %v later.\n", time.Since(start))

	}

	// select, example 2:
	c1 := make(chan interface{}); close(c1)
	c2 := make(chan interface{}); close(c2)

	var c1Count, c2Count, c3Count int
	for i:= 1000; i>=0; i--{
		select {
		case <-c1:
			c1Count++
		case <-c2:
			c2Count++
		default:
			c3Count++
		}
	}
	fmt.Printf("c1Count: %d\nc2Clount: %d\n", c1Count, c2Count)

	// the code block will be blocked
	//select {}

}