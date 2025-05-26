# Find if Path Exists in Graph

## Problem Description

You are given a bi-directional graph with `n` vertices, where each vertex is labeled from `0` to `n - 1` (inclusive). The edges in the graph are represented as a 2D integer array `edges`, where `edges[i] = [ui, vi]` denotes a bi-directional edge between vertex `ui` and vertex `vi`. Every vertex pair is connected by **at most one** edge, and no vertex has an edge to itself.

You need to determine if there is a **valid path** that exists from a given `source` vertex to a `destination` vertex.

Given `n`, `edges`, `source`, and `destination`, return `true` if there is a valid path from `source` to `destination`, or `false` otherwise.

## Constraints

* `1 <= n <= 2 * 10^5`
* `0 <= edges.length <= 2 * 10^5`
* `edges[i].length == 2`
* `0 <= ui, vi <= n - 1`
* `ui != vi`
* `0 <= source, destination <= n - 1`
* There are no duplicate edges.
* There are no self-edges.

## Examples

### Example 1:

Input: `n = 3`, `edges = [[0,1],[1,2],[2,0]]`, `source = 0`, `destination = 2`

Output: `true`

Explanation: There are two paths from vertex 0 to vertex 2:
- 0 -> 1 -> 2
- 0 -> 2

### Example 2:

Input: `n = 6`, `edges = [[0,1],[0,2],[3,5],[5,4],[4,3]]`, `source = 0`, `destination = 5`

Output: `false`

Explanation: There is no path from vertex 0 to vertex 5.

## Solution Approaches

This problem can be effectively solved using standard graph traversal algorithms and a Disjoint Set Union (Union-Find) data structure.

1.  **Breadth-First Search (BFS)**
2.  **Depth-First Search (DFS)** (Recursive and Iterative)
3.  **Union-Find (Disjoint Set Union)**

Each approach determines if the `source` and `destination` vertices belong to the same connected component.

### 1. Breadth-First Search (BFS)

**Concept:** BFS explores all the neighbor nodes at the present depth level before moving on to the nodes at the next depth level. It uses a queue to manage the nodes to visit.

**Algorithm:**
1.  Construct an adjacency list from the `edges`.
2.  Initialize a queue with the `source` vertex and a set to keep track of `visited` vertices.
3.  While the queue is not empty:
    * Dequeue a vertex.
    * If the dequeued vertex is the `destination`, return `true`.
    * Add all unvisited neighbors of the current vertex to the queue and mark them as visited.
4.  If the queue becomes empty and the `destination` was not reached, return `false`.

**Complexity:**
* **Time:** `O(V + E)`, where `V` is the number of vertices and `E` is the number of edges, as each vertex and edge is visited at most once.
* **Space:** `O(V + E)` for the adjacency list and `O(V)` for the queue and visited set.

### 2. Depth-First Search (DFS)

**Concept:** DFS explores as far as possible along each branch before backtracking. It can be implemented using recursion (which implicitly uses the call stack) or an explicit stack.

**Algorithm (Recursive):**
1.  Construct an adjacency list from the `edges`.
2.  Initialize a set to keep track of `visited` vertices.
3.  Define a recursive helper function `dfs(current_vertex)`:
    * Mark `current_vertex` as visited.
    * If `current_vertex` is the `destination`, return `true`.
    * For each unvisited `neighbor` of `current_vertex`, recursively call `dfs(neighbor)`. If any recursive call returns `true`, propagate `true` upwards.
    * If no path is found from `current_vertex`, return `false`.
4.  Call `dfs(source)` initially.

**Algorithm (Iterative):**
1.  Construct an adjacency list from the `edges`.
2.  Initialize a stack with the `source` vertex and a set to keep track of `visited` vertices.
3.  While the stack is not empty:
    * Pop a vertex from the stack.
    * If the popped vertex is the `destination`, return `true`.
    * For each unvisited `neighbor` of the current vertex, push it onto the stack and mark it as visited.
4.  If the stack becomes empty and the `destination` was not reached, return `false`.

**Complexity:**
* **Time:** `O(V + E)`, similar to BFS.
* **Space:** `O(V + E)` for the adjacency list and `O(V)` for the recursion stack (or explicit stack) and visited set.

### 3. Union-Find (Disjoint Set Union)

**Concept:** Union-Find is a data structure that maintains a collection of disjoint sets. It efficiently performs two operations: finding the representative (root) of a set and uniting two sets. If `source` and `destination` end up in the same set after processing all edges, a path exists.

**Algorithm:**
1.  Initialize a `parent` array where `parent[i] = i` for all `i` (each vertex is initially its own set).
2.  Implement a `find` operation with path compression: This function returns the representative of the set containing an element. It optimizes future lookups by flattening the tree structure.
3.  Implement a `union` operation: This function merges the sets of two elements by setting the parent of one set's root to the other set's root.
4.  Iterate through each `edge = [u, v]` in the `edges` array. For each edge, call `union(u, v)` to connect the sets containing `u` and `v`.
5.  After processing all edges, check if `find(source) == find(destination)`. If they share the same root, they are in the same connected component, so a path exists; otherwise, it doesn't.

**Complexity:**
* **Time:** `O((V + E) * α(V))`, where `α` is the inverse Ackermann function, which is practically a constant (very slow-growing). This makes the time complexity effectively `O(V + E)`.
* **Space:** `O(V)` for the parent array.

## Go Implementation

The provided Go solution demonstrates all three approaches.
