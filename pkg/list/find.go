package listutil

func Find[T any](items []T, predicate func(T) bool) (T, bool) {
	var zero T
	for _, item := range items {
		if predicate(item) {
			return item, true
		}
	}
	return zero, false
}
