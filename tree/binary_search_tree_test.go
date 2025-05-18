package tree_test

import (
	"math/rand"
	"testing"

	"github.com/josestg/dsa/adt/adttest"
	"github.com/josestg/dsa/tree"
)

func TestBinarySearchTree(t *testing.T) {
	c := tree.NewBinarySearchTree[int]
	g := func() int {
		return rand.Intn(128)
	}

	tests := []struct {
		name      string
		simulator adttest.Runner
	}{
		{name: "insert", simulator: adttest.BSTInsertSimulator(c, g)},
		{name: "in order", simulator: adttest.BSTInOrderSimulator(c, g)},
		{name: "min max", simulator: adttest.BSTMinMaxSimulator(c, g)},
		{name: "delete", simulator: adttest.BSTDeleteSimulator(c, g)},
	}

	for _, tt := range tests {
		t.Run(tt.name, tt.simulator)
	}
}
