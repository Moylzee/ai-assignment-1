package mutations

import "math/rand"

func Swap(tour []int) []int {
	i := rand.Intn(len(tour))
	j := rand.Intn(len(tour))
	tour[i], tour[j] = tour[j], tour[i]
	return tour
}

func InversionMutation(tour []int) []int {
	// Create a copy of the original tour to avoid modifying it directly
	mutatedTour := append([]int(nil), tour...)

	// Pick two random indices
	i := rand.Intn(len(mutatedTour))
	j := rand.Intn(len(mutatedTour))

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
