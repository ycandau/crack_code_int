package main

import (
	"reflect"
	"testing"
)

//------------------------------------------------------------------------------
// Problem 5.1

type testType51 struct {
	dest  int
	src   int
	begin int
	end   int
	exp   int
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
// Problem 5.3

type testType53 struct {
	n  int
	exp   int
}

var tests53 = []testType53{
	{0b_0000_0000_0000, 1},
	{0b_0000_0000_0001, 2},
	{0b_0000_0001_0000, 2},
	{0b_0000_0011_0000, 3},
	{0b_0000_0101_0000, 3},
	{0b_0101_0101_0101, 3},
	{0b_1110_1100_0101, 6},
}

func Test53(t *testing.T) {
	for _, test := range tests53 {
		output := FlipBit(test.n)
		if output != test.exp {
			t.Errorf("( %b ) => %v | %v", test.n, output, test.exp)
		} else {
			t.Log("Pass")
		}
	}
}

//------------------------------------------------------------------------------
// Problem 5.4

type testType54 struct {
	n    int
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
	n1  int
	n2  int
	exp int
}

var tests56 = []testType56{
	{0b_0, 0b_0, 0},
	{0b_1, 0b_0, 1},
	{0b_1111_1111, 0b_1011_0010, 4},
}

func Test56a(t *testing.T) {
	for _, test := range tests56 {
		output := Conversion_Naive(test.n1, test.n2)
		if output != test.exp {
			t.Errorf("( %b %b ) => %v | %v", test.n1, test.n2, output, test.exp)
		} else {
			t.Log("Pass")
		}
	}
}

func Test56b(t *testing.T) {
	for _, test := range tests56 {
		output := Conversion_Kernighan(test.n1, test.n2)
		if output != test.exp {
			t.Errorf("( %b %b ) => %v | %v", test.n1, test.n2, output, test.exp)
		} else {
			t.Log("Pass")
		}
	}
}

func Test56c(t *testing.T) {
	for _, test := range tests56 {
		output := Conversion_Table(test.n1, test.n2)
		if output != test.exp {
			t.Errorf("( %b %b ) => %v | %v", test.n1, test.n2, output, test.exp)
		} else {
			t.Log("Pass")
		}
	}
}

//------------------------------------------------------------------------------
// Problem 5.7

type testType57 struct {
	n   int
	exp int
}

var tests57 = []testType57{
	{0b_0, 0b_0},
	{0b_1, 0b_10},
	{0b_1111_1111, 0b_1111_1111},
	{0b_1001_1010, 0b_0110_0101},
}

func Test57a(t *testing.T) {
	for _, test := range tests57 {
		output := PairwiseSwap_Naive(test.n)
		if output != test.exp {
			t.Errorf("( %b ) => %b | %b", test.n, output, test.exp)
		} else {
			t.Log("Pass")
		}
	}
}

func Test57b(t *testing.T) {
	for _, test := range tests57 {
		output := PairwiseSwap_Mask(test.n)
		if output != test.exp {
			t.Errorf("( %b ) => %b | %b", test.n, output, test.exp)
		} else {
			t.Log("Pass")
		}
	}
}

//------------------------------------------------------------------------------
// Problem 5.8

type testType58 struct {
	screen []byte
	width  int
	x1     int
	x2     int
	y      int
	exp    []byte
}

func blank(dx, dy int) []byte {
	return make([]byte, (dx*dy)>>3)
}

var screen1 = []byte{0, 255, 0, 0, 0, 0, 0, 0, 0}
var screen2 = []byte{0, 0, 0, 240, 0, 0, 0, 0, 0}
var screen3 = []byte{0, 0, 0, 0, 0, 0, 0, 0, 15}

var tests58 = []testType58{
	{blank(24, 3), 24, 8, 15, 0, screen1},
	{blank(24, 3), 24, 0, 3, 1, screen2},
	{blank(24, 3), 24, 20, 23, 2, screen3},
}

func Test58(t *testing.T) {
	for _, test := range tests58 {
		output := DrawLine(test.screen, test.width, test.x1, test.x2, test.y)
		if !reflect.DeepEqual(output, test.exp) {
			t.Errorf("( %v %v %v ) => %v | %v",
				test.x1, test.x2, test.y, output, test.exp)
		} else {
			t.Log("Pass")
		}
	}
}
