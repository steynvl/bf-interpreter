// Package scanner provides a function
// for a small lexical analyser of bf.
package scanner

import (
	"fmt"
	"io/ioutil"
	"os"
	"unicode"
	"bf-interpreter/src/token"
)

// A tape represents a bf program.
type Tape struct {
	Token       token.Token
	StartOfLoop int
}

var srcFileName string
var position SourcePos

func Scan() []Tape {
	source := readFile()

	programCounter, i := 0, 0

	tape := make([]Tape, 0)
	programJumps := make([]int, 0)

	for _, el := range source {

		if unicode.IsSpace(el) {
			if string(el) == "\n" {
				position.line++
				position.col = -1
			}

			position.col++
			continue
		}

		position.col++
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

			i = programJumps[len(programJumps)-1]
			programJumps = removeLastEl(programJumps)

			tok.Token = token.JumpBackward
			tok.StartOfLoop = i
			tape = append(tape, tok)

			tape[i].StartOfLoop = programCounter
		default:
			programCounter--

		}
		programCounter++

	}

	return tape
}

func InitScanner(fileName string) {
	srcFileName = fileName
	position.line = 1
	position.col = 0
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
	return p[:len(p)-1]
}

func abortCompile(err Err) {
	switch err {
	case ErrFilePermission:
		fmt.Fprintf(os.Stderr, "Error: can't open file '%s', permission denied\n", srcFileName)
		os.Exit(2)
	case ErrFileDoesNotExist:
		fmt.Fprintf(os.Stderr, "Error: can't open '%s', file does not exist\n", srcFileName)
		os.Exit(3)
	case ErrFileOpen:
		fmt.Fprintf(os.Stderr, "Error: can't open file '%s'\n", srcFileName)
		os.Exit(4)
	case ErrBackJumpBeforeForward:
		fmt.Fprintf(os.Stderr, "Error in file \"%s\" at %d:%d: ']' occurs before '['\n",
			srcFileName, position.line, position.col)
		os.Exit(5)
	}
}
