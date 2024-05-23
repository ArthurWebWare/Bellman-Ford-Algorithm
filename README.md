# Ford's Algorithm Implementation in Go

Imagine you have a map with different cities connected by roads, each road having a certain distance. The Bellman–Ford algorithm is like a guide that helps you find the shortest path from one city to all other cities, even if some roads have negative lengths. It’s like a GPS for computers, useful for figuring out the quickest way to get from one point to another in a network. 

Sure, here is the translation of the explanation into English in markdown format for a README file:

This program implements Ford's algorithm for finding the shortest path in a weighted graph. Below is a detailed explanation of the code, along with an example of input data and the expected output.

## Example Input Data

```
Enter the number of vertices: 4

Enter the weighted graph:
0 10 0 30 100
0 0 50 0 0
0 0 0 20 10
0 0 0 0 60
0 0 0 0 0
```

## Explanation of Input Data

- The graph has 4 vertices (plus 1 added in the program for convenient indexing).
- The adjacency matrix for the graph, where `a[i][j]` represents the weight of the edge from vertex `i` to vertex `j`. If there is no edge, the weight is `0`.

### Adjacency Matrix Representation

```
    0   10   0   30  100
    0    0   50   0    0
    0    0    0   20   10
    0    0    0    0   60
    0    0    0    0    0
```

- Vertex 1 has edges to vertices 2, 4, and 5 with weights 10, 30, and 100 respectively.
- Vertex 2 has an edge to vertex 3 with weight 50.
- Vertex 3 has edges to vertices 4 and 5 with weights 20 and 10 respectively.
- Vertex 4 has an edge to vertex 5 with weight 60.

## Step-by-Step Program Output

1. **Input the number of vertices**:
    ```
    Enter the number of vertices: 4
    ```

2. **Input the weighted graph**:
    ```
    Enter the weighted graph:

    0 10 0 30 100
    0 0 50 0 0
    0 0 0 20 10
    0 0 0 0 60
    0 0 0 0 0
    ```

3. **Intermediate Results Output (Heights and Weights)**:
    ```
    __________________________________________________________
    |         |     |                    |                    |
    | (xi,xj) | Pij |     Hj - Hi        |     Hj - Hi        |
    |_________|_____|____________________|____________________|
    | ( 1, 2) |  10 | H2 - H1 =  10 = 10 | H2 - H1 =  10 = 10 |
    |_________|_____|____________________|____________________|
    | ( 1, 4) |  30 | H4 - H1 =  30 = 30 | H4 - H1 =  30 = 30 |
    |_________|_____|____________________|____________________|
    | ( 1, 5) | 100 | H5 - H1 = 100 = 100| H5 - H1 = 100 = 100|
    |_________|_____|____________________|____________________|
    | ( 2, 3) |  50 | H3 - H2 =  50 = 50 | H3 - H2 =  50 = 50 |
    |_________|_____|____________________|____________________|
    | ( 3, 4) |  20 | H4 - H3 =  20 = 20 | H4 - H3 =  20 = 20 |
    |_________|_____|____________________|____________________|
    | ( 3, 5) |  10 | H5 - H3 =  10 < 10 | H5 - H3 =  10 < 10 |
    |_________|_____|____________________|____________________|
    | ( 4, 5) |  60 | H5 - H4 =  60 = 60 | H5 - H4 =  60 = 60 |
    |_________|_____|____________________|____________________|
    ```

4. **Output the length of the shortest path**:
    ```
    The length of the shortest path = 60
    ```

5. **Output the shortest paths**:
    ```
    ***Shortest Paths***
    #Shortest path[ 1 ] : 5 4 1
    ```

## Conclusion

This example demonstrates how to input data into the program and what the output looks like for a weighted graph. The program builds the shortest paths using Ford's algorithm, prints a table of heights and weights for each edge, and shows the shortest paths from the starting vertex to all others.

## Source
https://www.geeksforgeeks.org/bellman-ford-algorithm-dp-23/