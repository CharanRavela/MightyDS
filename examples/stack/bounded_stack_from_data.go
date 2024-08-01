package main

import (
	"fmt"

	mds "github.com/CharanRavela/MightyDS/stack"
)

func bounded_stack_from_data() {
	// TODO: Need to implement Bounded stack from data and add usage examples
	stack := mds.NewStack()

	stack.Push(9)
	stack.Push([]int{5, 9, 4})
	stack.Push("MightyDS")

	fmt.Printf("Stack: %v, length: %v\n\n", stack.Data(), stack.Length())

	pop1, _ := stack.Pop()
	fmt.Printf("Stack: %v, length: %v\nLast popped element: %v\n\n", stack.Data(), stack.Length(), pop1)

	stack.Push("Let's Go MightyDS!")

	fmt.Printf("Stack before Peek: %v, length: %v\n\n", stack.Data(), stack.Length())
	peek1, _ := stack.Peek()
	fmt.Printf("Stack after peek: %v, length: %v\nPeek element result: %v\n\n", stack.Data(), stack.Length(), peek1)

	peekIndexValue, err := stack.PeekIndex(2)
	if err != nil {
		fmt.Println(err)
	}

	fmt.Printf("Peek Index 2 Yielded: %v\n\n", peekIndexValue)

	// TODO: Need to use PopN here once its implemented!
	stack.Pop()
	stack.Pop()
	stack.Pop()

	fmt.Printf("Stack after poping all the elements: %v, length: %v\n\n", stack.Data(), stack.Length())

	_, err = stack.Pop()
	fmt.Printf("Error due to empty stack pop: %v\n\n", err)
}
