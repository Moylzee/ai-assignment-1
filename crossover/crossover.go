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
