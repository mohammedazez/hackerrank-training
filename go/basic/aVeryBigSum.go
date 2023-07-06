package main

import "fmt"

func AveryBigSum(ar []int64) int64 {
	var sum int64

	// fmt.Println(ar)
	for i, s := range ar {
		fmt.Println(i, s)
		sum += s
	}

	fmt.Println("sum", sum)
	return sum
}
