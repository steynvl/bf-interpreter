package main

import (
	"os"
	"fmt"
	"bf-interpreter/scanner"
	"bufio"
	"bf-interpreter/token"
)

const tapeSize = 30000

func main() {
	args := os.Args

	if len(args) != 2 {
		fmt.Println("Usage: ./bf <program.bf>")
		os.Exit(0)
	}

	scanner.InitScanner(args[1])
	tape := scanner.Scan()

	execute(tape, bufio.NewReader(os.Stdin))
}

func execute(tape []scanner.Tape, reader *bufio.Reader) {
	pp := make([]int, tapeSize)
	pIndex := 0

	for i := 0; i < len(tape); i++ {

		switch tape[i].Token {

		case token.IncrementPointer:
			pIndex++
		case token.DecrementPointer:
			pIndex--
		case token.IncrementByte:
			pp[pIndex]++
		case token.DecrementByte:
			pp[pIndex]--
		case token.OutputByte:
			fmt.Printf("%c", pp[pIndex])
		case token.InputByte:
			val, _ := reader.ReadByte()
			pp[pIndex] = int(val)
		case token.JumpForward:
			if pp[pIndex] == 0 {
				i = int(tape[i].StartOfLoop)
			}
		case token.JumpBackward:
			if pp[pIndex] > 0 {
				i = int(tape[i].StartOfLoop)
			}
		}

	}
}
