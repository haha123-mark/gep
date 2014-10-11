// odd-3-parity is a simple experiment to run the GEP algorithm using the Boolean logic package.
// Given a set of input functions (Not, And, and Or), this solves how to create an
// odd 3 parity function from those basic building blocks.
package main

import (
	"fmt"
	"log"
	"math/rand"
	"os"
	"time"

	bn "github.com/gmlewis/gep/functions/bool_nodes"
	"github.com/gmlewis/gep/gene"
	"github.com/gmlewis/gep/genome"
	"github.com/gmlewis/gep/grammars"
	"github.com/gmlewis/gep/model"
)

var parityTests = []struct {
	in  []bool
	out bool
}{
	{[]bool{false, false, false}, false},
	{[]bool{false, false, true}, true},
	{[]bool{false, true, false}, true},
	{[]bool{false, true, true}, false},
	{[]bool{true, false, false}, true},
	{[]bool{true, false, true}, false},
	{[]bool{true, true, false}, false},
	{[]bool{true, true, true}, true},
}

func init() {
	rand.Seed(time.Now().UnixNano())
}

func validateParity(g *genome.Genome) float64 {
	correct := 0
	for _, n := range parityTests {
		r := g.EvalBool(n.in, bn.BoolAllGates)
		if r == n.out {
			correct++
		}
	}
	return 1000.0 * float64(correct) / float64(len(parityTests))
}

func main() {
	funcs := []gene.FuncWeight{
		{"Not", 10},
		{"And", 20},
		{"Or", 20},
	}
	numIn := len(parityTests[0].in)
	e := model.New(funcs, bn.BoolAllGates, 30, 7, 3, numIn, "And", validateParity)
	s := e.Evolve(1000)

	// Write out the Go source code for the solution.
	gr, err := grammars.LoadGoBooleanAllGatesGrammar()
	if err != nil {
		log.Printf("unable to load Boolean grammar: %v\n", err)
	}
	fmt.Printf("\n// gepModel is auto-generated Go source code for the\n")
	fmt.Printf("// odd-3-parity solution karva expression:\n// %q, score=%v\n", s, validateParity(s))
	s.Write(os.Stdout, gr)
}