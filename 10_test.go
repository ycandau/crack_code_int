package main

import (
	"reflect"
	"testing"
)

//------------------------------------------------------------------------------
// Problem 10.1

type testType101 struct {
	arr1 []int
	arr2 []int
	exp  []int
}

var tests101 = []testType101{
	{[]int{1, 2, 3, 4}, []int{}, []int{1, 2, 3, 4}},
	{[]int{}, []int{1, 2, 3, 4}, []int{1, 2, 3, 4}},
	{[]int{1, 2, 3}, []int{2, 3, 4}, []int{1, 2, 2, 3, 3, 4}},
	{[]int{1, 3, 5}, []int{2, 4, 6}, []int{1, 2, 3, 4, 5, 6}},
}

func Test101(t *testing.T) {
	for _, test := range tests101 {
		output := SortedMerge(test.arr1, test.arr2)
		if !reflect.DeepEqual(output, test.exp) {
			t.Errorf("( %v , %v ) => %v | %v", test.arr1, test.arr2, output, test.exp)
		} else {
			t.Log("Pass")
		}
	}
}

//------------------------------------------------------------------------------
// Problem 10.2

type testType102 struct {
	arr []string
	exp []string
}

var tests102 = []testType102{
	{
		[]string{"a", "abc", "abcd", "ab", "a", "bac", "dabc", "ba"},
		[]string{"a", "a", "ab", "ba", "abc", "bac", "abcd", "dabc"}},
}

func Test102(t *testing.T) {
	for _, test := range tests102 {
		output := GroupAnagrams(test.arr)
		if !reflect.DeepEqual(output, test.exp) {
			t.Errorf("( %v )\n=> %v\n!= %v", test.arr, output, test.exp)
		} else {
			t.Log("Pass")
		}
	}
}

//------------------------------------------------------------------------------
// Problem 10.3

type testType103 struct {
	arr []int
	val int
	exp int
}

var tests103 = []testType103{
	{[]int{8, 9, 10, 1, 2, 3, 4, 5, 6, 7}, 8, 0},
	{[]int{8, 9, 10, 1, 2, 3, 4, 5, 6, 7}, 9, 1},
	{[]int{8, 9, 10, 1, 2, 3, 4, 5, 6, 7}, 6, 8},
	{[]int{8, 9, 10, 1, 2, 3, 4, 5, 6, 7}, 7, 9},
	{[]int{4, 5, 6, 7, 8, 9, 10, 1, 2, 3}, 4, 0},
	{[]int{4, 5, 6, 7, 8, 9, 10, 1, 2, 3}, 5, 1},
	{[]int{4, 5, 6, 7, 8, 9, 10, 1, 2, 3}, 2, 8},
	{[]int{4, 5, 6, 7, 8, 9, 10, 1, 2, 3}, 3, 9},
	{[]int{4, 5, 6, 7, 8, 9, 10, 1, 2, 3}, 3, 9},
	{[]int{5, 5, 5, 1, 2}, 1, 3},
	{[]int{5, 5, 5, 1, 5}, 1, 3},
	{[]int{6, 7, 5, 5, 5}, 7, 1},
	{[]int{5, 7, 5, 5, 5}, 7, 1},
}

func Test103(t *testing.T) {
	for _, test := range tests103 {
		output, _ := RotatedSearch(test.arr, test.val)
		if output != test.exp {
			t.Errorf("( %v , %v ) => %v | %v", test.arr, test.val, output, test.exp)
		} else {
			t.Log("Pass")
		}
	}
}

//------------------------------------------------------------------------------
// Problem 10.4

type testType104 struct {
	arr []int
	val int
	exp int
}

var tests104 = []testType104{
	{[]int{1, 2, 3, 4, 5}, 3, 2},
	{[]int{1, 2, 3, 4, 5}, 6, -1},
}

func Test104(t *testing.T) {
	for _, test := range tests104 {
		output := ExpSearch(test.arr, test.val)
		if output != test.exp {
			t.Errorf("( %v , %v ) => %v | %v", test.arr, test.val, output, test.exp)
		} else {
			t.Log("Pass")
		}
	}
}

//------------------------------------------------------------------------------
// Problem 10.5

type testType105 struct {
	arr []string
	val string
	exp int
}

var tests105 = []testType105{
	{[]string{"ab", "", "", "bc", "", "cd", "", "", "", "de"}, "", -1},
	{[]string{"ab", "", "", "bc", "", "cd", "", "", "", "de"}, "ab", 0},
	{[]string{"ab", "", "", "bc", "", "cd", "", "", "", "de"}, "bc", 3},
	{[]string{"ab", "", "", "bc", "", "cd", "", "", "", "de"}, "cd", 5},
	{[]string{"ab", "", "", "bc", "", "cd", "", "", "", "de"}, "de", 9},
	{[]string{"ab", "", "", "", ""}, "ab", 0},
	{[]string{"ab", "", "", "", ""}, "xx", -1},
	{[]string{"", "", "", "", "ab"}, "ab", 4},
	{[]string{"", "", "", "", "ab"}, "xx", -1},
}

func Test105(t *testing.T) {
	for _, test := range tests105 {
		output := SparseSearch(test.arr, test.val)
		if output != test.exp {
			t.Errorf("( %v , %v ) => %v | %v", test.arr, test.val, output, test.exp)
		} else {
			t.Log("Pass")
		}
	}
}

//------------------------------------------------------------------------------
// Problem 10.7

type testType107 struct {
	arr []uint8
	exp uint8
}

func arr107(n uint8) []uint8 {
	arr := make([]uint8, (1 << 8) - 1)
	for i := range arr {
		if uint8(i) <= n {
			arr[i] = uint8(i)
		} else {
			arr[i] = uint8(i) + 1
		}
	}
	return arr
}

var tests107 = []testType107{
	{arr107(0), 0},
	{arr107(100), 100},
	{arr107(255), 255},
}

func Test107(t *testing.T) {
	for _, test := range tests107 {
		output := MissingInt(test.arr)
		if output != test.exp {
			t.Errorf("( ... ) => %v | %v", output, test.exp)
		} else {
			t.Log("Pass")
		}
	}
}

//------------------------------------------------------------------------------
// Problem 10.9

type testType109 struct {
	matr matrix
	val  int
	exp  [2]int
}

var matr109 = newMatrix([]int{1, 3, 5, 2, 4, 6, 7, 8, 9}, 3, 3)

var tests109 = []testType109{
	{matr109, 10, [2]int{-1, -1}},
	{matr109, 1, [2]int{0, 0}},
	{matr109, 5, [2]int{2, 0}},
	{matr109, 7, [2]int{0, 2}},
	{matr109, 9, [2]int{2, 2}},
}

func Test109(t *testing.T) {
	for _, test := range tests109 {
		output := MatrixSearch(test.matr, test.val)
		if !reflect.DeepEqual(output, test.exp) {
			t.Errorf("( %v , %v ) => %v | %v", test.matr, test.val, output, test.exp)
		} else {
			t.Log("Pass")
		}
	}
}

//------------------------------------------------------------------------------
// Problem 10.10

type testType110 struct {
	arr []int
	val int
	exp int
}

var stream110 = []int{5, 1, 4, 4, 5, 9, 7, 13, 3}

var tests110 = []testType110{
	{stream110, 0, 0},
	{stream110, 1, 0},
	{stream110, 2, 1},
	{stream110, 3, 1},
	{stream110, 4, 3},
	{stream110, 5, 5},
	{stream110, 6, 6},
	{stream110, 7, 6},
	{stream110, 8, 7},
	{stream110, 9, 7},
	{stream110, 10, 8},
	{stream110, 15, 9},
}

func Test110(t *testing.T) {
	for _, test := range tests110 {
		output := RankFromStream(test.arr, test.val)
		if output != test.exp {
			t.Errorf("( %v , %v ) => %v | %v", test.arr, test.val, output, test.exp)
		} else {
			t.Log("Pass")
		}
	}
}

//------------------------------------------------------------------------------
// Problem 10.11

type testType111 struct {
	arr []int
}

var tests111 = []testType111{
	{[]int{}},
	{[]int{1}},
	{[]int{1, 2}},
	{[]int{1, 2, 3, 4, 5}},
	{[]int{1, 2, 3, 4, 5, 6}},
}

func Test111(t *testing.T) {
	for _, test := range tests111 {
		output := PeaksAndValleys(test.arr)
		for i := 2; i < len(output); i++ {
			if (output[i-2] < output[i-1] && output[i-1] < output[i]) ||
				(output[i-2] > output[i-1] && output[i-1] > output[i]) {
				t.Errorf("( %v ) => %v | %v", test.arr, output, output[i-2:i+1])
				return
			}
		}
		t.Log("Pass")
	}
}
