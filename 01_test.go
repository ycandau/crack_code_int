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
