package geneticalgorithm

// A genotype is an ordered list of two dimensional vectors in our application
type genotype []gene

func (g *genotype) equals(another genotype) bool {
	if len(*g) != len(another) {
		return false
	}
	for i, gene := range *g {
		if !gene.equals(another[i]) {
			return false
		}
	}
	return true
}

type gene struct {
	tailNode int
	headNode int
}

func (g *gene) equals(another gene) bool {
	return g.tailNode == another.tailNode && g.headNode == another.headNode
}

func (g *gene) mutate(rand randomiser) {
	// return gene{tailNode: int(rand.random()), headNode: int(rand.random())}
}
