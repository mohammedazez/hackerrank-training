package main

import (
	"fmt"
)

func PlusMinus(arr []int32) {
	var positive []int32
	var negative []int32
	var zero []int32
	for i := 0; i < len(arr); i++ {
		if arr[i] < 0 {
			negative = append(negative, arr[i])
		}

		if arr[i] > 0 {
			positive = append(positive, arr[i])
		}

		if arr[i] == 0 {
			zero = append(zero, arr[i])
		}

	}

	fmt.Println(float64(len(positive)) / float64(len(arr)))
	fmt.Println(float64(len(negative)) / float64(len(arr)))
	fmt.Println(float64(len(zero)) / float64(len(arr)))
}
