package main

import (
	"golang.org/x/exp/constraints"
)

type Numeric interface {
	constraints.Signed | constraints.Unsigned | ~float32 | ~float64
}

// GenericInterpolationSearch takes a sorted array satisfying the Numeric interface
// and returns the index of element and true if found.
func NumericInterpolationSearch[T Numeric](arr []T, x T) (int, bool) {
	if len(arr) == 0 {
		return -1, false
	}

	low := 0
	high := len(arr) - 1

	for low <= high && x >= arr[low] && x <= arr[high] {
		if low == high {
			if arr[low] == x {
				return low, true
			}
			return -1, false
		}

		// interpolated position.
		pos := low + int(float64(x-arr[low])*float64(high-low)/float64(arr[high]-arr[low]))

		if pos < low {
			pos = low
		} else if pos > high {
			pos = high
		}

		if arr[pos] == x {
			return pos, true
		}

		if arr[pos] < x {
			low = pos + 1
		} else {
			high = pos - 1
		}
	}

	return -1, false
}
