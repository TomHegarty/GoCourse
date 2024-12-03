package lists

import "fmt"

// func main() {
// 	// var productNames [4]string = [4]string{"a book"}
// 	// prices := []float64{10.99, 9.99, 45.99, 20.00}
// 	// featuredPrices := prices[1:2]
// 	// highlightedPrices := prices[:1]

// 	// // fmt.Println(prices[0])
// 	// // fmt.Println(productNames)

// 	// fmt.Println(prices[:])
// 	// fmt.Println(len(prices))
// 	// fmt.Println(len(featuredPrices))
// 	// fmt.Println(len(highlightedPrices))

// 	// fmt.Println(cap(prices))
// 	// fmt.Println(cap(featuredPrices))
// 	// fmt.Println(cap(highlightedPrices))

// 	prices := []float64{10.99, 8.99, 45.99, 20.00}
// 	prices[1] = 9.99
// 	prices[2] = 11.99

// 	fmt.Println(prices)

// 	prices = append(prices, 5.99)

// 	fmt.Println(prices)

// 	prices = prices[1:]

// 	fmt.Println(prices)
// }

type Product struct {
	name        string
	description string
	price       float64
}

func main() {

	hobbies := [3]string{"biking", "magic", "counter-strike"}

	fmt.Println(hobbies)
	fmt.Println(hobbies[0])
	fmt.Println(hobbies[1:])

	firstSlice := hobbies[1:]

	hobbiesSlice := firstSlice[:2]
	// hobbiesSlice2 := hobbies[0:2]
	fmt.Println(hobbiesSlice)
	// fmt.Println(hobbiesSlice2)

	goals := []string{"build API,", "make magic App,"}

	goals[1] = "make BETTER magic app,"
	goals = append(goals, "get money")

	fmt.Println(goals)

	products := []Product{
		{
			"first-product",
			"A first Product",
			12.99,
		},
		{
			"second-product",
			"A second Product",
			8.99,
		},
	}

	fmt.Println(products)

	newProduct := Product{"third-product", "a new product", 3.99}

	products = append(products, newProduct)
	fmt.Println(products)

	extraProducts := []Product{
		{
			"first-product",
			"A first Product",
			12.99,
		},
		{
			"second-product",
			"A second Product",
			8.99,
		},
	}

	products = append(products, extraProducts...)
	fmt.Println(products)
}

// Time to practice what you learned!

// 1) Create a new array (!) that contains three hobbies you have
// 		Output (print) that array in the command line.
// 2) Also output more data about that array:
//		- The first element (standalone)
//		- The second and third element combined as a new list
// 3) Create a slice based on the first element that contains
//		the first and second elements.
//		Create that slice in two different ways (i.e. create two slices in the end)
// 4) Re-slice the slice from (3) and change it to contain the second
//		and last element of the original array.
// 5) Create a "dynamic array" that contains your course goals (at least 2 goals)
// 6) Set the second goal to a different one AND then add a third goal to that existing dynamic array
// 7) Bonus: Create a "Product" struct with title, id, price and create a
//		dynamic list of products (at least 2 products).
//		Then add a third product to the existing list of products.
