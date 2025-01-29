package selection

import (
	"math"
	"math/rand"
)

func RouletteWheelSelection(population [][]int, fitnessScores []float64, tournamentSize int) []int {
	totalFitness := 0.0

	// Calculate the total fitness
	for _, score := range fitnessScores {
		totalFitness += score
	}

	// Normalize the fitness scores
	normalizedFitnessScores := make([]float64, len(fitnessScores))
	for i, score := range fitnessScores {
		normalizedFitnessScores[i] = score / totalFitness
	}

	randomNumber := rand.Float64()

	// Select an individual based on the roulette wheel selection
	for i, score := range normalizedFitnessScores {
		randomNumber -= score
		if randomNumber <= 0 {
			return population[i]
		}
	}

	return population[len(population)-1]
}

func TournamentSelection(population [][]int, fitness []float64, tournamentSize int) []int {
	bestIndex := -1
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
