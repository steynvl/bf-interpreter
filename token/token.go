package token

type Token int

const (
	INCREMENT_THE_POINTER Token = 0
	DECREMENT_THE_POINTER Token = 1
	INCREMENT_THE_BYTE    Token = 2
	DECREMENT_THE_BYTE    Token = 3
	OUTPUT_BYTE           Token = 4
	INPUT_A_BYTE          Token = 5
	JUMP_FORWARD          Token = 6
	JUMP_BACKWARD         Token = 7
	TOKEN_EOF             Token = 8
)
