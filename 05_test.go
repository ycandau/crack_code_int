package main

import (
	"testing"
)

//------------------------------------------------------------------------------
// Problem 1.1

type testType51 struct {
	dest int
	src int
	pos int
	exp int
}

var tests51 = []testType51{
	{0, 0, 0, 0},
	{0, 7, 0, 7},
	{0, 7, 2, 27},

}

func Test51(t *testing.T) {
	for _, test := range tests51 {
		output := Insertion(test.dest, test.src, test.pos)
		if output != test.exp {
			t.Errorf("( %v %v %v ) => %v | %v",
				test.dest, test.src, test.pos, output, test.exp)
		} else {
			t.Log("Pass")
		}
	}
}
