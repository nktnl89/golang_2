package main

import (
	"fmt"
	"time"
)

func increase(result chan int, jobs <-chan int) {
	for j := range jobs {
		num := <-result
		result <- num + 1
		fmt.Printf("worker %d increment number to %d\n", j, num+1)
	}
}

func main() {

	workers := make(chan int, 1000)
	number := make(chan int, 1)
	number <- 0

	for i := 0; i < 1000; i++ {
		workers <- i

		go increase(number, workers)

	}

	time.Sleep(2 * time.Second)
	fmt.Println("итого: ", <-number)
}
