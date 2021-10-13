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
// Problem 10.3

func RotatedSearch(numbers []int, val int) (int, bool) {
	low, vlow := 0, numbers[0]
	high, vhigh := len(numbers)-1, numbers[len(numbers)-1]

	for low <= high {
		mid := (low + high) >> 1
		vmid := numbers[mid]

		if vmid == val {
			return mid, true
		}

		if vmid < vhigh {
			if vmid < val && val <= vhigh {
				low, vlow = mid+1, numbers[mid+1]
			} else {
				high, vhigh = mid-1, numbers[mid-1]
			}
		} else {
			if vlow <= val && val < vmid {
				high, vhigh = mid-1, numbers[mid-1]
			} else {
				low, vlow = mid+1, numbers[mid+1]
			}
		}
	}

	return 0, false
}

//------------------------------------------------------------------------------

func main() {
	strings := []string{"a", "abc", "abcd", "ab", "a", "bac", "dabc", "ba"}
	fmt.Println(GroupAnagrams_Map(strings))
}
