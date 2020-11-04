package array_hash_stringBuilder

import (
	"hash/maphash"
)

type HashTable struct {
	arr     []val
	hash    maphash.Hash
	factory int8
}

type val struct {
	exists bool
	key    string
	v      string
}

func NewHashTable(factory int8, size int) *HashTable {
	if factory <= 1 || factory > 10 {
		factory = 2
	}
	if size < 1 {
		size = 1
	}
	return &HashTable{
		arr:     make([]val, size),
		factory: factory,
	}
}

func (h *HashTable) Add(key, v string) error {
	i, err := h.getI(key)
	if err != nil {
		return err
	}
	if !h.arr[i].exists {
		h.arr[i] = val{
			exists: true,
			key:    key,
			v:      v,
		}
	} else {
		return h.reorder(key, v)
	}

	return nil
}

func (h *HashTable) Get(key string) (string, bool, error) {
	i, err := h.getI(key)
	if err != nil {
		return "", false, err
	}

	return h.arr[i].v, h.arr[i].exists, nil
}

func (h *HashTable) Delete(key string) error {
	i, err := h.getI(key)
	if err != nil {
		return err
	}
	h.arr[i].exists = false
	h.arr[i].key = ""
	h.arr[i].v = ""

	return nil
}

func (h *HashTable) reorder(lastKey, lastVal string) error {
	newH := NewHashTable(h.factory, len(h.arr)*int(h.factory))
	if err := newH.Add(lastKey, lastVal); err != nil {
		return err
	}
	for _, v := range h.arr {
		if v.exists {
			err := newH.Add(v.key, v.v)
			if err != nil {
				return err
			}
		}
	}
	h.arr = newH.arr

	return nil
}

func (h *HashTable) getI(key string) (int, error) {
	defer h.hash.Reset()

	if _, err := h.hash.WriteString(key); err != nil {
		return 0, err
	}

	return int(h.hash.Sum64() % uint64(len(h.arr))), nil
}
