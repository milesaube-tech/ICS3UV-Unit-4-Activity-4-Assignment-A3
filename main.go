/**
 * Author: Miles Aube
 * Version: 1.0.0
 * Date: 2025-12-13
 * Fileoverview: program to find user inputed word in user inputed sentence
 */

package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {

	// set variables
	var userSentence string
	var userWord string
	//create buffer reader
	reader := bufio.NewReader(os.Stdin)

	// get user input (sentence and word)
	fmt.Println("Please enter a sentence ")
	userSentence, _ = reader.ReadString('\n')

	fmt.Println("Please enter a word to search for in your sentence ")
	userWord, _ = reader.ReadString('\n')
	userWord = strings.TrimSpace(userWord)

	// create if statement for findiing word in sentence
	if strings.Contains(userSentence, userWord) {
		fmt.Println("Horray, the word", userWord, "was found in the sentence", userSentence)
	} else {
		fmt.Println("Sorry, the word", userWord, "was not found in the sentence", userSentence)
	}

	fmt.Println("\nDone.")
}
