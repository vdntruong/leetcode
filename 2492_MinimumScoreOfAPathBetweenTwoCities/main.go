package main

import (
	"fmt"
	"math"
)

func main() {
	// Example 1:
	n1 := 4
	roads1 := [][]int{{1, 2, 9}, {2, 3, 6}, {2, 4, 5}, {1, 4, 7}}
	fmt.Printf("n=%d, roads=%v -> Min Score: %d (Expected: 5)\n", n1, roads1, minScore(n1, roads1))

	// Example 2:
	n2 := 2
	roads2 := [][]int{{1, 2, 10}}
	fmt.Printf("n=%d, roads=%v -> Min Score: %d (Expected: 10)\n", n2, roads2, minScore(n2, roads2))

	// Example 3: Disconnected graph (N=5) but 1 is connected to 3.
	// If the destination N=5, then minScore for 1 to 5 would be MaxInt32 if 5 is not reachable.
	// However, the problem states "at least one path between 1 and n".
	// Let's use a case where N is reachable from 1.
	n3 := 3                                                  // Adjusted N for roads3_test to make it a valid test case per problem constraint
	roads3_test := [][]int{{1, 2, 2}, {2, 3, 5}, {4, 5, 10}} // Roads 4,5,10 are in a separate component
	fmt.Printf("n=%d, roads=%v -> Min Score: %d (Expected: 2)\n", n3, roads3_test, minScore(n3, roads3_test))

	// Example 4: Linear path
	n4 := 3
	roads4 := [][]int{{1, 2, 10}, {2, 3, 20}}
	fmt.Printf("n=%d, roads=%v -> Min Score: %d (Expected: 10)\n", n4, roads4, minScore(n4, roads4))
}

// minScore finds the minimum score of a path between city 1 and city n.
// The score is defined as the minimum distance of a road in that path.
// The problem statement guarantees at least one path between 1 and n.
func minScore(n int, roads [][]int) int {
	// Build an adjacency list to represent the graph.
	// Each entry adj[city] will be a slice of arrays/tuples: [[neighbor_city, distance]]
	adj := make(map[int][][2]int)
	for _, road := range roads {
		u, v, d := road[0], road[1], road[2]
		// Add bi-directional edges
		adj[u] = append(adj[u], [2]int{v, d})
		adj[v] = append(adj[v], [2]int{u, d})
	}

	// Initialize BFS.
	// We start BFS from city 1, as we need to find paths from city 1.
	queue := []int{1}
	// A map to keep track of visited cities to prevent cycles and redundant processing.
	visited := make(map[int]bool)
	visited[1] = true // Mark the starting city as visited.

	// Initialize minOverallScore to a very large value.
	// This variable will store the minimum distance found among all roads
	// within the connected component reachable from city 1.
	minOverallScore := math.MaxInt32

	// Perform BFS traversal.
	for len(queue) > 0 {
		// Dequeue the current city.
		curr := queue[0]
		queue = queue[1:]

		// Explore all neighbors of the current city.
		for _, neighborInfo := range adj[curr] {
			neighbor := neighborInfo[0]
			distance := neighborInfo[1]

			// The core insight: Any road encountered within the connected component
			// (reachable from city 1) is a candidate for the overall minimum score.
			// This is because if city 1 can reach 'curr', and 'curr' is connected to 'neighbor'
			// by 'distance', then this 'distance' can be the minimum edge on *some* path
			// from 1 to n (since n is guaranteed to be reachable).
			if distance < minOverallScore {
				minOverallScore = distance
			}

			// If the neighbor hasn't been visited yet, mark it as visited and enqueue it.
			if !visited[neighbor] {
				visited[neighbor] = true
				queue = append(queue, neighbor)
			}
		}
	}

	// After BFS completes, minOverallScore will hold the minimum distance
	// of any road in the connected component containing city 1 (and city n).
	return minOverallScore
}
