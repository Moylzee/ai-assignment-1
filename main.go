package main

import (
	"ai-assignment-1/crossover"
	"ai-assignment-1/mutations"
	"ai-assignment-1/selection"
	"ai-assignment-1/utilities"
	"ai-assignment-1/variables"
	"fmt"
	"log"
	"math"
	"math/rand"
	"strconv"
	"time"
)

// Variables for the main file
var (
	berlin = "data/berlin52.tsp"
	pr     = "data/pr1002.tsp"
	kr     = "data/kroA100.tsp"
	D      map[int]int
	F      map[int]float64
)

func main() {
	filename := kr // The file in which we will run | Change this if you want to run a different file
	log.Printf("Reading File: %s", filename)

	// Read The cities from the file
	cities, err := utilities.ReadTSPFile(filename)
	if err != nil {
		log.Fatalf("Error reading file: %v", err)
	}
	log.Printf("Successfully read %d cities", len(cities))

	// Load Variables for specific file
	vars := variables.LoadVariables(filename)

	// Run the genetic algorithm, tracking the computational time
	start := time.Now()
	bestTour := geneticAlgorithm(cities, vars.PopulationSize, vars.Generations, vars.TournamentSize, vars.CrossoverRate, vars.MutationRate, vars.ElitismCount, vars.CrossChance)
	// Output the best tour and its distance
	bestDistance := calculateTourDistance(bestTour, cities)
	elapsed := time.Since(start)

	log.Printf("Best Tour: %v", bestTour)
	log.Printf("Best Tour Distance: %f", bestDistance)
	log.Printf("Time taken: %s", elapsed)

	// Save the best tour and cities to a JSON file
	utilities.SaveBestTour(cities, bestTour, filename)
	utilities.SaveDistances(D, filename)
	utilities.SaveFitnesses(F, filename)
}

// Function to Calculate the Euclidean Distance between two cities
func calculateEuclideanDistance(cityA, cityB variables.City) float64 {
	x := math.Pow(cityA.X-cityB.X, 2)
	y := math.Pow(cityA.Y-cityB.Y, 2)
	distance := math.Sqrt(x + y)
	return distance
}

// calculateTourDistance will calculate the total distance of the tour
func calculateTourDistance(tour []int, cities []variables.City) float64 {
	toalDistance := 0.0
	for i := 0; i < len(tour)-1; i++ {
		toalDistance += calculateEuclideanDistance(cities[tour[i]], cities[tour[i+1]])
	}
	// Distance to return
	toalDistance += calculateEuclideanDistance(cities[tour[len(tour)-1]], cities[tour[0]])
	return toalDistance
}

func evaluatePopulation(population [][]int, cities []variables.City) []float64 {
	fitness := make([]float64, len(population))
	for i := 0; i < len(population); i++ {
		fitness[i] = calculateTourDistance(population[i], cities)
	}
	return fitness
}

func geneticAlgorithm(cities []variables.City, populationSize, generations, tournamentSize, crossoverRate, mutationRate, elitismCount int, cc float64) []int {
	// Generate the Initial Population
	population := utilities.GeneratePopulation(len(cities), populationSize)
	bestTour := population[0] // Start by assuming the first tour is the best
	bestDistance := math.MaxFloat64
	distances := make(map[int]int)
	fitnesses := make(map[int]float64)

	for gen := 0; gen < generations; gen++ {
		fitness := evaluatePopulation(population, cities)

		// Create the next generation
		nextGeneration := make([][]int, populationSize)

		// Elitism: carry over the best individuals
		for i := 0; i < elitismCount; i++ {
			bestIndex := findBestIndex(fitness)
			nextGeneration[i] = append([]int(nil), population[bestIndex]...)
			fitness[bestIndex] = math.MaxFloat64
		}

		for i := elitismCount; i < populationSize; i++ {
			parent1 := selection.TournamentSelection(population, fitness, tournamentSize)
			parent2 := selection.TournamentSelection(population, fitness, tournamentSize)
			// Selection
			var child []int
			// Crossover
			if rand.Float64() < float64(crossoverRate)/100.0 {
				// Choose between OX or PMX based on some probability
				if rand.Float64() < cc {
					child = crossover.OrderedCrossover(parent1, parent2)
				} else {
					child = crossover.PmxCrossover(parent1, parent2)
				}
			} else {
				child = append([]int(nil), parent1...) // No crossover, just copy parent1
			}

			// Mutation
			if rand.Float64() < float64(mutationRate)/100.0 {
				// Choose mutation
				if rand.Float64() < 0.5 {
					child = mutations.Swap(child)
				} else {
					child = mutations.InversionMutation(child)
				}
			}
			nextGeneration[i] = child
		}

		// Evaluate the next generation
		population = nextGeneration
		fitness = evaluatePopulation(population, cities)

		// Track the best fitness in this generation
		for i, f := range fitness {
			if f < bestDistance {
				bestDistance = f
				bestTour = population[i]
			}
		}

		// This code slice tracks every unique distance in a map for plotting
		exists := false
		for _, v := range distances {
			if v == int(bestDistance) {
				exists = true
				break
			}
		}
		if !exists {
			distances[gen] = int(bestDistance)
		}

		// Calculate the fitness (1 / Distance)
		fit := 1 / bestDistance
		formattedFit := fmt.Sprintf("%.6f", fit)
		fit, _ = strconv.ParseFloat(formattedFit, 64)

		// This code slice tracks every unique fitness in a map for plotting
		fitExists := false
		for _, v := range fitnesses {
			if v == fit {
				fitExists = true
				break
			}
		}
		if !fitExists {
			fitnesses[gen] = fit
		}

		log.Printf("Generation %d: Best Distance = %f | Fitness: %f", gen, bestDistance, fit)
	}

	// Assign the Map of fitnesses and Distances to the global Variables
	F = fitnesses
	D = distances
	return bestTour // Return the best solution found during all generations
}

func findBestIndex(fitness []float64) int {
	bestIndex := 0
	bestDistance := fitness[0]
	for i, f := range fitness {
		if f < bestDistance {
			bestDistance = f
			bestIndex = i
		}
	}
	return bestIndex
}
