package scanner

import (
	"io/ioutil"
	"os"
	"bf-interpreter/token"
	"unicode"
	"fmt"
)

type Tape struct {
	Token       token.Token
	StartOfLoop int
}

func ScanBf(fileName string) []Tape {
	source := readFile(fileName)

	programCounter, i := 0, 0

	tape := make([]Tape, 0)
	programJumps := make([]int, 0)

	for _, el := range source {

		if unicode.IsSpace(el) {
			continue
		}

		switch el {

		case '>':
			tape = append(tape, Tape{token.INCREMENT_THE_POINTER, 0})
		case '<':
			tape = append(tape, Tape{token.DECREMENT_THE_POINTER, 0})
		case '+':
			tape = append(tape, Tape{token.INCREMENT_THE_BYTE, 0})
		case '-':
			tape = append(tape, Tape{token.DECREMENT_THE_BYTE, 0})
		case '.':
			tape = append(tape, Tape{token.OUTPUT_BYTE, 0})
		case ',':
			tape = append(tape, Tape{token.INPUT_A_BYTE, 0})
		case '[':
			tape = append(tape, Tape{token.JUMP_FORWARD, 0})
			programJumps = append(programJumps, programCounter)
		case ']':

			if len(programJumps) == 0 {
				// TODO
				fmt.Println("] before [")
				os.Exit(11)
			}

			i = programJumps[len(programJumps) - 1]
			programJumps = removeLastEl(programJumps)

			tape = append(tape, Tape{token.JUMP_BACKWARD, i})
			tape[i].StartOfLoop = programCounter

		default:
			// TODO error
			os.Exit(10)

		}
		programCounter++

	}

	return tape
}

func removeLastEl(p []int) []int {
	return p[:len(p) - 1]
}

func readFile(fileName string) string {
	b, err := ioutil.ReadFile(fileName)

	if err != nil {
		// TODO
		os.Exit(9)
	}

	return string(b)
}
