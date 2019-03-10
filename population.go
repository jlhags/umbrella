package umbrella

import (
	"fmt"
	"math/rand"
	"time"
)

type Population struct {
	Organisms    []Organism
	MutationRate float64
	AlphaIndex   int
	Generation   int
	Size         int
}

func CreatePopulation(size int, organism Organism, mutationRate float64) *Population {

	rand.Seed(time.Now().UTC().UnixNano())
	var population Population
	population.Size = size
	population.MutationRate = mutationRate
	for i := 0; i < size; i++ {
		o := organism.New()
		population.Organisms = append(population.Organisms, o)
	}
	return &population
}

func (p *Population) PrintPopulation() {
	fmt.Printf("Generation %d (Mutation Rate: %f):\n", p.Generation, p.MutationRate)
	for _, o := range p.Organisms {
		fmt.Printf("\t%f - ", o.Fitness())
		o.DNA().Show()
	}
}

func (p *Population) PrintFitest() {
	fmt.Printf("Generation %d:\n\t", p.Generation)
	p.Organisms[p.AlphaIndex].DNA().Show()

}

func (p *Population) Survive() {
	alphaFitness := 0.0
	// Allow each organism to live
	for i, o := range p.Organisms {
		o.CalculateFitness()
		if o.Fitness() > alphaFitness {
			alphaFitness = o.Fitness()
			p.AlphaIndex = i
		}
	}
}

func chance(percent float64) bool {
	yeses := int(percent * 100)
	var wheel [100]bool
	for i := 0; i < yeses; i++ {
		wheel[i] = true
	}
	ret := wheel[rand.Intn(100)]
	return ret
}

func (p *Population) Selection() *Population {

	var p2 Population
	p2.MutationRate = p.MutationRate
	p2.Generation = p.Generation + 1
	p2.Size = p.Size
	p2.Organisms = append(p2.Organisms, p.Organisms[p.AlphaIndex])

	for len(p2.Organisms) < p.Size {
		for i, o := range p.Organisms {
			if i == p.AlphaIndex {
				continue
			}
			compatability := o.Fitness() / p.Organisms[p.AlphaIndex].Fitness()
			if o.Fitness() == 0 {
				// so you're saying there's a chance!
				compatability = p.MutationRate
			}
			if chance(compatability) == true {
				// fmt.Println("bow chikka bow wow")
				p2.Organisms = append(p2.Organisms, p.Mate(p.Organisms[p.AlphaIndex], o))

			} else {
				// fmt.Println("It's not you, it's me")
			}
			if len(p2.Organisms) == p.Size {
				break
			}
		}
	}
	return &p2
}

func (p *Population) Mate(o1 Organism, o2 Organism) Organism {
	d1 := o1.DNA()
	d2 := o2.DNA()
	var dna DNA
	for i := 0; i < len(d1.Genes); i++ {
		if rand.Intn(2) == 0 {

			dna.Genes = append(dna.Genes, d1.Genes[i].Copy())
		} else {
			dna.Genes = append(dna.Genes, d2.Genes[i].Copy())
		}
		if chance(p.MutationRate) == true {
			dna.Genes[i].Mutate()
		}

	}

	return o2.Conceive(&dna)
}
