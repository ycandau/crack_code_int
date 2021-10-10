package main

import (
	"fmt"
)

//------------------------------------------------------------------------------

func AllUnique(str string) bool {
	counts := map[rune]bool{}
	for _, ch := range str {
		if _, has := counts[ch]; has {
			return false
		}
		counts[ch] = true
	}
	return true
}

func AllUniqueWithoutMap(str string) bool {
	counts := make([]bool, 256)
	for _, ch := range str {
		index := int(ch)
		if has := counts[index]; has {
			return false
		}
		counts[ch] = true
	}
	return true
}

//------------------------------------------------------------------------------

func main() {
	fmt.Println(AllUnique("abc"))
	fmt.Println(AllUnique("abca"))
}
