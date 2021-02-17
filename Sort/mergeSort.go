package mySort

func mergeSort(l, r int, A []int) {
	if l >= r {
		return
	}
	m := (l + r) >> 1
	mergeSort(l, m, A)
	mergeSort(m+1, r, A)
	merge(l, r, m, A)

}

func merge(l, r, m int, A []int) {

	tmp := make([]int, r-l+1)
	i, j := l, m+1
	k := 0
	for ; i <= m && j <= r; k++ {
		if A[i] < A[j] {
			tmp[k] = A[i]
			i++
		} else {
			tmp[k] = A[j]
			j++
		}
	}
	for ; i <= m; i++ {
		tmp[k] = A[i]
		k++

	}
	for ; j <= r; j++ {
		tmp[k] = A[j]
		k++
	}

	for k := 0; k < len(tmp); k++ {
		A[l+k] = tmp[k]
	}
}
