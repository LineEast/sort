package sort

// |				Time Complexity			 		| Scpace Complexity |
// |------------------------------------------------|-------------------|
// |	Best		|	Average		|		Worst	| 		Worst		|
// |	Ω(n log(n))	|	Θ(n log(n))	|	O(n log(n))	|		O(n)		|

func MergeSort[T any](items []T, comp func(a, b T) bool) []T {
	if len(items) < 2 {
		return items
	}

	first := MergeSort(items[:len(items)/2], comp)
	second := MergeSort(items[len(items)/2:], comp)

	return merge(first, second, comp)
}

func merge[T any](a, b []T, comp func(a, b T) bool) []T {
	final := make([]T, 0, len(a)+len(b))

	i, j := 0, 0
	for i < len(a) && j < len(b) {
		if comp(a[i], b[j]) {
			final = append(final, a[i])
			i++
		} else {
			final = append(final, b[j])
			j++
		}
	}

	for ; i < len(a); i++ {
		final = append(final, a[i])
	}

	for ; j < len(b); j++ {
		final = append(final, b[j])
	}

	return final
}
