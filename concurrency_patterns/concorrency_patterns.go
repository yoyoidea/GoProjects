package main

import "fmt"

func main() {
	// 词法约束使词法的作用域仅公开用于多个并发进程的正确数据和并发原句
	// 负责写 channel
	chanOwner := func() <-chan int {
		results := make(chan int, 5)
		go func() {
			defer close(results)
			for i:=0; i <= 5; i++{
				results <- i
			}
		}()
		return results
	}
	// 负责消费 channel
	consumer := func(results <- chan int) {
		for result := range results{
			fmt.Printf("Received: %d\n", result)
		}
		fmt.Println("Done receiving")
	}
	results := chanOwner()
	consumer(results)

}
