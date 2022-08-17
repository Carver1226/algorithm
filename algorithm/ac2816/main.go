package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	in := bufio.NewReader(os.Stdin)
	var n, m int
	fmt.Fscan(in, &n, &m)
	a := make([]int, n)
	b := make([]int, m)
	for i := 0; i < n; i++ {
		fmt.Fscan(in, &a[i])
	}
	for i := 0; i < m; i++ {
		fmt.Fscan(in, &b[i])
	}

	i, j := 0, 0
	for j < m && i < n {
		if a[i] == b[j] {
			i++
		}
		j++
	}
	if i == n {
		fmt.Print("Yes")
	} else {
		fmt.Print("No")
	}
}