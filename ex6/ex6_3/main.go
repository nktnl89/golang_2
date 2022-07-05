package main

import (
	"fmt"
	"os"
	"runtime/trace"
)

func main() {
	trace.Start(os.Stderr)
	defer trace.Stop()

	m := make(map[string]string)

	go func() {
		m["1"] = "a"
	}()

	m["2"] = "b"

	for k, v := range m {
		fmt.Println(k, v)
	}

}
