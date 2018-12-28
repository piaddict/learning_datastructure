package bstree

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_New(t *testing.T) {
	// given
	bst := New()

	// when, then
	assert.NotNil(t, bst)
}

func Test_NewWithRoot(t *testing.T) {
	// given
	root := Node{key: 2}
	bst := NewWithRoot(&root)

	// when, then
	assert.NotNil(t, bst)
}

func Test_Inorder(t *testing.T) {
	// given
	rl := Node{key: 1}
	rr := Node{key: 3}
	root := Node{
		key:   2,
		left:  &rl,
		right: &rr,
	}
	bst := NewWithRoot(&root)

	// when
	list := bst.Inorder()

	// then
	assert.EqualValues(t, []int{1, 2, 3}, list)
}

func Test_Find(t *testing.T) {
	// given
	rll := Node{key: 1}
	rlr := Node{key: 3}
	rl := Node{
		key:   2,
		left:  &rll,
		right: &rlr,
	}
	rr := Node{key: 5}
	root := Node{
		key:   4,
		left:  &rl,
		right: &rr,
	}
	bst := NewWithRoot(&root)

	// when, then
	assert.Equal(t, &rlr, bst.Find(3))
	assert.Equal(t, &root, bst.Find(4))
	assert.Equal(t, &rr, bst.Find(5))
}

func Test_FindMin(t *testing.T) {
	// given
	rll := Node{key: 1}
	rlr := Node{key: 3}
	rl := Node{
		key:   2,
		left:  &rll,
		right: &rlr,
	}
	rr := Node{key: 5}
	root := Node{
		key:   4,
		left:  &rl,
		right: &rr,
	}
	bst := NewWithRoot(&root)

	// when, then
	assert.Equal(t, 1, bst.FindMin().key)

}

func Test_FindMax(t *testing.T) {
	// given
	rll := Node{key: 1}
	rlr := Node{key: 3}
	rl := Node{
		key:   2,
		left:  &rll,
		right: &rlr,
	}
	rr := Node{key: 5}
	root := Node{
		key:   4,
		left:  &rl,
		right: &rr,
	}
	bst := NewWithRoot(&root)

	// when, then
	assert.Equal(t, 5, bst.FindMax().key)
}

func Test_Insert(t *testing.T) {
	assert := assert.New(t)

	// given
	bst := New()

	// when, then
	bst.Insert(4)
	assert.EqualValues([]int{4}, bst.Inorder())
	bst.Insert(3)
	assert.EqualValues([]int{3, 4}, bst.Inorder())
	bst.Insert(7)
	assert.EqualValues([]int{3, 4, 7}, bst.Inorder())
	bst.Insert(1)
	assert.EqualValues([]int{1, 3, 4, 7}, bst.Inorder())
	bst.Insert(1)
	assert.Equal(1, bst.Find(1).duplicate)
}

func Test_Insert_노드연결여부확인(t *testing.T) {
	assert := assert.New(t)

	// given
	bst := New()

	// when, then
	bst.Insert(4)
	bst.Insert(3)
	bst.Insert(7)
	bst.Insert(1)

	// when
	node4 := bst.Find(4)
	node3 := bst.Find(3)
	node7 := bst.Find(7)
	node1 := bst.Find(1)

	// then
	assert.True(node4.left == node3)
	assert.True(node4.right == node7)
	assert.True(node3.left == node1)
}

func Test_Remove_자식이_없는_노드삭제(t *testing.T) {
	assert := assert.New(t)

	// given
	rrl := Node{key: 5} // target
	rrr := Node{key: 9}
	rl := Node{key: 2}
	rr := Node{
		key:   7,
		left:  &rrl,
		right: &rrr,
	}
	root := Node{
		key:   4,
		left:  &rl,
		right: &rr,
	}
	bst := BSTree{root: &root}

	// when
	result := bst.Remove(rrl.key)

	// then
	assert.Equal(&rrl, result)
	assert.Nil(rr.left)
	assert.EqualValues([]int{2, 4, 7, 9}, bst.Inorder())
}

func Test_Remove_자식이_하나인_노드삭제(t *testing.T) {
	assert := assert.New(t)

	// given
	rrl := Node{key: 5}
	rl := Node{key: 2}
	rr := Node{ // target
		key:  7,
		left: &rrl,
	}
	root := Node{
		key:   4,
		left:  &rl,
		right: &rr,
	}
	bst := BSTree{root: &root}

	// when
	result := bst.Remove(rr.key)

	// then
	assert.Equal(&rr, result)
	assert.EqualValues([]int{2, 4, 5}, bst.Inorder())
}

func Test_Remove_자식이_두개인_노드삭제(t *testing.T) {
	assert := assert.New(t)

	// given
	rrrl := Node{key: 8}
	rrrr := Node{key: 11}
	rrl := Node{key: 5}
	rrr := Node{
		key:   9,
		left:  &rrrl,
		right: &rrrr,
	}
	rl := Node{key: 2}
	rr := Node{ // target
		key:   7,
		left:  &rrl,
		right: &rrr,
	}
	root := Node{
		key:   4,
		left:  &rl,
		right: &rr,
	}
	bst := BSTree{root: &root}

	// when
	result := bst.Remove(rr.key)

	// then
	assert.Equal(&rr, result)
	assert.EqualValues([]int{2, 4, 5, 8, 9, 11}, bst.Inorder())
}
