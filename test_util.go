package main

import "fmt"

func squareArraysMatch(s1 []*Square, s2 []*Square) (bool, error) {
	if len(s1) != len(s2) {
		return false, fmt.Errorf("s1 of length %d, while s2 of length %d", len(s1), len(s2))
	}

	// Certainly a better way to do this, but here we are.
	for _, b := range s2 {
		foundMatch := false
		fmt.Printf("\nEvaluating %d, %d... %T", b.x, b.y, *b)
		for _, a := range s1 {
			fmt.Printf("\n  Checking against %d, %d... %T", a.x, a.y, *a)
			if *b == *a {
				fmt.Printf("    MATCH")
				foundMatch = true
				break
			}
		}

		if !foundMatch {
			return false, fmt.Errorf("Expected to find %d, %d in s1, but did not.", b.x, b.y)
		}
	}

	return true, nil
}
