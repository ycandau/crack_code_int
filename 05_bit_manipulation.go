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

func NextNumber(n int) (prev, next int) {
	return n, n
}
