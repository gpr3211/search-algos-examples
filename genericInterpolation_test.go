// Generic interpolation search returning value and bool

package main

import (
	"math/rand"
	"testing"
)

func BenchmarkGenericInterpolationSearch(b *testing.B) {
	var (
		items = make([]int, dataSize)
	)
	for i := range items {
		items[i] = i
	}
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		target := rand.Intn(dataSize)
		_, _ = NumericInterpolationSearch(items, target)
	}
}

func BenchmarkGenericInterpolationSearchRandom(b *testing.B) {
	items := randomSortedData(10000000)
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		target := rand.Intn(dataSize)
		_, _ = NumericInterpolationSearch(items, target)
	}
}

func BenchmarkGenericInterpolationSearchUniformJitter(b *testing.B) {
	items := uniformWithJitter(10000000)
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		target := rand.Intn(dataSize)
		_, _ = NumericInterpolationSearch(items, target)
	}
}

func BenchmarkGenericInterpolationSearchSkewed(b *testing.B) {
	items := skewedData(10000000)
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		target := rand.Intn(dataSize)
		_, _ = NumericInterpolationSearch(items, target)
	}
}
