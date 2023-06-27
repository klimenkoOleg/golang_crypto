package main

import "fmt"

type Node struct {
	Val    int
	Next   *Node
	Random *Node
}

func copyRandomList(head *Node) *Node {
	if head == nil {
		return nil
	}
	addrMap := make(map[*Node]*Node)

	curr := head
	var prev *Node
	var newHead *Node

	for ; curr != nil; curr = curr.Next {
		copyNode := &Node{Val: curr.Val, Next: nil, Random: curr.Random}
		if prev != nil {
			prev.Next = copyNode
		} else {
			newHead = copyNode
		}
		prev = copyNode
		addrMap[curr] = copyNode
	}

	for curr = newHead; curr != nil; curr = curr.Next {
		newAddr := addrMap[curr.Random]
		curr.Random = newAddr
		fmt.Println(curr.Random)
		fmt.Println(newAddr)
	}
	fmt.Println(addrMap)

	return newHead

}
