package mutations

// Package mutations provides mutation functions for the genetic algorithm

import "math/rand"

// Function Swap performs a simple swap mutation by exchanging two random elements in the tour
// params:
//   - tour: slice of integers representing the current tour
//
// returns:
//   - []int: modified tour with two elements swapped
func Swap(tour []int) []int {
	i := rand.Intn(len(tour))
	j := rand.Intn(len(tour))
	tour[i], tour[j] = tour[j], tour[i]
	return tour
}

// Function InversionMutation performs an inversion mutation by reversing a random subsection of the tour
// params:
//   - tour: slice of integers representing the current tour
//
// returns:
//   - []int: new tour with a subsection inverted
func InversionMutation(tour []int) []int {
	// Create a copy of the original tour to avoid modifying it directly
	mutatedTour := append([]int(nil), tour...)

	// Pick two random indices
	i := rand.Intn(len(mutatedTour))
	j := rand.Intn(len(mutatedTour))

	// Reverse the subsection of the tour
	if i > j {
		i, j = j, i
	}
	for i < j {
		mutatedTour[i], mutatedTour[j] = mutatedTour[j], mutatedTour[i]
		i++
		j--
	}
	return mutatedTour
}
