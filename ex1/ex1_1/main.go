package main

import "fmt"

func main() {
	implicitPanicCall()
}

func implicitPanicCall() int {
	defer func() {
		if v := recover(); v != nil {
			fmt.Println("Implicit panic was recovered", v)
		}
	}()

	zero := 0
	return 100 / zero
}
