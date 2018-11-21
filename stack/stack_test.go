package stack

import (
	"fmt"
)

func Example_New_스택생성() {
	stack := New(10)
	stack.Push(5)
	stack.Push(4)
	stack.Push(3)
	stack.Push(2)
	fmt.Println(stack != nil)
	fmt.Println(stack.Pop())
	fmt.Println(stack.Pop())
	// Output:
	// true
	// 2
	// 3
}

func Example_Push_빈스택_자료4개_추가() {
	stack := New(10)
	stack.Push(5)
	stack.Push(4)
	stack.Push(3)
	stack.Push(2)
	fmt.Println(stack.Peek())
	fmt.Println(stack.Peek())
	// Output:
	// 2
	// 2
}

func Example_Push_가득찬_스택_자료추가시_panic() {
	stack := New(3)
	stack.Push(5)
	stack.Push(4)
	fmt.Println(stack.IsFull())
	stack.Push(3)
	fmt.Println(stack.IsFull())
	stack.Push(2)
	// Output:
	// false
	// true
	// 스택이 가득 찼습니다
}

func Example_Pop_스택자료를_Pop() {
	stack := New(5)
	stack.Push(5)
	stack.Push(4)
	stack.Push(3)
	stack.Push(2)
	stack.Push(1)
	fmt.Println(stack.Pop())
	fmt.Println(stack.Pop())
	fmt.Println(stack.Pop())
	fmt.Println(stack.Pop())
	fmt.Println(stack.Pop())
	// Output:
	// 1
	// 2
	// 3
	// 4
	// 5
}

func Example_Pop_빈스택_Pop실행시_패닉() {
	stack := New(10)
	fmt.Println(stack.IsEmpty())
	stack.Pop()
	// Output:
	// true
	// 스택에 자료가 없습니다
}

func Example_Pop_스택자료를_Peek() {
	stack := New(5)
	stack.Push(5)
	stack.Push(4)
	stack.Push(3)
	stack.Push(2)
	stack.Push(1)
	fmt.Println(stack.Peek())
	fmt.Println(stack.Peek())
	fmt.Println(stack.Peek())
	fmt.Println(stack.Peek())
	fmt.Println(stack.Peek())
	// Output:
	// 1
	// 1
	// 1
	// 1
	// 1
}

func Example_Peek_빈스택_Peek실행시_패닉() {
	stack := New(10)
	fmt.Println(stack.IsEmpty())
	stack.Peek()
	// Output:
	// true
	// 스택에 자료가 없습니다
}
