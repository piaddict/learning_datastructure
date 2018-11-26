package queue

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNew_큐생성(t *testing.T) {
	// when
	queue := New(5)

	// then
	assert.NotNil(t, queue)
}

func TestAdd_큐_요소추가(t *testing.T) {
	// given
	queue := New(4) // actual size: 3

	// when
	queue.Add(5)
	queue.Add(10)

	// then
	assert.Equal(
		t,
		&ArrayQueue{queue: []interface{}{nil, 5, 10, nil}, size: 4, front: 0, rear: 2},
		queue,
	)
}

func TestAdd_가득찬_큐에_요소추가(t *testing.T) {
	// given
	queue := New(4) // actual size: 3

	// when
	queue.Add(5)
	queue.Add(6)
	queue.Add(7)
	queue.Add(8) // message: queue is full

	// then
	assert.Equal(
		t,
		&ArrayQueue{queue: []interface{}{nil, 5, 6, 7}, size: 4, front: 0, rear: 3},
		queue,
	)
}

func TestRemove_큐_요소삭제(t *testing.T) {
	// given
	queue := &ArrayQueue{queue: []interface{}{nil, 3, 7, nil}, size: 4, front: 0, rear: 2}

	// when
	removed := queue.Remove()

	// then
	assert.Equal(t, 3, removed)
	assert.Equal(
		t,
		&ArrayQueue{queue: []interface{}{nil, nil, 7, nil}, size: 4, front: 1, rear: 2},
		queue,
	)
}

func TestRemove_비어있는큐_요소삭제(t *testing.T) {
	// given
	queue := &ArrayQueue{queue: []interface{}{nil, 3, 7, nil}, size: 4, front: 0, rear: 2}

	// when
	one := queue.Remove()
	two := queue.Remove()
	three := queue.Remove()

	// then
	assert.Equal(t, 3, one)
	assert.Equal(t, 7, two)
	assert.Equal(t, nil, three)
	assert.Equal(
		t,
		&ArrayQueue{queue: []interface{}{nil, nil, nil, nil}, size: 4, front: 2, rear: 2},
		queue,
	)
}

func TestPeek_큐_현재요소_확인(t *testing.T) {
	// given
	queue := &ArrayQueue{queue: []interface{}{nil, 5, 4, 3, 2, nil}, size: 6, front: 0, rear: 4}

	// when
	peek := queue.Peek()

	// then
	assert.Equal(t, 5, peek)
}

func TestIsFull_큐_가득차있는지(t *testing.T) {
	// given
	queue := &ArrayQueue{queue: []interface{}{nil, 2, 4, 6, 8}, size: 5, front: 0, rear: 4}

	// when then
	assert.True(t, queue.IsFull())
}

func TestIsEmpty_큐_비어있는지(t *testing.T) {
	// given
	queueFoo := New(4)
	queueBar := &ArrayQueue{queue: []interface{}{nil, nil, nil}, size: 3, front: 0, rear: 0}

	// when then
	assert.True(t, queueFoo.IsEmpty())
	assert.True(t, queueBar.IsEmpty())
}
