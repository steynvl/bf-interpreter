package scanner

type Err int

// Errors that can occur while scanning bf file.
const (
	ErrFilePermission        Err = 0
	ErrFileDoesNotExist      Err = 1
	ErrFileOpen              Err = 2
	ErrBackJumpBeforeForward Err = 3
)

// SourcePos represents the current
// place (position) in the source file.
type SourcePos struct {
	line int
	col  int
}
