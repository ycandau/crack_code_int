package main

import (
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

func rotSearch(numbers []int, val, low, high int) (int, bool) {

	if low > high {
		return 0, false
	}

	mid := (low + high) >> 1
	if numbers[mid] == val {
		return mid, true
	}

	vmid := numbers[mid]
	if vmid > numbers[low] { // left side is increasing
		if numbers[low] <= val && val < vmid {
			return rotSearch(numbers, val, low, mid-1)
		} else {
			return rotSearch(numbers, val, mid+1, high)
		}
	} else if vmid < numbers[low] { // right side is increasing
		if vmid < val && val <= numbers[high] {
			return rotSearch(numbers, val, mid+1, high)
		} else {
			return rotSearch(numbers, val, low, mid-1)
		}
	} else {
		if vmid != numbers[high] {
			return rotSearch(numbers, val, mid+1, high)
		} else {
			if ind, found := rotSearch(numbers, val, low, mid-1); found {
				return ind, true
			}
			if ind, found := rotSearch(numbers, val, mid+1, high); found {
				return ind, true
			}
			return 0, false
		}
	}
}

func RotatedSearch(numbers []int, val int) (int, bool) {
	return rotSearch(numbers, val, 0, len(numbers)-1)
}

//------------------------------------------------------------------------------
// Problem 10.4

func get(numbers []int, ind int) int {
	if ind >= len(numbers) {
		return -1
	}
	return numbers[ind]
}

func ExpSearch(numbers []int, val int) int {
	low := 0
	high := 1
	vhigh := get(numbers, high)

	for vhigh != -1 && vhigh < val {
		low = high
		high = high << 1
		vhigh = get(numbers, high)
	}

	for low <= high {
		mid := (low + high) >> 1
		vmid := get(numbers, mid)
		if val < vmid || vmid == -1 {
			high = mid - 1
		} else if val > vmid {
			low = mid + 1
		} else {
			return mid
		}
	}
	return -1
}

//------------------------------------------------------------------------------
// Problem 10.5

func SparseSearch(arr []string, val string) int {
	low := 0
	high := len(arr) - 1

	for low <= high {
		mid := (low + high) >> 1
		next := mid

		for arr[next] == "" && next < len(arr) - 1 {
			next++
		}
		vnext := arr[next]

		if vnext == "" || val < vnext {
			high = mid - 1
		} else if val > vnext {
			low = next + 1
		} else {
			return next
		}
	}

	return -1
}


//------------------------------------------------------------------------------

func main() {
	// fmt.Println()
}
