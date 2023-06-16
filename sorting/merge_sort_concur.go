package main

import "fmt"

/*func Merge(left, right [] int) [] int{
	merged := make([] int, 0, len(left) + len(right))
	for len(left) > 0 || len(right) > 0{
		if len(left) == 0 {
			return append(merged,right...)
		}else if len(right) == 0 {
			return append(merged,left...)
		}else if left[0] < right[0] {
			merged = append(merged, left[0])
			left = left[1:]
		}else{
			merged = append(merged, right [0])
			right = right[1:]
		}
	}
	return merged
}*/

func MergeSortConcur(data []int, ch chan []int) {
	if len(data) <= 1 {
		ch <- data
	}
	mid := len(data) / 2

	leftCh := make(chan []int)
	rightCh := make(chan []int)
	go MergeSortConcur(data[:mid], leftCh)
	go MergeSortConcur(data[mid:], rightCh)

	left := <-leftCh
	right := <-rightCh
	//left := MergeSort(data[:mid])
	//right := MergeSort(data[mid:])
	ch <- Merge(left, right)
}

func main() {
	data := []int{9, 4, 3, 6, 1, 2, 10, 5, 7, 8}
	result := make(chan []int)
	go MergeSortConcur(data, result)
	dataSorted := <-result
	fmt.Printf("%v\n%v\n", dataSorted)

}
