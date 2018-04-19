// Package executor provides a function to execute
// a scanned bf tape.
package executor

import (
	"bufio"
	"fmt"
	"bf-interpreter/src/scanner"
	"bf-interpreter/src/token"
)

const tapeSize = 30000

func Run(tape []scanner.Tape, reader *bufio.Reader) {
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
