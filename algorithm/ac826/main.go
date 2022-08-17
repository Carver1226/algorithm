package main

import (
	"bufio"
	"fmt"
	"os"
)

const (
	N = 100010
)

type link struct {
	head int   //头结点，保存着下一个值的next指针
	e    []int //结点值
	ne   []int //next指针
	idx  int   //当前操作数的指针
}

func new() *link {
	return &link{
		head: -1,
		e:    make([]int, N),
		ne:   make([]int, N),
		idx:  0,
	}
}

//在头节点后面插入x
func (l *link) addToHead(x int) {
	//赋值
	l.e[l.idx] = x
	//当前结点指向头节点的next指针
	l.ne[l.idx] = l.head
	//头节点的next指针指向当前结点
	l.head = l.idx

	l.idx++
}

//在第k个插入的数后面插入x
func (l *link) add(k int, x int) {
	l.e[l.idx] = x
	l.ne[l.idx] = l.ne[k]
	l.ne[k] = l.idx
	l.idx++
}

//删除第k个插入的数后面的数
func (l *link) remove(k int) {
	l.ne[k] = l.ne[l.ne[k]]
}

func main() {
	in := bufio.NewReader(os.Stdin)
	l := new()
	var M int
	fmt.Fscan(in, &M)
	for ; M > 0; M-- {
		var oper string
		fmt.Fscan(in, &oper)
		switch oper {
		case "H":
			var x int
			fmt.Fscanf(in, "%d", &x)
			l.addToHead(x)
		case "D":
			var k int
			fmt.Fscanf(in, "%d", &k)
			if k == 0 {
				l.head = l.ne[l.head]
			} else {
				l.remove(k - 1)
			}
		case "I":
			var k, x int
			fmt.Fscanf(in, "%d %d", &k, &x)
			l.add(k - 1, x)
		}
	}

	for i := l.head; i != -1; i = l.ne[i] {
		fmt.Printf("%d ", l.e[i])
	}
}
