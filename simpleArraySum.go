package main

import "fmt"

func SimpleArraySum(ar []int32) int32 {
	// Write your code here
	var sum int32

	for i, s := range ar {
		fmt.Println(i, s)
		sum += s
	}
	return sum
}
