package main

import (
	"fmt"
	"hackerrank-training/basic"
)

func main() {

	// test
	// Test()

	// solve me first
	// var a, b, res uint32
	// fmt.Scanf("%v\n%v", &a, &b)
	// res = solveMeFirst(a, b)
	// fmt.Println(res)

	// simpleArraySum
	// primes := []int32{1, 2, 3, 4, 10, 11}
	// fmt.Println(simpleArraySum(primes))

	// A Very Big Sum
	// number := []int64{1000000001, 1000000002, 1000000003, 1000000004, 1000000005}
	// aVeryBigSum(number)

	// Compare The Triplets
	// alice := []int32{5, 6, 7}
	// boby := []int32{3, 6, 10}
	// alice1 := []int32{17, 28, 30}
	// boby1 := []int32{99, 16, 8}
	// fmt.Println(alice)
	// fmt.Println("var1 = ", reflect.ValueOf(alice).Kind())
	// fmt.Println(compareTriplets(alice1, boby1))

	// Diagonal Difference
	// a := [][]int32{{11, 2, 4}, {4, 5, 6}, {10, 8, -12}}
	// // fmt.Println("var1 = ", reflect.ValueOf(a).Kind())
	// // b := [][]int32{{-1, 1, -7, -8}, {-10, -8, -5, -2}, {0, 9, 7, -1}, {4, 4, -2, 1}}
	// DiagonalDifference(a)

	// Plus Minus
	// arr := []int32{1, 1, 0, -1, -1}
	// basic.PlusMinus(arr)

	// Mini max sum
	// arr := []int32{1, 2, 3, 4, 5}
	// basic.MiniMaxSum(arr)

	// warmup
	// staircase(6)

	// timeConversion

	// 12:00 am -> 00:00
	// 01:00 - 12:00 am -> keep just remove am
	// 01:00 - 11:00 pm -> sum to 12
	// a := "12:00:00AM"
	// b := "01:00:00AM"
	c := "12:45:54PM"
	// fmt.Println((basic.TimeConversion(a)))
	// basic.TimeConversion(b)
	fmt.Println((basic.TimeConversion(c)))

}
