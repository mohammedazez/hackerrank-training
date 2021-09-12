package main

import "fmt"

func simpleArraySum(ar []int32) int32 {
	// Write your code here
	var sum int32

	for i, s := range ar {
		fmt.Println(i, s)
		sum += s
	}
	return sum
}

func main() {
	primes := []int32{1, 2, 3, 4, 10, 11}

	fmt.Println(simpleArraySum(primes))
}
