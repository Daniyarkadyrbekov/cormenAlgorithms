package array_hash_stringBuilder

/*
One Away: There are three types of edits that can be performed on strings: insert a character, remove a character, or replace a character. Given two strings, write a function to check if they are one edit (or zero edits) away.
EXAMPLE
pale, ple -) true pales, pale - ) true pale, bale -) true pale, bae -) false
*/

func awayDistance(s1, s2 string) bool {
	distance := 1

	diffLen := len(s1) - len(s2)
	if diffLen > distance || diffLen < -distance {
		return false
	}

	if diffLen == 0 {
		c := 0
		for i := 0; i < len(s1); i++ {
			if s1[i] != s2[i] {
				c++
				if c > distance {
					return false
				}
			}
		}
		return true
	}
	if len(s1) > len(s2) {
		return checkInsert(s2, s1, distance)
	}
	return checkInsert(s1, s2, distance)
}

func checkInsert(s1, s2 string, distance int) bool {
	c := 0
	j := 0
	for i := 0; i < len(s1) && j < len(s2); i++ {
		if s1[i] != s2[j] {
			c++
			i--
			if c > distance {
				return false
			}
		}
		j++
	}
	return true
}
