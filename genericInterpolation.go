package main

import (
	"golang.org/x/exp/constraints"
)

type Numeric interface {
	constraints.Signed | constraints.Unsigned | ~float32 | ~float64
}

// Generic interpolation search returning value and bool
func GenericInterpolationSearch[T Numeric](arr []T, x T) (T, bool) {
	var zero T
	if len(arr) == 0 {
		return zero, false
	}

	low := 0
	high := len(arr) - 1

	for low <= high && x >= arr[low] && x <= arr[high] {
		if low == high {
			if arr[low] == x {
				return arr[low], true
			}
			return zero, false
		}

		pos := low + int(float64(x-arr[low])*float64(high-low)/float64(arr[high]-arr[low]))

		if arr[pos] == x {
			return arr[pos], true
		}

		if arr[pos] < x {
			low = pos + 1
		} else {
			high = pos - 1
		}
	}

	return zero, false
}
