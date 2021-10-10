package main

//------------------------------------------------------------------------------
// Problem 1.1

func AllUnique(str string) bool {
	counts := map[rune]bool{}
	for _, ch := range str {
		if _, has := counts[ch]; has {
			return false
		}
		counts[ch] = true
	}
	return true
}

func AllUniqueWithoutMap(str string) bool {
	counts := make([]bool, 256)
	for _, ch := range str {
		index := int(ch)
		if has := counts[index]; has {
			return false
		}
		counts[ch] = true
	}
	return true
}

//------------------------------------------------------------------------------
// Problem 1.2

func ArePermutations(str1, str2 string) bool {
	if len(str1) != len(str2) {
		return false
	}
	counts := map[rune]int{}
	for _, ch := range str1 {
		counts[ch]++
	}
	for _, ch := range str2 {
		val, found := counts[ch]
		if !found || val == 0 {
			return false
		}
		counts[ch]--
	}
	return true
}

//------------------------------------------------------------------------------
