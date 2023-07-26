# Kahn's Topological Sort Algorithm

# Algorithm Overview

1. Count the number of dependencies for each vertex.
2. Add vertices with no dependencies to a queue.
3. Pop vertices from queue.
3. Decrement dependencies from each neighboring vertex.
4. Repeat steps 2 - 4 until all vertices have been processed.
