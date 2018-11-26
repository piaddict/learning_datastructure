package queue

import "fmt"

type (
	// ArrayQueue 구조체는 내부 자료구조로 배열을 사용한다
	// 구현체는 원형 큐이다
	ArrayQueue struct {
		queue []interface{}
		size  int
		front int
		rear  int
	}

	// Queue 인터페이스
	Queue interface {
		IsFull() bool
		IsEmpty() bool
		Add(item interface{}) Queue
		Remove() interface{}
		Peek() interface{}
	}
)

func New(size int) Queue {
	queue := ArrayQueue{
		queue: make([]interface{}, size),
		size:  size,
		front: 0,
		rear:  0,
	}
	return &queue
}

func (this *ArrayQueue) IsFull() bool {
	return this.front == (this.rear+1)%this.size
}

func (this *ArrayQueue) IsEmpty() bool {
	return this.front == this.rear
}

func (this *ArrayQueue) Add(item interface{}) Queue {
	if this.IsFull() {
		fmt.Println("큐가 가득 찼습니다")
		return nil
	}
	this.rear = (this.rear + 1) % this.size
	this.queue[this.rear] = item
	return this
}

func (this *ArrayQueue) Remove() interface{} {
	if this.IsEmpty() {
		fmt.Println("큐가 비어있습니다")
		return nil
	}
	this.front = (this.front + 1) % this.size
	ret := this.queue[this.front]
	this.queue[this.front] = nil
	return ret
}

func (this *ArrayQueue) Peek() interface{} {
	if this.IsEmpty() {
		fmt.Println("큐가 비어있습니다")
		return nil
	}
	return this.queue[(this.front+1)%this.size]
}
