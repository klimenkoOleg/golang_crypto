package main

import (
	"fmt"
	"sort"
)

func main() {
	logs := [][]int{{20190101, 0, 1}, {20190104, 3, 4}, {20190107, 2, 3}, {20190211, 1, 5}, {20190224, 2, 4}, {20190301, 0, 3}, {20190312, 1, 2}, {20190322, 4, 5}}
	n := 6

	fmt.Println(earliestAcq(logs, n))
}

func earliestAcq(logs [][]int, n int) int {
	count = n

	root = make([]int, n)
	rank = make([]int, n)
	for i := 0; i < n; i++ {
		root[i] = i
		rank[i] = 1
	}
	sort.Slice(logs, func(i, j int) bool {
		return logs[i][0] < logs[j][0]
	})

	for i := 0; i < len(logs); i++ {
		union(logs[i][1], logs[i][2])
		if count == 1 {
			return logs[i][0]
		}
	}
	return -1
}

var (
	root  []int
	rank  []int
	count int
)

func find(x int) int {
	if x == root[x] {
		return x
	}
	root[x] = find(root[x])
	return root[x]
}

func union(x int, y int) {
	rootX := find(x)
	rootY := find(y)
	if rootX != rootY {
		count--
		if rank[rootX] > rank[rootY] {
			root[rootY] = rootX
		} else if rank[rootX] < rank[rootY] {
			root[rootX] = rootY
		} else {
			root[rootY] = rootX
			rank[rootX] += 1
		}
	}
}
