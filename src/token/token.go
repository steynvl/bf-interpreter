// Package token contains type Token that represent
// the types of tokens the scanner recognises.

package token

type Token int

const (
	IncrementPointer Token = 0
	DecrementPointer Token = 1
	IncrementByte    Token = 2
	DecrementByte    Token = 3
	OutputByte       Token = 4
	InputByte        Token = 5
	JumpForward      Token = 6
	JumpBackward     Token = 7
	None             Token = 8
)
