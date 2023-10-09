# Algorithms & Data Structures

## Sort
| Algorithm | Best Time | Avg Time | Worst Time | Worst Space | Description |
| --- | --- | ---| ---| --- | ---
| Heap Sort | O(n * log(n)) | O(n * log(n)) | O(n * log(n)) | O(1) | 
| Insertion Sort | O(n) | O(n^2) | O(n^2) | O(1) | 
| [Merge Sort](sort/merge.go) | O(n * log(n)) | O(n * log(n)) | O(n * log(n)) | O(n) | Uses divide and conquer strategy to recursively sort.
| Selection Sort | O(n^2) | O(n^2) | O(n^2) | O(1) | 
| Shell Sort | O(n) | O((n * log(n))^2)| O(n^(3/2))| O(1) |

## Topological Search
| Algorithm | Time Complexity | Description |
| --- | --- | ---
| Depth First Search | O(V + E) |
| Kahns| O(V + E) | 

## Minimum Spanning Tree
| Algorithm | Time Complexity | Description |
| --- | --- | ---
| Kruskals | | 
| Prims | | 


## Graph Search
| Algorithm | Time Complexity | Description |
| --- | --- | ---
| Breadth First Search | O(V + E) | Iterate through nodes using a queue.
| Depth First Search | O(V + E ) | Iterate through nodes using a stack.

## Minimum Spanning Tree

| Algorithm | Time Complexity | Description |
| --- | --- | ---
| Kruskal's | O(E * log(E)) | Better for sparse graphs.
| Prims's | O(V * log(V)) | Runs faster for dense graphs.

## Shortest Path

| Algorithm | Time Complexity | Description |
| --- | --- | ---
| Dijkstra’s | O(E*log(V)) | Shortest path from one vertex to all vertices. Requires non-negative edge weights.
| Bellman–Ford | O(V*E) | Finds the shortest path from one vertex to all vertices.
| Floyd Warshall | O(V^3) | Shortest path between all vertices. Can handle negative values.

## Sliding Window
| Algorithm | Description |
| --- | ---
| Dynamic | Finds smallest window greater than or equal to a given sum.
| Fixed | Finds the max sum of a window with a given size.
| Kadanes | Finds the greatest contiguous sum.

## Other
- [x] Linked List
- [x] Queue
- [x] Stack
- [ ] Union–Find Algorithm
- [ ] **Priority Queue**
- [ ] **K-ary Heap**

## Search
- [x] Binary Search

## Graph
- [x] Max Heap Tree
- [ ] Kruskal’s Algorithm
- [ ] Kahn’s Topological Sort Algorithm

### TODO
- quick sort