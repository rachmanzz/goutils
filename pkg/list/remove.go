package listutil

func RmItem[T any](arr *[]T, args ...int) {
	if arr == nil || len(*arr) == 0 {
		return
	}
	length := len(*arr)

	switch len(args) {
	case 2:
		start, end := args[0], args[1]

		// handle negatif
		if start < 0 {
			start = length + start
		}
		if end < 0 {
			end = length + end
		}

		// clamp
		if start < 0 {
			start = 0
		}
		if end > length {
			end = length
		}
		if start >= end {
			return
		}

		newArr := make([]T, 0, length-(end-start))
		newArr = append(newArr, (*arr)[:start]...)
		newArr = append(newArr, (*arr)[end:]...)
		*arr = newArr

	case 1:
		n := args[0]
		if n == 0 {
			return
		}
		if n > 0 {
			if n >= length {
				*arr = (*arr)[:0]
			} else {
				*arr = (*arr)[n:]
			}
		} else {
			n = length + n
			if n <= 0 {
				*arr = (*arr)[:0]
			} else {
				*arr = (*arr)[:n]
			}
		}
	}
}

func RmFunc[T any](arr *[]T, matchFn func(item T) bool) {
	if arr == nil || len(*arr) == 0 {
		return
	}

	newArr := make([]T, 0, len(*arr))
	for _, item := range *arr {
		if !matchFn(item) {
			newArr = append(newArr, item)
		}
	}
	*arr = newArr
}
