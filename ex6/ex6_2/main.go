package main

import (
	"log"
	"os"
	"runtime"
	"runtime/trace"
)

func main() {
	trace.Start(os.Stderr)
	defer trace.Stop()

	go log.Println("I'm working!")
	for i := 0; ; i += 1 {
		if i%1000 == 0 {
			runtime.Gosched()
		}
	}
}
