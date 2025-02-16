package selection

import (
	"math"
	"math/rand"
)

func TournamentSelection(population [][]int, fitness []float64, tournamentSize int) []int {
	bestIndex := 0
	bestFitness := math.MaxFloat64

	for i := 0; i < tournamentSize; i++ {
		randIndex := rand.Intn(len(population))
		if fitness[randIndex] < bestFitness {
			bestFitness = fitness[randIndex]
			bestIndex = randIndex
		}
	}
	return population[bestIndex]
}