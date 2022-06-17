package main

import (
	"fmt"
	"sync"
)

func main() {

	var n int
	fmt.Scanf("%d\n", &n)

	var count int32
	if n > 0 {
		wg := sync.WaitGroup{}
		wg.Add(n)

		for i := 0; i < n; i++ {
			go func() {
				count++
				wg.Done()
			}()
		}
		wg.Wait()
	}

	fmt.Println("Finished gorutines: ", count)
}
