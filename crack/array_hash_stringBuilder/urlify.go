package array_hash_stringBuilder

/*
URLify: Write a method to replace all spaces in a string with '%2e: You may assume that the string has sufficient space at the end to hold the additional characters, and that you are given the "true" length of the string. (Note: if implementing in Java, please use a character array so that you can perform this operation in place.)
EXAMPLE
Input: "Mr John Smith JJ, 13 Output: "Mr%2eJohn%2eSmith"
*/

func urlify1(str string) string {
	c := 0
	tmpCounter := 0

	for _, ch := range str {
		if ch == ' ' {
			tmpCounter++
		} else {
			c += tmpCounter
			tmpCounter = 0
		}
	}

	strByte := []byte(str)
	right, left := len(str)-1-2*c, len(str)-1

	for {
		if left <= right || right < 0 {
			return string(strByte)
		}

		if strByte[right] != ' ' {
			strByte[left] = strByte[right]
		} else {
			strByte[left] = '0'
			strByte[left-1] = '2'
			strByte[left-2] = '%'
			left -= 2
		}
		left--
		right--
	}
}
