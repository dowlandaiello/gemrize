// Package main is the gemrize entry point.
package main

import (
	"bufio"
	"fmt"
	"os"
	"os/exec"
	"strings"
)

// main is the gemrize cli.
func main() {
	reader := bufio.NewReader(os.Stdin)                     // Get a reader for the standard input
	p := prompt("What would you like to memorize?", reader) // Ask the user what they would like to memorize

	words := strings.Split(p, " ") // Get the words in the given paragraph

	lookForward := 10 // Go 10 words at a time

	// Check lookForward bigger than the number of inputted words
	if lookForward > len(words) {
		lookForward = len(words) // Set the look forward to the number of words
	}

	// 2 words at a time, at least
	for lookForward >= 1 {
		// Go 10 words at a time
		for i := 0; i < len(words); i += lookForward {
			repeatedCorrectly := false

			for !repeatedCorrectly {
				cutoff := i + lookForward // The last words to include

				// Check out of bounds
				if cutoff >= len(words)+1 {
					cutoff = len(words) // Just get the rest of the words
				}

				mergedWords := strings.Join(words[int(i):cutoff], " ") // Merge the individual words

				repeated := prompt(fmt.Sprintf("Please type the following: %s", mergedWords), reader) // Print the words to repeat

				repeatedCorrectly = repeated == mergedWords // Check repeated correctly
			}
		}

		lookForward-- // Decrement the look forward
	}
}

// prompt prompts the user with a given prompt.
func prompt(prompt string, reader *bufio.Reader) string {
	fmt.Println(prompt) // Print the prompt

	text, err := reader.ReadString('\n') // Prompt the user
	if err != nil {                      // Check for errors
		panic(err) // Panic
	}

	c := exec.Command("clear")
	c.Stdout = os.Stdout
	c.Run()

	return strings.Replace(strings.Replace(strings.Replace(strings.Replace(text, "\n", "", 1), "“", `"`, -1), "”", `"`, -1), "—", "-", -1) // Remove a \n character
}
