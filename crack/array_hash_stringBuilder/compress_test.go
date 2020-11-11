package array_hash_stringBuilder

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestCompress(t *testing.T) {
	require.Equal(t, "", compress(""))
	require.Equal(t, "t", compress("t"))
	require.Equal(t, "tt", compress("tt"))
	require.Equal(t, "tm", compress("tm"))
	require.Equal(t, "tmp", compress("tmp"))
	require.Equal(t, "t1m1p5", compress("tmppppp"))
	require.Equal(t, "t1m4p5", compress("tmmmmppppp"))
	require.Equal(t, "t4m4p5", compress("ttttmmmmppppp"))
	require.Equal(t, "t4", compress("tttt"))
}
