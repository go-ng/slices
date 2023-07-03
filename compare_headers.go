package slices

// CompareHeaders compares the headers (reflect.SliceHeader) of the given
// slices.
//
// Returns 0 if the headers contains absolutely equal values.
// Returns -1 value if the s0 header has a value lower than the s1's value.
// Returns 1 if the s0 header has a value higher than the s1's value.
//
// The values are compared in this order: Len, Cap, Data.
func CompareHeaders[T0, T1 any](s0 []T0, s1 []T1) int {
	h0 := Header(s0)
	h1 := Header(s1)

	if h0.Len != h1.Len {
		if h0.Len < h1.Len {
			return -1
		}
		return 1
	}

	if h0.Cap != h1.Cap {
		if h0.Cap < h1.Cap {
			return -1
		}
		return 1
	}

	if h0.Data != h1.Data {
		if h0.Data < h1.Data {
			return -1
		}
		return 1
	}

	return 0
}

// EqualHeaders returns true if the headers (reflect.SliceHeader) of given
// slices are equal.
func EqualHeaders[T0, T1 any](s0 []T0, s1 []T1) bool {
	return CompareHeaders(s0, s1) == 0
}
