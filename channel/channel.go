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

	// select
	start := time.Now()
	c := make(chan interface{})
	go func() {
		time.Sleep(5*time.Second)
		close(c)
	}()

	fmt.Println("Blocking on read....")

	select {
	case <-c:
		fmt.Printf("Unblocked %v later.\n", time.Since(start))

	}

}