package array_hash_stringBuilder

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestUrlify(t *testing.T) {
	require.Equal(t, "", urlify1(""))
	require.Equal(t, "Mr%20John%20Smith", urlify1("Mr John Smith    "))
	require.Equal(t, "MrJohnSmith", urlify1("MrJohnSmith"))
	require.Equal(t, "Mr%20%20%20John%20Smith", urlify1("Mr   John Smith        "))
}
