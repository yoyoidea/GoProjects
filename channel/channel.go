package main

import "fmt"

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


}