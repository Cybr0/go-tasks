package main

import (
	"fmt"
	"strings"
)

func main() {
	arrWords := []string{"Ясуми", "Ито", "Мацуно", "Хироюки", "Ито", "Хироси", "Минагава", "Акитоси", "Кавадзу", "Ито", "Ёити", "Вада"}
	searchWord1, searchWord2 := "Вада", "Ито"

	fmt.Println("ArrWords: ", arrWords)
	PrintDuplicateResult(searchWord1, arrWords)
	PrintDuplicateResult(searchWord2, arrWords)

}

func IsContainsDuplicateWord(word string, words []string) bool {
	var counter int
	for _, v := range words {
		if counter > 1 {
			return true
		}
		if strings.Compare(v, word) == 0 {
			counter++
		}
	}
	return false
}

func PrintDuplicateResult(searchWord string, words []string) {
	result := IsContainsDuplicateWord(searchWord, words)
	fmt.Printf("Is array has duplicate of '%v'? %v\n", searchWord, result)
}
