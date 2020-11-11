package array_hash_stringBuilder

import "strconv"

func compress(s string) string {

	if len(s) <= 1 {
		return s
	}

	res := make([]byte, 0)

	c := 1
	for i := 1; i <= len(s); i++ {
		if i < len(s) && s[i] == s[i-1] {
			c++
		} else {
			res = append(res, s[i-1])
			res = append(res, []byte(strconv.Itoa(c))...)
			c = 1
		}
	}

	if len(res) < len(s) {
		return string(res)
	}

	return s
}
