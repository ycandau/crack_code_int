package main

//------------------------------------------------------------------------------
// Problem 5.1

func Insertion(dest, src, begin, end int) int {
	mask := 1 << (end - begin + 1) // 0000 1000
	mask -= 1                      // 0000 0111
	mask <<= begin                 // 0001 1100

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
	for bit(n, i) == 0 {
		i++
	}
	for bit(n, i) == 1 {
		i++
	}
	next = n ^ (3 << (i - 1))

	i = 0
	for bit(n, i) == 1 {
		i++
	}
	for bit(n, i) == 0 {
		i++
	}
	prev = n ^ (3 << (i - 1))

	return
}

//------------------------------------------------------------------------------
// Problem 5.6

func Conversion_Naive(n1, n2 int) int {
	count := 0
	for diff := n1 ^ n2; diff != 0; diff >>= 1 {
		count += diff & 1
	}
	return count
}

func Conversion_Kernighan(n1, n2 int) int {
	count := 0
	for diff := n1 ^ n2; diff != 0; count++ {
		diff &= diff - 1 // clear least significant bit
	}
	return count
}

var counts = map[int]int{
	0: 0, 1: 1, 2: 1, 3: 2, 4: 1, 5: 2, 6: 2, 7: 3,
	8: 1, 9: 2, 10: 2, 11: 3, 12: 2, 13: 3, 14: 3, 15: 4,
}

func Conversion_Table(n1, n2 int) int {
	diff := n1 ^ n2
	return counts[(diff>>0)&0b_1111] +
		counts[(diff>>4)&0b_1111] +
		counts[(diff>>8)&0b_1111] +
		counts[(diff>>12)&0b_1111] +
		counts[(diff>>16)&0b_1111] +
		counts[(diff>>20)&0b_1111] +
		counts[(diff>>24)&0b_1111] +
		counts[(diff>>28)&0b_1111]
}

//------------------------------------------------------------------------------
// Problem 5.7

func PairwiseSwap_Naive(n int) int {
	output := 0
	for i := 0; i < 16; {
		output |= (n & (1 << i)) << 1
		i++
		output |= (n & (1 << i)) >> 1
		i++
	}
	return output
}

var mask57 = 0b_0101_0101_0101_0101_0101_0101_0101_0101

func PairwiseSwap_Mask(n int) int {
	return ((n & mask57) << 1) | ((n & (mask57 << 1)) >> 1)
}

//------------------------------------------------------------------------------
// Problem 5.8

func DrawLine(screen []byte, width, x1, x2, y int) []byte {

	return screen
}
