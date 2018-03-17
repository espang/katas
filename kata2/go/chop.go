package chop


// chop searches value in sorted. Returns a index
// of value when it is in sorted and -1 otherwise.
func chop(value int, sorted []int) int {
	if len(sorted) == 0 {
		return -1
	}

	idx := len(sorted)/2
	if pivot := sorted[idx]; pivot == value {
		return idx
	} else if pivot > value {
		return chop(value, sorted[:idx])
	} else {
		return chop(value, sorted[idx+1:])
	}
}