package zen

func Range(from, to int) []int {
	a := make([]int, to-from+1)
	for i := 0; i <= to-from; i++ {
		a[i] = i + from
	}
	return a
}

func In(val any, slice []any) bool {
	for _, v := range slice {
		if v == val {
			return true
		}
	}
	return false
}
