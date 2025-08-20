package main

func binarySearch(arr []int, x int) int {
	i := 0
	j := len(arr)
	for i != j {
		var m = (i + j) / 2
		if x == arr[m] {
			return m
		}
		if x < arr[m] {
			j = m
		} else {
			i = m + 1
		}
	}
	return -1
}
