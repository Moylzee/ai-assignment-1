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
		PopulationSize: 200,
		Generations:    2500,
		TournamentSize: 5,
		CrossoverRate:  95,
		MutationRate:   15,
		ElitismCount:   2,
	}
}
