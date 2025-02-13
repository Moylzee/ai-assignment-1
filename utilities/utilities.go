package utilities

import (
	"ai-assignment-1/variables"
	"bufio"
	"log"
	"os"
	"strconv"
	"strings"
)

// Niche little helper function
func Contains(tour []int, city int) bool {
	for _, c := range tour {
		if c == city {
			return true
		}
	}
	return false
}

func Contains64(fitnessArray []float64, fitness float64) bool {
	for _, f := range fitnessArray {
		if f == fitness {
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
