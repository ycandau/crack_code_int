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
// Problem 1.3

func countChar(str string, match rune) int {
	count := 0
	for _, ch := range str {
		if ch == match {
			count++
		}
	}
	return count
}

func ToURL(str string) string {
	count := countChar(str, ' ')
	output := make([]byte, len(str)+2*count)

	i := len(str) - 1
	j := len(output) - 1

	for ; i >= 0; i-- {
		ch := str[i]
		if ch == ' ' {
			output[j-2] = '%'
			output[j-1] = '2'
			output[j] = '0'
			j -= 3
		} else {
			output[j] = ch
			j--
		}
	}
	return string(output)
}
