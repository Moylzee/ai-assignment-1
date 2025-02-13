package variables

var BerlinVariables = vars{
	PopulationSize: 200,
	Generations:    2500,
	TournamentSize: 5,
	CrossoverRate:  95,
	MutationRate:   10,
	ElitismCount:   2,
}

var PcbVariables = vars{
	PopulationSize: 500,  // Increased population size
	Generations:    5000, // Increased number of generations
	TournamentSize: 10,   // Increased tournament size
	CrossoverRate:  90,   // Adjusted crossover rate
	MutationRate:   15,   // Adjusted mutation rate
	ElitismCount:   5,    // Increased elitism count
}

var AliVariables = vars{
	PopulationSize: 300,
	Generations:    5000,
	TournamentSize: 5,
	CrossoverRate:  95,
	MutationRate:   10,
	ElitismCount:   2,
}

var krVariables = vars{
	PopulationSize: 300,
	Generations:    3000,
	TournamentSize: 5,
	CrossoverRate:  95,
	MutationRate:   10,
	ElitismCount:   2,
}

var prVariables = vars{
	PopulationSize: 500,
	Generations:    5000,
	TournamentSize: 5,
	CrossoverRate:  95,
	MutationRate:   10,
	ElitismCount:   2,
}

type vars struct {
	PopulationSize int
	Generations    int
	TournamentSize int
	CrossoverRate  int
	MutationRate   int
	ElitismCount   int
}


func LoadVariables(filename string) vars {
	switch filename {
	case "berlin":
		return BerlinVariables
	case "pcb":
		return PcbVariables
	case "ali":
		return AliVariables
	case "kr":
		return krVariables
	case "pr":
		return prVariables
	}
	return BerlinVariables
}

// City represents a city with an ID and coordinates.
type City struct {
	ID int
	X  float64
	Y  float64
}
