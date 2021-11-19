package main

import "fmt"

func Test() {
	slice := []byte("abcdefgh")

	var arr [4]byte

	copy(arr[:], slice[:4])

	fmt.Println(arr)
}
