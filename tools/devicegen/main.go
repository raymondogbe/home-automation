package main

import (
	"os"
)

func main() {
	if len(os.Args) < 2 {
		println("usage: devicegen file.def")
		os.Exit(1)
	}

	path := os.Args[1]
	generate(path)
}
