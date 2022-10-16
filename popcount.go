package main

import "fmt"

// pc[i] is the population count of i.
var pc [256]byte

func init() {
	for i := range pc {
		pc[i] = pc[i/2] + byte(i&1)
		fmt.Printf("pc[%d]=%d\n", i, pc[i])
	}
}

// PopCount returns the population count (number of set bits) of x.
func PopCount(x uint64) int {
	sum := 0
	for i := 0; i < 64; i++ {
		sum += int(0b1 & x)
		x >>= 1
	}

	for x != 0 {
		x = x & (x - 1)
		sum++
	}
	return sum
	/*sum := 0
	for i := 0; i < 8; i++ {
		sum += int(pc[byte(x>>(i*8))])
	}
	return sum*/
	//return int(pc[byte(x>>(0*8))] +
	//	pc[byte(x>>(1*8))] +
	//	pc[byte(x>>(2*8))] +
	//	pc[byte(x>>(3*8))] +
	//	pc[byte(x>>(4*8))] +
	//	pc[byte(x>>(5*8))] +
	//	pc[byte(x>>(6*8))] +
	//	pc[byte(x>>(7*8))])
}

func main() {
	//for i := range [1..10] {
	//fmt.Println(PopCount(i))
	//}

	for i := 0; i < 10; i++ {
		fmt.Println(PopCount(uint64(0b10000000011)))
	}

	/*
		0
		1
		1
		2
		1
		2
		2
		3
		1
		2
	*/

}
