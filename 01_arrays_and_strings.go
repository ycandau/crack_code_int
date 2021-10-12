package main

import (
	"strconv"
	"strings"
)

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
		if count&1 == 1 {
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
	return bitVect == 0 || bitVect&(bitVect-1) == 0
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
// Problem 1.6

func StringCompression(str string) string {
	output := ""
	count := 0

	for i := 0; i < len(str); i++ {
		ch := str[i]
		count++
		if i+1 == len(str) || ch != str[i+1] {
			output += strconv.Itoa(count)
			output += string(ch)
			count = 0
		}
	}

	if len(str) <= len(output) {
		return str
	}

	return output
}

func StringCompression_SB(str string) string {
	var sb strings.Builder

	for count, i := 0, 0; i < len(str); i++ {
		ch := str[i]
		count++
		if i+1 == len(str) || ch != str[i+1] {
			sb.WriteString(strconv.Itoa(count))
			sb.WriteByte(ch)
			count = 0
		}
	}

	if len(str) <= sb.Len() {
		return str
	}

	return sb.String()
}

//------------------------------------------------------------------------------
// Problem 1.7

type matrix [][]int

func newMatrix(dx, dy int) matrix {
	matr := make([][]int, dy)
	for i := range matr {
		matr[i] = make([]int, dx)
	}
	return matr
}

func RotateMatrix(matr matrix) matrix {
	l := len(matr)

	for x := 0; x < (l+1)>>1; x++ {
		for y := 0; y < l>>1; y++ {
			x2 := l - x - 1
			y2 := l - y - 1

			tmp := matr[y][x]
			matr[y][x] = matr[x][y2]
			matr[x][y2] = matr[y2][x2]
			matr[y2][x2] = matr[x2][y]
			matr[x2][y] = tmp
		}
	}

	return matr
}

//------------------------------------------------------------------------------
// Problem 1.8

func ZeroMatrix(matr matrix) matrix {
	dx, dy := len(matr[0]), len(matr)
	cols := make([]bool, dx)
	rows := make([]bool, dy)

	for x := 0; x < dx; x++ {
		for y := 0; y < dy; y++ {
			if matr[y][x] == 0 {
				cols[x] = true
				rows[y] = true
			}
		}
	}

	for x := 0; x < dx; x++ {
		for y := 0; y < dy; y++ {
			if cols[x] || rows[y] {
				matr[y][x] = 0
			}
		}
	}

	return matr
}

func ZeroMatrix_IP(matr matrix) matrix {

	// Store zero presence in first row
	rowHasZero := false
	for _, n := range matr[0] {
		if n == 0 {
			rowHasZero = true
			break
		}
	}

	// Store zero presence in first col
	colHasZero := false
	for _, row := range matr {
		if row[0] == 0 {
			colHasZero = true
			break
		}
	}

	// Set zeros in first column and row
	dx, dy := len(matr[0]), len(matr)
	for x := 1; x < dx; x++ {
		for y := 1; y < dy; y++ {
			if matr[y][x] == 0 {
				matr[0][x] = 0
				matr[y][0] = 0
			}
		}
	}

	// Propagate zeros from first column and row
	for x := 1; x < dx; x++ {
		for y := 1; y < dy; y++ {
			if matr[0][x] == 0 || matr[y][0] == 0 {
				matr[y][x] = 0
			}
		}
	}

	// Set first row to zero if necessary
	for x := range matr[0] {
		if rowHasZero {
			matr[0][x] = 0
		}
	}

	// Set first column to zero if necessary
	for y := range matr {
		if colHasZero {
			matr[y][0] = 0
		}
	}

	return matr
}

//------------------------------------------------------------------------------
// Problem 1.9

func StringRotation(str, rot string) bool {
	if len(str) != len(rot) {
		return false
	}
	doubled := str + str
	return strings.Contains(doubled, rot)
}
