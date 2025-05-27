# Satisfiability of Equality Equations

## Problem Description

You are given an array of strings `equations` that represent relationships between variables where each string `equations[i]` is of length 4 and takes one of two forms: "x==y" or "x!=y". Here, `xi` and `yi` are lowercase letters (not necessarily different) that represent one-letter variable names.

Return `true` if it is possible to assign integers to variable names so as to satisfy all the given equations, or `false` otherwise.

## Constraints

* `1 <= equations.length <= 500`
* `equations[i].length == 4`
* `equations[i][0]` is a lowercase letter.
* `equations[i][1]` is either `'='` or `'!'`.
* `equations[i][2]` is `'='`.
* `equations[i][3]` is a lowercase letter.

## Examples

### Example 1:

Input: equations = ["a==b","b!=a"]

Output: false

Explanation: If we assign say, a = 1 and b = 1, then the first equation is satisfied, but not the second. There is no way to assign the variables to satisfy both equations.

### Example 2:

Input: equations = ["b==a","a==b"]

Output: true

Explanation: We could assign a = 1 and b = 1 to satisfy both equations.

## Solution using Union-Find Algorithm

This problem is a classic application of the Union-Find (Disjoint Set Union) algorithm. The core idea is to treat each variable as a node in a graph. An "==" (equality) equation implies that two variables belong to the same set (or connected component). An "!=" (inequality) equation implies that two variables *must not* belong to the same set.

The strategy is to process all equality equations first to establish connections between variables. Then, for each inequality equation, we check if the implied condition is violated.

### Algorithm Steps:

1.  **Initialize Union-Find Structure:**
    * Since variables are lowercase letters (`'a'` through `'z'`), there are 26 possible variables.
    * Create a `parent` array (or map) of size 26. Initialize `parent[i] = i` for all `i`, meaning each variable is initially in its own set.
    * (Optional but recommended for efficiency) Create a `rank` or `size` array for union-by-rank or union-by-size optimization. Initialize all ranks/sizes to 0 or 1.

2.  **Process Equality Equations (`==`):**
    * Iterate through the `equations` array.
    * If an equation is of the form "x==y":
        * Get the integer indices for `x` and `y` (e.g., `'a'` -> 0, `'b'` -> 1, ..., `'z'` -> 25).
        * Perform a `union` operation on the sets containing `x` and `y`. This merges their sets, indicating they must have the same value.

3.  **Process Inequality Equations (`!=`):**
    * After all equality equations are processed, iterate through the `equations` array *again*.
    * If an equation is of the form "x!=y":
        * Get the integer indices for `x` and `y`.
        * Find the representatives (roots) of the sets containing `x` and `y` using the `find` operation.
        * If `find(x)` is equal to `find(y)`, it means `x` and `y` are in the same set, but the equation `x!=y` states they must be different. This is a contradiction, so we cannot satisfy all equations. In this case, return `false`.

4.  **All Equations Satisfied:**
    * If the loop finishes without finding any contradictions, it means all equations can be satisfied. Return `true`.

### Union-Find Helper Functions:

* **`find(i int)`:** Returns the representative (root) of the set that `i` belongs to. For optimal performance, this function should include **path compression**, which flattens the tree structure by making all nodes along the path point directly to the root.

* **`union(i, j int)`:** Merges the sets containing `i` and `j`. For optimal performance, this function should include **union by rank** (or union by size), which attaches the root of the smaller tree to the root of the larger tree, minimizing tree height.

### Why this approach works:

* **Equality:** By performing `union` operations for "x==y", we ensure that `x` and `y` (and all other variables transitively equal to them) are grouped into the same connected component. This implicitly means they must all take the same integer value.
* **Inequality:** When we encounter "x!=y", we check if `x` and `y` are *already* in the same connected component due to previous equality constraints. If they are, it's a contradiction, because `x` and `y` are forced to be equal by `==` equations but must be different by a `!=` equation. If they are in different components, there's no conflict, and we can satisfy `x!=y` by assigning them different values.

### Complexity Analysis:

* **Time Complexity:** `O(L * α(N))`, where `L` is the number of equations (`equations.length`), `N` is the number of variables (26 in this case), and `α` is the inverse Ackermann function, which is practically a constant (less than 5 for any practical input size). The operations `find` and `union` take nearly constant time on average with path compression and union by rank.
* **Space Complexity:** `O(N)` for the `parent` and `rank` arrays, where `N` is the number of variables (26).
