package heap

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_PQ_New(t *testing.T) {
	// given, when
	pq := NewPriorityQueue()

	// then
	assert.NotNil(t, pq)
}
func Test_PQ_Enqueue_Dequeue(t *testing.T) {
	// given
	pq := NewPriorityQueue()

	// when
	pq.Enqueue(5)
	pq.Enqueue(3)
	pq.Enqueue(7)
	pq.Enqueue(2)

	// then
	assert.Equal(t, 2, pq.Dequeue())
	assert.Equal(t, 3, pq.Dequeue())
}

func Test_PQ_Peek(t *testing.T) {
	// given
	pq := NewPriorityQueue()

	// when
	pq.Enqueue(5)
	pq.Enqueue(3)

	// then
	assert.Equal(t, 3, pq.Peek())
	assert.Equal(t, 3, pq.Peek())
}
