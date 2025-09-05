package listutil

import (
	"cmp"
	"slices"
	"strings"
)

func Order[T cmp.Ordered](items []T, direction ...string) []T {
	if len(items) == 0 {
		return items
	}

	copied := make([]T, len(items))
	copy(copied, items)

	dir := "asc"
	if len(direction) > 0 && strings.ToLower(direction[0]) == "desc" {
		dir = "desc"
	}

	if dir == "asc" {
		slices.Sort(copied)
	} else {
		slices.Sort(copied)
		reverse(copied)
	}

	return copied
}

func reverse[T any](s []T) {
	for i, j := 0, len(s)-1; i < j; i, j = i+1, j-1 {
		s[i], s[j] = s[j], s[i]
	}
}

func OrderByFunc[T any](items []T, less func(a, b T) bool) []T {
	copied := make([]T, len(items))
	copy(copied, items)

	slices.SortFunc(copied, func(a, b T) int {
		if less(a, b) {
			return -1
		}
		if less(b, a) {
			return 1
		}
		return 0
	})

	return copied
}
