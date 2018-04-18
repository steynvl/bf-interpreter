package main

import (
	"os"
	"fmt"
	"bf-interpreter/scanner"
)

func main() {
	args := os.Args

	if len(args) != 2 {
		fmt.Println("Usage: ./bf <program.bf>")
		os.Exit(0)
	}

	tape := scanner.ScanBf(args[1])

	executeBf(tape)
}

func executeBf(tape []scanner.Tape) {

}
