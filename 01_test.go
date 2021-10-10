package main

import "testing"

//------------------------------------------------------------------------------

type addTest struct {
	str      string
	exp bool
}

var testsAllUnique = []addTest{
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