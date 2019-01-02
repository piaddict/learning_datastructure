package rbtree

type (
	RBTree struct {
		root *Node
	}

	Node struct {
		key       int
		duplicate int
		parent    *Node
		left      *Node
		right     *Node
		isBlack   bool
	}
)

func New() *RBTree {
	return &RBTree{}
}

func (rbt *RBTree) Inorder() []int {
	var list []int
	rbt.traversal(&list, rbt.root)
	return list
}

func (rbt *RBTree) traversal(list *[]int, node *Node) {
	if node == nil {
		return
	}

	rbt.traversal(list, node.left)
	*list = append(*list, node.key)
	rbt.traversal(list, node.right)
}

func (rbt *RBTree) Find(key int) *Node {
	node, _ := rbt.findNodeAndParent(key)
	return node
}

func (rbt *RBTree) FindMin() *Node {
	node, _ := rbt.findSubtreeMin(rbt.root)
	return node
}

func (rbt *RBTree) findSubtreeMin(node *Node) (*Node, *Node) {
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

func (rbt *RBTree) FindMax() *Node {
	node, _ := rbt.findSubtreeMax(rbt.root)
	return node
}

func (rbt *RBTree) findSubtreeMax(node *Node) (*Node, *Node) {
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

func (rbt *RBTree) findNodeAndParent(key int) (*Node, *Node) {
	node := rbt.root
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

func (rbt *RBTree) Insert(key int) {
	node := rbt.root
	if node == nil {
		rbt.root = &Node{
			key:     key,
			isBlack: true,
		}
		return
	}
	newNode := rbt.insertNewNode(key)
	rbt.rebalance(newNode)

	root := newNode
	for root.parent != nil {
		root = root.parent
	}
	rbt.root = root
}

func (rbt *RBTree) insertNewNode(key int) *Node {
	var nodePtr **Node
	var parent *Node
	node := rbt.root

	for node != nil {
		if node.key == key {
			node.duplicate++
			return nil
		} else if node.key > key {
			nodePtr = &node.left
		} else {
			nodePtr = &node.right
		}
		parent = node
		node = *nodePtr
	}
	newNode := &Node{
		key:     key,
		parent:  parent,
		isBlack: false,
	}
	*nodePtr = newNode
	return newNode
}

func (rbt *RBTree) rotateLeft(x *Node) {
	p := x.parent
	g := p.parent

	p.right = x.left
	x.left = p
	p.parent = x

	if p.right != nil {
		p.right.parent = p
	}
	if g != nil {
		ptr := rbt.findChildPointerOfParent(p, g)
		*ptr = x
	}
	x.parent = g
}

func (rbt *RBTree) rotateRight(x *Node) {
	p := x.parent
	g := p.parent

	p.left = x.right
	x.right = p
	p.parent = x

	if p.left != nil {
		p.left.parent = p
	}
	if g != nil {
		ptr := rbt.findChildPointerOfParent(p, g)
		*ptr = x
	}
	x.parent = g
}

func (rbt *RBTree) rebalance(x *Node) {
	p := x.parent

	if p == nil {
		x.isBlack = true
		return
	}
	if p.isBlack {
		return
	}

	g := p.parent
	u := rbt.findUncle(p, g)

	if u != nil && !u.isBlack {
		rbt.rebalanceIfRedParentRedUncle(x)
	} else if g != nil && (u == nil || u.isBlack) {
		rbt.rebalanceIfRedParentBlackUncle(x)
	}
}

func (rbt *RBTree) rebalanceIfRedParentRedUncle(x *Node) {
	p := x.parent
	g := p.parent
	u := rbt.findUncle(p, g)

	p.isBlack = true
	u.isBlack = true
	g.isBlack = false
	rbt.rebalance(g)
}

func (rbt *RBTree) rebalanceIfRedParentBlackUncle(x *Node) {
	p := x.parent
	g := p.parent

	if x == p.right && p == g.left {
		rbt.rotateLeft(x)
		x = x.left
	} else if x == p.left && p == g.right {
		rbt.rotateRight(x)
		x = x.right
	}

	p = x.parent
	g = p.parent

	if x == p.left {
		rbt.rotateRight(p)
	} else if x == p.right {
		rbt.rotateLeft(p)
	}

	p.isBlack = true
	g.isBlack = false
}

func (rbt *RBTree) Remove(key int) *Node {
	return nil
}

func (rbt *RBTree) findUncle(parent *Node, grandParent *Node) *Node {
	if grandParent == nil {
		return nil
	}
	if grandParent.left == parent {
		return grandParent.right
	}
	return grandParent.left
}

func (rbt *RBTree) findChildPointerOfParent(node *Node, parent *Node) **Node {
	if parent.left == node {
		return &parent.left
	}
	return &parent.right
}

func (rbt *RBTree) findOnlyChild(parent *Node) (child *Node) {
	if parent.left != nil {
		return parent.left
	}
	return parent.right
}
