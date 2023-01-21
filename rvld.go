package main

import (
	"fmt"
	"os"
)

func main() {
	//println("Hello, world.")
	if len(os.Args) < 2 {
		println("wrong args")
		os.Exit(1)
	}
	fmt.Printf("%v\n", os.Args)
}
