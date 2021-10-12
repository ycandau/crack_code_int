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
