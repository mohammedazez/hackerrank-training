package main

import (
	"fmt"
)

func MiniMaxSum(arr []int32) {

	for i := 0; i < len(arr); i++ {
		// fmt.Println(arr[i])

		for j := 0; j < len(arr); j++ {
			if arr[i] != arr[j] {
				fmt.Println(arr[i])
			}
		}

	}
}
