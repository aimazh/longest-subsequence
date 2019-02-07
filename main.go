package main

import (
	"fmt"
)

type Carcase struct {
	Chars map[rune]int
}

func asCarcase(str string) Carcase {
	c := Carcase{
		Chars: make(map[rune]int),
	}
	for i, char := range str {
		if _, ok := c.Chars[char]; !ok {
			c.Chars[char] = i	
		}
	}
	return c
}

func main() {
	// use any of: abpcplea, abppplee, abpcpleeaDChsssgtrsssmonkey, abpcplebaee
	base := "abpcplebaee"
	words := []string{"able", "ale", "apple", "bale", "monkey", "kangaroo", "applebee"}

	carcase := asCarcase(base)

	rightOrderWords := findRightOrderWords(words, carcase)
	posOfLongest := findPosOfLongest(rightOrderWords)

	if posOfLongest > -1 {
		fmt.Println("word found:", rightOrderWords[posOfLongest])
	} else {
		fmt.Println("word not found")
	}
}

func findPosOfLongest(words []string) int {
	maxLength := 0
	posOfLongest := -1
	for pos, word := range words {
		length := len(word)
		if length > maxLength {
			maxLength = length
			posOfLongest = pos
		}
	}
	return posOfLongest
}

func findRightOrderWords(words []string, carcase Carcase) []string {
	rightOrderWords := []string{}

	for _, word := range words {
		if isRightOrder(word, carcase) {
			rightOrderWords = append(rightOrderWords, word)
		}
	}

	return rightOrderWords
}

func isRightOrder(word string, carcase Carcase) bool {
	rightOrder := true
	for pos, char := range word {
		if pos == 0 {
			if carcase.Chars[char] != 0 {
				rightOrder = false
				break
			}
		} else if pos > 0 {
			charOrder, ok := carcase.Chars[char]
			if !ok {
				rightOrder = false
				break
			}

			prevPos := pos - 1
			prevChar := rune(word[prevPos])
			prevCharOrder, ok := carcase.Chars[prevChar]
			if !ok {
				rightOrder = false
				break
			}

			if prevCharOrder > charOrder {
				rightOrder = false
				break
			}
		}
	}

	return rightOrder
}
