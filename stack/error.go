package stack

type StackError string

func (e StackError) Error() string {
	return string(e)
}

const (
	ErrEmptyStack        = StackError("stack is empty")
	ErrFullStack         = StackError("stack has reached its maximum size")
	ErrDataSliceTooLarge = StackError("data slice length exceeds the maximum stack size")
)
