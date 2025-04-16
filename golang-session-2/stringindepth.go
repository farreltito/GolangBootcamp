package main

import (
	"fmt"
	"unicode/utf8"
)

/*
	This is a String in Depth project.
*/

func main() {
	//Looping Over String (byte-by-byte)
	//String to Byte
	kota1 := "Surabaya" //diambil per hurufnya = 1 byte

	for i := 0; i < len(kota1); i++ {
		fmt.Printf("Index: %d, byte: %d\n", i, kota1[i])
	}
	fmt.Println("")

	//Byte to String
	var kota2 []byte = []byte{83, 117, 114, 97, 98, 97, 121, 97}
	fmt.Println(string(kota2))
	fmt.Println("")

	//finding BYTE LENGTH from its String
	string1 := "dog"
	string2 := "cât"
	fmt.Printf("string1 byte LENGTH --> %d\n", len(string1))
	fmt.Printf("string1 byte LENGTH --> %d\n", len(string2))
	fmt.Println("")

	//Looping Over String (rune-by-rune)
	//finding CHARACTER LENGTH
	fmt.Printf("string1 char LENGTH --> %d\n", utf8.RuneCountInString(string1))
	fmt.Printf("string1 char LENGTH --> %d\n", utf8.RuneCountInString(string2))
	fmt.Println("")

	//Looping rune-rune or per Character
	for i, s := range string2 {
		fmt.Printf("Index --> %d, string --> %s\n", i, string(s))
	}
}

/*
	Note:

	Result:
	Index: 0, byte: 83
	Index: 1, byte: 117
	Index: 2, byte: 114
	Index: 3, byte: 97
	Index: 4, byte: 98
	Index: 5, byte: 97
	Index: 6, byte: 121
	Index: 7, byte: 97

	Surabaya

	string1 byte LENGTH --> 3
	string1 byte LENGTH --> 4

	string1 char LENGTH --> 3
	string1 char LENGTH --> 3

	Index --> 0, string --> c
	Index --> 1, string --> â
	Index --> 3, string --> t
*/
