package main

import (
	"fmt"
	"log"
	"os"
	"strconv"
)

func main() {
	path := os.Getenv("GOPATH") + "\\src\\golang_2\\ex1\\ex1_3\\files"
	fmt.Println(path)
	if _, err := os.Stat(path); os.IsNotExist(err) {
		os.Mkdir(path, os.ModePerm)
	}

	for i := 0; i < 1_000_000; i++ {
		fileName := path + "\\" + strconv.Itoa(i)

		os.Create(fileName)

		file, err := os.Open(fileName)
		if err != nil {
			log.Fatal(err)
		}
		defer func() {
			if err = file.Close(); err != nil {
				log.Fatal(err)
			}
		}()
		fmt.Println(i)
	}
}
