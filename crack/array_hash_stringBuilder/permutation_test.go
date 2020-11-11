package array_hash_stringBuilder

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestPermutation(t *testing.T) {

	for _, fn := range []func(s1, s2 string) bool{
		checkPermutation1,
		checkPermutation2,
	} {
		require.True(t, fn("abc", "abc"))
		require.True(t, fn("abqwertyc", "abcytrewq"))
		require.True(t, fn("😅😅😘👌🏻", "😘😅👌🏻😅"))
		require.False(t, fn("😅😅😘👌🏻", "😘😅👌🏻😅👌🏻😅"))
		require.False(t, fn("abqwertyce", "abcytrewq"))
	}
}
