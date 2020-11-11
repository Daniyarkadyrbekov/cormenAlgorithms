package array_hash_stringBuilder

/*
Palindrome Permutation: Given a string, write a function to check if it is a permutation of a palindrome. A palindrome is a word or phrase that is the same forwards and backwards. A permutation is a rearrangement of letters. The palindrome does not need to be limited to just dictionary words.
EXAMPLE
Input: Tact Coa
Output: True (permutations:"taco cat'; "atco cta '; etc.)
*/

func palindromePermutation(str string) bool {
	spaceCount := 0
	chars := make(map[int32]int)
	for _, c := range str {
		if c == ' ' {
			spaceCount++
			continue
		}
		chars[c]++
	}

	unPairedCount := 0
	for _, v := range chars {
		unPairedCount += v % 2
	}

	return unPairedCount == (len(str)-spaceCount)%2
}
