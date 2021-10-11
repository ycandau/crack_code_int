package main

import "testing"

//------------------------------------------------------------------------------
// Problem 1.1

type addTest1 struct {
	str string
	exp bool
}

var testsAllUnique = []addTest1{
	{"abcde", true},
	{"abcda", false},
	{"ebcde", false},
	{"abade", false},
}

func TestAllUnique(t *testing.T) {
	for _, test := range testsAllUnique {
		if output := AllUnique(test.str); output != test.exp {
			t.Errorf("( %v ) => %v | %v", test.str, output, test.exp)
		} else {
			t.Log("Pass")
		}
	}
}

func TestAllUniqueWithoutMap(t *testing.T) {
	for _, test := range testsAllUnique {
		if output := AllUniqueWithoutMap(test.str); output != test.exp {
			t.Errorf("( %v ) => %v | %v", test.str, output, test.exp)
		} else {
			t.Log("Pass")
		}
	}
}

//------------------------------------------------------------------------------
// Problem 1.2

type addTest2 struct {
	str1 string
	str2 string
	exp  bool
}

var testsArePermutations = []addTest2{
	{"xabc", "xabc", true},
	{"xabc", "abcx", true},
	{"xabc", "abc", false},
	{"abc", "xabc", false},
	{"aabc", "abcc", false},
	{"xabc", "abcy", false},
}

func TestArePermutations(t *testing.T) {
	for _, test := range testsArePermutations {
		if output := ArePermutations(test.str1, test.str2); output != test.exp {
			t.Errorf("( %v , %v ) => %v | %v",
				test.str1, test.str2, output, test.exp)
		} else {
			t.Log("Pass")
		}
	}
}

//------------------------------------------------------------------------------
// Problem 1.3

type addTest3 struct {
	str string
	exp string
}

var testsToURL = []addTest3{
	{"alpha", "alpha"},
	{"alpha beta", "alpha%20beta"},
	{"alpha beta gamma", "alpha%20beta%20gamma"},
}

func TestToURL(t *testing.T) {
	for _, test := range testsToURL {
		if output := ToURL(test.str); output != test.exp {
			t.Errorf("( %v ) => %v | %v",
				test.str, output, test.exp)
		} else {
			t.Log("Pass")
		}
	}
}

//------------------------------------------------------------------------------
// Problem 1.4

type addTest4 struct {
	str string
	exp bool
}

var testsIsPalindromePermutation = []addTest4{
	{"aa", true},
	{"aaa", true},
	{"aabb", true},
	{"aabbc", true},
	{"aabbcd", false},
}

func TestIsPalindromePermutation(t *testing.T) {
	for _, test := range testsIsPalindromePermutation {
		if output := IsPalindromePermutation(test.str); output != test.exp {
			t.Errorf("( %v ) => %v | %v", test.str, output, test.exp)
		} else {
			t.Log("Pass")
		}
	}
}

func TestIsPalindromePermutation_BV(t *testing.T) {
	for _, test := range testsIsPalindromePermutation {
		if output := IsPalindromePermutation_BV(test.str); output != test.exp {
			t.Errorf("( %v ) => %v | %v", test.str, output, test.exp)
		} else {
			t.Log("Pass")
		}
	}
}

//------------------------------------------------------------------------------
// Problem 1.5

type addTest5 struct {
	str1 string
	str2 string
	exp  bool
}

var testsOneAway = []addTest5{
	{"alpha", "alpha", true},
	{"alpha", "al_ha", true},
	{"alpha", "al__a", false},
	{"alpha", "alph_a", true},
	{"alpha", "al_ph_a", false},
	{"a_lpha", "alpha", true},
	{"a_lp_ha", "alpha", false},
}

func TestOneAway(t *testing.T) {
	for _, test := range testsOneAway {
		if output := OneAway(test.str1, test.str2); output != test.exp {
			t.Errorf("( %v , %v ) => %v | %v",
				test.str1, test.str2, output, test.exp)
		} else {
			t.Log("Pass")
		}
	}
}

//------------------------------------------------------------------------------
// Problem 1.6

type addTest6 struct {
	str string
	exp string
}

var testsStringCompression = []addTest6{
	{"alpha", "alpha"},
	{"aabbcc", "aabbcc"},
	{"aabbccc", "2a2b3c"},
	{"aaaaaaaaaa", "10a"},
}

func TestStringCompression(t *testing.T) {
	for _, test := range testsStringCompression {
		if output := StringCompression(test.str); output != test.exp {
			t.Errorf("( %v ) => %v | %v", test.str, output, test.exp)
		} else {
			t.Log("Pass")
		}
	}
}

func TestStringCompression_SB(t *testing.T) {
	for _, test := range testsStringCompression {
		if output := StringCompression_SB(test.str); output != test.exp {
			t.Errorf("( %v ) => %v | %v", test.str, output, test.exp)
		} else {
			t.Log("Pass")
		}
	}
}
