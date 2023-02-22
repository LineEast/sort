package sort

// |			Time Complexity			| Scpace Complexity |
// |------------------------------------|-------------------|
// |	Best	|	Average	|	Worst	| 		Worst		|
// |	Ω(n)	|	Θ(n^2)	|	Θ(n^2)	|		O(1)		|

func Bubble[T any](arr []T, comp func(T, T) bool) {
	for i := 0; i < len(arr); i++ {
		for j := i + 1; j < len(arr); j++ {
			if comp(arr[i], arr[j]) {
				arr[i], arr[j] = arr[j], arr[i]
			}
		}
	}
}
