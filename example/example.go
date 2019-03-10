package main

import (
	"flag"
	"math/rand"
	"umbrella"
)

type MyGene struct {
	Value byte
}

func (g *MyGene) Phenotype() string {
	return string(g.Value)
}

func (g *MyGene) Genotype() interface{} {
	return g.Value
}

func (g *MyGene) Copy() umbrella.Gene {
	return &MyGene{Value: g.Value}
}

func (g *MyGene) Mutate() {
	g.Value = RandByte()
}

type MyOrganism struct {
	dna     *umbrella.DNA
	fitness float64
}

func (m *MyOrganism) Fitness() float64 {
	return m.fitness
}

func (m *MyOrganism) DNA() *umbrella.DNA {
	return m.dna
}

const target = "Now is the time for all great men to come to the aid of their country. See the quick gray fox jump over the lazy dog."

func (m *MyOrganism) CalculateFitness() {
	m.fitness = 0
	for i, g := range m.DNA().Genes {
		if g.Genotype() == target[i] {
			m.fitness++
		}
	}
}

func (m *MyOrganism) Conceive(dna *umbrella.DNA) umbrella.Organism {
	return &MyOrganism{dna: dna}
}

func (m MyOrganism) New() umbrella.Organism {

	genes := []umbrella.Gene{}
	for i := 0; i < len(target); i++ {
		genes = append(genes, &MyGene{Value: RandByte()})
	}
	return &MyOrganism{dna: &umbrella.DNA{Genes: genes}}
}

const letterBytes = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ ."

func RandByte() byte {
	return letterBytes[rand.Intn(len(letterBytes))]
}

func main() {
	size := flag.Int("size", 100, "Population Size")
	mutationRate := flag.Float64("mutationRate", .01, "mutation rate")
	max := flag.Int("max", 100, "max generations")
	flag.Parse()
	var p2 *umbrella.Population
	m := &MyOrganism{}
	p := umbrella.CreatePopulation(*size, m, *mutationRate)
	for i := 0; i < *max; i++ {
		p.Survive()
		if p.AlphaIndex != 0 {
			p.PrintFitest()
		}
		//p.PrintPopulation()
		if p.Organisms[p.AlphaIndex].Fitness() == float64(len(target)) {
			break
		}
		p2 = p.Selection()
		p = p2
	}

}
