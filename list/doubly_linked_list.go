package list

import "fmt"

type (
	DoubleLinkedList struct {
		head *head
	}

	head struct {
		count int
		bNode *node
		eNode *node
	}

	node struct {
		data  interface{}
		lNode *node
		rNode *node
	}

	LinkedList interface {
		AddFirst(item interface{})
		Add(index int, item interface{})
		AddLast(item interface{})
		Get(index int) interface{}
		Size() int
		Remove(index int)
	}
)

func New() LinkedList {
	return &DoubleLinkedList{
		head: &head{
			count: 0,
		},
	}
}

func (this *DoubleLinkedList) AddFirst(item interface{}) {
	newNode := &node{data: item}

	begin := this.head.bNode
	if this.Size() == 0 {
		this.head.eNode = newNode
	} else {
		newNode.rNode = begin
		begin.lNode = newNode
	}
	this.head.bNode = newNode
	this.head.count++
}

func (this *DoubleLinkedList) Add(index int, item interface{}) {
	if this.checkIndexOutOfBounds(index) {
		return
	}

	switch this.head.count {
	case 0:
		this.AddFirst(item)
	case index + 1:
		this.AddLast(item)
	default:
		anteNode := this.getNode(index)
		postNode := this.getNode(index + 1)
		newNode := &node{
			data:  item,
			lNode: anteNode,
			rNode: postNode,
		}
		anteNode.rNode = newNode
		postNode.lNode = newNode
		this.head.count++
	}
}

func (this *DoubleLinkedList) AddLast(item interface{}) {
	newNode := &node{data: item}

	end := this.head.eNode
	if this.Size() == 0 {
		this.head.bNode = newNode
	} else {
		newNode.lNode = end
		end.rNode = newNode
	}
	this.head.eNode = newNode
	this.head.count++
}

func (this *DoubleLinkedList) Get(index int) interface{} {
	var node *node
	if node = this.getNode(index); node == nil {
		return nil
	}
	return node.data
}

func (this *DoubleLinkedList) getNode(index int) *node {
	if this.checkIndexOutOfBounds(index) {
		return nil
	}

	direction := index <= this.head.count/2

	var node *node
	if direction {
		node = this.head.bNode
		for i := 0; i < index; i++ {
			node = node.rNode
		}
	} else {
		node = this.head.eNode
		for i := 0; i < this.head.count-index-1; i++ {
			node = node.lNode
		}
	}

	return node
}

func (this *DoubleLinkedList) Size() int {
	return this.head.count
}

func (this *DoubleLinkedList) Remove(index int) {
	if this.checkIndexOutOfBounds(index) {
		return
	}

	node := this.getNode(index)

	if node.lNode != nil {
		node.lNode.rNode = node.rNode
	} else {
		this.head.bNode = node.rNode
	}

	if node.rNode != nil {
		node.rNode.lNode = node.lNode
	} else {
		this.head.eNode = node.lNode
	}

	node = nil
	this.head.count--
}

func (this *DoubleLinkedList) checkIndexOutOfBounds(index int) bool {
	if index < 0 || index >= this.head.count {
		fmt.Println("error: out of bounds")
		return true
	}
	return false
}
