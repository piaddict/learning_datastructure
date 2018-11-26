package stack

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNew_스택생성(t *testing.T) {
	// when
	stack := New(10)

	// then
	assert.NotNil(t, stack)
}

func TestPush_빈스택_자료4개_추가(t *testing.T) {
	// given
	stack := New(10)

	// when
	stack.Push(5)
	stack.Push(4)
	stack.Push(3)
	stack.Push(2)

	// then
	assert.Equal(t, 2, stack.Peek())
	assert.Equal(t, 2, stack.Peek())
}

func TestPush_가득찬_스택_자료추가시_panic(t *testing.T) {
	// given
	stack := New(3)

	// when
	stack.Push(5)
	stack.Push(4)
	stack.Push(3)

	// then
	assert.True(t, stack.IsFull())
	assert.Nil(t, stack.Push(2))
}

func TestPop_스택자료를_Pop(t *testing.T) {
	// given
	stack := &ArrayStack{
		stack:  []interface{}{5, 4, 3, 2, nil},
		size:   5,
		cursor: 3,
	}

	// when
	one := stack.Pop()
	two := stack.Pop()
	three := stack.Pop()
	four := stack.Pop()
	five := stack.Pop()

	// then
	assert.Equal(t, 2, one)
	assert.Equal(t, 3, two)
	assert.Equal(t, 4, three)
	assert.Equal(t, 5, four)
	assert.Equal(t, nil, five)
}

func TestPop_빈스택_Pop실행시_nil_반환(t *testing.T) {
	// given
	stack := New(10)

	//when
	ret := stack.Pop() // message: stack is null

	// then
	assert.Nil(t, ret)
}

func TestPop_스택자료를_Peek(t *testing.T) {
	// given
	stack := &ArrayStack{
		stack:  []interface{}{5, 4, 3, nil, nil},
		size:   5,
		cursor: 2,
	}

	// when then
	assert.Equal(t, 3, stack.Peek())
	assert.Equal(t, 3, stack.Peek())
}

func TestPeek_빈스택_Peek실행시_nil_반환(t *testing.T) {
	// given
	stack := New(10)

	// when
	ret := stack.Peek() // message: stack is null

	// then
	assert.Nil(t, ret)
}
