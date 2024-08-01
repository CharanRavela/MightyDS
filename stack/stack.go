// TODO: Need to add comments for functions
package stack

import "fmt"

type Stack struct {
	data []any

	pos     int
	maxSize int

	stats StackStats
}

func NewStack() *Stack {
	return &Stack{
		data: make([]any, 0),

		pos:     -1,
		maxSize: -1,

		stats: StackStats{},
	}
}

func NewBoundedStack(n int) *Stack {
	return &Stack{
		data:    make([]any, 0),
		pos:     -1,
		maxSize: n,
	}
}

func NewStackFromData(d []any) *Stack {
	return &Stack{
		data:    d,
		pos:     len(d) - 1,
		maxSize: -1,
	}
}

func NewBoundedStackFromData(n int, d []any) (*Stack, error) {
	if len(d) > n {
		return nil, fmt.Errorf("%w: data slice size: %v, stack size : %v", ErrDataSliceTooLarge, len(d), n)
	}

	return &Stack{
		data:    d,
		pos:     len(d) - 1,
		maxSize: n,
	}, nil
}

func (s *Stack) Push(i any) {
	s.data = append(s.data, i)
	s.pos += 1
}

func (s *Stack) Pop() (any, error) {
	if s.pos < 0 {
		return nil, ErrEmptyStack
	}

	value := s.data[s.pos]
	s.data = s.data[:s.pos]
	s.pos -= 1
	return value, nil
}

func (s *Stack) PopN(n int) (any, error) {
	// TODO: Implement PopN method
	// Takes number of pop operations to be performed as input
	// Perfoms `n` pops if stacks has those many elements left in the stack
	return nil, nil
}

func (s *Stack) Peek() (any, error) {
	if s.pos < 0 {
		return nil, ErrEmptyStack
	}

	return s.data[s.pos], nil
}

func (s *Stack) PeekIndex(i int) (any, error) {
	// Calculate the index from the bottom of the stack
	bottomIndex := s.pos - i
	if bottomIndex < 0 || bottomIndex > s.pos {
		return nil, fmt.Errorf("index %d out of bounds. Current stack size is %d", i, s.pos+1)
	}
	return s.data[bottomIndex], nil
}

func (s *Stack) Clear() {
	s.data = make([]any, 0)
	s.pos = -1
}

func (s *Stack) Data() []any {
	return s.data[:s.pos+1]
}

func (s *Stack) Length() int {
	return s.pos + 1
}

func (s *Stack) Stats() StackStats {
	return s.stats
}

func (s *Stack) MaxStackSize() int {
	return s.maxSize
}
