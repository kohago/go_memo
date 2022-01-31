package main

import (
	"fmt"
)

func main() {
	input := "The quick brown fox jumped over the lazy dog"
	rev := reverse(input)
	doubleRev := reverse(rev)
	fmt.Printf("original: %q\n", input)
	fmt.Printf("reversed: %q\n", rev)
	fmt.Printf("reversed again: %q\n", doubleRev)
}

//every character was a single byte. However, characters such as 泃 can require several bytes.
//so will crash
func reverse(s string) string {
	//if !utf8.ValidString(s) {
	//	return s, errors.New("input is not valid UTF-8")
	//}
	//b:=[]rune(s)
	b := []byte(s)
	for i, j := 0, len(b)-1; i < len(b)/2; i, j = i+1, j-1 {
		b[i], b[j] = b[j], b[i]
	}
	return string(b)
}
