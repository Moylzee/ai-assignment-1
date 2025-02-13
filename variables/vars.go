package variables

type vars struct {
	PopulationSize int
	Generations    int
	TournamentSize int
	CrossoverRate  int
	MutationRate   int
	ElitismCount   int
}

func LoadVariables() vars {
	return vars{
		PopulationSize: 500,
		Generations:    1000,
		TournamentSize: 5,
		CrossoverRate:  95,
		MutationRate:   5,
		ElitismCount:   2,
	}
}

// City represents a city with an ID and coordinates.
type City struct {
	ID int
	X  float64
	Y  float64
}
