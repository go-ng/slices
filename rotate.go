package slices

// Rotate rotates the slice by shift items.
func Rotate[E any](s []E, shift int) {
	rotateThreeReversals(s, shift)
}

func rotateThreeReversals[E any](s []E, shift int) {
	length := len(s)
	if length == 0 {
		return
	}
	shift %= length
	if shift < 0 {
		shift += length
	}
	if shift == 0 {
		return
	}
	Reverse(s)
	Reverse(s[:shift])
	Reverse(s[shift:])
}
