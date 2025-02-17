package crossover

// Package crossover provides functions for performing crossover operations in genetic algorithms

import (
	"math/rand"
)

// OrderedCrossover performs ordered crossover (OX) between two parent arrays
// The algorithm:
// 1. Randomly selects a subsequence from p1
// 2. Copies that subsequence to the child
// 3. Fills remaining positions with elements from p2 that aren't already used
// Parameters:
//
//	p1: First parent array
//	p2: Second parent array
//
// Returns:
//
//	A new child array
func OrderedCrossover(p1, p2 []int) []int {
	size := len(p1)
	start := rand.Intn(size)
	end := rand.Intn(size)

	if start > end {
		start, end = end, start
	}

	child := make([]int, size)
	copy(child[start:end], p1[start:end])

	// Create a set to track elements already in the child
	existing := make(map[int]bool)
	for i := start; i < end; i++ {
		existing[child[i]] = true
	}

	// Fill the rest of the child array
	index := end
	for i := 0; i < size; i++ {
		if !existing[p2[i]] {
			if index == size {
				index = 0
			}
			if index == start {
				index = end
			}
			child[index] = p2[i]
			index++
		}
	}

	return child
}

// PmxCrossover performs partially mapped crossover (PMX) between two parent arrays
// The algorithm:
// 1. Randomly selects a crossover segment
// 2. Copies the segment from p1 to child
// 3. Fills remaining positions with unused elements from p2
// Parameters:
//
//	p1: First parent array
//	p2: Second parent array
//
// Returns:
//
//	A new child array
func PmxCrossover(p1, p2 []int) []int {
	n := len(p1)
	child := make([]int, n)

	start, end := rand.Intn(n), rand.Intn(n)
	if start > end {
		start, end = end, start
	}
	used := make(map[int]bool, n)

	copy(child[start:end], p1[start:end])
	for i := start; i < end; i++ {
		used[p1[i]] = true
	}

	j := 0
	for i := 0; i < n; i++ {
		if i < start || i >= end {
			// Find next unused value from p2
			for used[p2[j]] {
				j++
			}
			child[i] = p2[j]
			used[p2[j]] = true
			j++
		}
	}

	return child
}
