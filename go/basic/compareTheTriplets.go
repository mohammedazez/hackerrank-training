package main

func CompareTriplets(a []int32, b []int32) []int32 {

	var temp1 []int32
	var temp2 []int32
	var result []int32
	var one int
	var two int

	if len(a) == len(b) {
		for i := range a {
			if a[i] > b[i] {
				temp1 = append(temp1, a[i])
				one = len(temp1)
			}

			if b[i] > a[i] {
				temp2 = append(temp2, b[i])
				two = len(temp2)
			}

		}
	}

	result = append(result, int32(one), int32(two))

	return result
}
