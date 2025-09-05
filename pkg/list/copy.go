package listutil

func Copy[T any](arr []T, idx ...int) []T {
	if arr == nil {
		return nil
	}

	n := len(arr)
	if n == 0 {
		return nil
	}

	switch len(idx) {
	case 0:
		// Salin semua
		copied := make([]T, n)
		copy(copied, arr)
		return copied

	case 1:
		i := idx[0]
		if i >= 0 {
			if i > n {
				i = n
			}
			copied := make([]T, i)
			copy(copied, arr[:i])
			return copied
		} else {
			i = n + i
			if i < 0 {
				i = 0
			}
			copied := make([]T, n-i)
			copy(copied, arr[i:])
			return copied
		}

	case 2:
		start, end := idx[0], idx[1]
		if start < 0 {
			start = n + start
		}
		if end < 0 {
			end = n + end
		}
		if start < 0 {
			start = 0
		}
		if end > n {
			end = n
		}
		if start >= end {
			return nil
		}
		copied := make([]T, end-start)
		copy(copied, arr[start:end])
		return copied

	default:
		return Copy(arr, idx[0], idx[1])
	}
}
