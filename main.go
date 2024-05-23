package main

import (
	"fmt"
)

const max = 1000

var (
	adjacencyMatrix       [100][100]int
	vertices, k, num, loc int
	h, h2, b, c           [100]int
	elem                  = 1001
	q                     = 1
)

func calculateMinimumPath(j2 int) {
	var i2 int

	if elem < 1000 {
		for i2 = 0; i2 < num; i2++ {
			if c[i2] == elem {
				loc = i2
			}
		}
		for i2 = num - 1; i2 >= loc; i2-- {
			b[k] = c[i2]
			k++
		}
		elem = 1001
	}

	b[k] = j2
	k++

	if j2 == 1 {
		fmt.Printf("\n#Минимальный путь[ %d ] : ", q)
		q++
		for i2 = k - 1; i2 >= 0; i2-- {
			fmt.Printf("%3d", b[i2])
		}

		num = k
		for i2 = k - 1; i2 >= 0; i2-- {
			c[k-i2-1] = b[i2]
		}
		k = 0
	}

	for i2 = vertices - 1; i2 >= 1; i2-- {
		if adjacencyMatrix[i2][j2] > 0 {
			if h2[j2]-h2[i2] == adjacencyMatrix[i2][j2] {
				calculateMinimumPath(i2)
				elem = j2
			}
		}
	}
}

func outputAlgorithmResult(a [100][100]int, n int) {
	fmt.Println("\n __________________________________________________________")
	fmt.Println("|         |     |                    |                    |")
	fmt.Println("| (xi,xj) | Pij |     Hj - Hi        |     Hj - Hi        |")
	fmt.Println("|_________|_____|____________________|____________________|")

	for i := 1; i < n; i++ {
		for j := 1; j < n; j++ {
			if a[i][j] > 0 {
				fmt.Printf("\n| (%2d,%2d) | %3d |", i, j, a[i][j])
				if h[j]-h[i] == a[i][j] {
					fmt.Printf(" H%d - H%d = %3d = %2d |", j, i, a[i][j], a[i][j])
				}
				if h[j]-h[i] < a[i][j] {
					fmt.Printf(" H%d - H%d = %3d < %2d |", j, i, h[j]-h[i], a[i][j])
				}
				if h[j]-h[i] > a[i][j] {
					if h[j]-h[i] > 500 {
						fmt.Printf(" H%d - H%d = inf > %2d |", j, i, a[i][j])
					} else {
						fmt.Printf(" H%d - H%d = %3d > %2d |", j, i, h[j]-h[i], a[i][j])
					}
					h[j] = h[i] + a[i][j]
				}
				if h2[j]-h2[i] == a[i][j] {
					fmt.Printf(" H%d - H%d = %3d = %2d |", j, i, a[i][j], a[i][j])
				}
				if h2[j]-h2[i] < a[i][j] {
					fmt.Printf(" H%d - H%d = %3d < %2d |", j, i, h2[j]-h2[i], a[i][j])
				}
				fmt.Println("\n|_________|_____|____________________|____________________|")
			}
		}
	}
}

func main() {
	fmt.Print("\nВведите количество вершин : ")
	fmt.Scan(&vertices)
	vertices++

	fmt.Println("\nВведите взвешанный граф : \n")

	for i := 1; i < vertices; i++ {
		for j := 1; j < vertices; j++ {
			fmt.Scan(&adjacencyMatrix[i][j])
		}
	}

	h[1] = 0
	for i := 2; i < vertices; i++ {
		h[i] = max
	}

	h2[1] = 0
	for i := 2; i < vertices; i++ {
		h2[i] = max
	}

	for i := 1; i < vertices; i++ {
		for j := 1; j < vertices; j++ {
			if adjacencyMatrix[i][j] > 0 {
				if h2[j]-h2[i] > adjacencyMatrix[i][j] {
					h2[j] = h2[i] + adjacencyMatrix[i][j]
				}
			}
		}
	}

	outputAlgorithmResult(adjacencyMatrix, vertices)

	fmt.Printf("\n\n\n\tДлина минимального пути = %d", h2[vertices-1])
	fmt.Println("\n\n***Минимальные пути***")
	calculateMinimumPath(vertices - 1)
}
