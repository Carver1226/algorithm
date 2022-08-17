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
	tt = 0
)

func main() {
	in := bufio.NewReader(os.Stdin)
	var M int
	fmt.Fscan(in, &M)
	for ; M > 0; M-- {
		var oper string
		fmt.Fscan(in, &oper)
		switch oper {
		case "push":
			var x int
			fmt.Fscan(in, &x)
			tt++
			stk[tt] = x
		case "pop":
			tt--
		case "empty":
			if tt == 0 {
				fmt.Println("YES")
			} else {
				fmt.Println("NO")
			}
		case "query":
			fmt.Println(stk[tt])
		}
	}
}