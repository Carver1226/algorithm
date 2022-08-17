package main

type MyCircularDeque struct {
	q []int
	front int
	rear int
}

func Constructor(k int) MyCircularDeque {
	return MyCircularDeque{make([]int, k + 1), 0, 0}
}

//获取队列在物理上的下标
func (this *MyCircularDeque) get(index int) int {
	return (index + len(this.q)) % len(this.q)
}


func (this *MyCircularDeque) InsertFront(value int) bool {
	if this.IsFull() {
		return false
	}
	this.front = this.get(this.front - 1)
	this.q[this.front] = value
	return true
}


func (this *MyCircularDeque) InsertLast(value int) bool {
	if this.IsFull() {
		return false
	}
	this.q[this.rear] = value
	this.rear = this.get(this.rear + 1)
	return true
}


func (this *MyCircularDeque) DeleteFront() bool {
	if this.IsEmpty() {
		return false
	}
	this.front =  this.get(this.front + 1)
	return true
}


func (this *MyCircularDeque) DeleteLast() bool {
	if this.IsEmpty() {
		return false
	}
	this.rear = this.get(this.rear - 1)
	return true
}


func (this *MyCircularDeque) GetFront() int {
	if this.IsEmpty() {
		return -1
	}
	return this.q[this.front]
}


func (this *MyCircularDeque) GetRear() int {
	if this.IsEmpty() {
		return -1
	}
	return this.q[this.get(this.rear - 1)]
}


func (this *MyCircularDeque) IsEmpty() bool {
	return this.front == this.rear
}


func (this *MyCircularDeque) IsFull() bool {
	return this.get(this.rear + 1) == this.front
}