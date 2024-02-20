package main

import (
	"os"
)

func main() {
	if len(os.Args) < 2 {
		println("Incorrect usage: go run main.go main.ygl")
		os.Exit(1)
	}

	println(os.Args[1])
	os.Exit(0)
}
