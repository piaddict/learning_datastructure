package list

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_New(t *testing.T) {
	// given, when

	list := New()

	// then
	assert.NotNil(t, list)
}

func Test_AddFirst_Get(t *testing.T) {
	assert := assert.New(t)

	// given
	list := New()
	list.AddFirst("피카츄")
	list.AddFirst("라이츄")
	list.AddFirst("치코리타")

	// when
	data1 := list.Get(0)
	data2 := list.Get(1)
	data3 := list.Get(2)
	data4 := list.Get(3)

	// then
	assert.Equal("치코리타", data1)
	assert.Equal("라이츄", data2)
	assert.Equal("피카츄", data3)
	assert.Nil(data4)
}

func Test_Add_특정노드_뒤_새노드를_추가한다(t *testing.T) {
	assert := assert.New(t)

	// given
	list := New()
	list.AddFirst("피카츄")
	list.AddFirst("라이츄")

	// when
	list.Add(0, "파이리")

	// then
	assert.Equal("라이츄", list.Get(0))
	assert.Equal("파이리", list.Get(1))
	assert.Equal("피카츄", list.Get(2))
}

func Test_AddLast_Get(t *testing.T) {
	assert := assert.New(t)

	// given
	list := New()
	list.AddLast("피카츄")
	list.AddLast("라이츄")
	list.AddLast("치코리타")

	// when
	data1 := list.Get(0)
	data2 := list.Get(1)
	data3 := list.Get(2)
	data4 := list.Get(3)

	// then
	assert.Equal("피카츄", data1)
	assert.Equal("라이츄", data2)
	assert.Equal("치코리타", data3)
	assert.Nil(data4)
}

func Test_Size(t *testing.T) {
	assert := assert.New(t)

	// given
	list := New()
	list.AddLast("피카츄")
	list.AddLast("라이츄")

	// when, then
	assert.Equal(2, list.Size())
	list.AddLast("파이리")
	assert.Equal(3, list.Size())
}

func Test_Remove(t *testing.T) {
	assert := assert.New(t)

	// given
	list := New()
	list.AddFirst("피카츄")
	list.AddLast("라이츄")
	list.AddLast("치코리타")
	list.Add(1, "푸린") // [피카츄, 라이츄, 푸린, 치코리타]

	// when
	list.Remove(1) // [피카츄, 푸린, 치코리타]

	// then
	assert.Equal("피카츄", list.Get(0))
	assert.Equal("푸린", list.Get(1))
	assert.Equal("치코리타", list.Get(2))
}

func Test_Remove_첫요소_AddFirst_새요소_순서검사(t *testing.T) {
	assert := assert.New(t)

	// given
	list := New()
	list.AddFirst("피카츄")
	list.AddLast("라이츄")
	list.AddLast("치코리타")
	list.Add(1, "푸린") // [피카츄, 라이츄, 푸린, 치코리타]

	// when
	list.Remove(0)       // [라이츄, 푸린, 치코리타]
	list.AddFirst("아보크") // [아보크, 라이츄, 푸린, 치코리타]

	// then
	assert.Equal("아보크", list.Get(0))
	assert.Equal("라이츄", list.Get(1))
	assert.Equal("푸린", list.Get(2))
	assert.Equal("치코리타", list.Get(3))
}

func Test_Remove_마지막요소_AddLast_새요소_순서검사(t *testing.T) {
	assert := assert.New(t)

	// given
	list := New()
	list.AddFirst("피카츄")
	list.AddLast("라이츄")
	list.AddLast("치코리타")
	list.Add(1, "푸린") // [피카츄, 라이츄, 푸린, 치코리타]

	// when
	list.Remove(3)       // [피카츄, 라이츄, 푸린]
	list.AddLast("갸라도스") // [피카츄, 라이츄, 푸린, 갸라도스]

	// then
	assert.Equal("피카츄", list.Get(0))
	assert.Equal("라이츄", list.Get(1))
	assert.Equal("푸린", list.Get(2))
	assert.Equal("갸라도스", list.Get(3))
}
