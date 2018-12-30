package stack

import (
	"fmt"
)

type (
	// ArrayStack 구조체는 내부 자료구조로 슬라이스를 사용한다
	ArrayStack struct {
		stack  []interface{}
		size   int
		cursor int
	}

	// Stack 인터페이스
	Stack interface {
		IsFull() bool
		IsEmpty() bool
		Push(item interface{}) Stack
		Pop() interface{}
		Peek() interface{}
	}
)

func New(size int) Stack {
	s := ArrayStack{
		stack:  make([]interface{}, size),
		size:   size,
		cursor: -1,
	}
	return &s
}

func (this *ArrayStack) IsFull() bool {
	if this.cursor == this.size-1 {
		return true
	}
	return false
}

func (this *ArrayStack) IsEmpty() bool {
	if this.cursor == -1 {
		return true
	}
	return false
}

func (this *ArrayStack) Push(item interface{}) Stack {
	if this.IsFull() {
		fmt.Println("스택이 가득 찼습니다")
		return nil
	}
	this.stack[this.cursor+1] = item
	this.cursor++
	return this
}

func (this *ArrayStack) Pop() interface{} {
	if this.IsEmpty() {
		fmt.Println("스택에 자료가 없습니다")
		return nil
	}
	ret := this.stack[this.cursor]
	this.stack[this.cursor] = nil
	this.cursor--
	return ret
}

func (this *ArrayStack) Peek() interface{} {
	if this.IsEmpty() {
		fmt.Println("스택에 자료가 없습니다")
		return nil
	}
	return this.stack[this.cursor]
}
