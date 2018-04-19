package main

import (
	"bufio"
	"fmt"
	"os"
	"bf-interpreter/src/scanner"
	"bf-interpreter/src/executor"
)

func main() {
	args := os.Args

	if len(args) != 2 {
		fmt.Println("Usage: ./bf <program.bf>")
		os.Exit(0)
	}

	scanner.InitScanner(args[1])
	tape := scanner.Scan()

	executor.Run(tape, bufio.NewReader(os.Stdin))
}
