package main

// import "fmt"

// func main() {
// 	numbers := []int{1, 2, 3}

// 	double := createTransformer(2)
// 	triple := createTransformer(3)
// 	quadruple := createTransformer(4)

// 	transformed := transformNumbers(&numbers, func(number int) int {
// 		return number * 2
// 	})

// 	fmt.Println(transformed)

// 	doubled := double(transformed[0])
// 	tripled := triple(transformed[0])
// 	quadrupled := quadruple(transformed[0])

// 	fmt.Println(doubled)
// 	fmt.Println(tripled)
// 	fmt.Println(quadrupled)
// }

// func transformNumbers(numbers *[]int, transform func(int) int) []int {
// 	dNumbers := []int{}

// 	for _, val := range *numbers {
// 		dNumbers = append(dNumbers, transform(val))
// 	}

// 	return dNumbers
// }

// // factory function
// func createTransformer(factor int) func(int) int {
// 	return func(number int) int {
// 		return number * factor
// 	}
// }
