package main

func interpolationSearch(arr []int, x int) int {
	low := 0
	high := len(arr) - 1
	for (arr[low] < x) && (x < arr[high]) {

		var mid = low + ((x-arr[low])*(high-low))/(arr[high]-arr[low])

		if arr[mid] < x {
			low = mid + 1
		} else if arr[mid] > x {
			high = mid - 1
		} else {
			return mid
		}
	}
	if arr[low] == x {
		return low
	}
	if arr[high] == x {
		return high
	}
	return -1
}
