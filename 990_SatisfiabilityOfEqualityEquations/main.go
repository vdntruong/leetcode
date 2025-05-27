package main

import (
	"fmt"
)

// UnionFind structure
type UnionFind struct {
	parent []int
	rank   []int // Used for union by rank optimization
}

// NewUnionFind creates a new UnionFind instance for 'n' elements
func NewUnionFind(n int) *UnionFind {
	parent := make([]int, n)
	rank := make([]int, n)
	for i := 0; i < n; i++ {
		parent[i] = i // Each element is initially its own parent
		rank[i] = 0   // Initial rank is 0
	}
	return &UnionFind{parent: parent, rank: rank}
}

// Find operation with path compression
func (uf *UnionFind) Find(i int) int {
	if uf.parent[i] == i {
		return i
	}
	uf.parent[i] = uf.Find(uf.parent[i]) // Path compression
	return uf.parent[i]
}

// Union operation with union by rank
func (uf *UnionFind) Union(i, j int) {
	rootI := uf.Find(i)
	rootJ := uf.Find(j)

	if rootI != rootJ {
		// Attach a smaller rank tree under the root of a larger rank tree
		if uf.rank[rootI] < uf.rank[rootJ] {
			uf.parent[rootI] = rootJ
		} else if uf.rank[rootI] > uf.rank[rootJ] {
			uf.parent[rootJ] = rootI
		} else {
			// Ranks are equal, pick one as root and increment its rank
			uf.parent[rootJ] = rootI
			uf.rank[rootI]++
		}
	}
}

// equationsPossible solves the Satisfiability of Equality Equations problem
func equationsPossible(equations []string) bool {
	// There are 26 possible lowercase variables ('a' through 'z')
	uf := NewUnionFind(26)

	// Phase 1: Process all equality equations ("x==y")
	for _, eq := range equations {
		if eq[1] == '=' { // Check for '=='
			// Convert characters to 0-25 integer indices
			var1 := int(eq[0] - 'a')
			var2 := int(eq[3] - 'a')
			uf.Union(var1, var2) // Union the sets of equal variables
		}
	}

	// Phase 2: Process all inequality equations ("x!=y")
	for _, eq := range equations {
		if eq[1] == '!' { // Check for '!='
			// Convert characters to 0-25 integer indices
			var1 := int(eq[0] - 'a')
			var2 := int(eq[3] - 'a')

			// If two variables that must be unequal are found in the same set,
			// then the equations are unsatisfiable.
			if uf.Find(var1) == uf.Find(var2) {
				return false // Contradiction found
			}
		}
	}

	// If no contradictions are found, all equations can be satisfied
	return true
}

func main() {
	// Example 1
	equations1 := []string{"a==b", "b!=a"}
	fmt.Printf("Equations: %v -> Possible: %t\n", equations1, equationsPossible(equations1)) // Expected: false

	// Example 2
	equations2 := []string{"b==a", "a==b"}
	fmt.Printf("Equations: %v -> Possible: %t\n", equations2, equationsPossible(equations2)) // Expected: true

	// Example 3: a==b, b==c, c!=a
	equations3 := []string{"a==b", "b==c", "c!=a"}
	fmt.Printf("Equations: %v -> Possible: %t\n", equations3, equationsPossible(equations3)) // Expected: false (a,b,c are all equal, contradicts c!=a)

	// Example 4: a==b, c==d, a!=c
	equations4 := []string{"a==b", "c==d", "a!=c"}
	fmt.Printf("Equations: %v -> Possible: %t\n", equations4, equationsPossible(equations4)) // Expected: true (a,b are equal; c,d are equal; a,c are different)

	// Example 5: a==a, b!=c, c==b
	equations5 := []string{"a==a", "b!=c", "c==b"}
	fmt.Printf("Equations: %v -> Possible: %t\n", equations5, equationsPossible(equations5)) // Expected: false (b!=c and c==b is contradiction)
}
