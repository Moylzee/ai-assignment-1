package main

import (
	"ai-assignment-1/crossover"
	"ai-assignment-1/mutations"
	"ai-assignment-1/selection"
	"ai-assignment-1/utilities"
	"ai-assignment-1/variables"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"math"
	"math/rand"
	"strconv"
)

var (
	berlin = "data/berlin52.tsp"
	pr     = "data/pr1002.tsp"
	kr     = "data/kroA100.tsp"
	D      map[int]int
	F      map[int]float64
)

func main() {
	filename := kr
	log.Printf("Reading File: %s", filename)

	// Read The cities from the file
	cities, err := utilities.ReadTSPFile(filename)
	if err != nil {
		log.Fatalf("Error reading file: %v", err)
	}

	log.Printf("Successfully read %d cities", len(cities))

	vars := variables.LoadVariables(filename)

	// Run the genetic algorithm
	log.Println(vars)
	bestTour := geneticAlgorithm(cities, vars.PopulationSize, vars.Generations, vars.TournamentSize, vars.CrossoverRate, vars.MutationRate, vars.ElitismCount, vars.CrossChance)
	// Output the best tour and its distance
	bestDistance := calculateTourDistance(bestTour, cities)

	log.Printf("Best Tour: %v", bestTour)
	log.Printf("Best Tour Distance: %f", bestDistance)

	// Save the best tour and cities to a JSON file
	saveBestTour(cities, bestTour)
	saveDistances(D)
	saveFitnesses(F)
}

// generateRandomTour generates a random tour of cities
// The Order of the cities in which the salesman will visit
func generateRandomTour(numCities int) []int {
	tour := make([]int, numCities)
	for i := 0; i < numCities; i++ {
		tour[i] = i
	}
	rand.Shuffle(len(tour), func(i, j int) {
		tour[i], tour[j] = tour[j], tour[i]
	})
	return tour
}

func generatePopulation(numCities, populationSize int) [][]int {
	log.Println("Generating Population")
	population := make([][]int, populationSize)
	for i := 0; i < populationSize; i++ {
		population[i] = generateRandomTour(numCities)
	}
	log.Println("Population Generated")
	return population
}

func calculateEuclideanDistance(cityA, cityB variables.City) float64 {
	x := math.Pow(cityA.X-cityB.X, 2)
	y := math.Pow(cityA.Y-cityB.Y, 2)
	distance := math.Sqrt(x + y)
	return distance
}

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
	population := generatePopulation(len(cities), populationSize)
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

		// Check and update distances map
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

		fit := 1 / bestDistance
		formattedFit := fmt.Sprintf("%.6f", fit)
		fit, _ = strconv.ParseFloat(formattedFit, 64)

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

func saveBestTour(cities []variables.City, bestTour []int) {
	data := map[string]interface{}{
		"cities":    cities,
		"best_tour": bestTour,
	}

	jsonData, err := json.MarshalIndent(data, "", "  ")
	if err != nil {
		log.Fatalf("Error marshalling JSON: %v", err)
	}

	err = ioutil.WriteFile("results/best_tour.json", jsonData, 0644)
	if err != nil {
		log.Fatalf("Error writing JSON file: %v", err)
	}

	log.Println("Best tour and cities saved to best_tour.json")
}

func saveDistances(distances map[int]int) {
	data := map[string]interface{}{
		"distances": distances,
	}

	jsonData, err := json.MarshalIndent(data, "", "  ")
	if err != nil {
		log.Fatalf("Error marshalling JSON: %v", err)
	}

	err = ioutil.WriteFile("results/distances.json", jsonData, 0644)
	if err != nil {
		log.Fatalf("Error writing JSON file: %v", err)
	}

	log.Println("Distances saved to distances.json")
}

func saveFitnesses(fitnesses map[int]float64) {
	data := map[string]interface{}{
		"fitnesses": fitnesses,
	}

	jsonData, err := json.MarshalIndent(data, "", "  ")
	if err != nil {
		log.Fatalf("Error marshalling JSON: %v", err)
	}

	err = ioutil.WriteFile("results/fitnesses.json", jsonData, 0644)
	if err != nil {
		log.Fatalf("Error writing JSON file: %v", err)
	}

	log.Println("Fitnesses saved to fitnesses.json")
}
