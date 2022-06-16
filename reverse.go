package slices

// Reverse reverses the slice
func Reverse[E any](s []E) {
	lastIdx := len(s) - 1
	lenDiv2 := len(s) / 2
	for idx := 0; idx < lenDiv2; idx++ {
		s[idx], s[lastIdx-idx] = s[lastIdx-idx], s[idx]
	}
}
