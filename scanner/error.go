package scanner

type Error int

const (
	ERR_OPENING_SOURCE_FILE      = 0
	ERR_BACK_JUMP_BEFORE_FORWARD = 1
	ERR_UNKNOWN_CHARACTER        = 2
)
