package geneticalgorithm

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGenotypeEqualityCheck(t *testing.T) {
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

func TestBreedWithoutCrossoverOrMutation(t *testing.T) {

	// We test breeding with no crossover or mutation
	crossoverRate := 0.0
	mutationRate := 0.0
	genotypeLen := 2

	org1 := organism{}
	org2 := organism{}
	parents := [2]organism{org1, org2}

	// The crossover rate and mutation rate will always be lower than the random numbers generated, so no crossover or
	// mutation will occur
	rand := fakeRandomiser{}
	rand.prepopFloats([]float64{0.5})

	children := breed(parents, &rand, crossoverRate, mutationRate, genotypeLen)

	firstChildIsCloneOfParent := children[0].genotype.equals(parents[0].genotype)
	secondChildIsCloneOfParent := children[1].genotype.equals(parents[1].genotype)

	assert.True(t, firstChildIsCloneOfParent)
	assert.True(t, secondChildIsCloneOfParent)
}

func TestBreedWithCrossover(t *testing.T) {

	t.Skip("Test skipped because crossover not implemented yet")

	// We test breeding with crossover. The 100% crossover probability will always be higher than the random number generated
	crossoverRate := 1.0
	mutationRate := 0.0
	genotypeLen := 2

	org1 := organism{genotype: []gene{{1, 2}, {3, 4}}}
	org2 := organism{genotype: []gene{{5, 6}, {7, 8}}}
	parents := [2]organism{org1, org2}

	// The crossover rate will always be higher than the random number generated, so crossover will occur
	rand := fakeRandomiser{}
	rand.prepopFloats([]float64{0.5})

	children := breed(parents, &rand, crossoverRate, mutationRate, genotypeLen)

	firstChildIsCorrect := children[0].genotype.equals([]gene{{1, 2}, {7, 8}})
	secondChildIsCorrect := children[1].genotype.equals([]gene{{5, 6}, {3, 4}})

	assert.True(t, firstChildIsCorrect)
	assert.True(t, secondChildIsCorrect)
}
