package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"
)

var minVowels = int(3)
var vowels = "aeiou"
var totalNice = int(0)

// Nice (5a) =
// 1. Contains at least 3 vowels - 'aeiou' - DONE
// 2. Contains one letter that appears twice in a row. DONE.
// 3. Does NOT contain 'ab', 'cd', 'pq', 'xy' - DONE

func main() {
	if file, err := os.Open("input.txt"); err == nil {
		defer file.Close()

		scanner := bufio.NewScanner(file)
		for scanner.Scan() {
			word := strings.TrimSpace(scanner.Text())
			if checkWord(word) {
				totalNice++
			}
		}
		if err = scanner.Err(); err != nil {
			log.Fatal(err)
		}
	} else {
		log.Fatal(err)
	}
	fmt.Printf("\n\n%d total Nice strings\n", totalNice)
}

func checkWord(word string) bool {
	enoughVowels := vowelCount(word)
	clean := noBadStrings(word)
	duplicateLetter := haveDuplicates(word)

	if enoughVowels && clean && duplicateLetter {
		fmt.Printf("'%s' is nice\n", word)
		return true
	}
	fmt.Printf("'%s' is naughty\n", word)
	return false
}

func haveDuplicates(word string) bool {
	length := len(word)
	for i := 1; i < length; i++ {
		if word[i] == word[i-1] {
			fmt.Printf("%s has a duplicate %s\n", word, string(word[i]))
			return true
		}
	}
	fmt.Printf("%s does NOT have a duplicate\n", word)
	return false
}

func vowelCount(word string) bool {
	count := int(0)
	length := len(vowels)
	for i := 0; i < length; i++ {
		letter := vowels[i]
		letterCount := strings.Count(word, string(letter))
		if letterCount > 0 {
			count += letterCount
		}
	}
	fmt.Printf("%s has %d vowels\n", word, count)
	if count >= minVowels {
		return true
	}
	return false
}

func noBadStrings(word string) bool {
	result := true
	if stringMatch(word, "ab") {
		return false
	}
	if stringMatch(word, "cd") {
		return false
	}
	if stringMatch(word, "pq") {
		return false
	}
	if stringMatch(word, "xy") {
		return false
	}
	return result
}

func stringMatch(haystack string, needle string) bool {
	if strings.Contains(haystack, needle) {
		fmt.Printf("%s contains %s\n", haystack, needle)
		return true
	}
	fmt.Printf("%s does NOT contain %s\n", haystack, needle)
	return false
}
