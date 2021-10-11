package main

import (
	"reflect"
	"testing"
)

//------------------------------------------------------------------------------
// Problem 1.1

type testType1 struct {
	str string
	exp bool
}

var tests1 = []testType1{
	{"abcde", true},
	{"abcda", false},
	{"ebcde", false},
	{"abade", false},
}

func Test1a(t *testing.T) {
	for _, test := range tests1 {
		output := AllUnique(test.str)
		if output != test.exp {
			t.Errorf("( %v ) => %v | %v", test.str, output, test.exp)
		} else {
			t.Log("Pass")
		}
	}
}

func Test1b(t *testing.T) {
	for _, test := range tests1 {
		output := AllUniqueWithoutMap(test.str)
		if output != test.exp {
			t.Errorf("( %v ) => %v | %v", test.str, output, test.exp)
		} else {
			t.Log("Pass")
		}
	}
}

//------------------------------------------------------------------------------
// Problem 1.2

type testType2 struct {
	str1 string
	str2 string
	exp  bool
}

var tests2 = []testType2{
	{"xabc", "xabc", true},
	{"xabc", "abcx", true},
	{"xabc", "abc", false},
	{"abc", "xabc", false},
	{"aabc", "abcc", false},
	{"xabc", "abcy", false},
}

func Test2(t *testing.T) {
	for _, test := range tests2 {
		output := ArePermutations(test.str1, test.str2)
		if output != test.exp {
			t.Errorf("( %v , %v ) => %v | %v",
				test.str1, test.str2, output, test.exp)
		} else {
			t.Log("Pass")
		}
	}
}

//------------------------------------------------------------------------------
// Problem 1.3

type testType3 struct {
	str string
	exp string
}

var tests3 = []testType3{
	{"alpha", "alpha"},
	{"alpha beta", "alpha%20beta"},
	{"alpha beta gamma", "alpha%20beta%20gamma"},
}

func Test3(t *testing.T) {
	for _, test := range tests3 {
		output := ToURL(test.str)
		if output != test.exp {
			t.Errorf("( %v ) => %v | %v",
				test.str, output, test.exp)
		} else {
			t.Log("Pass")
		}
	}
}

//------------------------------------------------------------------------------
// Problem 1.4

type testType4 struct {
	str string
	exp bool
}

var tests4 = []testType4{
	{"aa", true},
	{"aaa", true},
	{"aabb", true},
	{"aabbc", true},
	{"aabbcd", false},
}

func Test4a(t *testing.T) {
	for _, test := range tests4 {
		if output := IsPalindromePermutation(test.str); output != test.exp {
			t.Errorf("( %v ) => %v | %v", test.str, output, test.exp)
		} else {
			t.Log("Pass")
		}
	}
}

func Test4b(t *testing.T) {
	for _, test := range tests4 {
		if output := IsPalindromePermutation_BV(test.str); output != test.exp {
			t.Errorf("( %v ) => %v | %v", test.str, output, test.exp)
		} else {
			t.Log("Pass")
		}
	}
}

//------------------------------------------------------------------------------
// Problem 1.5

type testType5 struct {
	str1 string
	str2 string
	exp  bool
}

var tests5 = []testType5{
	{"alpha", "alpha", true},
	{"alpha", "al_ha", true},
	{"alpha", "al__a", false},
	{"alpha", "alph_a", true},
	{"alpha", "al_ph_a", false},
	{"a_lpha", "alpha", true},
	{"a_lp_ha", "alpha", false},
}

func Test5(t *testing.T) {
	for _, test := range tests5 {
		output := OneAway(test.str1, test.str2)
		if output != test.exp {
			t.Errorf("( %v , %v ) => %v | %v",
				test.str1, test.str2, output, test.exp)
		} else {
			t.Log("Pass")
		}
	}
}

//------------------------------------------------------------------------------
// Problem 1.6

type testType6 struct {
	str string
	exp string
}

var tests6 = []testType6{
	{"alpha", "alpha"},
	{"aabbcc", "aabbcc"},
	{"aabbccc", "2a2b3c"},
	{"aaaaaaaaaa", "10a"},
}

func Test6a(t *testing.T) {
	for _, test := range tests6 {
		if output := StringCompression(test.str); output != test.exp {
			t.Errorf("( %v ) => %v | %v", test.str, output, test.exp)
		} else {
			t.Log("Pass")
		}
	}
}

func Test6b(t *testing.T) {
	for _, test := range tests6 {
		output := StringCompression_SB(test.str)
		if output != test.exp {
			t.Errorf("( %v ) => %v | %v", test.str, output, test.exp)
		} else {
			t.Log("Pass")
		}
	}
}

//------------------------------------------------------------------------------
// Problem 1.7

func copyMatrix(matr matrix) matrix {
	output := newMatrix(len(matr[0]), len(matr))
	for x := range matr[0] {
		for y := range matr {
			output[y][x] = matr[y][x]
		}
	}
	return output
}

type testType7 struct {
	inp matrix
	exp matrix
}

var tests7 = []testType7{
	{matrix{{1, 2}, {3, 4}}, matrix{{2, 4}, {1, 3}}},
	{
		matrix{{1, 2, 3}, {4, 5, 6}, {7, 8, 9}},
		matrix{{3, 6, 9}, {2, 5, 8}, {1, 4, 7}}},
}

func Test7(t *testing.T) {
	for _, test := range tests7 {
		input := copyMatrix(test.inp)
		output := RotateMatrix(test.inp)
		if !reflect.DeepEqual(output, test.exp) {
			t.Errorf("( %v ) => %v | %v", input, output, test.exp)
		} else {
			t.Log("Pass")
		}
	}
}

//------------------------------------------------------------------------------
// Problem 1.8

type testType8 struct {
	inp matrix
	exp matrix
}

var tests8 = []testType7{
	{
		matrix{{1, 0, 3}, {4, 5, 6}, {7, 8, 9}},
		matrix{{0, 0, 0}, {4, 0, 6}, {7, 0, 9}}},
	{
		matrix{{1, 0, 3}, {4, 5, 6}, {7, 8, 0}},
		matrix{{0, 0, 0}, {4, 0, 0}, {0, 0, 0}}},
}

func Test8(t *testing.T) {
	for _, test := range tests8 {
		input := copyMatrix(test.inp)
		output := ZeroMatrix(test.inp)
		if !reflect.DeepEqual(output, test.exp) {
			t.Errorf("( %v ) => %v | %v", input, output, test.exp)
		} else {
			t.Log("Pass")
		}
	}
}
