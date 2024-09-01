package geneticalgorithm

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

// We define a fake randomiser for testing purposes

type fakeRandomiser struct {
	vals []float64
}

func (r *fakeRandomiser) random() float64 {
	if len(r.vals) == 0 {
		panic("No more random numbers")
	}
	rand := r.vals[0]
	r.vals = r.vals[1:]
	return rand
}

func (r *fakeRandomiser) prepop(rands []float64) {
	r.vals = rands
}

func testGenotypeEqualityCheck(t *testing.T) {
	// We test the equality check for genotypes
	gene1 := gene{tailNode: 1, headNode: 2}
	gene2 := gene{tailNode: 1, headNode: 4}
	gene3 := gene{tailNode: 5, headNode: 6}

	var genotype1 genotype = []gene{gene1, gene2}
	var genotype2 genotype = []gene{gene1, gene2}
	var genotype3 genotype = []gene{gene1, gene3}

	assert.True(t, genotype1.equals(genotype2))
	assert.False(t, genotype1.equals(genotype3))
}

func testBreedWithoutCrossoverOrMutation(t *testing.T) {

	// We test breeding with no crossover or mutation
	crossoverRate := 0.0
	mutationRate := 0.0

	org1 := organism{}
	org2 := organism{}
	parents := [2]organism{org1, org2}

	// The crossover rate and mutation rate will always be lower than the random numbers generated, so no crossover or
	// mutation will occur
	rand := fakeRandomiser{vals: []float64{0.5, 0.5}}

	children := breed(parents, &rand, crossoverRate, mutationRate)

	firstChildIsCloneOfParent := children[0].genotype.equals(parents[0].genotype)
	secondChildIsCloneOfParent := children[1].genotype.equals(parents[1].genotype)

	assert.True(t, firstChildIsCloneOfParent)
	assert.True(t, secondChildIsCloneOfParent)

}
