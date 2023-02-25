package dynamicArray

type DynamicArray struct {
	A        []int
	currSize int
	capacity int
}

func (d DynamicArray) resizeArray() {
	d.capacity = d.capacity * 2

	temp := make([]int, d.capacity)
	copy(temp, d.A)

	d.A = temp
}

func (d DynamicArray) shrinkArray() {
	d.capacity = d.capacity / 2

	temp := make([]int, d.capacity)
	copy(temp, d.A)

	d.A = temp
}

func (d DynamicArray) insertAtEnd(x int) {
	if d.currSize == d.capacity {
		d.resizeArray()
	}

	d.A[d.currSize] = x
	d.currSize = d.currSize + 1
}

func (d DynamicArray) insertAtMiddle(index int, x int) {
	if d.currSize == d.capacity {
		d.resizeArray()
	}

	for i := d.currSize - 1; i > index; i-- {
		d.A[i] = d.A[i-1]
	}

	d.A[index] = x
}

func (d DynamicArray) deleteAtEnd() {
	if d.currSize == 0 {
		return
	}

	d.A[d.currSize-1] = 0
	d.currSize = d.currSize - 1

	if d.currSize == d.capacity/4 {
		d.shrinkArray()
	}
}

func (d DynamicArray) deleteAtMiddle(index int) {
	if d.currSize == 0 {
		return
	}

	for i := 0; i < d.currSize; i = i + 1 {
		d.A[i] = d.A[i+i]
	}

	d.A[d.currSize-1] = 0
	d.currSize = d.currSize - 1

	if d.currSize == d.capacity/4 {
		d.shrinkArray()
	}
}
