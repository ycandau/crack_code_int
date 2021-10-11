package main

import "fmt"

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

	// Work backwards which could allow in place copying
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

//------------------------------------------------------------------------------
// Problem 1.4

func IsPalindromePermutation(str string) bool {
	// Count the letter occurences
	counts := map[rune]int{}
	for _, ch := range str {
		counts[ch]++
	}

	// Odd length strings can have one odd count
	// Even length strings can have 0 odd counts
	// Equivalent to: 0 or 1 odd counts
	
	// Check the counts
	oddCount := 0
	for _, count := range counts {
		if (count & 1 == 1) {
			oddCount++
		}
		if oddCount > 1 {
			return false
		}
	}

	return true
}

//------------------------------------------------------------------------------
// Problem 1.4 with bit vector

func toggle(bitVec int, ch rune) int {
	i := ch - 'a'
	return bitVec ^ (1 << i)
}

func IsPalindromePermutation_BV(str string) bool {
	// We only need to know if counts are odd or not
	bitVect := 0
	for _, ch := range str {
		bitVect = toggle(bitVect, ch)
	}

	// Should have 0 or 1 bit set to 1
	return bitVect == 0 || bitVect & (bitVect - 1) == 0
}

//------------------------------------------------------------------------------
// Problem 1.5

func oneReplace(str1, str2 string) bool {
	count := 0
	for i := range str1 {
		if str1[i] != str2[i] {
			count++
		}
		if count > 1 {
			return false
		}
	}
	return true
}

func oneEdit(short, long string) bool {
	i := 0
	j := 0
	for i < len(short) {
		if short[i] == long[j] {
			i++
			j++
		} else {
			if j != i {
				return false
			}
			j++
		}
	}
	return true
}

func OneAway(str1, str2 string) bool {
	diff := len(str1) - len(str2)
	switch diff {
	case 0:
		return oneReplace(str1, str2)
	case -1:
		return oneEdit(str1, str2)
	case 1:
		return oneEdit(str2, str1)
	default:
		return false
	}
}

//------------------------------------------------------------------------------

func main() {
	fmt.Println("|", OneAway("alpha", "alpha"), "|")
	fmt.Println("|", OneAway("alpha", "al_ha"), "|")
	fmt.Println("|", OneAway("alpha", "al__a"), "|")

	fmt.Println("|", OneAway("alpha", "alph_a"), "|")
	fmt.Println("|", OneAway("alpha", "al_ph_a"), "|")
	fmt.Println("|", OneAway("a_lpha", "alpha"), "|")
	fmt.Println("|", OneAway("a_lp_ha", "alpha"), "|")
}
