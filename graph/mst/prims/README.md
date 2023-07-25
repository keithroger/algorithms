# Prim's Algorithm

## Algorithm Overview

1. Add edges from first vertex to a minimum priority queue that sorts using the edge weights.

2. Until the queue is empty, pop off vertices. Every time a new edge is popped, the vertex being pointed can be added to he minimum spanning tree.
After adding a new vertex to the spanning tree, add its edges to the minimum priority queue.
Skip edges that point to already visited vertices.
