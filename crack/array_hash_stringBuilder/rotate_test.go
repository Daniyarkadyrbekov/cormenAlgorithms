package array_hash_stringBuilder

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestRotate(t *testing.T) {
	require.Len(t, rotate(nil), 0)
	require.Len(t, rotate([][]int{}), 0)
	require.Len(t, rotate([][]int{{1, 2}, {1}}), 0)
	require.Equal(t, [][]int{{1}}, rotate([][]int{{1}}))
	require.Equal(t, [][]int{{3, 1}, {4, 2}}, rotate([][]int{{1, 2}, {3, 4}}))
	require.Equal(t, [][]int{{7, 4, 1}, {8, 5, 2}, {9, 6, 3}}, rotate([][]int{{1, 2, 3}, {4, 5, 6}, {7, 8, 9}}))
}
