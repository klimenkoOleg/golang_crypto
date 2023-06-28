package main

import "fmt"

type MyCircularQueue struct {
	queue []int
	head  int
	tail  int
	len   int
}

func Constructor(k int) MyCircularQueue {
	return MyCircularQueue{make([]int, k), -1, -1, 0}
}

func (this *MyCircularQueue) EnQueue(value int) bool {
	if this.len == len(this.queue) {
		return false
	}
	if this.len == 0 {
		this.tail = 0
		this.head = 0
	} else {
		this.tail++
	}
	if this.tail == len(this.queue) {
		this.tail = 0
	}
	this.queue[this.tail] = value
	this.len++
	return true
}

func (this *MyCircularQueue) DeQueue() bool {
	if this.len == 0 {
		return false
	}
	this.head++
	if this.head == len(this.queue) {
		this.head = 0
	}
	this.len--
	if this.len == 0 {
		this.head = -1
		this.tail = -1
	}
	return true
}

func (this *MyCircularQueue) Front() int {
	if this.len == 0 {
		return -1
	}
	return this.queue[this.head]
}

func (this *MyCircularQueue) Rear() int {
	if this.len == 0 {
		return -1
	}
	return this.queue[this.tail]
}

func (this *MyCircularQueue) IsEmpty() bool {
	return this.len == 0
}

func (this *MyCircularQueue) IsFull() bool {
	return this.len == len(this.queue)
}

func main() {
	myCircularQueue := Constructor(3)
	fmt.Println(myCircularQueue.EnQueue(1)) // return True
	fmt.Println(myCircularQueue.EnQueue(2)) // return True
	fmt.Println(myCircularQueue.EnQueue(3)) // return True
	fmt.Println(myCircularQueue.EnQueue(4)) // return False
	fmt.Println(myCircularQueue.Rear())     // return 3
	fmt.Println(myCircularQueue.IsFull())   // return True
	fmt.Println(myCircularQueue.DeQueue())  // return True
	fmt.Println(myCircularQueue.EnQueue(4)) // return True
	fmt.Println(myCircularQueue.Rear())     // return 4

	/*myCircularQueue := Constructor(3)
	fmt.Println(myCircularQueue.EnQueue(2))
	fmt.Println(myCircularQueue.Rear())
	fmt.Println(myCircularQueue.Front())
	fmt.Println(myCircularQueue.DeQueue())
	fmt.Println(myCircularQueue.Front())
	fmt.Println(myCircularQueue.DeQueue())
	fmt.Println(myCircularQueue.Front())
	fmt.Println(myCircularQueue.EnQueue(4))
	fmt.Println(myCircularQueue.EnQueue(2))
	fmt.Println(myCircularQueue.EnQueue(2))
	fmt.Println(myCircularQueue.EnQueue(3)) */

	//["MyCircularQueue","enQueue","Rear","Front","deQueue","Front","deQueue","Front","enQueue","enQueue","enQueue","enQueue"]
	//	[[3],[2],[],[],[],[],[],[],[4],[2],[2],[3]]

	/*myCircularQueue := Constructor(3)
	fmt.Println(myCircularQueue.EnQueue(1))
	fmt.Println(myCircularQueue.EnQueue(2))
	fmt.Println(myCircularQueue.EnQueue(3))
	fmt.Println(myCircularQueue.EnQueue(4))
	fmt.Println(myCircularQueue.Rear())
	fmt.Println(myCircularQueue.IsFull())
	fmt.Println(myCircularQueue.DeQueue())
	fmt.Println(myCircularQueue.EnQueue(4))
	fmt.Println(myCircularQueue.Rear())
	*/

	//myCircularQueue := Constructor(3)

	//["MyCircularQueue","enQueue","enQueue","Front","enQueue","deQueue","enQueue","enQueue","Rear","isEmpty","Front","deQueue"]
	//	[[2],[8],[8],[],[4],[],[1],[1],[],[],[],[]]

	//myCircularQueue := Constructor(3)
	//
	//
	//["MyCircularQueue","enQueue","Rear","enQueue","deQueue","Front","deQueue","deQueue","isEmpty","deQueue","enQueue","enQueue"]
	//[[2],[4],[],[9],[],[],[],[],[],[],[6],[4]]

	//["MyCircularQueue","enQueue","enQueue","enQueue","enQueue","Rear","isFull","deQueue","enQueue","Rear"]
	//	[[3],[1],[2],[3],[4],[],[],[],[4],[]]
}

/**
 * Your MyCircularQueue object will be instantiated and called as such:
 * obj := Constructor(k);
 * param_1 := obj.EnQueue(value);
 * param_2 := obj.DeQueue();
 * param_3 := obj.Front();
 * param_4 := obj.Rear();
 * param_5 := obj.IsEmpty();
 * param_6 := obj.IsFull();
 */
