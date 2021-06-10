package main

import (
	"fmt"

	"github.com/fishwin/mylib/data_struct/stack"
)

func main() {
	s := stack.NewStack()
	s.Push(1)
	s.Push(2)
	s.Push(3)

	fmt.Println(s.Pop())
	fmt.Println(s.Pop())
	fmt.Println(s.Length())
	fmt.Println(s.Pop())
	fmt.Println(s.Length())
}
