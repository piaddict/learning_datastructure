package stack

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_New_스택생성(t *testing.T) {
	stack := New(10)
	stack.Push(5)
	stack.Push(4)
	stack.Push(3)
	stack.Push(2)

	assert.NotNil(t, stack)
	assert.Equal(t, 2, stack.Pop())
	assert.Equal(t, 3, stack.Pop())
}

func Test_Push_빈스택_자료4개_추가(t *testing.T) {
	stack := New(10)
	stack.Push(5)
	stack.Push(4)
	stack.Push(3)
	stack.Push(2)

	assert.Equal(t, 2, stack.Peek())
	assert.Equal(t, 2, stack.Peek())
}

func Test_Push_가득찬_스택_자료추가시_panic(t *testing.T) {
	stack := New(3)
	stack.Push(5)
	stack.Push(4)
	assert.False(t, stack.IsFull())
	stack.Push(3)
	assert.True(t, stack.IsFull())
	assert.Nil(t, stack.Push(2))
}

func Test_Pop_스택자료를_Pop(t *testing.T) {
	stack := New(5)
	stack.Push(5)
	stack.Push(4)
	stack.Push(3)
	stack.Push(2)
	stack.Push(1)

	assert.Equal(t, 1, stack.Pop())
	assert.Equal(t, 2, stack.Pop())
	assert.Equal(t, 3, stack.Pop())
	assert.Equal(t, 4, stack.Pop())
	assert.Equal(t, 5, stack.Pop())
}

func Test_Pop_빈스택_Pop실행시_패닉(t *testing.T) {
	stack := New(10)
	assert.True(t, stack.IsEmpty())
	assert.Nil(t, stack.Pop())
}

func Test_Pop_스택자료를_Peek(t *testing.T) {
	stack := New(5)
	stack.Push(5)
	stack.Push(4)
	stack.Push(3)
	stack.Push(2)
	stack.Push(1)

	assert.Equal(t, 1, stack.Peek())
	assert.Equal(t, 1, stack.Peek())
	assert.Equal(t, 1, stack.Peek())
	assert.Equal(t, 1, stack.Peek())
	assert.Equal(t, 1, stack.Peek())
}

func Test_Peek_빈스택_Peek실행시_패닉(t *testing.T) {
	stack := New(10)
	assert.True(t, stack.IsEmpty())
	assert.Nil(t, stack.Peek())
}
