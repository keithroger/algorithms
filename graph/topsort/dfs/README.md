# Topological Sort Using Depth First Search

## Algorithm Overview


1. Pick a node and use a recursive post-order depth first search to visit nodes.
2. As vertices are processed they can be added to the topological ordering in reverse.
3. Continue using dfs on unvisited nodes until they have been processed.