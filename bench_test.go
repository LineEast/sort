package sort

import (
	"testing"
)

func BenchmarkMergeSort(b *testing.B) {
	arr := []int{12, 6, 5, 3, 1, 8, 9, 10, 7, 2, 4, 11}
	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		MergeSort(arr, func(a, b int) bool {
			return a < b
		})
	}
}

func BenchmarkQsort(b *testing.B) {
	arr := []int{12, 6, 5, 3, 1, 8, 9, 10, 7, 2, 4, 11}
	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		Qsort(arr, func(a, b int) bool {
			return a < b
		})
	}
}
