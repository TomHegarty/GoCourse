package main

import "fmt"

func main() {
	age := 32 // only created once

	agePointer := &age

	fmt.Println("Age: ", age)
	editAgeToAdultYears(agePointer)
	fmt.Println("Adult Years: ", age)
}

func editAgeToAdultYears(age *int) {
	// return *age - 18
	*age = *age - 18
}
