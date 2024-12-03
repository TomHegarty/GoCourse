package main

import "fmt"

func main() {
	fact := factorial(5)
	fmt.Println(fact)

	numbers := []int{1, 10, 15}
	sum := sumup(numbers...)

	fmt.Println(sum)

}

func factorial(number int) int {

	// recursive approach
	if number == 0 {
		return 1
	}
	return number * factorial(number-1)

	// result := 1

	// for i := 1; 1 <= number; i++ {
	// 	result = result * i
	// }

	// return result
}

func sumup(numbers ...int) int {
	sum := 0

	for _, val := range numbers {
		sum += val
	}

	return sum
}
