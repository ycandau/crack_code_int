package main

import (
	"testing"
)

//------------------------------------------------------------------------------
// Problem 5.1

type testType51 struct {
	dest int
	src int
	begin int
	end int
	exp int
}

var tests51 = []testType51{
	{0, 0, 0, 2, 0},
	{255, 0, 0, 2, 255 - 7},
	{255, 0, 2, 4, 255 - 28},
	{0, 7, 0, 5, 7},
	{0, 7, 2, 7, 28},
	{255, 4, 0, 3, 255 - 11},
	{255, 4, 2, 5, 255 - 44},

}

func Test51(t *testing.T) {
	for _, test := range tests51 {
		output := Insertion(test.dest, test.src, test.begin, test.end)
		if output != test.exp {
			t.Errorf("( %v %v %v %v ) => %v | %v",
				test.dest, test.src, test.begin, test.end, output, test.exp)
		} else {
			t.Log("Pass")
		}
	}
}

//------------------------------------------------------------------------------
// Problem 5.4

type testType54 struct {
	n int
	prev int
	next int
}

var tests54 = []testType54{
	{2, 1, 4},
	{10, 9, 12},
}

func Test54(t *testing.T) {
	for _, test := range tests54 {
		prev, next := NextNumber(test.n)
		if prev != test.prev || next != test.next {
			t.Errorf("( %v ) => ( %v %v ) | ( %v %v )",
				test.n, prev, next, test.prev, test.next)
		} else {
			t.Log("Pass")
		}
	}
}

//------------------------------------------------------------------------------
// Problem 5.6

type testType56 struct {
	n1 int
	n2 int
	exp int
}

var tests56 = []testType56{
	{0b_0, 0b_0, 0},
	{0b_1, 0b_0, 1},
	{0b_1111_1111, 0b_1011_0010, 4},
}

func Test56(t *testing.T) {
	for _, test := range tests56 {
		output := Conversion(test.n1, test.n2)
		if output != test.exp {
			t.Errorf("( %b %b ) => %v | %v", test.n1, test.n2, output, test.exp)
		} else {
			t.Log("Pass")
		}
	}
}
