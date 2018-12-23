package stack

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_New(t *testing.T) {
	// when
	stack := New(10)

	// then
	assert.NotNil(t, stack)
}

func Test_Push_Pop(t *testing.T) {
	assert := assert.New(t)

	// given
	stack := New(10)

	// when
	stack.Push(5)
	stack.Push(4)
	stack.Push(3)
	stack.Push(2)

	// then
	assert.Equal(2, stack.Pop())
	assert.Equal(3, stack.Pop())
}

func Test_스택이_가득차면_추가불가(t *testing.T) {
	assert := assert.New(t)

	// given
	stack := New(3)
	stack.Push(5)
	stack.Push(4)
	stack.Push(3)

	// when
	pushed := stack.Push(0) // message: stack is full

	// then
	assert.Nil(pushed)
}

func Test_IsFull(t *testing.T) {
	assert := assert.New(t)

	// given
	stack := New(3)

	// when, then
	stack.Push(5)
	stack.Push(4)
	assert.False(stack.IsFull())

	stack.Push(3)
	assert.True(stack.IsFull())
}

func Test_스택에_자료넣고_모두확인(t *testing.T) {
	assert := assert.New(t)

	// given
	stack := New(5)
	stack.Push(5)
	stack.Push(4)
	stack.Push(3)
	stack.Push(2)

	// when
	one := stack.Pop()
	two := stack.Pop()
	three := stack.Pop()
	four := stack.Pop()
	five := stack.Pop()

	// then
	assert.Equal(2, one)
	assert.Equal(3, two)
	assert.Equal(4, three)
	assert.Equal(5, four)
	assert.Nil(five)
}

func Test_Pop_빈스택은_nil_반환(t *testing.T) {
	// given
	stack := New(10)

	//when
	ret := stack.Pop() // message: stack is null

	// then
	assert.Nil(t, ret)
}

func Test_Peek_자료를_꺼내지않고_확인만한다(t *testing.T) {
	assert := assert.New(t)

	// given
	stack := New(5)
	stack.Push(5)
	stack.Push(4)
	stack.Push(3)

	// when then
	assert.Equal(3, stack.Peek())
	assert.Equal(3, stack.Peek())
}

func Test_Peek_실행시_빈스택은_nil_반환(t *testing.T) {
	// given
	stack := New(10)

	// when
	ret := stack.Peek() // message: stack is null

	// then
	assert.Nil(t, ret)
}
