package basic

import "fmt"

func DiagonalDifference(arr [][]int32) int32 {

	var diag1Sum int32 //TL > BR
	var diag2Sum int32 //TR > BL
	var col int

	for row := 1; row < len(arr); row++ {
		diag1Sum += arr[row][col]
		fmt.Println("diag1Sum", diag1Sum)

		diag2Sum += arr[row][len(arr[row])-1-col]
		fmt.Println("diag2Sum", diag2Sum)
		col++
	}

	fmt.Println(diag1Sum - diag2Sum)
	return diag1Sum - diag2Sum
}
