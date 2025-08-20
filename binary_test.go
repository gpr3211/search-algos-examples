package main

import (
	"math/rand"
	"testing"
)

func BenchmarkBinarySorted(b *testing.B) {

	items := make([]int, 10_000_000)
	for i := range items {
		items[i] = i
	}

	b.ResetTimer()

	for i := 0; i < b.N; i++ {

		target := rand.Intn(dataSize)
		_ = binarySearch(items, target)
	}
}
func BenchmarkBinaryRandom(b *testing.B) {

	items := randomSortedData(10000000)

	b.ResetTimer()

	for i := 0; i < b.N; i++ {

		target := rand.Intn(dataSize)
		_ = binarySearch(items, target)
	}
}
func BenchmarkBinaryUniformJitter(b *testing.B) {

	items := uniformWithJitter(10000000)

	b.ResetTimer()

	for i := 0; i < b.N; i++ {

		target := rand.Intn(dataSize)
		_ = binarySearch(items, target)
	}
}
func BenchmarkBinarySkewed(b *testing.B) {

	items := skewedData(10000000)

	b.ResetTimer()

	for i := 0; i < b.N; i++ {

		target := rand.Intn(dataSize)
		_ = binarySearch(items, target)
	}
}
