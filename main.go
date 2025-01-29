package main

import (
	"ai-assignment-1/selection"
	"ai-assignment-1/variables"
	"bufio"
	"log"
	"math"
	"math/rand"
	"os"
	"strconv"
	"strings"
)

// City represents a city with an ID and coordinates.
type City struct {
	ID int
	X  float64
	Y  float64
}

func main() {
	filename := "data/berlin52.tsp"
	log.Printf("Reading File: %s", filename)
	cities, err := readTSPFile(filename)
	if err != nil {
		log.Fatalf("Error reading file: %v", err)
	}
	log.Printf("Successfully read %d cities", len(cities))

	vars := variables.LoadVariables()

	// Run the genetic algorithm
	bestTour := geneticAlgorithm(cities, vars.PopulationSize, vars.Generations, vars.TournamentSize, vars.CrossoverRate, vars.MutationRate, vars.ElitismCount)

	// Output the best tour and its distance
	bestDistance := calculateTourDistance(bestTour, cities)
	log.Printf("Best Tour: %v", bestTour)
	log.Printf("Best Tour Distance: %f", bestDistance)
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

// readTSPFile parses a TSPLIB file and returns a slice of City structs.
func readTSPFile(filename string) ([]City, error) {
	file, err := os.Open(filename)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	var cities []City
	scanner := bufio.NewScanner(file)

	// Read header and skip until "NODE_COORD_SECTION"
	for scanner.Scan() {
		line := strings.TrimSpace(scanner.Text())
		if line == "NODE_COORD_SECTION" {
			break
		}
	}

	// Parse city coordinates
	for scanner.Scan() {
		line := strings.TrimSpace(scanner.Text())
		if line == "EOF" {
			break
		}

		fields := strings.Fields(line)
		if len(fields) < 3 {
			continue
		}

		id, err := strconv.Atoi(fields[0])
		if err != nil {
			log.Printf("Skipping line due to invalid ID: %s", line)
			continue
		}
		x, err := strconv.ParseFloat(fields[1], 64)
		if err != nil {
			log.Printf("Skipping line due to invalid X coordinate: %s", line)
			continue
		}
		y, err := strconv.ParseFloat(fields[2], 64)
		if err != nil {
			log.Printf("Skipping line due to invalid Y coordinate: %s", line)
			continue
		}

		cities = append(cities, City{ID: id, X: x, Y: y})
	}

	// Check for any scanner errors
	if err := scanner.Err(); err != nil {
		return nil, err
	}

	return cities, nil
}

func calculateEuclideanDistance(cityA, cityB City) float64 {
	x := math.Pow(cityA.X-cityB.X, 2)
	y := math.Pow(cityA.Y-cityB.Y, 2)
	distance := math.Sqrt(x + y)
	return distance
}

func calculateTourDistance(tour []int, cities []City) float64 {
	toalDistance := 0.0
	for i := 0; i < len(tour)-1; i++ {
		toalDistance += calculateEuclideanDistance(cities[tour[i]], cities[tour[i+1]])
	}
    // Distance to return 
	toalDistance += calculateEuclideanDistance(cities[tour[len(tour)-1]], cities[tour[0]])
	return toalDistance
}

func evaluatePopulation(population [][]int, cities []City) []float64 {
	fitness := make([]float64, len(population))
	for i := 0; i < len(population); i++ {
		fitness[i] = calculateTourDistance(population[i], cities)
	}
	return fitness
}

// CROSSOVER FUNCTIONS

func orderedCrossover(p1, p2 []int) []int {
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
				if !contains(child, p2[j]) {
					child[i] = p2[j]
					break
				}
			}
		}
	}
	return child
}

func pmxCrossover(p1, p2 []int) []int {
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
				if !contains(child, p2[j]) {
					child[i] = p2[j]
					break
				}
			}
		}
	}
	return child
}

// MUTATIONS

func swap(tour []int) []int {
	i := rand.Intn(len(tour))
	j := rand.Intn(len(tour))
	tour[i], tour[j] = tour[j], tour[i]
	return tour
}

func inversionMutation(tour []int) []int {
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

// UTILS

func contains(tour []int, city int) bool {
	for _, c := range tour {
		if c == city {
			return true
		}
	}
	return false
}

func geneticAlgorithm(cities []City, populationSize, generations, tournamentSize, crossoverRate, mutationRate, elitismCount int) []int {
	population := generatePopulation(len(cities), populationSize)
	bestTour := population[0] // Start by assuming the first tour is the best
	bestFitness := math.MaxFloat64

	for gen := 0; gen < generations; gen++ {
		fitness := evaluatePopulation(population, cities)

		// Create the next generation
		nextGeneration := make([][]int, populationSize)

		// Elitism: carry over the best individuals
		for i := 0; i < elitismCount; i++ {
			bestIndex := findBestIndex(fitness)
			nextGeneration[i] = append([]int(nil), population[bestIndex]...)
			fitness[bestIndex] = math.MaxFloat64 // Mark as used
		}

		for i := elitismCount; i < populationSize; i++ {
			// Selection
			parent1 := selection.TournamentSelection(population, fitness, tournamentSize)
			parent2 := selection.TournamentSelection(population, fitness, tournamentSize)

			var child []int
			// Crossover
			if rand.Float64() < float64(crossoverRate)/100.0 {
				// Choose between OX or PMX based on some probability
				if rand.Float64() < 0.5 {
					child = orderedCrossover(parent1, parent2)
				} else {
					child = pmxCrossover(parent1, parent2)
				}
			} else {
				child = append([]int(nil), parent1...) // No crossover, just copy parent1
			}

			// Mutation
			if rand.Float64() < float64(mutationRate)/100.0 {
				// Choose mutation
				if rand.Float64() < 0.5 {
					child = swap(child)
				} else {
					child = inversionMutation(child)
				}
			}

			nextGeneration[i] = child
		}

		// Evaluate the next generation
		population = nextGeneration
		fitness = evaluatePopulation(population, cities)

		// Track the best fitness in this generation
		for i, f := range fitness {
			if f < bestFitness {
				bestFitness = f
				bestTour = population[i]
			}
		}

		log.Printf("Generation %d: Best Fitness = %f", gen, bestFitness)
	}

	return bestTour // Return the best solution found during all generations
}

func findBestIndex(fitness []float64) int {
	bestIndex := 0
	bestFitness := fitness[0]
	for i, f := range fitness {
		if f < bestFitness {
			bestFitness = f
			bestIndex = i
		}
	}
	return bestIndex
}
