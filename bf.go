package main

import (
	"os"
	"fmt"
)

func main() {
	args := os.Args

	if len(args) != 2 {
		fmt.Println("Usage: ./bf <program.bf>")
		os.Exit(0)
	}

	

}