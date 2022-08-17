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
	q = make([]int, N)
	tt = -1 //队尾
	hh = 0  //队头
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
			q[tt] = x
		case "pop":
			hh++
		case "empty":
			if tt >= hh {
				fmt.Println("NO")
			} else {
				fmt.Println("YES")
			}
		case "query":
			fmt.Println(q[hh])
		}
	}
}