package main

import (
	"fmt"
	"golang_2/ex2/ex2_2/panic_func"
)

func main() {
	_, err := panic_func.CallImplicitPanic()

	if err != nil {
		fmt.Println(err)
	}

}
