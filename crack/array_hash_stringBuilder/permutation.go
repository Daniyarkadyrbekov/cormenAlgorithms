package array_hash_stringBuilder

import "sort"

//Given two strings, write a method to decide if one is a permutation of the other

type sortRunes []rune

func (s sortRunes) Less(i, j int) bool {
	return s[i] < s[j]
}

func (s sortRunes) Swap(i, j int) {
	s[i], s[j] = s[j], s[i]
}

func (s sortRunes) Len() int {
	return len(s)
}

func SortString(s string) string {
	r := []rune(s)
	sort.Sort(sortRunes(r))
	return string(r)
}

func checkPermutation1(s1, s2 string) bool {
	s1 = SortString(s1)
	s2 = SortString(s2)

	return s1 == s2
}

func checkPermutation2(s1, s2 string) bool {
	chars := make(map[int32]int)
	for _, ch := range s1 {
		chars[ch] += 1
	}
	for _, ch := range s2 {
		if chars[ch] == 0 {
			return false
		}
		chars[ch] -= 1
	}

	for _, v := range chars {
		if v != 0 {
			return false
		}
	}

	return true
}
