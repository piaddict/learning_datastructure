package heap

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_New(t *testing.T) {
	// given
	heap := New()

	// when, then
	assert.NotNil(t, heap)
}

func Test_Insert(t *testing.T) {
	// given
	heap := New(1, 2, 3, 4, 5, 6, 7)

	// when
	heap.Insert(0)

	// then
	assert.EqualValues(t, []int{0, 1, 3, 2, 5, 6, 7, 4}, heap.Print())
}

func Test_Remove(t *testing.T) {
	assert := assert.New(t)

	// given
	heap := New(1, 2, 3, 4, 5, 6, 7)

	// when
	ret := heap.Remove()

	// then
	assert.Equal(1, ret)
	assert.EqualValues([]int{2, 4, 3, 7, 5, 6}, heap.Print())
}

func Test_Remove_오른쪽_child가_작은경우(t *testing.T) {
	assert := assert.New(t)

	// given
	heap := New(1, 3, 2, 4, 5, 6, 7)

	// when
	heap.Remove()

	// then
	assert.EqualValues([]int{2, 3, 6, 4, 5, 7}, heap.Print())
}

func Test_Remove_leaf까지_도달하지않음(t *testing.T) {
	assert := assert.New(t)

	// given
	heap := New(3, 10, 4, 12, 22, 7, 5)

	// when
	heap.Remove()

	// then
	assert.EqualValues([]int{4, 10, 5, 12, 22, 7}, heap.Print())
}

func Test_Peek(t *testing.T) {
	assert := assert.New(t)

	// given
	heap := New(20, 12, 4, 1, 0, 33)

	// when
	peek1 := heap.Peek()
	peek2 := heap.Peek()

	// then
	assert.Equal(20, peek1)
	assert.Equal(20, peek2)
}

func Test_IsEmpty(t *testing.T) {
	assert := assert.New(t)

	// given
	heap := New()

	// when, then
	assert.True(heap.IsEmpty())
	heap.Insert(100)
	assert.False(heap.IsEmpty())
	heap.Remove()
	assert.True(heap.IsEmpty())
}

func Test_Size(t *testing.T) {
	// given
	heap := New(20, 12, 4, 1, 0, 33)

	// when, then
	assert.Equal(t, 6, heap.Size())
}
