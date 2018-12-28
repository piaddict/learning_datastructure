package bstree

type (
	BSTree struct {
		root *Node
	}

	Node struct {
		key       int
		duplicate int
		left      *Node
		right     *Node
	}
)

func New() *BSTree {
	return &BSTree{}
}

func NewWithRoot(root *Node) *BSTree {
	return &BSTree{root: root}
}

func (bst *BSTree) Inorder() []int {
	var list []int
	bst.traversal(&list, bst.root)
	return list
}

func (bst *BSTree) traversal(list *[]int, node *Node) {
	if node == nil {
		return
	}
	bst.traversal(list, node.left)
	*list = append(*list, node.key)
	bst.traversal(list, node.right)
}

func (bst *BSTree) Find(key int) *Node {
	node, _ := bst.findNodeAndParent(key)
	return node
}

func (bst *BSTree) FindMin() *Node {
	node, _ := bst.findSubtreeMin(bst.root)
	return node
}

func (bst *BSTree) findSubtreeMin(node *Node) (*Node, *Node) {
	var parent *Node
	for {
		if node.left == nil {
			break
		}
		parent = node
		node = node.left
	}

	return node, parent
}

func (bst *BSTree) FindMax() *Node {
	node, _ := bst.findSubtreeMax(bst.root)
	return node
}

func (bst *BSTree) findSubtreeMax(node *Node) (*Node, *Node) {
	var parent *Node
	for {
		if node.right == nil {
			break
		}
		parent = node
		node = node.right
	}

	return node, parent
}

func (bst *BSTree) findNodeAndParent(key int) (*Node, *Node) {
	node := bst.root
	var parent *Node

	for node != nil {
		if node.key == key {
			return node, parent
		}

		parent = node
		if node.key > key {
			node = node.left
		} else {
			node = node.right
		}
	}

	return nil, nil
}

func (bst *BSTree) Insert(key int) {
	node := bst.root
	if node == nil {
		bst.root = &Node{key: key}
	} else {
		bst.insertSubtree(key)
	}
}

func (bst *BSTree) insertSubtree(key int) {
	var nodePtr **Node
	node := bst.root
	newNode := &Node{key: key}

	for node != nil {
		if node.key == key {
			node.duplicate++
			return
		} else if node.key > key {
			nodePtr = &node.left
		} else {
			nodePtr = &node.right
		}
		node = *nodePtr
	}
	*nodePtr = newNode
}

func (bst *BSTree) Remove(key int) *Node {
	node, parent := bst.findNodeAndParent(key)
	if node == nil {
		return nil
	}

	leftOrRight := bst.findChildPointerOfParent(node, parent)
	if node.left != nil && node.right != nil {
		max, parent := bst.findSubtreeMin(node.right)
		max.left = node.left
		max.right = node.right
		*leftOrRight = max

		if parent != nil {
			*bst.findChildPointerOfParent(max, parent) = nil
		}
	} else {
		child := bst.findOnlyChild(node)
		*leftOrRight = child
	}

	return node
}

func (bst *BSTree) findChildPointerOfParent(node *Node, parent *Node) **Node {
	if parent.left == node {
		return &parent.left
	}
	return &parent.right
}

func (bst *BSTree) findOnlyChild(parent *Node) (child *Node) {
	if parent.left != nil {
		return parent.left
	}
	return parent.right
}
