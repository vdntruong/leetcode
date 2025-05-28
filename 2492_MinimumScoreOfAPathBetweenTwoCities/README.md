# Minimum Score of a Path Between Two Cities

## Problem Type:
Graph Traversal / Connected Components

## Problem Statement:
You are given a bi-directional graph consisting of `n` cities, numbered from 1 to `n`. The connections between cities are defined by a 2D integer array `roads`, where each `roads[i] = [ai, bi, distancei]` represents a road between city `ai` and city `bi` with a given `distancei`. The graph is not necessarily connected, but there is a guarantee that **at least one path exists between city 1 and city `n`**.

## Definitions:
* **Path:** A sequence of roads connecting two cities.
* **Score of a Path:** The minimum distance of a road encountered *in that path*.

## Goal:
Return the **minimum possible score** of a path between city 1 and city `n`.

## Key Notes & Clarifications:
* Paths can contain the same road and visit cities (including 1 and `n`) multiple times. This implies that the specific path chosen doesn't need to be simple (no repeated vertices/edges); only reachability matters for considering which roads are part of a connected component.
* The crucial guarantee: "The test cases are generated such that there is **at least one path between 1 and `n`**." This simplifies the problem as it means city `n` will always be in the same connected component as city 1.

## Core Insight for Solution:
Since the problem asks for the *minimum possible score* of *a* path between city 1 and city `n`, and given that a path can contain roads multiple times, any road that is part of the **connected component** containing both city 1 and city `n` can effectively be considered "on a path" between them.

Therefore, the problem simplifies to: **Find the minimum distance among all roads that belong to the same connected component as city 1 (and implicitly, city `n`).**

You don't need to find a specific shortest path; you just need to explore all reachable roads from city 1 and find the smallest distance among all such roads.

## Constraints:
* `2 <= n <= 10^5` (Number of cities)
* `1 <= roads.length <= 10^5` (Number of roads)
* `roads[i].length == 3`
* `1 <= ai, bi <= n`
* `ai != bi` (No self-loops)
* `1 <= distancei <= 10^4` (Distance on a road)
* There are no repeated edges.
* There is at least one path between 1 and `n`.

## Example 1:

Input: n = 4, roads = [[1,2,9],[2,3,6],[2,4,5],[1,4,7]]

Output: 5

Explanation: The path from city 1 to 4 with the minimum score is: 1 -> 2 -> 4. The score of this path is min(9,5) = 5. It can be shown that no other path has less score.

## Example 2:

Input: n = 4, roads = [[1,2,2],[1,3,4],[3,4,7]]

Output: 2

Explanation: The path from city 1 to 4 with the minimum score is: 1 -> 2 -> 1 -> 3 -> 4. The score of this path is min(2,2,4,7) = 2.

## Common Algorithms for Solution:
* **Breadth-First Search (BFS)**
* **Depth-First Search (DFS)**
* **Union-Find**

All these algorithms can be used to traverse the connected component containing city 1 and simultaneously find the minimum distance among all roads encountered in that component.