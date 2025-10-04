package util

func Filter[T any](items []T, keep func(T) bool) []T {
	result := make([]T, 0)
	for _, item := range items {
		if keep(item) {
			result = append(result, item)
		}
	}
	return result
}
