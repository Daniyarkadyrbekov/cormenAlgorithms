package array_hash_stringBuilder

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestPalindromePermutation(t *testing.T) {
	require.True(t, palindromePermutation(""))
	require.True(t, palindromePermutation("nadomecemodan"))
	require.True(t, palindromePermutation("nadomed  emo can"))
	require.True(t, palindromePermutation("tact coa"))

	require.False(t, palindromePermutation("nadomechemodan"))
}
