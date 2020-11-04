package array_hash_stringBuilder

import (
	"strconv"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestDataStructs(t *testing.T) {
	for _, info := range []struct {
		Name string
		Func func(*testing.T)
	}{
		{
			Name: "arrList",
			Func: testArrList,
		},
		{
			Name: "hashTable",
			Func: testHashTable,
		},
	} {
		if !t.Run(info.Name, info.Func) {
			return
		}
	}
}

func testArrList(t *testing.T) {
	customArr := NewArrList(0, 2)
	for i := -10; i <= 100; i++ {
		customArr.Add(i)
	}

	for i := 100; i >= -10; i-- {
		require.Equal(t, i, customArr.Pop())
	}
}

func testHashTable(t *testing.T) {
	desc := "i = %v"

	hash := NewHashTable(2, 0)
	const from, to = 1, 4
	for i := from; i <= to; i++ {
		kv := strconv.Itoa(i)
		require.NoError(t, hash.Add(kv, kv), desc, i)
	}

	for i := from; i <= to; i++ {
		if i%3 == 0 {
			kv := strconv.Itoa(i)
			require.NoError(t, hash.Delete(kv), desc, i)
		}

	}

	for i := from; i <= to; i++ {
		kv := strconv.Itoa(i)
		v, exists, err := hash.Get(kv)
		require.NoError(t, err)
		if i%3 == 0 {
			require.False(t, exists, desc, i)
			require.Equal(t, "", v, desc, i)
		} else {
			require.True(t, exists, desc, i)
			require.Equal(t, kv, v, desc, i)
		}
	}
}
