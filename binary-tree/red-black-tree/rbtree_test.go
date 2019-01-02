package rbtree

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_New(t *testing.T) {
	// given
	tree := New()

	// when, then
	assert.NotNil(t, tree)
}

func Test_Insert_Inorder(t *testing.T) {
	assert := assert.New(t)

	// given
	tree := New()
	tree.Insert(5)
	tree.Insert(2)
	tree.Insert(1)
	tree.Insert(4)
	tree.Insert(7)
	tree.Insert(22)
	tree.Insert(33)
	tree.Insert(11)
	tree.Insert(12)
	tree.Insert(13)
	tree.Insert(14)
	tree.Insert(15)
	tree.Insert(34)
	tree.Insert(35)
	tree.Insert(36)
	tree.Insert(37)
	tree.Insert(38)

	// when
	n5 := tree.Find(5)
	n2 := tree.Find(2)
	n1 := tree.Find(1)
	n4 := tree.Find(4)
	n7 := tree.Find(7)
	n22 := tree.Find(22)
	n33 := tree.Find(33)
	n11 := tree.Find(11)
	n12 := tree.Find(12)
	n13 := tree.Find(13)
	n14 := tree.Find(14)
	n15 := tree.Find(15)
	n34 := tree.Find(34)
	n35 := tree.Find(35)
	n36 := tree.Find(36)
	n37 := tree.Find(37)
	n38 := tree.Find(38)

	// then
	assert.Equal(n2, n5.left)
	assert.Equal(n1, n2.left)
	assert.Equal(n4, n2.right)
	assert.Equal(n13, n5.right)
	assert.Equal(n11, n13.left)
	assert.Equal(n7, n11.left)
	assert.Equal(n12, n11.right)
	assert.Equal(n34, n13.right)
	assert.Equal(n22, n34.left)
	assert.Equal(n14, n22.left)
	assert.Equal(n15, n14.right)
	assert.Equal(n33, n22.right)
	assert.Equal(n36, n34.right)
	assert.Equal(n35, n36.left)
	assert.Equal(n37, n36.right)
	assert.Equal(n38, n37.right)

	assert.True(n5.isBlack)
	assert.True(n2.isBlack)
	assert.True(n1.isBlack)
	assert.True(n4.isBlack)
	assert.False(n13.isBlack)
	assert.True(n11.isBlack)
	assert.True(n7.isBlack)
	assert.True(n12.isBlack)
	assert.True(n34.isBlack)
	assert.False(n22.isBlack)
	assert.True(n14.isBlack)
	assert.False(n15.isBlack)
	assert.True(n33.isBlack)
	assert.False(n36.isBlack)
	assert.True(n35.isBlack)
	assert.True(n37.isBlack)
	assert.False(n38.isBlack)

	assert.EqualValues([]int{1, 2, 4, 5, 7, 11, 12, 13, 14, 15, 22, 33, 34, 35, 36, 37, 38}, tree.Inorder())
}
