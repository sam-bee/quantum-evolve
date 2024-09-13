package geneticalgorithm

// This is just a pseudo-random number generator. Genetic algoritms often need random numbers. Testing requires
// deterministic behaviour. We define three things in this file:

// 1. An interface for a random number generator;
// 2. A default implementation, being a thin wrapper round rand.Float64() and rand.IntN() from the standard library;
// 3. A fake implementation, which is only used in unit tests.

import (
	"math/rand/v2"
)

// 1. Main interface for a PRNG

type randomiser interface {
	randFloat64() float64 // Returns a random number between 0 and 1
	randIntN(n int) int   // Returns a random integer between 0 and n
}

// 2. Default (real) implementation of the 'randomiser' interface just calls rand.Float64() and rand.IntN()

type defaultRandomiser struct{}

var _ randomiser = &defaultRandomiser{}

func (r *defaultRandomiser) randFloat64() float64 {
	return rand.Float64()
}

func (r *defaultRandomiser) randIntN(n int) int {
	return rand.IntN(n)
}

// 3. Fake implemenation of the 'randomiser' interface. This is used in unit tests. It can be prepopulated with fake
// random numbers to return.

type fakeRandomiser struct {
	ints   []int
	floats []float64
}

var _ randomiser = &fakeRandomiser{}

func (r *fakeRandomiser) prepopFloats(rands []float64) {
	r.floats = rands
}

func (r *fakeRandomiser) prepopInts(rands []float64) {
	r.floats = rands
}

func (r *fakeRandomiser) randFloat64() float64 {
	if len(r.floats) == 0 {
		panic("No more random numbers")
	}
	rand := r.floats[0]
	r.floats = r.floats[1:]
	return rand
}

func (r *fakeRandomiser) randIntN(n int) int {
	if len(r.ints) == 0 {
		panic("No more random numbers")
	}
	rand := r.ints[0]
	r.ints = r.ints[1:]
	return rand
}
