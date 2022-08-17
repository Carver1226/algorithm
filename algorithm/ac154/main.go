package main

import (
	"bufio"
	"fmt"
	"os"
)

const (
	N = 1000010
)

var (
	q = make([]int, N)
	t = -1
	h = 0
)

func main() {
	in := bufio.NewReader(os.Stdin)
	var n, k int
	fmt.Fscan(in, &n, &k)
	a := make([]int, N)
	for i := 0; i < n; i++ {
		fmt.Fscan(in, &a[i])
	}
	for i := 0; i < n; i++ {
		if h <= t && i - k >= q[h] {
			h++
		}
		for h <= t && a[q[t]] >= a[i] {
			t--
		}
		t++
		q[t] = i
		if i + 1 >= k {
			fmt.Printf("%d ", a[q[h]])
		}
	}
	fmt.Println()
	t = -1
	h = 0
	for i := 0; i < n; i++ {
		if h <= t && i - k >= q[h] {
			h++
		}
		for h <= t && a[q[t]] <= a[i] {
			t--
		}
		t++
		q[t] = i
		if i + 1 >= k {
			fmt.Printf("%d ", a[q[h]])
		}
	}
}