package basic

import (
	"strconv"
)

func TimeConversion(s string) string {
	// Write your code here
	a, _ := strconv.Atoi(string(s[0]) + string(s[1]))
	if a == 12 && string(s[8])+string(s[9]) == "AM" {
		s = "00" + string(s[2]) + string(s[3]) + string(s[4]) + string(s[5]) + string(s[6]) + string(s[7])
	} else if a >= 1 && a <= 12 && string(s[8])+string(s[9]) == "AM" {
		s = string(s[0]) + string(s[1]) + string(s[2]) + string(s[3]) + string(s[4]) + string(s[5]) + string(s[6]) + string(s[7])
	} else if a >= 1 && a <= 11 && string(s[8])+string(s[9]) == "PM" {
		sum := a + 12
		str := strconv.Itoa(sum)
		s = str + string(s[2]) + string(s[3]) + string(s[4]) + string(s[5]) + string(s[6]) + string(s[7])
	} else if a == 12 && string(s[8])+string(s[9]) == "PM" {
		s = string(s[0]) + string(s[1]) + string(s[2]) + string(s[3]) + string(s[4]) + string(s[5]) + string(s[6]) + string(s[7])
	}
	return s
}
