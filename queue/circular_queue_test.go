package queue

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_New(t *testing.T) {
	// when
	queue := New(5)

	// then
	assert.NotNil(t, queue)
}

func Test_Add_Remove(t *testing.T) {
	assert := assert.New(t)

	// given
	queue := New(4) // actual size: 3

	// when
	queue.Add(5)
	queue.Add(10)

	// then
	assert.Equal(5, queue.Remove())
	assert.Equal(10, queue.Remove())
}

func Test_가득찬_큐_요소추가_불가(t *testing.T) {
	assert := assert.New(t)

	// given
	queue := New(4) // actual size: 3

	// when
	queue.Add(5)
	queue.Add(6)
	queue.Add(7)
	queue.Add(8) // message: queue is full

	// then
	assert.Equal(5, queue.Remove())
	assert.Equal(6, queue.Remove())
	assert.Equal(7, queue.Remove())
	assert.Equal(nil, queue.Remove())
}

func Test_비어있는_큐_요소삭제_불가(t *testing.T) {
	assert := assert.New(t)

	// given
	queue := New(4) // actual size: 3

	// when
	one := queue.Remove() // message: queue is empty
	two := queue.Remove() // message: queue is empty

	// then
	assert.Equal(nil, one)
	assert.Equal(nil, two)
}

func Test_Peek_현재요소_꺼내지않고_확인(t *testing.T) {
	assert := assert.New(t)

	// given
	queue := New(6) // actual size: 5
	queue.Add(5)
	queue.Add(4)
	queue.Add(3)
	queue.Add(2)

	// when
	peek1 := queue.Peek()
	peek2 := queue.Peek()

	// then
	assert.Equal(5, peek1)
	assert.Equal(5, peek2)
}

func Test_Peek_현재요소_꺼내지않고_확인_비어있으면_nil(t *testing.T) {
	assert := assert.New(t)

	// given
	queue := New(6) // actual size: 5

	// when
	peek := queue.Peek() // message: queue is empty

	// then
	assert.Nil(peek)
}

func Test_IsFull(t *testing.T) {
	// given
	queue := New(5) // actual size: 4

	// when then
	queue.Add(2)
	queue.Add(4)
	queue.Add(6)
	assert.False(t, queue.IsFull())
	queue.Add(8)
	assert.True(t, queue.IsFull())
	queue.Remove()
	assert.False(t, queue.IsFull())
}

func Test_IsEmpty(t *testing.T) {
	// given
	queue := New(4) // actual size: 3

	// when then
	assert.True(t, queue.IsEmpty())
	queue.Add(110)
	assert.False(t, queue.IsEmpty())
	queue.Remove()
	assert.True(t, queue.IsEmpty())
}
