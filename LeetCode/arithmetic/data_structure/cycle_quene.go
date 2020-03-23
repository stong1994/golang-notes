package data_structure

/*
MyCircularQueue(k): 构造器，设置队列长度为 k 。
Front: 从队首获取元素。如果队列为空，返回 -1 。
Rear: 获取队尾元素。如果队列为空，返回 -1 。
enQueue(value): 向循环队列插入一个元素。如果成功插入则返回真。
deQueue(): 从循环队列中删除一个元素。如果成功删除则返回真。
isEmpty(): 检查循环队列是否为空。
isFull(): 检查循环队列是否已满。
*/
type MyCircularQueue struct {
	queue []int
	front int
	end   int
	lens  int
}

/** Initialize your data structure here. Set the size of the queue to be k. */
func Constructor(k int) MyCircularQueue {
	return MyCircularQueue{
		queue: make([]int, k),
		front: -1,
		end:   -1,
		lens:  k,
	}
}

/** Insert an element into the circular queue. Return true if the operation is successful. */
func (this *MyCircularQueue) EnQueue(value int) bool {
	if this.queue == nil || this.IsFull() {
		return false
	}
	if this.end < this.lens-1 {
		this.end++
	} else {
		this.end = 0
	}

	if this.front == -1 {
		this.front = 0
	}

	this.queue[this.end] = value
	return true
}

/** Delete an element from the circular queue. Return true if the operation is successful. */
func (this *MyCircularQueue) DeQueue() bool {
	if this.queue == nil || this.IsEmpty() {
		return false
	}
	if this.front <= this.lens-1 {
		this.front++
	} else {
		this.front = 0
	}
	return true
}

/** Get the front item from the queue. */
func (this *MyCircularQueue) Front() int {
	if this.queue == nil || this.IsEmpty() {
		return -1
	}
	if !this.checkIndex(this.front) {
		panic("index out of range")
	}
	return this.queue[this.front]
}

/** Get the last item from the queue. */
func (this *MyCircularQueue) Rear() int {
	if this.queue == nil || this.IsEmpty() {
		return -1
	}
	if !this.checkIndex(this.end) {
		panic("index out of range")
	}
	return this.queue[this.end]
}

/** Checks whether the circular queue is empty or not. */
func (this *MyCircularQueue) IsEmpty() bool {
	if len(this.queue) <= 0 {
		return true
	}
	return false
}

/** Checks whether the circular queue is full or not. */
func (this *MyCircularQueue) IsFull() bool {
	if (this.front == 0 && this.end == this.lens-1) || (this.front != -1 && this.end+1 == this.front) {
		return true
	}
	return false
}

func (this *MyCircularQueue) checkIndex(index int) bool {
	if len(this.queue) <= index {
		return false
	}
	return true
}
