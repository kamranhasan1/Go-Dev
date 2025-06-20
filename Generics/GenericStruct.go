package main

import "fmt"

type stack[T any] struct {
	elements T
}

func main() {
	myStack := stack[[]string]{
		elements: []string{"golang"},
	}
	fmt.Println(myStack)
}