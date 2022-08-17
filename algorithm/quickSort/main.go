package main

import (
	"fmt"
)

func quickSort(arr []int, left int, right int) {
	if left >= right {
		return
	}
	x := arr[(left + right) / 2]
	l := left - 1
	r := right + 1
	for l < r {
		l++
		for arr[l] < x {
			l++
		}
		r--
		for arr[r] > x {
			r--
		}
		if l < r {
			arr[l], arr[r] = arr[r], arr[l]
		}
	}
	quickSort(arr, left, r)
	quickSort(arr, r + 1, right)
}

func qSort(arr []int) {
	quickSort(arr, 0, len(arr) - 1)
}

func main() {
	arr := []int{2, 2}
	qSort(arr)
	fmt.Printf("%v\n", arr)
}
