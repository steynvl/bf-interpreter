package scanner

type Err int

const (
	ErrFilePermission        Err = 0
	ErrFileDoesNotExist      Err = 1
	ErrFileOpen              Err = 2
	ErrBackJumpBeforeForward Err = 3
	ErrUnknownCharacter      Err = 4
)

type SourcePos struct {
	line int
	col  int
}
