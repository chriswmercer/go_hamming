//Package hamming provides functions to calculate hamming differences
package hamming

import "fmt"

//Distance calculates the Hamming difference between two DNA strands, which are represented as strings
func Distance(a, b string) (int, error) {
	if len(a) != len(b) {
		return -1, fmt.Errorf("hamming: Distance can only be calculated for strings of equal length")
	}

	charA := []rune(a)
	charB := []rune(b)
	difference := 0

	for index, element := range charA {
		if element != charB[index] {
			difference++
		}
	}

	return difference, nil
}
