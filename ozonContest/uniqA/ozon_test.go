package main

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestFindUniq(t *testing.T) {
	for _, data := range []struct {
		arr []int
		exp int
	}{
		{arr: []int{}, exp: 0},
		{arr: []int{1, 2, 3, 1, 2}, exp: 3},
		{arr: []int{1, 2, 2, 1, 2, 3, 2}, exp: 3},
		{arr: []int{1, 2, 2, 1, 2, 3, 2}, exp: 3},
	} {
		require.Equal(t, data.exp, findUniq(data.arr))
	}
}
