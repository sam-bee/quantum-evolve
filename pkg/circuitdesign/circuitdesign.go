package circuitdesign

// A Circuit Design is an organism in the genetic algorithm. Its Phenotype should be a valid circuit diagram
// representing a quantum algorithm which could be run on a quantum computer, or simulated on classical hardware. All
// circuit designs in the system share a common Template of quantum logic gates. The Genotype of a Circuit Design is a
// set of directed wires, each plugging one quantum logic gate into another. The Phenotype is, therefore, a directed
// graph, where the nodes are quantum logic gates and the edges are wires connecting the gates.

// There is a single background template of quantum logic gates, common to all circuit designs. An individual circuit
// design is just the wiring instructions on how to connect the gates.

// Glossary of domain terms for the package:
// - Circuit design: An organism in the genetic algorithm. Represents a circuit of quantum logic gates.
// - Node: a quantum logic gate
// - Edge: a directed link between two nodes, acting as a wire in the circuit design
// - Orphan node: a node which can be safely deleted from the circuit design without changing behaviour, such as a node
//   with no incoming edges
// - Template: the background template of quantum logic gates. A large list of gates gates, with many duplicates. There
//   is one template in the whole system.
// - Genotype: a set of directed edges, representing the wiring of quantum logic gates in the circuit design
// - Phenotype: the actual circuit diagram. A directed graph of nodes and edges. The phenotype is derived from the
//   Genotype and Template. Orphan nodes are removed.

var templateSize = 200

type quantumGatesTemplate struct {
	nodes []node
}

var template = quantumGatesTemplate{
	nodes: []node{
		{name: "H"},
		{name: "X"},
		{name: "Y"},
		{name: "Z"},
		{name: "S"},
		{name: "T"},
		{name: "RX"},
		{name: "RY"},
		{name: "RZ"},
		{name: "CNOT"},
		{name: "CZ"},
		{name: "SWAP"},
	},
}


type circuitDesign struct {
	genotype genotype
	phenotype phenotype
}

type node struct {
	name string
}

type genotype []edge

type edge struct {
	tailNode int
	headNode int
}

type phenotype struct {
	nodes map[int]node
	edges []edge
}
