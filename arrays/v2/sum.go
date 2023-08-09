package main

func Sum(numbers [5]int) int {
	total := 0
	for _, number := range numbers {
		total += number
	}
	return total
	// returns total, which is the output declared in func declaration.
}
