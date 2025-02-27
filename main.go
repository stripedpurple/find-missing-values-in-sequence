package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strconv"
	"strings"
)

func main() {
	// Create a scanner to read from stdin
	scanner := bufio.NewScanner(os.Stdin)

	// Read all input until EOF
	var input strings.Builder
	for scanner.Scan() {
		input.WriteString(scanner.Text() + " ") // Add a space to separate lines
	}

	// Split the input string into individual tokens using any whitespace
	tokens := strings.Fields(input.String())

	// Convert valid number tokens to integers, ignoring non-numeric tokens
	var intNumbers []int
	for _, token := range tokens {
		num, err := strconv.Atoi(token)
		if err == nil { // Only append if the token is a valid number
			intNumbers = append(intNumbers, num)
		}
	}

	// If no valid numbers are provided, exit
	if len(intNumbers) == 0 {
		fmt.Println("No valid numbers provided")
		return
	}

	// Sort the numbers
	sort.Ints(intNumbers)

	// Find the min and max values
	min := intNumbers[0]
	max := intNumbers[len(intNumbers)-1]

	// Create a map to store the existing numbers for quick lookup
	existingNumbers := make(map[int]bool)
	for _, num := range intNumbers {
		existingNumbers[num] = true
	}

	// Find the missing numbers between min and max
	var missingNumbers []int
	for i := min; i <= max; i++ {
		if !existingNumbers[i] {
			missingNumbers = append(missingNumbers, i)
		}
	}

	// Print the missing numbers
	for _, num := range missingNumbers {
		fmt.Printf("%d ", num)
	}
	fmt.Println() // Print a newline at the end
}
