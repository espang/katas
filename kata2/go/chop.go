package chop

// chopRec searches value in sorted. Returns a index
// of value when it is in sorted and -1 otherwise.
func chopRec(value int, sorted []int) int {
	if len(sorted) == 0 {
		return -1
	}

	idx := len(sorted) / 2
	if pivot := sorted[idx]; pivot == value {
		return idx
	} else if pivot > value {
		return chopRec(value, sorted[:idx])
	}

	res := chopRec(value, sorted[idx+1:])
	if res == -1 {
		return res
	}
	return idx + 1 + res
}

func chopIter(value int, sorted []int) int {
	left, right := 0, len(sorted)-1
	for {
		m := (left + right) / 2

		if left > right {
			return -1
		}

		if sorted[m] > value {
			right = m - 1
		} else if sorted[m] < value {
			left = m + 1
		} else {
			return m
		}
	}
}
