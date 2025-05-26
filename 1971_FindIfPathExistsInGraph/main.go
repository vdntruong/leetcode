package main

import "fmt"

func main() {
	// Example 1
	n1 := 3
	edges1 := [][]int{{0, 1}, {1, 2}, {2, 0}}
	source1 := 0
	destination1 := 2
	fmt.Printf("Example 1: n=%d, edges=%v, source=%d, destination=%d -> Path Exists: %t\n",
		n1, edges1, source1, destination1, validPath(n1, edges1, source1, destination1))

	// Example 2
	n2 := 6
	edges2 := [][]int{{0, 1}, {0, 2}, {3, 5}, {5, 4}, {4, 3}}
	source2 := 0
	destination2 := 5
	fmt.Printf("Example 2: n=%d, edges=%v, source=%d, destination=%d -> Path Exists: %t\n",
		n2, edges2, source2, destination2, validPath(n2, edges2, source2, destination2))

	// Additional test case: Disconnected graph
	n3 := 4
	edges3 := [][]int{{0, 1}, {2, 3}}
	source3 := 0
	destination3 := 3
	fmt.Printf("Example 3: n=%d, edges=%v, source=%d, destination=%d -> Path Exists: %t\n",
		n3, edges3, source3, destination3, validPath(n3, edges3, source3, destination3))

	// Additional test case: Self loop (should still work)
	n4 := 2
	edges4 := [][]int{{0, 1}, {1, 1}}
	source4 := 0
	destination4 := 1
	fmt.Printf("Example 4: n=%d, edges=%v, source=%d, destination=%d -> Path Exists: %t\n",
		n4, edges4, source4, destination4, validPath(n4, edges4, source4, destination4))

	// Additional test case: Source equals destination
	n5 := 5
	edges5 := [][]int{{0, 1}, {1, 2}, {2, 3}, {3, 4}}
	source5 := 2
	destination5 := 2
	fmt.Printf("Example 5: n=%d, edges=%v, source=%d, destination=%d -> Path Exists: %t\n",
		n5, edges5, source5, destination5, validPath(n5, edges5, source5, destination5))
}

// validPath returns true if there is a path between source and destination, false otherwise.
func validPath(n int, edges [][]int, source int, destination int) bool {
	// Uncomment the desired solution to use:
	return unionFindObjSolution(n, edges, source, destination)
	//return unionFindSolution(n, edges, source, destination)
	//return bfsSolution(n, edges, source, destination)
	//return dfsIterativeSolution(n, edges, source, destination)
	//return dfsRecursiveSolution(n, edges, source, destination)
}

// unionFindSolution implements the Union-Find algorithm.
func unionFindSolution(n int, edges [][]int, source int, destination int) bool {
	if source == destination {
		return true
	}

	parent := make(map[int]int)
	for i := 0; i < n; i++ {
		parent[i] = i
	}

	// 1. Union-Find Helper Function
	var find func(x int) int
	find = func(x int) int {
		if parent[x] != x {
			parent[x] = find(parent[x]) // Path Compression
		}
		return parent[x]
	}

	// 2. Union-Find Union Function
	union := func(x, y int) {
		rootX := find(x)
		rootY := find(y)
		if rootX != rootY {
			parent[rootX] = rootY // Union by Rank
		}
	}

	// 3. Union-Find Traversal
	for _, edge := range edges {
		s, d := edge[0], edge[1]
		union(s, d)
	}

	// 4. Check if the destination is found
	return find(source) == find(destination)
}

// unionFindOptimized implements an optimized Union-Find algorithm with both
// path compression and union by rank in a functional style.
func unionFindOptimized(n int, edges [][]int, source int, destination int) bool {
	if source == destination {
		return true
	}

	parent := make(map[int]int)
	rank := make(map[int]int)

	// Initialize each node as its own parent with rank 0
	for i := 0; i < n; i++ {
		parent[i] = i
		rank[i] = 0
	}

	// Find with path compression
	var find func(x int) int
	find = func(x int) int {
		if parent[x] != x {
			parent[x] = find(parent[x]) // Recursive path compression
		}
		return parent[x]
	}

	// Union by rank for balanced trees
	union := func(x, y int) {
		rootX := find(x)
		rootY := find(y)

		if rootX == rootY {
			return // Already in the same set
		}

		// Union by rank: attach a smaller rank tree under the root of a higher rank tree
		if rank[rootX] < rank[rootY] {
			parent[rootX] = rootY
		} else if rank[rootX] > rank[rootY] {
			parent[rootY] = rootX
		} else {
			// Equal ranks: make one root and increment its rank
			parent[rootY] = rootX
			rank[rootX]++
		}
	}

	// Process all edges
	for _, edge := range edges {
		s, d := edge[0], edge[1]
		union(s, d)
	}

	// Check connectivity
	return find(source) == find(destination)
}

// bfsSolution implements the Breadth-First Search algorithm.
func bfsSolution(n int, edges [][]int, source int, destination int) bool {
	if source == destination {
		return true
	}

	// 1. Create Adjacency List
	adj := make(map[int][]int)
	for _, edge := range edges {
		s, d := edge[0], edge[1]
		adj[s] = append(adj[s], d)
		adj[d] = append(adj[d], s) // bi-directional
	}

	// 2. Initialize Queue and Visited Map
	queue := []int{source}
	visited := make(map[int]bool)
	visited[source] = true

	// 3. BFS Traversal
	for len(queue) > 0 {
		curr := queue[0]
		queue = queue[1:] // Dequeue

		// 4. Check if the destination is found
		if curr == destination {
			return true
		}

		// 5. Add neighbours to the queue
		for _, neighbour := range adj[curr] {
			if !visited[neighbour] {
				visited[neighbour] = true
				queue = append(queue, neighbour) // Enqueue
			}
		}
	}

	// 6. Return false if the destination is not found
	return false
}

// dfsIterativeSolution implements the Depth-First Search algorithm (iterative).
func dfsIterativeSolution(n int, edges [][]int, source int, destination int) bool {
	if source == destination {
		return true
	}

	// 1. Create Adjacency List
	adj := make(map[int][]int)
	for _, edge := range edges {
		s, d := edge[0], edge[1]
		adj[s] = append(adj[s], d)
		adj[d] = append(adj[d], s) // bi-directional
	}

	// 2. Initialize Stack and Visited Map
	stack := []int{source}
	visited := make(map[int]bool)
	visited[source] = true

	// 3. DFS Traversal
	for len(stack) > 0 {
		curr := stack[len(stack)-1]  // Peek
		stack = stack[:len(stack)-1] // Pop

		// 4. Check if the destination is found
		if curr == destination {
			return true
		}

		// 5. Recurse for neighbours
		for _, neighbour := range adj[curr] {
			if !visited[neighbour] {
				visited[neighbour] = true
				stack = append(stack, neighbour) // Push
			}
		}
	}

	// 6. Return false if the destination is not found
	return false
}

// dfsRecursiveSolution implements the Depth-First Search algorithm (recursive).
func dfsRecursiveSolution(n int, edges [][]int, source int, destination int) bool {
	if source == destination {
		return true
	}

	// 1. Create Adjacency List
	adj := make(map[int][]int)
	for _, edge := range edges {
		s, d := edge[0], edge[1]
		adj[s] = append(adj[s], d)
		adj[d] = append(adj[d], s) // bi-directional
	}

	// 2. Initialize Visited Map
	visited := make(map[int]bool)

	// 3. DFS Helper Function
	var dfs func(curr int) bool
	dfs = func(curr int) bool {
		// 4. Check if the destination is found
		visited[curr] = true
		if curr == destination {
			return true
		}

		// 5. Recurse for neighbours
		for _, neighbour := range adj[curr] {
			if !visited[neighbour] {
				if dfs(neighbour) {
					return true
				}
			}
		}
		return false
	}

	// 6. DFS Traversal
	return dfs(source)
}

type ParentData struct {
	arr  []int
	rank []int
}

func (p *ParentData) find(i int) int {
	for p.arr[i] != i {
		i = p.arr[i]
	}
	return i
}

func (p *ParentData) union(i, j int) {
	iRoot, jRoot := p.find(i), p.find(j)

	if iRoot == jRoot {
		return
	}

	if p.rank[iRoot] < p.rank[jRoot] {
		p.arr[iRoot] = jRoot
	} else if p.rank[iRoot] > p.rank[jRoot] {
		p.arr[jRoot] = iRoot
	} else {
		p.arr[jRoot] = iRoot
		p.rank[iRoot]++
	}
}

// unionFindObjSolution implements the Union-Find algorithm using an object-oriented approach.
func unionFindObjSolution(n int, edges [][]int, source int, destination int) bool {
	var parent ParentData
	parent.arr = make([]int, n)
	parent.rank = make([]int, n)

	for i := 0; i < n; i++ {
		parent.arr[i] = i
	}
	for _, edge := range edges {
		s, d := edge[0], edge[1]
		parent.union(s, d)
	}

	return parent.find(source) == parent.find(destination)
}
