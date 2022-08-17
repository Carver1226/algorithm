package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"sort"
)

type Interval struct {
	l float64
	r float64
}

func main() {
	in := bufio.NewReader(os.Stdin)
	var n int
	fmt.Fscan(in, &n)
	is := make([]Interval, 0)
	for i := 1; i <= n; i++ {
		var l, r float64
		fmt.Fscan(in, &l)
	
		fmt.Fscan(in, &r)
		is = append(is, Interval{l, r})
	}
	res := merge(is)
	fmt.Print(len(res))
}

func merge(is []Interval) []Interval {
	result := make([]Interval, 0)
	sort.Slice(is, func(i, j int) bool {
		return is[i].l < is[j].l
	})
	st := -2e9
	ed := -2e9
	for _, interval := range is {
		if ed < interval.l {
			if st != -2e9 {
				result = append(result, Interval{st, ed})
			}
			st = interval.l
			ed = interval.r
		} else {
			ed = math.Max(ed, interval.r)
		}
	}
	if st != -2e9 {
		result = append(result, Interval{st, ed})
	}
	return result
}
