package geneticalgorithm

import (
	"fmt"
	"math/rand/v2"
)

// We define types for the Gene, Organism, and Population.

type gene struct {
	tailNode int
	headNode int
}

func (g *gene) equals(another gene) bool {
	return g.tailNode == another.tailNode && g.headNode == another.headNode
}

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

type organism struct {
	genotype genotype
	fitness  float64
}

type population struct {
	organisms []organism
}

// We define a random number generator, injectable for testing

type randomiser interface {
	random() float64
}

type defaultRandomiser struct{}

func (r *defaultRandomiser) random() float64 {
	return rand.Float64()
}

// Defining the shape of a fitness function. We will not be using tournament selection, so the function will not need
// the rest of the population.

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
	return [2]organism{organism{}, organism{}}
}
