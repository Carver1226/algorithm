package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	in := bufio.NewReader(os.Stdin)
	var n, m, x int
	fmt.Fscan(in, &n, &m, &x)
	A := make([]int, n)
	B := make([]int, m)

	for i := 0; i < n; i++ {
		fmt.Fscan(in, &A[i])
	}
	for i := 0; i < m; i++ {
		fmt.Fscan(in, &B[i])
	}

	for i, j := 0, m-1; i < n; i++ {
		for j >= 0 && A[i] + B[j] > x {
			j--
		}
		if j >= 0 && A[i] + B[j] == x {
			fmt.Printf("%d %d", i, j)
		}
	}
}