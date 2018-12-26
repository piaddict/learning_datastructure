package heap

type PriorityQueue struct {
	heap *Heap
}

func NewPriorityQueue() *PriorityQueue {
	return &PriorityQueue{
		heap: New(),
	}
}

func (this *PriorityQueue) Enqueue(element int) {
	this.heap.Insert(element)
}

func (this *PriorityQueue) Dequeue() int {
	return this.heap.Remove()
}

func (this *PriorityQueue) Peek() int {
	return this.heap.Peek()
}
