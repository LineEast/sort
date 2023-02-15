package sort

func Qsort[T any](a []T, com func(T, T) bool) {
	if len(a) < 2 {
		return
	}

	left, right := 0, len(a)-1

	pivot := len(a) / 2

	a[pivot], a[right] = a[right], a[pivot]

	for i := range a {
		if com(a[i], a[right]) {
			a[left], a[i] = a[i], a[left]
			left++
		}
	}

	a[left], a[right] = a[right], a[left]

	Qsort(a[:left], com)
	Qsort(a[left+1:], com)
}
