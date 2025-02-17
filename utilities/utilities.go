package utilities

import (
	"ai-assignment-1/variables"
	"bufio"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"math/rand"
	"os"
	"strconv"
	"strings"
)

// generateRandomTour generates a random tour of cities
// The Order of the cities in which the salesman will visit
func generateRandomTour(numCities int) []int {
	log.Println("Generating Random Tour")
	tour := make([]int, numCities)
	for i := 0; i < numCities; i++ {
		tour[i] = i
	}
	rand.Shuffle(len(tour), func(i, j int) {
		tour[i], tour[j] = tour[j], tour[i]
	})
	log.Println("Successfully Generated Random Tour")
	return tour
}

func GeneratePopulation(numCities, populationSize int) [][]int {
	log.Println("Generating Population")
	population := make([][]int, populationSize)
	for i := 0; i < populationSize; i++ {
		population[i] = generateRandomTour(numCities)
	}
	log.Println("Population Generated")
	return population
}

// Niche little helper function
func Contains(tour []int, city int) bool {
	for _, c := range tour {
		if c == city {
			return true
		}
	}
	return false
}

// readTSPFile parses a TSPLIB file and returns a slice of City structs.
func ReadTSPFile(filename string) ([]variables.City, error) {
	file, err := os.Open(filename)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	var cities []variables.City
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

		// Validity Checks Per Line
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

		cities = append(cities, variables.City{ID: id, X: x, Y: y})
	}

	// Check for any scanner errors
	if err := scanner.Err(); err != nil {
		return nil, err
	}

	return cities, nil
}

// Write Results to Files

func SaveDistances(distances map[int]int, fileDir string) {
	data := map[string]interface{}{
		"distances": distances,
	}

	jsonData, err := json.MarshalIndent(data, "", "  ")
	if err != nil {
		log.Fatalf("Error marshalling JSON: %v", err)
	}

	// Build the FilePath
	var dir string
	switch fileDir {
	case "data/berlin52.tsp":
		dir = "br"
	case "data/pr1002.tsp":
		dir = "pr"
	case "data/kroA100.tsp":
		dir = "kr"
	}
	filepath := fmt.Sprintf("results/%s/best_tour.json", dir)

	err = ioutil.WriteFile(filepath, jsonData, 0644)
	if err != nil {
		log.Fatalf("Error writing JSON file: %v", err)
	}

	log.Println("Distances saved to distances.json")
}

func SaveFitnesses(fitnesses map[int]float64, fileDir string) {
	data := map[string]interface{}{
		"fitnesses": fitnesses,
	}

	jsonData, err := json.MarshalIndent(data, "", "  ")
	if err != nil {
		log.Fatalf("Error marshalling JSON: %v", err)
	}

	// Build the FilePath
	var dir string
	switch fileDir {
	case "data/berlin52.tsp":
		dir = "br"
	case "data/pr1002.tsp":
		dir = "pr1"
	case "data/kroA100.tsp":
		dir = "kr"
	}
	filepath := fmt.Sprintf("results/%s/best_tour.json", dir)

	err = ioutil.WriteFile(filepath, jsonData, 0644)
	if err != nil {
		log.Fatalf("Error writing JSON file: %v", err)
	}

	log.Println("Fitnesses saved to fitnesses.json")
}

func SaveBestTour(cities []variables.City, bestTour []int, fileDir string) {
	data := map[string]interface{}{
		"cities":    cities,
		"best_tour": bestTour,
	}

	jsonData, err := json.MarshalIndent(data, "", "  ")
	if err != nil {
		log.Fatalf("Error marshalling JSON: %v", err)
	}

	// Build the FilePath
	var dir string
	switch fileDir {
	case "data/berlin52.tsp":
		dir = "br"
	case "data/pr1002.tsp":
		dir = "pr"
	case "data/kroA100.tsp":
		dir = "kr"
	}
	filepath := fmt.Sprintf("results/%s/best_tour.json", dir)

	err = ioutil.WriteFile(filepath, jsonData, 0644)
	if err != nil {
		log.Fatalf("Error writing JSON file: %v", err)
	}

	log.Println("Best tour and cities saved to best_tour.json")
}
