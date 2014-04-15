// qsort.go
package qsort

func Qsort(values []int) {
	partition(values, 0, len(values) - 1)
}

func partition(values []int, lh int, rh int) {
	if rh <= lh {
		return
	}

	l, r := lh, rh
	p := values[l]

	for ; l < r; {
		for ; values[r] >= p && l < r; r-- {
		}
		for ; values[l] < p && l < r; l++ {
		}

		if l < r {
			values[l], values[r] = values[r], values[l]

		}
	}
	values[l] = p;

	if l != r {
		panic("l not equal to r")
	}

	partition(values, lh, l - 1)
	partition(values, l + 1, rh)
}
