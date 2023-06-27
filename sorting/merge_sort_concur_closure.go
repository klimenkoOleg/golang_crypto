package main

import (
	"fmt"
	"sync"
)

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

func MergeSortConcurClosure(data []int) []int {
	if len(data) <= 1 {
		return data
	}
	mid := len(data) / 2

	//leftCh := make(chan []int)
	//rightCh := make(chan []int)

	var left []int
	var right []int

	ch := make(chan bool)

	go func() {
		left = MergeSortConcurClosure(data[:mid])
		ch <- true
	}()

	go func() {
		right = MergeSortConcurClosure(data[mid:])
		ch <- true
	}()

	<-ch
	<-ch

	//go MergeSortConcurClosure(data[:mid], leftCh)
	//go MergeSortConcurClosure(data[mid:], rightCh)

	//left := <-leftCh
	//right := <-rightCh
	//left := MergeSort(data[:mid])
	//right := MergeSort(data[mid:])
	return Merge(left, right)
}

func Merge(left, right []int) []int {
	merged := make([]int, 0, len(left)+len(right))
	for len(left) > 0 || len(right) > 0 {
		if len(left) == 0 {
			return append(merged, right...)
		} else if len(right) == 0 {
			return append(merged, left...)
		} else if left[0] < right[0] {
			merged = append(merged, left[0])
			left = left[1:]
		} else {
			merged = append(merged, right[0])
			right = right[1:]
		}
	}
	return merged
}

func main() {
	data := []int{9, 4, 3, 6, 1, 2, 10, 5, 7, 8}
	//result := make(chan []int)
	result := MergeSortConcurClosure(data)
	//dataSorted := <-result
	fmt.Printf("%v\n%v\n", data, result)

	var wg sync.WaitGroup

	wg.Add(2) //called before running the goroutines

	go func() {
		// Do work.
		wg.Done() //goroutine reached its completion; calling Done() to signal
	}()
	go func() {
		// Do work.
		wg.Done() //goroutine reached its completion; calling Done() to signal

	}()

	wg.Wait() //blocking the code until all the .Done() statements are executed
}
