package umbrella

import "fmt"

type DNA struct {
	Genes []Gene
}

func (d *DNA) Show() {
	for _, g := range d.Genes {
		fmt.Print(g.Phenotype())
	}
	fmt.Println()
}
