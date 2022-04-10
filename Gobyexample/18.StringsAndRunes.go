package main

import (
	"fmt"
	"unicode/utf8"
)

/*
A go string is a read only slice of bytes. The language and the standard library treat strings 
specially - as containers of text encoded in UTF-8. In other languages, strings are made of
characters. In go, the concept of a character is called a rune - it's an integer that 
represents a Unicode point. 

s is a string assigned a literal value representing the word "hello" in the Thai language. 
Go string literals are UTF-8 encoded text.

Since strings are equivalent to []byte, this will produce the length of the raw bytes stored within.

Indexing into a string produces the raw byte values at each index. This loop generates the hex values
of all the bytes that constitute the code points in s.

To count how many runes are in a string, we can use the utf8 package. Note that the run time of
RuneCountInString depends on the size of the string, because it has to decode each utf8 rune 
sequentially. Some thai characters are represented by multiple UTF-8 code points, so the result of
this count may be surprising.

A range loop handles strings specially and decodes each rune along with its offset in the string.

We can achieve the same iteration by using the utf8.DecodeRuneString function explicitly.

This demonstrates passing a rune value to a function.

Values enclosed in single quotes are rune literals. We can compare a rune value to a
rune literal directly.
*/

func main() {
	const s = "สวัสดี"

	fmt.Println("Len:", len(s))

	for i := 0; i < len(s); i++ {
		fmt.Printf("%x ", s[i])
	}

	for i := 0; i < len(s); i++ {
		fmt.Printf("%x ", s[i])
	}
	fmt.Println()

	fmt.Println("Rune Count: ", utf8.RuneCountInString(s))
}