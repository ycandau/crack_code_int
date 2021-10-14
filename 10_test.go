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
