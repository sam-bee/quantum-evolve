package geneticalgorithm

import (
	"fmt"
	"math/rand/v2"
)

type organism struct {
	genotype genotype
	fitness  float64
}

type population struct {
	organisms []organism
}

// We define a random number generator, injectable for testing. Random numbers are float64s between 0 and 1.

type randomiser interface {
	random() float64
}

type defaultRandomiser struct{}

func (r *defaultRandomiser) random() float64 {
	return rand.Float64()
}

// Defining the shape of a fitness function. We will not be using tournament selection, so the function will not need
// the rest of the population as a parameter.

type FitnessFunction func(org organism) float64

// The Genetic Algorithm

type GeneticAlgorithm struct {
	populationSize int
	mutationRate   float64
	crossoverRate  float64
	fitnessFunc    FitnessFunction
	randomiser     randomiser
}

func NewGeneticAlgorithm(popSize int, mutationRate float64, crossoverRate float64, fitnessFunc FitnessFunction) GeneticAlgorithm {

	rand := &defaultRandomiser{}

	return GeneticAlgorithm{
		populationSize: popSize,
		mutationRate:   mutationRate,
		crossoverRate:  crossoverRate,
		fitnessFunc:    fitnessFunc,
		randomiser:     rand,
	}
}

func (ga *GeneticAlgorithm) Run() {
	fmt.Println("Running Genetic Algorithm")
}

// Functions of a Genetic Algorithm

func breed(parents [2]organism, rand randomiser, crossoverRate float64, mutationRate float64) [2]organism {

	var child [2]organism

	// Either do the crossover, or just clone the parents
	trigger := rand.random()

	if trigger < crossoverRate {
		child[0] = crossover(parents, rand)
		child[1] = crossover(parents, rand)
	} else {
		child[0] = parents[0]
		child[1] = parents[1]
	}

	// Mutation happens to all organisms
	geneticMutation(&child[0], rand, mutationRate)
	geneticMutation(&child[1], rand, mutationRate)

	return child
}

func crossover(parents [2]organism, rand randomiser) organism {
	return organism{}
}

func geneticMutation(org *organism, rand randomiser, mutationRate float64) {

	// use a for loop with range to iterate over the genotype of the organism
	for i, gene := range org.genotype {
		if rand.random() < mutationRate {
			org.genotype[i] = gene.mutate(rand)
		}
	}

}
