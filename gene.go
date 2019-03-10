package umbrella

type Gene interface {
	Phenotype() string
	Genotype() interface{}
	Mutate()
	Copy() Gene
}
