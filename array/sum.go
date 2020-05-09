package main

func Sum(numbers []int) int {
	sum := 0
	for _, number := range numbers {
		sum += number
	}
	return sum
}

func SumAll(numbersToSum ...[]int) []int {

	// lengthOfNumbers := len(numbersToSum)
	// sums := make([]int, lengthOfNumbers)
	var sums []int //starts off with an empty slice

	for _, numbers := range numbersToSum {
		// sums[i] = Sum(numbers)
		sums = append(sums, Sum(numbers)) //adds a new index to the sums slice
	}

	return sums

}
