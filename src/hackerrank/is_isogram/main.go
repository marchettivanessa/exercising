package main

import (
	"fmt"
	"strings"
	"unicode"
)

// Isogram checks if the word is an isogram
func IsIsogram(word string) bool {
	// stores if the letter has appeared
	checksWord := make(map[rune]bool)

	// removes spaces and hiphens
	cleanedWord := strings.ReplaceAll(word, " ", "")
	cleanedWord = strings.ReplaceAll(cleanedWord, "-", "")
	// iterate througb the word
	for _, char := range cleanedWord {
		char = unicode.ToLower(char)
		// checks if the char is repeated
		if checksWord[char] {
			return false // letter is repeated, not an isogram
		}
		checksWord[char] = true
	}
	return true
}

func main() {

	notAnIsogram := "Vamos testar se é isogram"
	itIsAnIsogram := "Jaquelin!"

	fmt.Println(IsIsogram(notAnIsogram))
	fmt.Println(IsIsogram(itIsAnIsogram))
}
