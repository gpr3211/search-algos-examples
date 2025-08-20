package main

import (
	"math/rand"
	"sort"
	"time"
)

func randomSortedData(n int) []int {
	r := rand.New(rand.NewSource(time.Now().UnixNano()))
	data := make([]int, n)
	for i := range data {
		data[i] = r.Intn(n * 10) // allow duplicates
	}
	sort.Ints(data)
	return data
}
func uniformWithJitter(n int) []int {
	data := make([]int, n)
	for i := 0; i < n; i++ {
		data[i] = i*10 + rand.Intn(n/2000) // slight randomness
	}
	return data
}

func skewedData(n int) []int {
	data := make([]int, n)
	val := 0
	for i := 0; i < n; i++ {
		// Growth accelerates like i^2, but add some noise
		step := (i*i)/n + rand.Intn(10) // random jitter in steps
		val += step
		data[i] = val
	}
	return data
}
