package main

import "fmt"

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
			arr[l], arr[r] = swap(arr[l], arr[r])
		}
	}
	quickSort(arr, left, r)
	quickSort(arr, r + 1, right)
}

func swap(a int, b int) (int, int){
	return b, a
}

func main() {
	arr := []int{2, 5, 6, 1, 4, 8}
	quickSort(arr, 0, 5)
	fmt.Printf("%v", arr)
}
