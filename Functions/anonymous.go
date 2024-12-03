package main

// import "fmt"

// func main() {
// 	numbers := []int{1, 2, 3, 4}

// 	doubled := transformNumbers(&numbers, double)
// 	tripled := transformNumbers(&numbers, triple)

// 	moreNumbers := []int{5, 1, 2}

// 	fmt.Println(doubled)
// 	fmt.Println(tripled)

// 	transformerFn1 := getTransformFunction(&numbers)
// 	transformerFn2 := getTransformFunction(&moreNumbers)

// 	transformedNumbers := transformNumbers(&numbers, transformerFn1)
// 	moreTransformedNumbers := transformNumbers(&numbers, transformerFn2)

// 	fmt.Println(transformedNumbers)
// 	fmt.Println(moreTransformedNumbers)
// }

// // function passed
// // define args + types & return types + nice to add custom type
// type TransformFn func(int) int

// func transformNumbers(numbers *[]int, transform TransformFn) []int {
// 	dNumbers := []int{}

// 	for _, val := range *numbers {
// 		dNumbers = append(dNumbers, transform(val))
// 	}

// 	return dNumbers
// }

// func double(number int) int {
// 	return number * 2
// }

// func triple(number int) int {
// 	return number * 3
// }

// // returning a function
// func getTransformFunction(numbers *[]int) TransformFn {
// 	if (*numbers)[0] == 1 {
// 		return double
// 	} else {
// 		return triple
// 	}
// }
