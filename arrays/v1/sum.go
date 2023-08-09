package main

func Sum(numbers [5]int) (total int) {

	for i := 0; i < len(numbers); i++ {
		total += numbers[i]
	}
	return
	// returns total, which is the output declared in func declaration.
}
