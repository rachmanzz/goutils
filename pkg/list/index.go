package listutil

func IndexOf[T comparable](items []T, target T) int {
	for i, item := range items {
		if item == target {
			return i
		}
	}
	return -1
}

func IndexBy[T any](items []T, predicate func(T) bool) int {
	for i, item := range items {
		if predicate(item) {
			return i
		}
	}
	return -1
}
