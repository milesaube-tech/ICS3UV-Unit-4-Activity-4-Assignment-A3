/**
 * @author Miles Aube
 * @version 1.0.0
 * @date 2025-12-13
 * @fileoverview program for user to input sentence then find word in sentence 
 */

// get user input 
const userSentence = prompt("Please enter a sentence") || "No sentence entered";
const userWord = prompt("Please enter a word to search for in your sentence") || "No word entered";

// create if statement to find word in sentence 
if (userSentence.includes(userWord)){
  console.log(`Horray, the word ${userWord} was found in the sentence ${userSentence}`)
}

// else statement if word not found 
else {
  console.log(`Sorry, the word ${userWord} was not found in the sentence ${userSentence}`)
}

console.log("\nDone.");