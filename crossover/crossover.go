package crossover

import (
	"ai-assignment-1/utilities"
	"math/rand"
)

func OrderedCrossover(p1, p2 []int) []int {
	size := len(p1)
	start := rand.Intn(size)
	end := rand.Intn(size)

	if start > end {
		start, end = end, start
	}

	child := make([]int, size)
	copy(child[start:end], p1[start:end])

	for i := 0; i < size; i++ {
		if i < start || i >= end {
			for j := 0; j < size; j++ {
				if !utilities.Contains(child, p2[j]) {
					child[i] = p2[j]
					break
				}
			}
		}
	}
	return child
}

func PmxCrossover(p1, p2 []int) []int {
	start := rand.Intn(len(p1))
	end := rand.Intn(len(p1))

	if start > end {
		start, end = end, start
	}

	child := make([]int, len(p1))
	copy(child[start:end], p1[start:end])

	for i := 0; i < len(p1); i++ {
		if i < start || i >= end {
			for j := 0; j < len(p2); j++ {
				if !utilities.Contains(child, p2[j]) {
					child[i] = p2[j]
					break
				}
			}
		}
	}
	return child
}
