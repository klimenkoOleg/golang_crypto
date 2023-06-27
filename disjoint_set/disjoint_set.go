package main

import "fmt"

func main() {
	//[[1,0,0],[0,1,0],[0,0,1]]
	//isConnected := [][]int{{1, 0, 0}, {0, 1, 0}, {0, 0, 1}}
	//isConnected := [][]int{{1, 1, 0}, {1, 1, 0}, {0, 0, 1}}
	//isConnected := [][]int{{1, 0, 0, 1}, {0, 1, 1, 0}, {0, 1, 1, 1}, {1, 0, 1, 1}}
	//isConnected := [][]int{{1, 0, 0, 0, 0, 0, 0, 0, 0, 1, 0, 0, 0, 0, 0}, {0, 1, 0, 1, 0, 0, 0, 0, 0, 0, 0, 0, 0, 1, 0}, {0, 0, 1, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0}, {0, 1, 0, 1, 0, 0, 0, 1, 0, 0, 0, 1, 0, 0, 0}, {0, 0, 0, 0, 1, 0, 0, 0, 0, 0, 0, 0, 1, 0, 0}, {0, 0, 0, 0, 0, 1, 0, 0, 0, 0, 0, 0, 0, 0, 0}, {0, 0, 0, 0, 0, 0, 1, 0, 0, 0, 0, 0, 0, 0, 0}, {0, 0, 0, 1, 0, 0, 0, 1, 1, 0, 0, 0, 0, 0, 0}, {0, 0, 0, 0, 0, 0, 0, 1, 1, 0, 0, 0, 0, 0, 0}, {1, 0, 0, 0, 0, 0, 0, 0, 0, 1, 0, 0, 0, 0, 0}, {0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 1, 0, 0, 0, 0}, {0, 0, 0, 1, 0, 0, 0, 0, 0, 0, 0, 1, 0, 0, 0}, {0, 0, 0, 0, 1, 0, 0, 0, 0, 0, 0, 0, 1, 0, 0}, {0, 1, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 1, 0}, {0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 1}}
	//fmt.Println(findCircleNum(isConnected))
	//edges := [][]int{{0, 1}, {0, 2}, {0, 3}, {1, 4}}
	edges := [][]int{{0, 1}, {1, 2}, {3, 4}}
	//n := 5
	n := 5

	//fmt.Println(validTree(n, edges))
	fmt.Println(countComponents(n, edges))

}

/*
func validTree(n int, edges [][]int) bool {
	if n == 1 {
		return true
	}
	root = make([]int, n)
	for i := 0; i < n; i++ {
		root[i] = i
	}
	for _, edge := range edges {
		var x, y int
		if edge[0] < edge[1] {
			x = edge[0]
			y = edge[1]
		} else {
			x = edge[1]
			y = edge[0]
		}
		if !union(x, y) {
			return false
		}
	}

	singleRoot := find(0)
	for i := 1; i < n; i++ {
		if find(i) != singleRoot {
			return false
		}
	}
	return true
}

var root []int

func find(x int) int {
	for x != root[x] {
		x = root[x]
	}
	return x
}

func union(x int, y int) bool {

	if root[y] != x && root[y] != y {
		return false
	}

	rootX := find(x)
	rootY := find(y)
	if rootX != rootY {
		root[rootY] = rootX
	}
	return true
}

func connected(x int, y int) bool {
	return find(x) == find(y)
}*/

/*
// The find function – optimized with path compression:
func find(x int) int {
	if x == root[x] {
		return x
	}
	root[x] = find(root[x])
	return root[x]
}

// The union function – Optimized by union by rank:
func union(x int, y int) {

	rootX := find(x)
	rootY := find(y)
	if rootX != rootY {
		if rank[rootX] > rank[rootY] {
			root[rootY] = rootX
		} else if rank[rootX] < rank[rootY] {
			root[rootX] = rootY
		} else {
			//if x < y { // this sorts roots while union
			root[rootY] = rootX
			rank[rootX] += 1
			//} else {
			//	root[rootX] = rootY
			//	rank[rootY] += 1
			//}
		}
	}
}*/

var (
	root []int
)

func find(x int) int {
	return root[x]
}

func union(x int, y int) {
	rootX := find(x)
	rootY := find(y)
	if rootX != rootY {
		for i := 0; i < len(root); i++ {
			if root[i] == rootY {
				root[i] = rootX
			}
		}
	}
}

func countComponents(n int, edges [][]int) int {
	if n == 0 || edges == nil {
		return 0
	}
	root = make([]int, n)
	for i := 0; i < n; i++ {
		root[i] = i
	}
	for i := 0; i < len(edges); i++ {
		union(edges[i][0], edges[i][1])
	}
	components := make(map[int]bool)
	for i := 0; i < n; i++ {
		root := find(i)
		components[root] = true
	}
	return len(components)
}

/*func findCircleNum(isConnected [][]int) int {
	if isConnected == nil || len(isConnected) == 0 {
		return 0
	}
	root = make([]int, len(isConnected))
	for i := 0; i < len(isConnected); i++ {
		root[i] = i
	}
	for i := 0; i < len(isConnected); i++ {
		for j := i + 1; j < len(isConnected[i]); j++ {
			if isConnected[i][j] == 1 {
				union(i, j)
			}
		}
	}

	cnt := make(map[int]bool)
	count := 0
	for i := 0; i < len(root); i++ {
		_, ok := cnt[root[i]]
		if !ok {
			cnt[root[i]] = true
			count++
		}

	}
	return count
}*/
