package stack

/*
Package stack provides a simple stack data structure implementation.

A stack is a collection of elements with two principal operations:
- Push: Adds an element to the top of the stack.
- Pop: Removes the element from the top of the stack and returns it.

Stacks operate on a Last-In-First-Out (LIFO) principle, meaning the most recently added element is the first one to be removed.

Types:
- Stack: Represents a stack data structure with methods to manipulate and query the stack.

Functions:
- NewStack: Creates and returns a new, empty stack.
- NewStackWithCapacity: Creates and returns a new stack with a specified initial capacity.
- NewStackFromData: Creates and returns a new stack initialized with a slice of elements.

Methods on *Stack:
- Push(i any): Adds the element `i` to the top of the stack.
- Pop() (any, error): Removes and returns the element from the top of the stack. Returns an error if the stack is empty.
- Data() []any: Returns a slice containing the elements currently in the stack from bottom to top.
- Length() int: Returns the number of elements currently in the stack.

Example usage:

    stack := NewStack()
    stack.Push(1)
    stack.Push(2)
    stack.Push(3)

    value, err := stack.Pop()
    if err != nil {
        fmt.Println("Error:", err)
    } else {
        fmt.Println("Popped value:", value)
    }

    fmt.Println("Current stack data:", stack.Data())
    fmt.Println("Current stack length:", stack.Length())
*/
