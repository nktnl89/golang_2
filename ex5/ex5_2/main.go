package main

import (
	"fmt"
	"sync"
)

func main() {

	wg := sync.WaitGroup{}
	wg.Add(2)

	var counter int
	mut := sync.Mutex{}

	go func(c *int) {
		mut.Lock()
		defer mut.Unlock()
		*c++
		fmt.Println(*c)
		wg.Done()

	}(&counter)

	go func(c *int) {
		mut.Lock()
		defer mut.Unlock()
		*c--
		fmt.Println(*c)
		wg.Done()

	}(&counter)

	wg.Wait()

	fmt.Println(counter)
}
