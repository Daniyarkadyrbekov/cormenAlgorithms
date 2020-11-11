package array_hash_stringBuilder

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestOneAway(t *testing.T) {
	require.True(t, awayDistance("", ""))
	require.True(t, awayDistance("", "q"))
	require.True(t, awayDistance("pale", "pale"))
	require.True(t, awayDistance("pales", "pale"))
	require.True(t, awayDistance("ple", "pale"))
	require.True(t, awayDistance("apale", "pale"))

	require.False(t, awayDistance("pa", "pale"))
	require.False(t, awayDistance("bake", "pale"))
	require.False(t, awayDistance("pabls", "pale"))
	require.False(t, awayDistance("pables", "pale"))
}
