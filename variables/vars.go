package variables

var BerlinVariables = vars{
	PopulationSize: 200,
	Generations:    2500,
	TournamentSize: 5,
	CrossoverRate:  95,
	MutationRate:   10,
	ElitismCount:   2,
	CrossChance:    100,
}

var krVariables = vars{
	PopulationSize: 200,
	Generations:    2500,
	TournamentSize: 5,
	CrossoverRate:  95,
	MutationRate:   10,
	ElitismCount:   2,
	CrossChance:    0.95,
}

var prVariables = vars{
	PopulationSize: 200,
	Generations:    100000,
	TournamentSize: 3,
	CrossoverRate:  95,
	MutationRate:   10,
	ElitismCount:   2,
	CrossChance:    0.0,
}

type vars struct {
	PopulationSize int
	Generations    int
	TournamentSize int
	CrossoverRate  int
	MutationRate   int
	ElitismCount   int
	CrossChance    float64
}

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
