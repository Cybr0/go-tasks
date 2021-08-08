package main

import (
	"fmt"
)

func main() {
	fmt.Print("Введите слово: ")
	var word string
	fmt.Scan(&word)

	if IsPalindrome(word) {
		fmt.Println("yes")
	} else {
		fmt.Println("no")
	}
}

func IsPalindrome(word string) bool {
	reversedWord := Reverse(word)
	if reversedWord == word {
		return true
	}
	return false
}

func Reverse(s string) string {
	runes := []rune(s)
	for i, j := 0, len(runes)-1; i < j; i, j = i+1, j-1 {
		runes[i], runes[j] = runes[j], runes[i]
	}
	return string(runes)
}
