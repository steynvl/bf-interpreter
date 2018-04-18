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

var srcFileName string

func ScanBf(fileName string) []Tape {
	srcFileName = fileName
	source := readFile()

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
				abortCompile(ERR_BACK_JUMP_BEFORE_FORWARD)
			}

			i = programJumps[len(programJumps) - 1]
			programJumps = removeLastEl(programJumps)

			tape = append(tape, Tape{token.JUMP_BACKWARD, i})
			tape[i].StartOfLoop = programCounter

		default:
			abortCompile(ERR_UNKNOWN_CHARACTER)

		}
		programCounter++

	}

	return tape
}

func removeLastEl(p []int) []int {
	return p[:len(p) - 1]
}

func readFile() string {
	b, err := ioutil.ReadFile(srcFileName)

	if err != nil {
		abortCompile(ERR_OPENING_SOURCE_FILE)
	}

	return string(b)
}

func abortCompile(error Error) {

	switch error {
	case ERR_OPENING_SOURCE_FILE:
	case ERR_BACK_JUMP_BEFORE_FORWARD:
	case ERR_UNKNOWN_CHARACTER:
	}

	// TODO
	fmt.Println("abort compile")
	os.Exit(2)

}
