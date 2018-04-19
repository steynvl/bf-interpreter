package scanner

import (
	"io/ioutil"
	"bf-interpreter/token"
	"unicode"
	"fmt"
	"os"
)

type Tape struct {
	Token       token.Token
	StartOfLoop int
}

var srcFileName string
var ch int32

func Scan() []Tape {
	source := readFile()

	programCounter, i := 0, 0

	tape := make([]Tape, 0)
	programJumps := make([]int, 0)

	for _, el := range source {

		if unicode.IsSpace(el) {
			continue
		}
		ch = el
		tok := Tape{token.None, 0}

		switch el {

		case '>':
			tok.Token = token.IncrementPointer
			tape = append(tape, tok)
		case '<':
			tok.Token = token.DecrementPointer
			tape = append(tape, tok)
		case '+':
			tok.Token = token.IncrementByte
			tape = append(tape, tok)
		case '-':
			tok.Token = token.DecrementByte
			tape = append(tape, tok)
		case '.':
			tok.Token = token.OutputByte
			tape = append(tape, tok)
		case ',':
			tok.Token = token.InputByte
			tape = append(tape, tok)
		case '[':
			tok.Token = token.JumpForward
			tape = append(tape, tok)
			programJumps = append(programJumps, programCounter)
		case ']':
			if len(programJumps) == 0 {
				abortCompile(ErrBackJumpBeforeForward)
			}

			i = programJumps[len(programJumps) - 1]
			programJumps = removeLastEl(programJumps)

			tok.Token = token.JumpBackward
			tok.StartOfLoop = i
			tape = append(tape, tok)

			tape[i].StartOfLoop = programCounter
		default:
			abortCompile(ErrUnknownCharacter)

		}
		programCounter++

	}

	return tape
}

func InitScanner(fileName string) {
	srcFileName = fileName
}

func readFile() string {
	b, err := ioutil.ReadFile(srcFileName)

	if err != nil {
		switch err {
		case os.ErrPermission:
			abortCompile(ErrFilePermission)
		case os.ErrNotExist:
			abortCompile(ErrFileDoesNotExist)
		default:
			abortCompile(ErrFileOpen)
		}
	}
	return string(b)
}

func removeLastEl(p []int) []int {
	return p[:len(p) - 1]
}

func abortCompile(err Err) {
	switch err {
	case ErrFilePermission:
		fmt.Printf("Err: can't open file '%s', permission denied\n", srcFileName)
		os.Exit(2)
	case ErrFileDoesNotExist:
		fmt.Printf("Err: can't open '%s', file does not exist\n", srcFileName)
		os.Exit(3)
	case ErrFileOpen:
		fmt.Printf("Err: can't open file '%s'\n", srcFileName)
	case ErrBackJumpBeforeForward:
		fmt.Printf("Err: ']' occurs before '['\n")
		os.Exit(5)
	case ErrUnknownCharacter:
		fmt.Printf("Err: Unknown character '%c'\n", ch)
		os.Exit(6)
	}
}
