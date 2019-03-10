package umbrella

type Organism interface {
	Fitness() float64
	DNA() *DNA
	CalculateFitness()
	New() Organism
	Conceive(dna *DNA) Organism
}
