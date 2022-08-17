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
	e = make([]int, N) 	//值
	l = make([]int, N) 	//指向结点左端点的指针
	r = make([]int, N) 	//指向结点右端点的指针
	idx int   			//指向当前操作数的指针
)

func new() {
	//左端点坐标为0，右端点为1
	r[0] = 1
	l[1] = 0
	idx = 2
}

//向第k个插入的数的右边插入一个数
func insert(k int, x int) {
	e[idx] = x
	l[idx] = k
	l[r[k]] = idx
	r[idx] = r[k]
	r[k] = idx
	idx++
}

//表示将第k个插入的数删除
func remove(k int) {
	r[l[k]] = r[k]
	l[r[k]] = l[k]
}

func main() {
	new()
	in := bufio.NewReader(os.Stdin)
	var M int
	fmt.Fscan(in, &M)
	for ; M > 0; M-- {
		var oper string
		fmt.Fscan(in, &oper)
		switch oper {
		case "L" :
			var x int
			fmt.Fscan(in, &x)
			insert(0, x)
		case "R" :
			var x int
			fmt.Fscan(in, &x)
			insert(l[1], x)
		case "D" :
			var k int
			fmt.Fscan(in, &k)
			remove(k+1)
		case "IL" :
			var k, x int
			fmt.Fscan(in, &k, &x)
			insert(l[k+1], x)
		case "IR" :
			var k, x int
			fmt.Fscan(in, &k, &x)
			insert(k+1, x)
		}
	}
	for i := r[0]; i != 1; i = r[i] {
		fmt.Printf("%d ", e[i])
	}
}