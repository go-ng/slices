package slices

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestCompareHeaders(t *testing.T) {
	b := []int{1, 2, 3}

	c := make([]int, len(b))
	copy(c, b)
	require.Equal(t, b, c)
	require.True(t, CompareHeaders(b, c) != 0)

	d := b
	require.Equal(t, b, d)
	require.True(t, CompareHeaders(b, d) == 0)
}
