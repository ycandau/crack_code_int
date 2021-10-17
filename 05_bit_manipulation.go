package main

//------------------------------------------------------------------------------
// Problem 5.1

func Insertion(dest, src, begin, end int) int {
	mask := 1 << (end - begin + 1) // 0000 1000
	mask -= 1 // 0000 0111
	mask <<= begin // 0001 1100

	return (dest &^ mask) | (src << begin)
}

//------------------------------------------------------------------------------
// Problem 5.4

func bit(n, i int) int {
	return (n >> i) & 1
}

func NextNumber(n int) (prev, next int) {
	if n == 0 {
		return 0, 0
	}

	i := 0
	for bit(n, i) == 0 { i++ }
	for bit(n, i) == 1 { i++ }
	next = n ^ (3 << (i-1))

	i = 0
	for bit(n, i) == 1 { i++ }
	for bit(n, i) == 0 { i++ }
	prev = n ^ (3 << (i-1))

	return
}

//------------------------------------------------------------------------------
// Problem 5.6

func Conversion(n1, n2 int) int {
	return 0
}
