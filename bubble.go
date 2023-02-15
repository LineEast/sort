package sort

func Bubble[T any](arr []T, comp func(T, T) bool) {
	for i := 0; i < len(arr); i++ {
		for j := i + 1; j < len(arr); j++ {
			if comp(arr[i], arr[j]) {
				arr[i], arr[j] = arr[j], arr[i]
			}
		}
	}
}
