package main

import (
	"strconv"
)

func TimeConversion(s string) string {
	time := string(s[0]) + string(s[1])
	format := string(s[8]) + string(s[9])
	i, _ := strconv.Atoi(time)

	if i == 12 && format == "AM" {
		s = "00" + string(s[2]) + string(s[3]) + string(s[4]) + string(s[5]) + string(s[6]) + string(s[7])

	} else if i <= 11 && i >= 1 && format == "PM" {
		i += 12
		s = strconv.Itoa(i) + string(s[2]) + string(s[3]) + string(s[4]) + string(s[5]) + string(s[6]) + string(s[7])
	} else {
		s = time + string(s[2]) + string(s[3]) + string(s[4]) + string(s[5]) + string(s[6]) + string(s[7])
	}
	return s
}
