package main

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

func main() {
	SortedMerge([]int{1, 2, 3}, []int{4, 5, 6})

}
