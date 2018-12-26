package heap

type Heap struct {
	data []int
}

// New 함수 인자들은 입력시 정렬되지 않음
func New(args ...int) *Heap {
	return &Heap{
		data: args,
	}
}

func (this *Heap) Print() []int {
	return this.data
}

func (this *Heap) Insert(element int) {
	this.data = append(this.data, element)

	child := len(this.data) - 1
	parent := this.findParent(child)

	for child != 0 || this.data[child] < this.data[parent] {
		this.data[child], this.data[parent] = this.data[parent], this.data[child]
		child, parent = parent, this.findParent(parent)
	}
}

func (this *Heap) Remove() int {
	parent, child := 0, 1
	last := len(this.data) - 1

	ret := this.data[parent]
	this.data[parent] = this.data[last]
	this.data = this.data[:last]

	for child < last {
		if child+1 < last && this.data[child] > this.data[child+1] {
			child++
		}
		if this.data[parent] <= this.data[child] {
			break
		}
		this.data[parent], this.data[child] = this.data[child], this.data[parent]
		parent, child = child, this.findFirstChild(child)
	}

	return ret
}

func (this *Heap) findParent(index int) int {
	return (index - 1) / 2
}

func (this *Heap) findFirstChild(index int) int {
	return index*2 + 1
}

func (this *Heap) Peek() int {
	return this.data[0]
}

func (this *Heap) IsEmpty() bool {
	return len(this.data) == 0
}

func (this *Heap) Size() int {
	return len(this.data)
}
