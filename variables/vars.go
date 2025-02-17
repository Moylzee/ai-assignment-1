package variables

// This file is used to store the variables for the genetic algorithm
// The variables are stored in a struct and can be accessed by importing this file

// Variables to use with the Berlin52 file
var BerlinVariables = vars{
	PopulationSize: 1000,
	Generations:    2500,
	TournamentSize: 5,
	CrossoverRate:  95,
	MutationRate:   5,
	ElitismCount:   2,
	CrossChance:    100,
}

// Variables to use with the kroA100 file
var krVariables = vars{
	PopulationSize: 1000,
	Generations:    2500,
	TournamentSize: 5,
	CrossoverRate:  95,
	MutationRate:   10,
	ElitismCount:   2,
	CrossChance:    0.95,
}

// Variables to use with the pr1002 file
var prVariables = vars{
	PopulationSize: 50,
	Generations:    100000,
	TournamentSize: 3,
	CrossoverRate:  95,
	MutationRate:   2,
	ElitismCount:   2,
	CrossChance:    0.0,
}

// Struct for the Variables
type vars struct {
	PopulationSize int
	Generations    int
	TournamentSize int
	CrossoverRate  int
	MutationRate   int
	ElitismCount   int
	CrossChance    float64
}

// Return the Variables depending on what file is being used
func LoadVariables(filename string) vars {
	switch filename {
	case "data/berlin52.tsp":
		return BerlinVariables
	case "data/kroA100.tsp":
		return krVariables
	case "data/pr1002.tsp":
		return prVariables
	default:
		return BerlinVariables
	}
}

// City represents a city with an ID and coordinates.
type City struct {
	ID int
	X  float64
	Y  float64
}
