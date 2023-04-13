package main

import(
	"fmt"
	"sample-app/stack"
)

func main(){
	var stack stack.IStack = &stack.Stack{}
	stack.Init(8)
	fmt.Println(stack.Empty()) //true
	stack.Push(1)
	stack.Push(2)
	stack.Push(3)
	stack.Push(1)
	stack.Push(2)
	stack.Push(3)
	
	stack.Push(1)
	stack.Push(2)
	stack.Push(3)
	fmt.Println(stack.Peek()) // 3, nil
	fmt.Println(stack.Pop()) // 3, nil
	fmt.Println(stack.Peek()) // 2, nil
	fmt.Println(stack.Pop()) // 2, nil
	fmt.Println(stack.Pop()) // 1, nil
	fmt.Println(stack.Pop()) // 0, stack is empty
}