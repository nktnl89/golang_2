package main

import (
	"fmt"
	"log"
	"os"
	"runtime/trace"
	"sync"
)

var mut = sync.Mutex{}

func main() {
	trace.Start(os.Stderr)
	defer trace.Stop()

	gorutineNum := 16

	wg := sync.WaitGroup{}
	filename := "test.txt"

	for i := 0; i < gorutineNum; i++ {
		wg.Add(1)
		go writeToFile(filename, i, &wg)
	}
	wg.Wait()
}

func writeToFile(filename string, i int, wg *sync.WaitGroup) {
	defer wg.Done()
	mut.Lock()
	defer mut.Unlock()
	f, _ := os.OpenFile(filename, os.O_APPEND, 0644)
	log.Printf(fmt.Sprintf("gorutine №%d opened file\n", i))
	if _, err := f.Write([]byte(fmt.Sprintf("gorutine №%d was here\n", i))); err != nil {
		log.Fatal(err)
	}
	if err := f.Close(); err != nil {
		log.Fatal(err)
	}
	log.Printf(fmt.Sprintf("gorutine №%d saved changes\n", i))

}
