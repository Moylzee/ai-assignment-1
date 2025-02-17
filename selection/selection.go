package selection

// Package selection provides a selection mechanism for the genetic algorithms
import (
	"math"
	"math/rand"
)


// Function TournamentSelection performs tournament selection on a population
// Parameters:
//   population: 2D slice containing the population of solutions
//   fitness: slice containing fitness values for each solution
//   tournamentSize: number of individuals to compete in each tournament
// Returns:
//   The selected solution as a slice of integers 
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