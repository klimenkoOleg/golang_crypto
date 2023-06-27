package main

import (
	"fmt"
)

type Node struct {
	Val  int
	Next *Node
}

func InsertTail(head *Node, val int) *Node {
	//head := &headRef
	newNode := &Node{val, nil}
	//if head == nil {
	//	return newNode
	//}
	cur := head
	for ; cur.Next != nil; cur = cur.Next {
	}
	cur.Next = newNode
	return newNode
}

func PrintList(head *Node) {
	cur := head
	fmt.Println(cur.Val)
	cur = cur.Next
	for ; cur != nil && cur != head; cur = cur.Next {
		fmt.Println(cur.Val)
	}
}

func insert(aNode *Node, x int) *Node {
	if aNode == nil {
		head := &Node{x, nil}
		head.Next = head
		return head
	}
	if aNode == aNode.Next {
		aNode.Next = &Node{x, aNode}
		return aNode
	}
	//fakeNode := &Node{aNode, 0}
	for cur := aNode; ; {
		b1 := x >= cur.Val && x <= cur.Next.Val
		b2 := x >= cur.Val && cur.Val > cur.Next.Val
		b3 := x <= cur.Val && x <= cur.Next.Val && cur.Val > cur.Next.Val
		if (b1) ||
			(b2) ||
			(b3) {
			nextItem := cur.Next
			cur.Next = &Node{x, nextItem}
			return aNode
		}
		cur = cur.Next
		if cur.Next == aNode {
			nextItem := cur.Next
			cur.Next = &Node{x, nextItem}
			return aNode
		}
	}
	//return aNode
}

func main() {
	//arr := []int{3, 4, 1}
	//arr := []int{3, 3, 3}
	//arr := []int{1, 3, 5}
	//arr := []int{3, 5, 1}
	arr := []int{3, 3, 5}
	var head, last *Node
	head = &Node{} // this is fake head
	for _, val := range arr {
		last = InsertTail(head, val)
	}
	last.Next = head.Next
	PrintList(head.Next)
	fmt.Println()
	//insert(head.Next, 2)
	//insert(head.Next, 0)
	//insert(head.Next, 0)
	insert(head.Next, 0)
	PrintList(head.Next)
	//fmt.Println(last)
}
