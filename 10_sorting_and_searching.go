package main

import (
	"fmt"
	"sort"
)

//------------------------------------------------------------------------------
// Problem 10.1

func SortedMerge(arr1, arr2 []int) []int {
	output := make([]int, len(arr1)+len(arr2))
	for i := range arr1 {
		output[i] = arr1[i]
	}

	i1, i2, j := len(arr1)-1, len(arr2)-1, len(arr1)+len(arr2)-1

	for j >= 0 {
		if i2 == -1 || (i1 != -1 && arr1[i1] >= arr2[i2]) {
			output[j] = arr1[i1]
			i1--
		} else {
			output[j] = arr2[i2]
			i2--
		}
		j--
	}

	return output
}

//------------------------------------------------------------------------------
// Problem 10.2

func sortString(str string) string {
	runes := []rune(str)
	sort.Slice(runes, func(i, j int) bool { return runes[i] < runes[j] })
	return string(runes)
}

func GroupAnagrams(strings []string) []string {
	copies := make([]string, len(strings))
	copy(copies, strings)

	sort.Slice(copies, func(i, j int) bool {
		s1 := sortString(copies[i])
		s2 := sortString(copies[j])
		if s1 == s2 {
			return copies[i] < copies[j]
		}
		return sortString(copies[i]) < sortString(copies[j])
	})

	return copies
}

func GroupAnagrams_Map(strings []string) []string {
	anagramMap := make(map[string][]string)
	for _, str := range strings {
		sorted := sortString(str)
		anagramMap[sorted] = append(anagramMap[sorted], str)
	}

	anagrams := make([]string, 0, len(strings))
	for _, strs := range anagramMap {
		anagrams = append(anagrams, strs...)
	}

	return anagrams
}

//------------------------------------------------------------------------------

func main() {
	strings := []string{"a", "abc", "abcd", "ab", "a", "bac", "dabc", "ba"}
	fmt.Println(GroupAnagrams_Map(strings))
}
