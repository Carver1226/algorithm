package quickSort

func quickSort(arr []int, left int, right int) {
	if left >= right {
		return
	}
	x := arr[(left + right) / 2]
	l := left
	r := right
	for l < r {
		for arr[l] < x {
			l++
		}
		for arr[r] > x {
			r--
		}
		if l < r {
			swap(&arr[l], &arr[r])
		}
	}
	quickSort(arr, left, (left + right) / 2)
	quickSort(arr, (left + right) / 2 + 1, right)
}

func swap(a *int, b *int) {
	a, b = b, a
}