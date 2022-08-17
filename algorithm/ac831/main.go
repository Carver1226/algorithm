package main

import (
	"bufio"
	"fmt"
	"os"
)

const (
	N = 100010
	M = 1000010
)

var (
	next = make([]int, N)
	P = make([]byte, N)
	S = make([]byte, M)
)

func main() {
	in := bufio.NewReader(os.Stdin)
	var n, m int
	//s[]是长文本，p[]是模式串，n是s的长度，m是p的长度
	fmt.Fscan(in, &n, &P, &m, &S)
	//将P和S后移一位
	P, S = append([]byte{' '}, P...), append([]byte{' '}, S...)
	//P和S下标从1开始, 求模式串的next数组
	for i, j := 2, 0; i <= n; i++ {
		//如果前缀和后缀相等就最长公共前后缀的长度就++，不相等就回退，直到相等或者为0
		for j != 0 && P[i] != P[j+1] {
			j = next[j]
		}
		if P[i] == P[j+1] {
			j++
		}
		next[i] = j
	}
	for i, j := 1, 0; i <= m; i++ {
		//如果s串和p串匹配则继续，不匹配则回退
		for j != 0 && S[i] != P[j+1] {
			j = next[j]
		}
		if S[i] == P[j+1] {
			j++
		}
		if j == n {
			fmt.Printf("%d ", i - n)
			j = next[j]
		}
	}
}