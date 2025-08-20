package main

import (
	"math/rand"
	"testing"
)

const dataSize = 10_000_000

func BenchmarkInterpolationSearch(b *testing.B) {
	var (
		items = make([]int, dataSize)
	)
	for i := range items {
		items[i] = i
	}

	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		target := rand.Intn(dataSize)
		_ = interpolationSearch(items, target)
	}
}
func BenchmarkInterpolationSearchRandom(b *testing.B) {

	items := randomSortedData(10000000)

	b.ResetTimer()

	for i := 0; i < b.N; i++ {

		target := rand.Intn(dataSize)
		_ = interpolationSearch(items, target)
	}
}
func BenchmarkInterpolationSearchUniformJitter(b *testing.B) {

	items := uniformWithJitter(10000000)

	b.ResetTimer()

	for i := 0; i < b.N; i++ {

		target := rand.Intn(dataSize)
		_ = interpolationSearch(items, target)
	}
}
func BenchmarkInterpolationSearchSkewed(b *testing.B) {

	items := skewedData(10000000)

	b.ResetTimer()

	for i := 0; i < b.N; i++ {

		target := rand.Intn(dataSize)
		_ = interpolationSearch(items, target)
	}
}
