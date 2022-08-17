package main

import (
	"bufio"
	"fmt"
	"os"
)

const (
	N = 100010
)

var (
	stk = make([]int, N)
	t = 0
)

func main() {
	in := bufio.NewReader(os.Stdin)
	var n int
	fmt.Fscan(in, &n)
	for i := 1; i <= n; i++ {
		var x int
		fmt.Fscan(in, &x)
		for t > 0 && stk[t] >= x {
			t--
		}
		if t == 0 {
			fmt.Print("-1 ")
		} else {
			fmt.Printf("%d ", stk[t])
		}
		t++
		stk[t] = x
	}
}